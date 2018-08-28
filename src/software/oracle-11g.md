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



#### 分类

- DML锁（data locks，数据锁），用于保护数据的完整性
- DDL锁（dictionary locks，字典锁），用于保护数据库对象的结构，如表、索引等的结构定义
- 内部锁和闩（internal locks and latches），保护数据库的内部结构





## 参照

- [Oracle 锁机制探究](https://www.cnblogs.com/leohahah/p/7039907.html)

