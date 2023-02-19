

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



*通过 `set` 修改的参数只作用于当前的会话中；`set global` 再服务重启以后就会失效；`my.ini/my.conf` 的配置应该是永久的。*



### Zip file(8.0)

> Windows 版本下安装

1. 目录指向`$/bin`
2. 初始化 `$ mysqld --initialize --console` 来初始化安装。
3. 安装 `$ mysqld install [<server_name>]`, *server_name 默认为 mysql*
4. 启动 `$ net start <server_name>` 启动服务

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



### mysql8.0 rpm 版本安装

操作系统环境：`centOS8 x86_64` MySQL8.0.27



**依赖关系**

- mysql-community-server-8.0.27-1.el8.x86_64.rpm
    - mysql-community-8.0.27-1.el8.src.rpm
    - mysql-community-client-8.0.27-1.el8.x86_64.rpm
        - mysql-community-client-plugins
        - mysql-community-libs

```shell
# 依赖安装示例
wget https://repo.mysql.com/yum/mysql-8.0-community/el/8/x86_64/mysql-community-client-8.0.27-1.el8.x86_64.rpm
wget https://repo.mysql.com/yum/mysql-8.0-community/el/8/x86_64/mysql-community-common-8.0.27-1.el8.x86_64.rpm

# rpm 安装
rpm -ivh ./mysql-community-server-8.0.27-1.el8.x86_64.rpm

# mysqld 初始化
mysqld --initialize --console

# 启动服务
service mysqld start
# 服务状态查询
service mysqld status

# 查看 mysql 安装情况
rpm -qa |grep mysql

# mysql 无权限的问题。对资源文件进行授权
chown mysql.mysql -R /var/lib/mysql
# 查看原始密码
 cat /var/log/mysqld.log
```




### 5.7 

> Windows 版本安装；

*下载官网 msi，根据向导安装需要的应用。安装后需要执行命令：*

1.  初始化 `$ /bin/mysqld --initialize-insecure`
2. 安装 `$ /bin/mysqld install`
3. 启动 `$ net start mysql`



```mysql
# 环境初始化
mysqld --initialize --console


-- 安装项目完成是，密码为空；这是需要自行设置密码
alter user 'root'@'localhost' identified by 'password';
```





#### mysql5.7 rpm Centos7

CentOs7 项目安装：

```shell
#CentOS7 mysql 安装
# 下载相关的 rpm 包文件
wget https://repo.mysql.com/yum/mysql-5.7-community/el/7/x86_64/mysql-community-server-5.7.38-1.el7.x86_64.rpm
wget https://repo.mysql.com/yum/mysql-5.7-community/el/7/x86_64/mysql-community-common-5.7.38-1.el7.x86_64.rpm
wget https://repo.mysql.com/yum/mysql-5.7-community/el/7/x86_64/mysql-community-client-5.7.38-1.el7.x86_64.rpm
wget https://repo.mysql.com/yum/mysql-5.7-community/el/7/x86_64/mysql-community-libs-5.7.38-1.el7.x86_64.rpm


# 执行安装
rpm -ivh mysql-community-common-5.7.38-1.el7.x86_64.rpm
rpm -ivh mysql-community-libs-5.7.38-1.el7.x86_64.rpm
rpm -ivh mysql-community-client-5.7.38-1.el7.x86_64.rpm
rpm -ivh libaio-0.3.109-13.el7.x86_64.rpm

# mysql 初始化
mysqld --initialize --console

# mysql 权限设置
chmod -R 777 /var/lib/mysql

# 启动服务
service mysqld start
```



依赖关系


- mysql-community-server
    - mysql-community-client
        - mysql-community-common
    - mysql-community-libs
    - libaio-0.3.109-13.el7.x86_64.rpm





### 附带的工具/命令集

- mysql				*数据 SQL 交互对话，用于表查询等*
- mysqldump, mysqlpump     *数据库备份工具*
- mysqlimport     *数据导入*
- mysqlbinlog      数据库二进制日志





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

检查慢日志记录是找出扫描行数过多查询的好方法。



### 连接

*shell 数据库连接。*

```shell
# 数据库登录连接选项；可配置项： host/user/database
# host 默认为 localhost
mysql -h host -u user -p [database]
mysql --host=localhost --user=uname --password=password [database]

# 其他选项
--port,-P  							# 端口号配置
--default-character-set=name		# 设置默认的编码
--xml,-X							# 输出为 XML 字符串格式
--html,-H							# 输出为 HTML 字符串格式

# 连接数据库并执行SQL
mysql [-e sql|--execute="sql"]	



# 查看 shell 的连接信息
status
help	# 查看mysql shell 内部的命令
```



### 属性相关的函数

```mysql
user()		-- 当前的用户信息
version()	-- 版本信息
```



### 事务

ACID 事务原子性：

- atomicity       原子性
- consistency   一致性
- isolation        隔离性
- duration        持久性



数据库死锁可能导致非常慢的查询，即使查询变慢。



> 事务

MySQL 默认事务自动提交(Autocommit)

```mysql
-- 启动事务
start transaction;
-- 执行数据库操作

-- 提交事务
commit;

-- 事务回滚
rollback;
```



MySQL将每个数据库(也称作schema)持久化在数据目录的子目录下。

- 创建表时生成一个与表名同名的“.frm”文件，其保存表定义。



```mysql
# 查看表-x_user-状态
show table status like 'x_user';
```





### 语句

```mysql
-- 创建数据库
create database conero DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

-- 数据库删除
drop database conero;
```





#### 类型



