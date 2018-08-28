# Oracle-11g

> Joshua Conero
>
> 2018年8月28日 星期二



## 基础



### 查询

> select 中包含 *ddl* 会报错

``错误ORA-14551: 无法在查询中执行 DML 操作``





### 锁

> 查询 *DDL* 锁

```sql
-- 所有锁
SELECT * FROM dba_ddl_locks;
-- session sid=dba_ddl_locks.session_id
select * from v$session where sid=192;

-- 通过杀死 session 解锁
--杀掉进程 sid,serial#
alter system kill session'10,15';

```



