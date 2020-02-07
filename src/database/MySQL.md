

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




### 5.7 

> Windows 版本安装；

*下载官网 msi，根据向导安装需要的应用。安装后需要执行命令：*

1.  初始化 `$ /bin/mysqld --initialize-insecure`
2. 安装 `$ /bin/mysqld install`
3. 启动 `$ net start mysql`



```mysql
-- 安装项目完成是，密码为空；这是需要自行设置密码
alter user 'root'@'localhost' identified by 'password';
```





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





### 语句

#### 类型

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
	curtime() as 当前时间戳, now(), current_timestamp as `常量同义-时间戳`
;
```



##### 字符串

**文本**

- *TinyText*           最大长度255个字节(2^8-1)
- *Text*                  最大长度65535个字节(2^16-1)



#### show

可用于查询对象： *databases, tables, variables, columns, server status* 等

```mysql
-- 通用查询方法
show databases [like ''];		-- 显示所有的数据库，可使用 like 查询；

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

-- 修改用户密码
update mysql.user set password=password('新密码') where User='phplam' and Host='localhost';

-- 授权
grant all privileges on zhangsanDb.* to zhangsan@'%' identified by 'zhangsan';
flush privileges;
```





### 实例

_**information_schema**_   提供数据库元数据，包括数据库、数据表、列、插件、函数、存储过程、触发器以及权限等；其包含一些只读数据表，`show` 的提供数据



_**performance_schema**_   提供数据库服务运行时监视器，包含实时的信息。





## issue

### time_zone 错误，jdbc 连接解决

```
# java.lang.RuntimeException: com.mysql.cj.exceptions.InvalidConnectionAttributeException: The server time zone value 'ÖÐ¹ú±ê×¼Ê±¼ä' is unrecognized or represents more than one time zone. You must configure either the server or JDBC driver (via the serverTimezone configuration property) to use a more specifc time zone value if you want to utilize time zone support.
```



> *设置时间区就可以*

```sql
set global time_zone ='+8:00';
```



### #1607 时间类型错误，更新字段。

> **出现版本 5.7.27**

由于 5.7 中不同的修复版本(小版本)中时间类型的兼容问题，导时间类型表修改时报错。但是可以通过 `sql_mode` 来兼容该问题。

```mysql
show variables like 'sql_mode';
# ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION

# 删除 【NO_ZERO_IN_DATE】，【NO_ZERO_DATE】即可修改以上问题
# ONLY_FULL_GROUP_BY group by 的兼容问题
set global sql_mdoe='STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION';
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

