/**
  2018年8月15日 星期三
  oracle 辅助程序包： oracle 11g 测试
  Joshua Conero
 */


-- 数据类型
    -- 先删除程序包，再 删除关联的类型
        jc_ref_obj;
        DROP PACKAGE jc;
          DROP TYPE jc_msg_obj;
          DROP TYPE jc_ref_obj;
-- type
/
   -- 无类型引用
    CREATE OR REPLACE TYPE jc_msg_obj AS OBJECT(
        "error" NUMBER,                  -- 错误， 0/1
        "msg" VARCHAR2(100),              -- 错误消息
        "remark" VARCHAR2(254)
    )
    ;
/
    -- 无类型引用
    CREATE OR REPLACE TYPE jc_ref_obj AS OBJECT(
        "table" VARCHAR2(80),
        "colname" VARCHAR2(100),
        "r_table" VARCHAR2(80),
        "r_colname" VARCHAR2(100),
        "const_name" VARCHAR2(200),              -- 错误消息
        "owner" VARCHAR2(30)
    )
    ;

-- 包头
/
    CREATE OR REPLACE PACKAGE jc IS

      -- 外部质量信息管控表
      type msg_table is table of jc_msg_obj;

      -- 删除所有数据, (1 存在, 不存在)
      FUNCTION clearData(tname VARCHAR2, whstr VARCHAR2 := '', isStr CHAR := 'N') RETURN msg_table PIPELINED;

      -- 表格是否存在, (1 存在, 不存在)
      FUNCTION table_exist(tname VARCHAR2) RETURN NUMBER;

      -- 表依赖关系
      TYPE reftable IS TABLE OF jc_ref_obj;
      FUNCTION ref_table(tname VARCHAR2 := '') RETURN reftable PIPELINED;
      -- 表所有子表引用查询
      FUNCTION ref_ctable(tname VARCHAR2) RETURN reftable PIPELINED;

      descript VARCHAR2(200) := 'oracle 辅助程序包';
      version VARCHAR2(10) := '0.0.1';
      release DATE := to_date('20180815', 'YYYYMMDD');
      build DATE := to_date('20180815', 'YYYYMMDD');

    END jc;

/

-- 定义包程序
    CREATE OR REPLACE PACKAGE BODY jc IS
        FUNCTION clearData(tname VARCHAR2, whstr VARCHAR2 := '', isStr CHAR := 'N') RETURN msg_table PIPELINED IS
            v_r msg_table;
            v_sql VARCHAR2(1000);
            v_msg VARCHAR2(100);
          BEGIN
            IF table_exist(tname) < 1 THEN
              v_msg := tname || '-数据表不存在';
              PIPE ROW (jc_msg_obj(1, v_msg, NULL ));
            -- ELSE
            END IF;
            RETURN;
          END clearData
          ;

        FUNCTION table_exist(tname VARCHAR2) RETURN NUMBER IS
            v_count NUMBER := 0;
          BEGIN
              SELECT COUNT(1) INTO v_count FROM USER_TABLES where table_name = tname;
          RETURN v_count;
          end table_exist
          ;

        FUNCTION ref_table(tname VARCHAR2 := '') RETURN reftable PIPELINED IS
            TYPE t_refcursor IS REF CURSOR;
            c_src t_refcursor;
            v_r jc_ref_obj;
          BEGIN
            -- 游标处理
            OPEN c_src FOR -- 存在，但是未保存的数据
              select
                  -- a.table_name as "table", a.column_name as "column", a.owner as "owner", a.CONSTRAINT_NAME as "const_name",
                  -- b.CONSTRAINT_TYPE as "const_type",
                  -- c.table_name as "r_table", d.COLUMN_NAME as "r_col"
                  jc_ref_obj(a.table_name, a.column_name, c.table_name, d.COLUMN_NAME, a.CONSTRAINT_NAME, a.owner)
                from USER_CONS_COLUMNS a
                  left join USER_CONSTRAINTS b on a.CONSTRAINT_NAME = b.CONSTRAINT_NAME and a.TABLE_NAME=b.TABLE_NAME
                  join USER_CONSTRAINTS b on a.CONSTRAINT_NAME = b.CONSTRAINT_NAME and a.TABLE_NAME=b.TABLE_NAME
                  left join USER_CONSTRAINTS c on c.CONSTRAINT_NAME = b.R_CONSTRAINT_NAME
                  left join USER_CONS_COLUMNS d on c.CONSTRAINT_NAME = d.CONSTRAINT_NAME and c.TABLE_NAME = d.TABLE_NAME
                  where b.CONSTRAINT_TYPE = 'R'
                  and nvl2(tname, a.table_name, '1') = nvl2(tname, tname, '1')
                ;
            LOOP
                FETCH c_src INTO v_r;
                EXIT WHEN c_src%NOTFOUND;
                PIPE ROW (v_r);
            END LOOP;

            CLOSE c_src;
          END ref_table
          ;
        FUNCTION ref_ctable(tname VARCHAR2) RETURN reftable PIPELINED IS
            TYPE t_refcursor IS REF CURSOR;
            c_src t_refcursor;
            v_r jc_ref_obj;
          BEGIN
            -- 游标处理
            OPEN c_src FOR -- 存在，但是未保存的数据
              select
                jc_ref_obj(a."table", a."colname", a."r_table", a."r_colname", a."const_name", a."owner")
              from table(JC.REF_TABLE()) a
                 start with a."r_table"=tname
                CONNECT BY a."r_table" = PRIOR a."table"
              ;
            LOOP
                FETCH c_src INTO v_r;
                EXIT WHEN c_src%NOTFOUND;
                PIPE ROW (v_r);
            END LOOP;

            CLOSE c_src;
          END ref_ctable
          ;
    END jc;
