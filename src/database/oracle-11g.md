# Oracle

> Joshua Conero
>
> 2018年8月28日 星期二



知识基于版本 oracle 11g 或 oracle 12c。



## 基础

默认端口：1521



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



### 用户

默认管理员账号：

*system*



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



### 命令或工具

```shell
# 查看网络连接状态
lsnrctl status
```





### 备份及恢复

导入sql文件的方法

- 使用 Oracle SQL Developer，以sql查询文件打开到查询控制台中，执行运行脚本即可
- 使用 sqlplus 工具，连接到数据库中，使用命令如 `@e:/apps/DZZZEX.sql` 导入数据



```shell
# 创建 DUMP_DIR，sql 对话中执行
create directory DUMP_DIR as 'E:\dbData';

# 使用 impdb 导入sql
impdp system/password@dzzzex25 directory=DUMP_DIR dumpfile=ORACLE_dzzzex.dbmd full=y
```





#### sql plus

按照提供登录用户，相关信息查看

执行用户和密码进行数据库登录，中文乱码时可使用Windows 记事本保存为 `ANSI` 编码格式

```shell
# 指定数据库及用户密码登录数据库
.\sqlplus.exe 'system/passwd@sid'

# 查看是否支持 utf8
SELECT value AS database_charset FROM nls_database_parameters WHERE parameter = 'NLS_CHARACTERSET';

# 设置 utf8 编码支持
# windows pwsh 环境变量
$env:NLS_LANG='AMERICAN_AMERICA.AL32UTF8'
```



sql plus 相关操作

```sql
# 查看当前用户信息
SELECT name, dbid, created, log_mode FROM v$database;

# 执行sql
@E:\dbData\DZZZEX.sql
# 或使用 start 执行sql
start E:\dbData\DZZZEX.sql
# 提交
commit
```







### 版本介绍

- Oracle 11g                      发布于2007.7 月，15年前后接触的大多数工厂都使用该版本。广泛应用于企业级应用、数据仓库、OLTP（在线事务处理）等领域。
- Oracle 12c                      发布于2013.6 月，引入多租户架构、云数据库的概念。适用于企业级应用、云计算、大数据等领域。
- Oracle 18c                      发布于2018.9 月，进一步增强了多租户架构和性能，优化自动化和云集成功能。
- Oracle 19c                      发布于2020.6 月，强调稳定性和性能优化，增强对人工智能领域的支持。
- Oracle 21c                      发布于2021
- Oracle 22c                      发布于2022，集成最新的AI、机器学习和云技术，适应现代分布式和边缘计算需求。
- Oracle 23ai                     发布于2024.5，结合人工智能（AI）功能，进一步增强数据库的智能化管理和数据分析能力。



## 参照

- [Oracle 锁机制探究](https://www.cnblogs.com/leohahah/p/7039907.html)
- https://edelivery.oracle.com   Oracle 数据库等下载地址
- [解决Oracle 本地可以连接，远程不能连接问题](https://www.cnblogs.com/xyt-0412/p/9897180.html) 
- [DBeaver添加阿里maven镜像](https://developer.aliyun.com/article/1551241)
- [oracle-base 安装包下载](https://oracle-base.com/)