> 选择优化的数据类型

- 更小的通常更好。比如确定数据的取值范围后选择最匹配的类型，不要选择过大于它的值。
- 尽可能用简单类型。
- 尽量避免NULL值。



##### 时间类型

```mysql
tear 		-- 年。       字节数为1,  取值范围: [1901-2155]
date 		-- 日期。     字节数为4,  取值范围: [1000-01-01] ~ [9999-12-31]
time		-- 时间。     字节数为3,  取值范围: [-838:59:59] ~ [838:59:59]
datetime	-- 日期时间。  字节数为8,  取值范围: [1000-01-01 00:00:00] ~ [9999-12-31 23:59:59]
timestamp	-- 时间戳。    字节数为4,  取值范围: [19700101080001 - 2038011911407]


-- date
-- curdate() 同义函数 current_date(), current_date
-- now() 同义函数 current_timestamp(), current_timestamp;
select
	curdate() as 当前日期, current_date as `常量同义-日期`,
	curtime() as 当前时间, current_time as `常量同义-时间`,
	curtime() as 当前时间戳, now(), current_timestamp as `常量同义-时间戳`,
	UNIX_TIMESTAMP(now()) as `时间转时间戳`,
	UNIX_TIMESTAMP('2006-01-02 15:04:05') as `时间转时间戳2`,
	FROM_UNIXTIME(UNIX_TIMESTAMP(NOW())) as `时间戳转时间`
;


-- 计算时间查：datetime 类型
select r.lock_tm, current_timestamp() `now`, TIMESTAMPDIFF(second, r.lock_tm, current_timestamp()) diff_sec from r_beffq_video_attach_queue r;

-- 日期格式
select date_format(now(), '%Y-%m-%d %h:%i:%s');
select date_format(now(), '%Y/%m/%d %h:%i:%s') v_time;
```



##### 字符串

**文本**

- *TinyText*           最大长度255个字节(2^8-1)
- *Text*                  最大长度65535个字节(2^16-1)



##### 数字

分为：整数(int) 和 实数(decimal)。



**int**

| 类型        | 有符号范围     | 无符号范围 | 存储字节(1Byte=8 bit) |
| ----------- | -------------- | ---------- | --------------------- |
| bigint      | -2^63 ~ 2^63-1 | 0 ~ 2^64-1 | 8                     |
| int/integer | -2^31 ~ 2^31-1 | 0 ~ 2^32-1 | 4                     |
| mediumint   | -2^23 ~ 2^23-1 | 0 ~ 2^24-1 | 3                     |
| smallint    | -2^15 ~ 2^15-1 | 0 ~ 2^16-1 | 2                     |
| tinyint     | -2^7 ~ 2^7-1   | 0 ~ 2^8-1  | 1                     |

**unsigned**       非负属性，提交数据的保存范围。



#### show

可用于查询对象： *databases, tables, variables, columns, server status* 等

```mysql
-- 通用查询方法
show databases [like ''];		-- 显示所有的数据库，可使用 like 查询；

-- 数据表查看
show tables;

-- 查询某个数据的 DDL
show create table <tableName>;

-- 查询某个数据表的列
-- 显示全部列
show columns from <tableName>;
-- 含条件查询
show columns from <tableName> where field like 'i%';

-- 二进制日志
show {binary | master} logs;
```



> 参考

```mysql

-- 所有参考
-- 二进制日志
SHOW {BINARY | MASTER} LOGS
SHOW BINLOG EVENTS [IN 'log_name'] [FROM pos] [LIMIT [offset,] row_count]
-- 字符集
SHOW CHARACTER SET [like_or_where]
SHOW COLLATION [like_or_where]			
-- 查询列
SHOW [FULL] COLUMNS FROM tbl_name [FROM db_name] [like_or_where]
-- 查询 DDL
SHOW CREATE DATABASE db_name
SHOW CREATE EVENT event_name
SHOW CREATE FUNCTION func_name
SHOW CREATE PROCEDURE proc_name
SHOW CREATE TABLE tbl_name
SHOW CREATE TRIGGER trigger_name
SHOW CREATE VIEW view_name
SHOW DATABASES [like_or_where]
SHOW ENGINE engine_name {STATUS | MUTEX}
SHOW [STORAGE] ENGINES
SHOW ERRORS [LIMIT [offset,] row_count]
SHOW EVENTS
SHOW FUNCTION CODE func_name
SHOW FUNCTION STATUS [like_or_where]
SHOW GRANTS FOR user
SHOW INDEX FROM tbl_name [FROM db_name]
SHOW MASTER STATUS
SHOW OPEN TABLES [FROM db_name] [like_or_where]
SHOW PLUGINS
SHOW PROCEDURE CODE proc_name
SHOW PROCEDURE STATUS [like_or_where]
SHOW PRIVILEGES
SHOW [FULL] PROCESSLIST
SHOW PROFILE [types] [FOR QUERY n] [OFFSET n] [LIMIT n]
SHOW PROFILES
SHOW RELAYLOG EVENTS [IN 'log_name'] [FROM pos] [LIMIT [offset,] row_count]
SHOW SLAVE HOSTS
SHOW SLAVE STATUS [FOR CHANNEL channel]
SHOW [GLOBAL | SESSION] STATUS [like_or_where]
SHOW TABLE STATUS [FROM db_name] [like_or_where]
SHOW [FULL] TABLES [FROM db_name] [like_or_where]
SHOW TRIGGERS [FROM db_name] [like_or_where]
SHOW [GLOBAL | SESSION] VARIABLES [like_or_where]
SHOW WARNINGS [LIMIT [offset,] row_count]

like_or_where:
 LIKE 'pattern'
 | WHERE expr
```



