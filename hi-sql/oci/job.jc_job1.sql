-- link: https://www.cnblogs.com/ulex/p/4415478.html
-- 2018年12月28日 星期五
-- Joshua Conero

-- 创建测试表
DROP TABLE jc_job1_tb;

-- CREATE OR REPLACE TABLE jc_job1_tb(
CREATE TABLE jc_job1_tb(
    name varchar(30) not null,
    int_x NUMBER(10) NOT NULL,
    year NUMBER(4) NOT NULL,
    score NUMBER(3) NOT NULL,
    hash VARCHAR2(100) NOT NULL,
    mtime date not null
);

/
-- 创建 jc_job1_proc 存储过程
CREATE OR REPLACE PROCEDURE jc_job1_proc IS
    BEGIN
        INSERT INTO jc_job1_tb VALUES(
            'JOSHUA CONERO',
            DBMS_RANDOM.RANDOM,            
            CEIL(DBMS_RANDOM.VALUE(1949, 2949)),
            FLOOR(DBMS_RANDOM.VALUE()*100),
            dbms_random.string('P',100),
            SYSDATE
        );
    END;

/
-- 创建 JOB
DECLARE job_n NUMBER;
    BEGIN
        DBMS_JOB.SUBMIT(
            job_n,
            'jc_job1_proc;',
            SYSDATE,
            'TRUNC(sysdate, ''mi'')+1/(24*60)'           -- 每分钟执行
        );
        COMMIT;
    END;

/*
-- 删除 JOB

-- 查找所有 JOB
SELECT J.* FROM DBA_JOBS J;
/
DECLARE 
    job_id NUMBER := 0;
BEGIN
    SELECT J.JOB INTO job_id FROM DBA_JOBS J WHERE J.WHAT = 'jc_job1_proc;';
    IF job_ID <> 0 THEN
        DBMS_JOB.REMOVE(job_id);
        COMMIT;
    END IF;
END;
/
DROP PROCEDURE jc_job1_proc;
DROP TABLE jc_job1_tb;
*/    