## postgresql

> 2024年5月24日 星期五
>
> Joshua Conero



先进的开源关系数据库



官网地址：[https://www.postgresql.org/](https://www.postgresql.org/)

中文手册：https://github.com/postgres-cn/pgdoc-cn



### 安装

#### 龙蜥 8.8 安装 pg12

```shell
dnf install -y https://download.postgresql.org/pub/repos/yum/reporpms/EL-8-x86_64/pgdg-redhat-repo-latest.noarch.rpm

# 关闭内置postgresql模块
dnf -qy module disable postgresql

# 安装数据库
dnf install -y postgresql12-server

# 初始化数据
/usr/pgsql-12/bin/postgresql-12-setup initdb

# 创建 postgresql-12 服务
systemctl enable postgresql-12

# 启动服务
systemctl start postgresql-12


# 使用 postgress 用户
su - prostgress
# 使用 psql 登陆服务即可
psql
```



设置 pg12 外网可访问

```shell
# 修改配置文件 /var/lib/pgsql/12/data/postgresql.conf
listen_addresses = '*'

# 设置 pg_hba.conf，允许所有ip可访问
host    replication     all             0.0.0.0/0               ident
host    all             all             0.0.0.0/0               md5
```



### SQL

#### create

```sql
-- 创建数据表
CREATE DATABASE ksj_gysjj_xyjgxt with OWNER = postgres ENCODING = 'UTF8';
COMMENT ON DATABASE ksj_gysjj_xyjgxt IS '数据库名称';

-- scheme 创建
CREATE SCHEMA IF NOT EXISTS scjg_xyjg AUTHORIZATION postgres;
```





#### 其他

数据表对象等基本查询，常用系统表 pg_database, pg_class, pg_namespace等

```sql
-- 查询数据对象
select * from pg_class c;

-- 查询数据库命名空间
select * from pg_namespace n;

-- 表名及字段名查询
select c.oid, c.relname, a.attname, n.nspname, d.description, c.* from pg_class c 
	join pg_namespace n on c.relnamespace=n.oid
	join pg_attribute a on a.attrelid = c.oid
	left join pg_description d on d.objoid = c.oid
	where n.nspname = 'scjg_xyjg'
	;
```



数据集表字段信息查询

```sql
WITH RECURSIVE all_schemas AS (
    SELECT 
        nspname AS schema_name,
        oid AS schema_oid
    FROM 
        pg_namespace
    WHERE 
        nspname NOT LIKE 'pg_%' AND nspname != 'information_schema'
),
all_tables AS (
    SELECT 
        all_schemas.schema_name,
        tbl.relname AS table_name,
        tbl.oid AS table_oid
    FROM 
        all_schemas
    JOIN 
        pg_class tbl ON tbl.relnamespace = all_schemas.schema_oid
    WHERE 
        tbl.relkind IN ('r', 'p') -- 'r' for regular tables, 'p' for partitioned tables
),
all_columns AS (
    SELECT 
        all_tables.schema_name,
        all_tables.table_name,
        all_tables.table_oid,
        att.attname AS column_name,
        att.attnum AS column_num,
        att.atttypid::regtype AS data_type,
        att.attnotnull AS not_null,
        pg_get_expr(def.adbin, def.adrelid) AS default_value,
        pg_description.description AS comment
    FROM 
        all_tables
    JOIN 
        pg_attribute att ON att.attrelid = all_tables.table_oid
    LEFT JOIN 
        pg_attrdef def ON att.attrelid = def.adrelid AND att.attnum = def.adnum
    LEFT JOIN 
        pg_description ON pg_description.objoid = all_tables.table_oid AND pg_description.objsubid = att.attnum
    WHERE 
        att.attnum > 0 AND NOT att.attisdropped
)
SELECT 
    schema_name,
    table_name,
    column_name,
    data_type,
    CASE WHEN not_null THEN 'NOT NULL' ELSE '' END AS is_not_null,
    default_value,
    comment
FROM 
    all_columns
ORDER BY 
    schema_name, 
    table_name, 
    column_num;
```









### psql

sql 查询工具

```shell
# 查看客户端编码格式
select name,setting,context from pg_settings where name like '%encoding%';

# 设置编码
\encoding UTF8

# 导入sql
\i D:/conero/xxx/test_data.sql

# 将日志信息写入到文件中
\o /home/scjg_xyjg-db-pg12/import.log

# 查看数据库类别
\l
# 更详细的查看数据库列表
\l+

# 切换数据库
\c mydb

# 查看用户列表
\du

# 查看数据库服务器版本号
select version();

show server_version;
```





### 备份/恢复

可使用 pgAdmin 工具，选择数据库使用“Restore“、”Backup”工具恢复或备份数据库。（实际使用图形化界面，错误信息不太好看）



或使用 psql 命令行工具导入SQL

```shell
# 导入sql
\i D:/conero/xxx/test_data.sql

# 编码信息查看
\encoding

# 设置当前客户端编码信息
\encoding UTF8

# 显示服务器编码
show server_encoding;
```