#### insert

*新增语句*

```mysql
-- 可用于新增前做数据检测
INSERT INTO tbl_temp2 (fld_id)
 SELECT tbl_temp1.fld_order_id
 FROM tbl_temp1 WHERE tbl_temp1.fld_order_id > 100;
 
 

-- select insert
insert into __admin_rule (name, parentid, title, icon, listor, `type`, is_show, status) 
	select
		od.name, 
		ifnull((select id from __admin_rule where name = (
		select od.name from __admin_rule_online_dev where id = od.parentid and od.parentid > 0 ) and is_del = 0), 0) as parentid,
		od.title, od.icon, od.listor, od.`type`, od.is_show, od.status 
	from __admin_rule_online_dev od
	left join __admin_rule ar on od.name = ar.name 
	where ar.id is null 
	;
```



#### delete 

*delete join 语句*

```mysql
-- delete 删除 join 语句
delete dd from fg_article_data dd
	inner join fg_article ac on dd.id=ac.id
	where ac.catid=470
;


delete from fg_article where catid=470;

-- 删除数据，全部数据
-- truncate删除后将重新水平线和索引(id从零开始) ,delete不会删除索引 
truncate (table) tb_name;
```



#### select

```mysql
-- find_in_set
select find_in_set(1, '2,5,1,9,2');   -- 3
```



自定义字段信息排序

```mysql
-- 自定义字段排序信息排序(显示排序)
SELECT * from member bm where id in (472, 98, 519) 
	order by field(id, 472, 98, 519)
;

# 条件排序
select * from guest order by batch_code != 0 desc, batch_code;
```





##### 随机字段

*使用 `rand()` 随机函数查询随机行*





*查询 `table` 表中的随机一行。*

```mysql
-- 效率较慢
select * from `table` order by rand() limit 1;

-- 改进
select * from `table` where `id` >= (
	select floor(rand() * (SELECT MAX(id) FROM ``table``))
);
```





#### update

*更新数据*



**更新连表查询**

```mysql
update __guest_invite_notes gin join (
	select * from
	(select gi.exhibition_id, gi.guest_id,
		(select count(1) from __guest_invite where guest_id=gi.guest_id) as counter
		from __guest_invite gi
		where gi.exhibition_id != 2
	) z
	where z.counter = 1
) tt on gin.guest_id = tt.guest_id
	set gin.exhibition_id=tt.exhibition_id
	where gin.guest_id=tt.guest_id
;
```



#### create

快速建表

```mysql
-- 创建 test_tb 的附表，并将数据转入表 copy_test_tb
create table copy_test_tb as select * from test_tb;

-- 创建 test_tb 结构相等的附表
create table copy_test_tb_struct like test_tb;
```



#### alter

数据表更改

```mysql
-- 修改 tp_auth_rule 当前自动增加主键的值
alter table tp_auth_rule auto_increment=330;

-- 新增唯一键索引
alter table v_module add unique key AK_module_unique (module);
```



#### 条件/where

默认情况下：`like`, `in`, `=` 都不限定大小写。

utf8 编码下：

- utf8_bin    将字符串中的每一个字符以十六进制方式存储数据，区分大小写。

- utf8_general_ci  不区分大小写，ci为case insensitive的缩写，即大小写不敏感。



可由 “binary” 通过SQL层面限制，或者改变表或字段编码。

```sql
-- binary
select * from hx_people hp where binary contact_source_name = 'chu';
```



##### like

like 查询默认不区分大小

```mysql
-- 查询拼音不区分大小写
select * from member where pinyin like 'z%';

-- 在 like 后加入 <binary> 则区分大小写
select * from member where pinyin like binary 'z%';

-- 亦可建表是指定字段区分大小写
create table member(pinyin varchar(20) binary);

-- 查询等于时也无法区分大小写
-- 不区分大小写
select * from bigexp_config bc where `key` = 'main';
-- 区分大小写
select * from bigexp_config bc where binary `key` = 'main';
```



#### @/@@ 的区别

- `@x `            *用户自定义变量（用户变量）。如两者相同: `set @value = last_insert_id(); 或 @value := last_insert_id(); ` 和 `set @value = @@identity; 或 @value := @@identity  `*

- `@@x`            *global/session变量（系统变量）*



> 可用于 SQL 运行时自增id关联键的写入

