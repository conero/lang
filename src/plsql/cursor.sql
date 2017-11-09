-- 2017年11月9日 星期四
-- Oracle 学习
-- 游标


-- 动态删除表格/1
    DECLARE
    v_sql VARCHAR2(1500);
    TYPE tb_records IS RECORD(
        TABLE_NAME  VARCHAR2(100)
    );
    -- 游标
    CURSOR tb_cur IS
        SELECT TABLE_NAME FROM all_tables WHERE OWNER='SA' AND table_name LIKE '%-%' and LENGTH(table_name) > 10
        ;
    BEGIN
    FOR tb_records IN tb_cur LOOP
        v_sql := 'DROP TABLE "'||tb_records.TABLE_NAME||'"';
        -- 输出结果集
        -- SYS.DBMS_OUTPUT.PUT_LINE(v_sql);
        -- DBMS_OUTPUT.PUT_LINE(sysdate);
        EXECUTE IMMEDIATE v_sql;
        DBMS_OUTPUT.PUT_LINE(v_sql);
    END LOOP;
    END;


--- 游标测试/2  -- 死循环!!
    DECLARE
    tname VARCHAR2(100);
    v_sql VARCHAR2(1500);
    -- 游标
    CURSOR tb_cur IS
        SELECT TABLE_NAME FROM all_tables WHERE 
        OWNER='SA' AND table_name LIKE '%-%' 
        -- and LENGTH(table_name) > 10
        ;
    BEGIN
    OPEN tb_cur;    -- 打开游标
    -- FETCH tb_cur INTO tname;
    FETCH tb_cur INTO tname;
    LOOP
        EXIT WHEN NOT tb_cur%FOUND;
        v_sql := 'DROP TABLE "'||tname||'"';
        -- 输出结果集
        -- SYS.DBMS_OUTPUT.PUT_LINE(v_sql);
        DBMS_OUTPUT.PUT_LINE(v_sql);
    END LOOP;
    CLOSE tb_cur;
    END;    