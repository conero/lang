

# MySQL

> 2019年3月18日 星期一
>
> Joshua Conero



## 关于

*最流行的关系型数据库之一，常用于中小型 WEB应用。*

*瑞典MySQL AB 公司开发，后被 Oracle 收购。分支数据库 `MariaDB`*

*由 C/C++ 编写而成*

*容量：例子有支持 `5亿` 行记录的表。如含20万数据库表、50亿条记录。*





## 安装

### 配置

数据库表：*performance_schema*

```mysql
# MySQL ，通过 SQL 修改 数据库的配置信息
# performance_schema.global_variables 全局配置
set global max_allowed_packet = 200*1024*1024;

# 查看配置
show variables like '%like%';
```





### Zip file(8.0)

> Windows 版本下安装

1. 目录指向`$/bin`
2. 初始化 `$ mysqld --initialize --console` 来初始化安装。
3. 安装 `$ mysqld install`
4. 启动 `$ net start mysql`

*初始化时可获得随机密码。使用 cmd 登录 mysql 需要重置密码：*

```sql
alter user 'root'@'localhost' identified by 'password';
```



> *版本问题*

- *v8.0 myAdmin 登录失败*

  - 可通过SQL更改密码方式:

    ```sql
    ALTER USER 'root'@'localhost' IDENTIFIED BY 'password' PASSWORD EXPIRE NEVER; 
    ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';
    ```




### 5.7 

> Windows 版本安装；

*下载官网 msi，根据向导安装需要的应用。安装后需要执行命令：*

1.  初始化 `$ /bin/mysqld --initialize-insecure`
2. 安装 `$ /bin/mysqld install`
3. 启动 `$ net start mysql`



## SQL

### 主键/常量等

```mysql
-- 删除表格的主键
alter table TABLENAME drop primary key;
```



### drop

```mysql
DROP TABLE [IF EXISTS] tbl_name;		-- 删除存在数据表
```





## 基础

### sql_mode

```sql
-- 查询配置信息
show variables like '%sql_mode%'; 
-- 默认结果
/*
ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION
*/


-- 删除 ONLY_FULL_GROUP_BY
set global `sql_mode`='STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';

```



### 数据库引擎

*数据库存储引擎是数据库底层软件组织，数据库管理系统（DBMS）使用数据引擎进行创建、查询、更新和删除数据。不同的存储引擎提供不同的存储机制、索引技巧、锁定水平等功能，使用不同的存储引擎，还可以 获得特定的功能。*

```mysql
-- 查看MySQL数据库引擎
show engines;
```



#### InnoDB

*MySQL5.7/8.0 等版本的默认引擎，支持事务(Transcations)，行锁定和外键。*

- *提供了具有提交、回滚和崩溃恢复能力的事物安全（ACID兼容）存储引擎：事务(Transcations)*
- *InnoDB是为处理巨大数据量的最大性能设计*
- *InnoDB存储引擎完全与MySQL服务器整合，InnoDB存储引擎为在主内存中缓存数据和索引而维持它自己的缓冲池。InnoDB将它的表和索引在一个逻辑表空间中，表空间可以包含数个文件（或原始磁盘文件）*
- *InnoDB支持外键完整性约束，存储表中的数据时，每张表的存储都按主键顺序存放，如果没有显示在表定义时指定主键，InnoDB会为每一行生成一个6字节的ROWID，并以此作为主键*
- *InnoDB被用在众多需要高性能的大型数据库站点上*



*数据库引擎是用于存储、处理和保护数据的核心服务。利用数据库引擎可控制访问权限并快速处理事务，从而满足企业内大多数需要处理大量数据的应用程序的要求。   使用数据库引擎创建用于联机事务处理或联机分析处理数据的关系数据库。这包括创建用于存储数据的表和用于查看、管理和保护数据安全的数据库对象（如索引、视图和存储过程）。*





#### MyISAM

*基于ISAM存储引擎，并对其进行扩展；它是在Web、数据仓储和其他应用环境下最常使用的存储引擎之一。MyISAM拥有较高的插入、查询速度，但不支持事物。*



#### MEMORY

*内存引擎，为查询和引用其他表数据提供快速访问。*



## issue

### time_zone 错误，jdbc 连接解决

```
# java.lang.RuntimeException: com.mysql.cj.exceptions.InvalidConnectionAttributeException: The server time zone value 'ÖÐ¹ú±ê×¼Ê±¼ä' is unrecognized or represents more than one time zone. You must configure either the server or JDBC driver (via the serverTimezone configuration property) to use a more specifc time zone value if you want to utilize time zone support.
```



> *设置时间区就可以*

```sql
set global time_zone ='+8:00';
```





## 备份

> mysql命令行导入数据，*数据库管理软件/工具中使用 source 导入数据无效*。

```powershell
# 以 utf8 的编码登录到 mysql 中
mysql -u root -p --default-character-set=utf8 

# 字符集可通过 sql 执行替换
set names 'uft8';
# 显示当前的状态，包括字符集
status

# 选择数据库
use dbname;
# 使用 source 导入数据库十分的缓慢
# 使用 source 导入脚本
source d:/.../name.sql


# 数据库复制 oldDbname => <toNewDb> 数据库已经存在的的数据库
mysqldump <oldDbname> -u root -p<password> --add-drop-table | mysql <toNewDb> -u root -p<password>
```





### mysqldump

> 本地导出数据为SQL

```powershell
# mysqldump -uroot -p --default-character-set=utf8 dbname > d:/.../file.sql
mysqldump -uroot -p --default-character-set=utf8 dataset > d:/tmp/dataset.sql

# 经测试只导出一行
# needtoKnowMore
mysqldump -uroot -p --default-character-set=utf8 dataset tablename > d:/tmp/tablename.sql
mysqldump -uroot -p --default-character-set=utf8 dataset tablename --result-file=d:/tmp/tablename.sql
```





> 单实例： mysqldump -h192.168.48.131 -uroot -proot --opt dbname1 | mysql -h192.168.48.1 -uroot -proot -C dbname2   从131服务器上将dbname1数据库备份到1上。



//@TODO <https://blog.csdn.net/u011302734/article/details/74936470>



### binlog

> 查看数据库 bin 日志

```mysql
-- 查看binlog日志是否开启
show variables like 'log_%'; 

-- 查看所有 binlog 日志列表
show master logs;

-- 查看二进制状态
show master status;
-- 日志刷新
flush logs;
```



//@TODO *window 下如开启 binlog ，实现数据备份的解决方案。*