```mysql
delete from `__menu` where `name` like 'jc_%';

-- 独立句子(分句)定义 
INSERT INTO `__menu` (`id`, `parentid`, `app`, `model`, `action`, `data`, `type`, `status`, `name`, `remark`, `listorder`) VALUES 
	(null, 0, 'Admin', 'Menu', 'add', '', 1, 1, 'jc_A0', '', 7) ;
set @id=last_insert_id();
INSERT INTO `__menu` (`id`, `parentid`, `app`, `model`, `action`, `data`, `type`, `status`, `name`, `remark`, `listorder`) VALUES (null, @id, 'Visit', 'Visit', 'app_column', '', 1, 1, 'jc_B1', '', 0);
set @id2 := @@identity;
INSERT INTO `__menu` (`id`, `parentid`, `app`, `model`, `action`, `data`, `type`, `status`, `name`, `remark`, `listorder`) VALUES (null, @id2, 'Visit', 'Visit', 'app_column', 'type=shujufuwu', 1, 1, 'jc_CC1', '', 0);
INSERT INTO `__menu` (`id`, `parentid`, `app`, `model`, `action`, `data`, `type`, `status`, `name`, `remark`, `listorder`) VALUES (null, @id2, 'Visit', 'Visit', 'app_column', 'type=banshizhinan', 1, 1, 'jc_CC0', '', 0);



-- 非独立句子定义 
INSERT INTO `__menu` (`id`, `parentid`, `app`, `model`, `action`, `data`, `type`, `status`, `name`, `remark`, `listorder`) VALUES (null, 0, 'Admin', 'Menu', 'add', '', 1, 1, 'jc_A0', '', 7) ;
INSERT INTO `__menu` (`id`, `parentid`, `app`, `model`, `action`, `data`, `type`, `status`, `name`, `remark`, `listorder`) VALUES (null, @id := @@identity, 'Visit', 'Visit', 'app_column', '', 1, 1, 'jc_B1', '', 0);
INSERT INTO `__menu` (`id`, `parentid`, `app`, `model`, `action`, `data`, `type`, `status`, `name`, `remark`, `listorder`) VALUES (null, @id2 :=@@identity, 'Visit', 'Visit', 'app_column', 'type=shujufuwu', 1, 1, 'jc_CC1', '', 0);
INSERT INTO `__menu` (`id`, `parentid`, `app`, `model`, `action`, `data`, `type`, `status`, `name`, `remark`, `listorder`) VALUES (null, @id2, 'Visit', 'Visit', 'app_column', 'type=banshizhinan', 1, 1, 'jc_CC0', '', 0);
```



*查询数据赋值*

```mysql
select 1, 109 into @id, @id2;
select @id, @id2;
```



#### 函数

**FROM_UNIXTIME**    *时间戳转时间格式*

```mysql
SELECT 
	FROM_UNIXTIME(UNIX_TIMESTAMP(NOW())) as 时间戳转时间,
	UNIX_TIMESTAMP(NOW()) as 时间转为时间戳,
	NOW() as 当前时间
;


-- 当前数据库版本
select VERSION();
```





#### information_schema

*MySQL 数据表相信*

```mysql
select * from information_schema.TABLES t where TABLE_SCHEMA = DATABASE();
```



#### 其他工具

##### explain

*mysql 执行SQL详情，语法:*

```mysql
{EXPLAIN | DESCRIBE | DESC}
 	tbl_name [col_name | wild]
 
 explain_type: {
     EXTENDED
     | PARTITIONS
     | FORMAT = format_name
}

format_name: {
 TRADITIONAL
 | JSON
}

explainable_stmt: {
 SELECT statement
 | DELETE statement
 | INSERT statement
 | REPLACE statement
 | UPDATE statement
}
```



*示例*

```mysql
explain select * from member;

describe member;
show columns from member;

explain select * from member;
explain member;
desc member;
```







### 自定义函数(Function)

*自定义函数；基本语法*

```mysql
CREATE FUNCTION metaphon
 RETURNS STRING
 SONAME 'udf_example.so';

DROP FUNCTION metaphon;
```



*如： `dev_eid` 函数，返回整形数据*

```mysql
drop function if exists dev_eid;
delimiter //
 -- 数博会当前的项目id
 create function dev_eid()
 	returns int
   	begin
	   	declare eid int;
    	select id into eid from __exhibition e where e.is_default = 1 limit 1;
    	return eid;
 end
//
```







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
-- 其设置与 my.cnf/my.ini 配置文件下应用应置于 "[mysqld]" 下

```



### 约束

#### 外键

外键重建

```mysql
alter table _video drop foreign key FK_video__category_id;
alter table _video add constraint FK_video__category_id foreign key (category_id)
      references _video_category (id) on delete restrict on update restrict;
```



### 数据库引擎

*数据库存储引擎是数据库底层软件组织，数据库管理系统（DBMS）使用数据引擎进行创建、查询、更新和删除数据。不同的存储引擎提供不同的存储机制、索引技巧、锁定水平等功能，使用不同的存储引擎，还可以 获得特定的功能。*



*存储引擎就相当于是数据存储的发动机，来驱动数据在磁盘层面进行存储。*

```mysql
-- 查看MySQL数据库引擎
show engines;

# 修改表的引擎
alter table v_table_name engine = myisam;

