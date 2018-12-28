-- 创建测试表
CREATE OR REPLACE jc_job1_tb(
    name varchar(30) not null,
    int_x NUMBER(10) NOT NULL,
    year NUMBER(4) NOT NULL,
    score NUMBER(3) NOT NULL,
    hash VARCHAR2(100) NOT NULL,
    mtime date not null
);

/
-- 创建 jc_job1_proc 存储过程
CREATE OR REPLACE PROCEDURE jc_job1_proc
    BEGIN
        INSERT INTO jc_job1_tb VALUES(
            'JOSHUA CONERO',
            DBMS_RANDOM.RANDOM,            
            CEIL(DBMS_RANDOM.VALUE(1949, 2949),
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
            '',
            SYSDATE,
            'TRUNC(sysdate,''''mi'''')+1/(24*60)'           -- 每分钟执行
        );
        COMMIT;
    END;