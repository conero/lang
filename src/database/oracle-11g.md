# Oracle-11g

> Joshua Conero
>
> 2018年8月28日 星期二



## 基础



### 查询

> select 中包含 *DML* 会报错

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



## 系统/DBA



> session

```sql
alter session set nls_date_format='yyyy-mm-dd hh24:mi:ss';
```



### objects

> 数据库对象

- 相关表
  - *all_objects*     所有用户下对象
  - *dba_objects*  dba所有对象
  - *user_objects*  当前用户所有对象





## PL/SQL

### 循环

> 死循环可能操作 “锁”现象



### 常见问题



#### NO_DATA_FOUND

> PL/SQL 内部查询数据为空时，会抛出异常 *NO DATA FIND*， 而是用 NVL/NVL2 无效(整体记录不存在时)



## oracle /Sys

### Pseudocolumns/ 伪列

> Hierarchical Query   (级联查询)   connect by。 级联伪列

1. `CONNECT_BY_ISCYCLE=1/0`   		是否存在循环
2. `CONNECT_BY_ISLEAF=1/0`                  是否为*页*
3. `LEVEL`                                                    层级数



> Sequence Pseudocolumns    序列伪列

*用于触发器等*

1. `sequence.CURRVAL`   序列当前的值
2. `sequence.NEXTVAL` 序列下一个值







### DUAL

> 单行文件



### 随机数

> *dbms_random* /   ``DBMS_RANDOM``

- dbms_random.value     小数(0-1)
- select trunc(dbms_random.value(0,100)) from dual        （0， 100） 的随机数
- dbms_random.normal    正态分布随机数
- dbms_random.string(type, len)



> dbms_random.string(type)

 /* type可取值如下：
      'u','U'    :    大写字母
      'l','L'    :    小写字母
      'a','A'    :    大、小写字母
      'x','X'    :    数字、大写字母
      'p','P'    :    可打印字符

*/



### 用户密码过期

错误信息： `SQLSTATE[HY000]: OCISessionBegin: OCI_SUCCESS_WITH_INFO: ORA-28002: the password will expire within 7 days`



> 解决方法

```sql
-- 查询密码状态
SELECT * FROM dba_profiles s WHERE s.profile='DEFAULT' AND resource_name='PASSWORD_LIFE_TIME';

-- 修改密码为： 将密码设置成永不过期
ALTER PROFILE DEFAULT LIMIT PASSWORD_LIFE_TIME UNLIMITED;

-- 更新密码
alter user uname identified by pswd;
```





## 参照

- [Oracle 锁机制探究](https://www.cnblogs.com/leohahah/p/7039907.html)