# 查看 innodb 引擎状态
show engine InnoDB status;
```



#### InnoDB

*MySQL5.7/8.0 等版本的默认引擎，支持事务(Transcations)，行锁定和外键。*

- *提供了具有提交、回滚和崩溃恢复能力的事物安全（ACID兼容）存储引擎：事务(Transcations)*
- *InnoDB是为处理巨大数据量的最大性能设计*
- *InnoDB存储引擎完全与MySQL服务器整合，InnoDB存储引擎为在主内存中缓存数据和索引而维持它自己的缓冲池。InnoDB将它的表和索引在一个逻辑表空间中，表空间可以包含数个文件（或原始磁盘文件）*
- *InnoDB支持外键完整性约束，存储表中的数据时，每张表的存储都按主键顺序存放，如果没有显示在表定义时指定主键，InnoDB会为每一行生成一个6字节的ROWID，并以此作为主键*
- *InnoDB被用在众多需要高性能的大型数据库站点上*



*数据库引擎是用于存储、处理和保护数据的核心服务。利用数据库引擎可控制访问权限并快速处理事务，从而满足企业内大多数需要处理大量数据的应用程序的要求。   使用数据库引擎创建用于联机事务处理或联机分析处理数据的关系数据库。这包括创建用于存储数据的表和用于查看、管理和保护数据安全的数据库对象（如索引、视图和存储过程）。*



自从 MySQL 5.1 之后，默认的存储引擎变成了 InnoDB 存储引擎，相对于 MyISAM，InnoDB 存储引擎有了较大的改变，它的主要特点是

- 支持事务操作，具有事务 ACID 隔离特性，默认的隔离级别是`可重复读(repetable-read)`、通过`MVCC（并发版本控制）`来实现的。能够解决`脏读`和`不可重复读`的问题。
- InnoDB 支持外键操作。
- InnoDB 默认的锁粒度`行级锁`，并发性能比较好，会发生死锁的情况。
- 和 MyISAM 一样的是，InnoDB 存储引擎也有 `.frm文件存储表结构` 定义，但是不同的是，InnoDB 的表数据与索引数据是存储在一起的，都位于 B+ 数的叶子节点上，而 MyISAM 的表数据和索引数据是分开的。
- InnoDB 有安全的日志文件，这个日志文件用于恢复因数据库崩溃或其他情况导致的数据丢失问题，保证数据的一致性。
- InnoDB 和 MyISAM 支持的索引类型相同，但具体实现因为文件结构的不同有很大差异。
- 增删改查性能方面，果执行大量的增删改操作，推荐使用 InnoDB 存储引擎，它在删除操作时是对行删除，不会重建表。





#### MyISAM

*基于ISAM存储引擎，并对其进行扩展；它是在Web、数据仓储和其他应用环境下最常使用的存储引擎之一。MyISAM拥有较高的插入、查询速度，但不支持事物。*

崩溃后无法安全恢复。

其将表保存为".MYD" 数据文件 和 ".MYI" 索引文件。



在 5.1 版本之前，MyISAM 是 MySQL 的默认存储引擎，MyISAM 并发性比较差，使用的场景比较少，主要特点是

- 不支持`事务`操作，ACID 的特性也就不存在了，这一设计是为了性能和效率考虑的。

- 不支持`外键`操作，如果强行增加外键，MySQL 不会报错，只不过外键不起作用。

- MyISAM 默认的锁粒度是`表级锁`，所以并发性能比较差，加锁比较快，锁冲突比较少，不太容易发生死锁的情况。

- MyISAM 会在磁盘上存储三个文件，文件名和表名相同，扩展名分别是 `.frm(存储表定义)`、`.MYD(MYData,存储数据)`、`MYI(MyIndex,存储索引)`。这里需要特别注意的是 MyISAM 只缓存`索引文件`，并不缓存数据文件。

- MyISAM 支持的索引类型有 `全局索引(Full-Text)`、`B-Tree 索引`、`R-Tree 索引`

  Full-Text 索引：它的出现是为了解决针对文本的模糊查询效率较低的问题。

  B-Tree 索引：所有的索引节点都按照平衡树的数据结构来存储，所有的索引数据节点都在叶节点

  R-Tree索引：它的存储方式和 B-Tree 索引有一些区别，主要设计用于存储空间和多维数据的字段做索引,目前的 MySQL 版本仅支持 geometry 类型的字段作索引，相对于 BTREE，RTREE 的优势在于范围查找。

- 数据库所在主机如果宕机，MyISAM 的数据文件容易损坏，而且难以恢复。

- 增删改查性能方面：SELECT 性能较高，适用于查询较多的情况



#### MEMORY

*内存引擎，为查询和引用其他表数据提供快速访问。*

MEMORY 存储引擎使用存在内存中的内容来创建表。每个 MEMORY 表实际只对应一个磁盘文件，格式是 `.frm`。MEMORY 类型的表访问速度很快，因为其数据是存放在内存中。默认使用 `HASH 索引`。



#### merge

MERGE 存储引擎是一组 MyISAM 表的组合，MERGE 表本身没有数据，对 MERGE 类型的表进行查询、更新、删除的操作，实际上是对内部的 MyISAM 表进行的。MERGE 表在磁盘上保留两个文件，一个是 `.frm` 文件存储表定义、一个是 `.MRG` 文件存储 MERGE 表的组成等。





#### 选择合适的存储引擎

在实际开发过程中，我们往往会根据应用特点选择合适的存储引擎。

- MyISAM：如果应用程序通常以检索为主，只有少量的插入、更新和删除操作，并且对事物的完整性、并发程度不是很高的话，通常建议选择 MyISAM 存储引擎。
- InnoDB：如果使用到外键、需要并发程度较高，数据一致性要求较高，那么通常选择 InnoDB 引擎，一般互联网大厂对并发和数据完整性要求较高，所以一般都使用 InnoDB 存储引擎。
- MEMORY：MEMORY 存储引擎将所有数据保存在内存中，在需要快速定位下能够提供及其迅速的访问。MEMORY 通常用于更新不太频繁的小表，用于快速访问取得结果。
- MERGE：MERGE 的内部是使用 MyISAM 表，MERGE 表的优点在于可以突破对单个 MyISAM 表大小的限制，并且通过将不同的表分布在多个磁盘上， 可以有效地改善 MERGE 表的访问效率。



### 索引

所有的 MySQL 类型都可以进行索引，对相关列使用索引是提高 `SELECT` 查询性能的最佳途径。MyISAM 和 InnoDB 都是使用 `BTREE` 作为索引，MySQL 5 不支持`函数索引`，但是支持 `前缀索引`。

前缀索引顾名思义就是对列字段的前缀做索引，前缀索引的长度和存储引擎有关系。MyISAM 前缀索引的长度支持到 1000 字节，InnoDB 前缀索引的长度支持到 767 字节，索引值重复性越低，查询效率也就越高。

MySQL只能高效地使用索引的最左前缀列。



```mysql
-- 查询表，v_guest_invite——的索引情况
show index from v_guest_invite;
```



在 MySQL 中，主要有下面这几种索引

- `全局索引(FULLTEXT)`：全局索引，目前只有 MyISAM 引擎支持全局索引，它的出现是为了解决针对文本的模糊查询效率较低的问题，并且只限于 CHAR、VARCHAR 和 TEXT 列。
- `哈希索引(HASH)`：哈希索引是 MySQL 中用到的唯一 key-value 键值对的数据结构，很适合作为索引。HASH 索引具有一次定位的好处，不需要像树那样逐个节点查找，但是这种查找适合应用于查找单个键的情况，对于范围查找，HASH 索引的性能就会很低。默认情况下，MEMORY 存储引擎使用 HASH 索引，但也支持 BTREE 索引。
- `B-Tree 索引`：B 就是 Balance 的意思，BTree 是一种平衡树，它有很多变种，最常见的就是 B+ Tree，它被 MySQL 广泛使用。
- `R-Tree 索引`：R-Tree 在 MySQL 很少使用，仅支持 geometry 数据类型，支持该类型的存储引擎只有MyISAM、BDb、InnoDb、NDb、Archive几种，相对于 B-Tree 来说，R-Tree 的优势在于范围查找。



我们使用 `explain` 进行sql查询分析，检测数据库表的索引使用情况。



#### 索引设计原则

创建索引的时候，要尽量考虑以下原则，便于提升索引的使用效率。

- 选择`索引位置`，选择索引最合适的位置是出现在 `where` 语句中的列，而不是 `select` 关键字后的选择列表中的列。
- 选择使用`唯一索引`，顾名思义，唯一索引的值是唯一的，可以更快速的确定某条记录，例如学生的学号就适合使用唯一性索引，而学生的性别则不适合使用，因为不管搜索哪个值，都差不多有一半的行。
- 为经常使用的字段建立索引，如果某个字段经常用作查询条件，那么这个字段的查询速度在极大程度上影响整个表的查询速度，因此为这样的字段建立索引，可以提高整个表的查询速度。
- 不要过度索引，限制索引数目，索引的数目不是越多越好，每个索引都会占据磁盘空间，索引越多，需要的磁盘空间就越大。
- 尽量使用`前缀索引`，如果索引的值很长，那么查询速度会受到影响，这个时候应该使用前缀索引，对列的某几个字符进行索引，可以提高检索效率。
- 利用最左前缀，在创建一个 n 列的索引时，实际上是创建了 MySQL 可利用的 n 个索引。多列索引可以起到几个索引的作用，利用索引最左边的列来匹配行，这样的列称为最左前缀。
- 对于使用 InnoDB 存储引擎的表来说，记录会按照一定的顺序保存。如果有明确的主键定义，那么会按照主键的顺序进行保存；如果没有主键，但是有唯一索引，那么就按照唯一索引的顺序进行保存。如果既没有主键又没有唯一索引，那么表中会自动生成一个内部列，按照这个列的顺序进行保存。一般来说，使用主键的顺序是最快的
- 删除不再使用或者很少使用的索引





### 视图

视图的英文名称是 `view`，它是一种虚拟存在的表。视图对于用户来说是透明的，它并不在数据库中实际存在，视图是使用数据库行和列动态组成的表。



视图相对于普通的表来说，优势包含下面这几项

- 使用视图可以简化操作：使用视图我们不用关注表结构的定义，我们可以把经常使用的数据集合定义成视图，这样能够简化操作。
- 安全性：用户对视图不可以随意的更改和删除，可以保证数据的安全性。
- 数据独立性：一旦视图的结构 确定了， 可以屏蔽表结构变化对用户的影响， 数据库表增加列对视图没有影响；具有一定的独立性



**视图的作用**：简化用户（或代码）的查询操作、提供数据安全（比如屏蔽敏感字段）等。SQL查询性能主要与SQL使用索引以及其他查询优化有关



### 游标/cursor

> 基本语法

```mysql
begin
    -- cursor 定义语法
    declare cursor_name cursor from select_statement;
    
    -- continue 条件标记
    -- condition_name: mysql_error_code, SQLSTATE [VALUE] sqlstate_value;
    declare condition_name continue for condition_value;
	
	-- 条件处理语句
	-- handler_action: CONTINUE, EXIT, UNDO
	-- conditon_value: mysql_error_code, SQLSATE [VALUE] sqlstate_value, condition_name, SQLWANRING, NOT FOUND, SQLEXCEPTION
	declare handler_action handler for conditon_value [, conditon_value] statement;
	
	
    open cursor_name;
        -- fetch 获取游标中的参数
        fetch cursor_name into var_name[, var_name]...;


    -- 关闭游标
    close cursor_name;
end;
```

> 样例

```mysql
-- drop table if exists jc_user;
-- 删除已经存在的过程
drop procedure if exists base_cursor_p;

-- 创建测试表
create table if not exists jc_user(
	id int primary key auto_increment,
    name varchar(50),
    cur_mark_time datetime,
    cur_idx int
);
-- 插入测试数据
insert into jc_user(name) values ('JC'), ('Emma'), ('Sr'), ('Core'), ('Ben');

-- 定于过程块
delimiter //
create procedure base_cursor_p()
begin
	declare not_more_record int default false;
	
	declare v_id int;
	declare v_name varchar(50);
	declare v_idx int default -1;
	declare cur cursor for select id, name from jc_user;
	declare continue handler for not found set not_more_record= true;
	
	update jc_user set cur_mark_time=null, cur_idx=0;
	
	open cur;
	-- 循环语句： loop, repeat, while
	while not_more_record<>true do
		fetch cur into v_id, v_name;
        
        -- 数据处理/业务逻辑
        update jc_user set cur_mark_time=current_time(), cur_idx=v_idx where id=v_id;
        set v_idx = v_idx + 1;
	end while;
	-- 关闭游标
	close cur;
end //


-- 恢复通用的 ";" 分隔符
delimiter ;
-- 执行过程
call base_cursor_p();
drop procedure if exists base_cursor_p;
```



### 模拟匿名块执行(delimiter)

> 执行SQL脚本

```mysql
-- 创建临时存储过程并执行(如： _jc_tmp_untitle_sqlblock)
-- 删除过程过程，词语句可以去除点
drop procedure if exists _jc_tmp_untitle_sqlblock;

-- 创建
-- 修改分割符号
delimiter //
create procedure _jc_tmp_untitle_sqlblock()
begin
	select 'Example MySQL UNtitle SQL block by Joshua Conero(古丞秋).' as msg;
end;
//

-- 恢复分割符号
delimiter ;

-- 执行匿名块
call _jc_tmp_untitle_sqlblock();
-- 删除匿名块
drop procedure if exists _jc_tmp_untitle_sqlblock;
```



### 存储过程

**「存储过程是在数据库系统中完成一组特定功能的 SQL 语句集」**，它存储在数据库系统中，一次编译后永久有效。

> 优点

- 使用存储过程具有可封装性，能够隐藏复杂的 SQL 逻辑。
- 存储过程可以接收参数，并返回结果
- 存储过程性能非常高，一般用于批量执行语句

> 缺点

- 存储过程编写复杂
- 存储过程对数据库的依赖性比较强，可移植性比较差



#### 变量

在 MySQL 中，变量可分为两大类，即`系统变量`和`用户变量`，这是一种粗略的分法。但是根据实际应用又被细化为四种类型，即局部变量、用户变量、会话变量和全局变量。



MySQL 中的局部变量与 Java 很类似 ，Java 中的局部变量是 Java 所在的方法或者代码块，而 MySQL 中的局部变量作用域是所在的存储过程。MySQL 局部变量使用 `declare` 来声明。



当服务启动时，它将所有全局变量初始化为默认值。其作用域为 server 的整个生命周期。

```mysql
show global variables;

-- 修改全局变量
set global sql_warnings=ON;  

-- 查看
select @@global.version;
show global variables like '%version%';
```





##### 会话变量

用户变量是基于`会话变量`实现的，可以暂存，用户变量与连接有关，也就是说一个客户端定义的变量不能被其他客户端使用看到。当客户端退出时，链接会自动释放。我们可以使用 `set` 语句设置一个变量

```mysql
set @name = "Joshua Conero";
select @name, VERSION();

-- 查看当前的会话变量
show session variables;
```





*实例，mysql 5.7 运行正确*

```mysql
-- 二级目录移动
delimiter //
create procedure _jc_tmp_untitle_move_by_catid(v_catid int)
begin
    declare not_more_record int default false;
    declare v_bm_catid int;

    -- cursor 定义语法
    declare bumen_cs cursor for select `catid` from fg_category where parentid=v_catid;
	declare continue handler for not found set not_more_record= true;

    open bumen_cs;
        -- 循环语句： loop, repeat, while
        while not_more_record<>true do
            fetch bumen_cs into v_bm_catid;
            call jc_tmp_untitle_move_by_catid(v_bm_catid);
        end while;
    -- 关闭游标
    close bumen_cs;

    -- 导文章
    insert into cp_article_191015 select * from fg_article where catid=v_catid;
    insert into cp_article_data_191015 select dd.* from fg_article_data dd
        inner join fg_article fat on dd.id=fat.id and fat.catid=v_catid
    ;
    -- 数据参数
    delete dd from fg_article_data dd 
        inner join fg_article fat on dd.id=fat.id and fat.catid=v_catid
    ;
    delete from fg_article where catid=v_catid;
    
    -- 菜单导入
    insert into cp_category_191015 select * from fg_category where parentid=v_catid;
    delete from fg_category where parentid=v_catid;
end;
delimiter ;
```



### 用户

*用户表对应 `mysql.user`，可对此表对用户信息进行维护。*

```mysql
-- 创建用户[conero](password)
--  if exists 错误
CREATE USER 'conero'@'localhost' IDENTIFIED BY 'password';
create user 'conero'@'%' identified by '151009_170512';

-- 修改用户密码（8.0 可能报错）
update mysql.user set password=password('新密码') where User='phplam' and Host='localhost';

-- 授权
grant all privileges on zhangsanDb.* to zhangsan@'%' identified by 'zhangsan';
-- 授权给数据库无须更变密码
grant all privileges on zhangsanDb.* to zhangsan@'%';
GRANT ALL PRIVILEGES ON *.* TO 'root'@'192.168.199.99';

flush privileges;

-- 查看用户权限
SELECT DISTINCT CONCAT('User: ''',user,'''@''',host,''';') AS query FROM mysql.user;
-- 查看指定用户
show grants for 'usermy'@'%'; 

-- 删除权限
delete from mysql.user where user="usermy" and host="%";

-- 查看用户密码
select user,host from mysql.user;
```



### 数据库锁

InnoDB支持表、行(默认)级锁，而MyISAM支持表级锁

```mysql
-- 查看数据库锁
show OPEN TABLES where In_use > 0;
```







### 实例

_**information_schema**_   提供数据库元数据，包括数据库、数据表、列、插件、函数、存储过程、触发器以及权限等；其包含一些只读数据表，`show` 的提供数据



_**performance_schema**_   提供数据库服务运行时监视器，包含实时的信息。



### 优化方案

> where执行顺序

where执行顺序是从左往右执行的，在数据量小的时候不用考虑，但数据量多的时候要考虑条件的先后顺序，此时应遵守一个原则：排除越多的条件放在第一个。





## issue

### time_zone 错误，jdbc 连接解决

```
# java.lang.RuntimeException: com.mysql.cj.exceptions.InvalidConnectionAttributeException: The server time zone value 'ÖÐ¹ú±ê×¼Ê±¼ä' is unrecognized or represents more than one time zone. You must configure either the server or JDBC driver (via the serverTimezone configuration property) to use a more specifc time zone value if you want to utilize time zone support.
```



> *设置时间区就可以*

```sql
set global time_zone ='+8:00';
```



### 1607 时间类型错误，更新字段。

> **出现版本 5.7.27**

由于 5.7 中不同的修复版本(小版本)中时间类型的兼容问题，导时间类型表修改时报错。但是可以通过 `sql_mode` 来兼容该问题。

```mysql
show variables like 'sql_mode';
# ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION

# 删除 【NO_ZERO_IN_DATE】，【NO_ZERO_DATE】即可修改以上问题
# ONLY_FULL_GROUP_BY group by 的兼容问题
set global sql_mdoe='STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';
```



### 201030 mysql root密码重置

*方式是：在 `my.conf` 等之类的加入 `skip-grant-tables`，使之跳过密码已修改root密码。最后再去除：skip-grant-tables配置。*



但是mysql 中，单纯加入以上配置依然报错。

```ini
[mysqld]
skip-grant-tables

# mysql8中依然报错，此时需要在加上配置项
 shared-memory
```



修改root密码：

```mysql
-- 用过免密登录后修改密码：
select host, user, authentication_string, plugin from user;
update user set authentication_string='' where user='root';

-- 再修改
ALTER user 'root'@'localhost' IDENTIFIED BY 'pswd-any';
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
# 导入缓慢可设置值系统不同的参数实现，max_allowed_packet 的值为数字
#		set global interactive_timeout = 120;
#		set global wait_timeout = 120;
#		set global max_allowed_packet = 32M;
source d:/.../name.sql


# 数据库复制 oldDbname => <toNewDb> 数据库已经存在的的数据库
mysqldump <oldDbname> -u root -p<password> --add-drop-table | mysql <toNewDb> -u root -p<password>
```





### mysqldump

> 本地导出数据为SQL

```powershell
# mysqldump -uroot -p --default-character-set=utf8 dbname > d:/.../file.sql
mysqldump -uroot -p --default-character-set=utf8 dataset > d:/tmp/dataset.sql

# 导出的字符格式，utf8. ">" 与系统的语言编码一致 
mysqldump -uroot -p --default-character-set=utf8 dataset --result-file=dump.sql

# 经测试只导出一行
# needtoKnowMore
mysqldump -uroot -p --default-character-set=utf8 dataset tablename > d:/tmp/tablename.sql
mysqldump -uroot -p --default-character-set=utf8 dataset tablename --result-file=d:/tmp/tablename.sql

# 选项
# --no-data 无数据

# 远程测试，可能导致其他链接缓慢。
# 导出测试的服务器所在数据库 "207.12.24.56"
mysqldump -h "207.12.24.56" -uroot -p --default-character-set=utf8 dataset tablename --result-file=d:/tmp/tablename.sql
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

-- 重置清空日志
reset master;

# 使用 mysqlbinlog 读取日志
mysqlbinlog binlog.000111
```



//@TODO *window 下如开启 binlog ，实现数据备份的解决方案。*





## 组件

*mysql 基本组件描述：*

- **mysql**     客服端软件，对任何mysql服务端进行访问（SQL 命令行工具）
- **mysqld**   SQL后台程序，服务端程序。其英文名为 *mysql deamon*。(主程序)





## example

### 游标

*游标测试*

```mysql
-- 创建测试数据表
drop table if exists tx;
create table if not exists tx (
 	idx int,
 	pid int,
 	vtname varchar(30)
);
-- 删除旧的数据
delete from tx where 1;
-- 删除存在的测试过程
drop procedure if exists tx_tmp1;

-- 临时过程定义
delimiter //
create procedure tx_tmp1()
begin
	declare v_flag BOOLEAN DEFAULT 0;
	declare v_idx int default 0;
	declare v_pid int default 0;
	declare v_cursor CURSOR FOR select id from fg_visit;
	-- 未发现查询到数据时，修改 v_flag 的值；类似于异常处理
	declare continue handler for not found SET v_flag=1;
	
	open v_cursor;
		repeat
			FETCH v_cursor INTO v_pid;
		-- 业务逻辑
		IF v_flag!=1 THEN
			insert into tx(idx, pid, vtname) values(v_pid, v_idx, 'fg_visit');
			set v_idx = v_idx + 1;
		END IF;
	UNTIL v_flag END REPEAT;
	close v_cursor;
end;
//

delimiter ;

-- 临时过程调用
call tx_tmp1();
-- 临时过程删除
drop procedure if exists tx_tmp1;
```





## 附录

### 参考

- [47 张图带你 MySQL 进阶](https://mp.weixin.qq.com/s/mJ_rCfIxsltZYKDtXZx82g)