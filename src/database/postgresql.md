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



#### ubuntu 24.04 版本安装

设置系统的deb 源后，使用 apt 安装，[参考](https://blog.junxworks.cn/articles/2025/03/31/1743412091261.html)

```shell
sudo apt install -y postgresql-15

# 检查服务状态
sudo systemctl status postgresql@15-main

# 修改密码
sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD '你的密码';"

# 相关配置文件 /etc/postgresql/15/main/
# 重启服务
sudo systemctl restart postgresql@15-main

# 设置开机启动
sudo systemctl enable postgresql@15-main
```



### SQL



#### 数据类型



数字整形：支持类型 int、int2/smallint、int4/integer/int、int8/bigint

- int
- int2    2字节整形，16位 (`2^16`)
- int4    4字节整形，32位
- int8    8字节整形，64位





##### 字符串 



字符串换行

```sql
-- standard_conforming_strings = on（PostgreSQL 9.1+ 默认为 on）
select E'line1\nline2\n  line3' as test;
```





#### create

```sql
-- 创建数据表
CREATE DATABASE ksj_gysjj_xyjgxt with OWNER = postgres ENCODING = 'UTF8';
COMMENT ON DATABASE ksj_gysjj_xyjgxt IS '数据库名称';

-- scheme 创建
CREATE SCHEMA IF NOT EXISTS scjg_xyjg AUTHORIZATION postgres;
```



表格备份生成

```sql
-- 表格快速生成
-- 表结构复制/PG 基础写法（类似 MySQL CREATE TABLE LIKE）
create table sys_user_bk (like sys_user);

-- 完整复刻（含索引 / 约束 / 序列等）
create table sys_user_bk (like sys_user INCLUDING ALL);

-- 临时表创建
CREATE TEMP TABLE temp_guest_code (LIKE guest_invite_code INCLUDING ALL);

-- 快速插入数据/查询插入数据
insert into temp_guest_code select * from guest_invite_code;
```





#### alter

用户密码修改，默认安装后远程工具无法登录，需要修改密码

```sql
ALTER USER postgres WITH PASSWORD 'blockPswd';
```



列字段新增

```sql
	alter table guest_invite add column recept_level int2 DEFAULT 0;
COMMENT ON COLUMN public.guest_invite.recept_level IS '嘉宾接待级别';
```



列修改

```sql
-- 已知 sex 性别为varchar，将其修改为 int

ALTER TABLE public.guest_user ALTER COLUMN sex DROP DEFAULT;
ALTER TABLE public.guest_user ALTER COLUMN sex TYPE int2 USING sex::int2;
-- 合成一句
ALTER TABLE public.guest_user 
	ALTER COLUMN pw_type DROP default,
	ALTER COLUMN pw_type TYPE int2 USING pw_type::int2;

-- 设置默认值
ALTER TABLE public.guest_user ALTER COLUMN sex SET DEFAULT 0;
-- 设置必填
ALTER TABLE public.guest_user ALTER COLUMN sex SET NOT NULL;

-- 合并写法
ALTER TABLE public.guest_user
  ALTER COLUMN sex SET DEFAULT 0,
  ALTER COLUMN sex SET NOT NULL;
```



带值转换

```sql
-- 1. 转换类型时确保无 NULL（如前所述）
ALTER TABLE public.guest_user
ALTER COLUMN sex TYPE INT2
USING COALESCE(
  CASE 
    WHEN sex = '男' THEN 1
    WHEN sex = '女' THEN 2
    WHEN sex ~ '^\d+$' THEN sex::INT2
    ELSE 0
  END,
  0
);
```



### update

```sql
-- 类似 join 数据更新
-- b.pw_type::integer   字符串转数字
UPDATE member_user u
SET pw_type = b.pw_type::integer
FROM member_user_bk b  
WHERE 
  u.user_id = b.user_id  
  AND length(b.pw_type) > 0; 
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





### 常用函数

#### 逻辑函数

**coalesce**  类似与mysql的 *ifnull* 函数，可用于非空判断

```sql
-- 如SQL查询
select c.code_id, c.invite_code, c.status, g.guest_name, g.company, g.post, i.invite_id  from guest_invite_code c	
	left join guest_invite i on i.invite_code = c.invite_code and i.pro_code = c.pro_code 
	left join guest_user g on g.guest_id = i.guest_id 
	-- where c.apply_id = 1994227078059667458
	order by coalesce(i.invite_id, 0) desc
;
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

pgsql 复制分为物理复制、逻辑复制。

可使用 pgAdmin 工具，选择数据库使用“Restore“、”Backup”工具恢复或备份数据库。（实际使用图形化界面，错误信息不太好看）



命令行导出

```shell
# 命令行导出sql
pg_dump -U postgres -h 123.57.63.25 -p 39751 sbh > sbh_backup.sql

# 命令格式
pg_dump -U 用户名 -h 主机地址 -p 端口 数据库名 > 备份文件.sql


# 使用 psql 导入存文本 sql 脚本
# 切换用户:  sudo su - postgres
psql -U postgres -d sbh_25 -f /home/edat/conero/repository/conero/e-birdnest/expo25rw/sbh_backup.sql
```





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



### 附录

#### 源设置

Ubuntu 22.04/24.04 使用[阿里源](https://developer.aliyun.com/mirror/postgresql/)安装 PostgreSQL, 参考 https://zhuanlan.zhihu.com/p/698519180

```shell
# bash 命令
sudo apt update && \
sudo apt install -y curl gpg gnupg2 software-properties-common apt-transport-https lsb-release ca-certificates && \
curl -fsSL https://mirrors.aliyun.com/postgresql/repos/apt/ACCC4CF8.asc | sudo gpg --dearmor -o /etc/apt/trusted.gpg.d/postgresql.gpg && \
echo "deb https://mirrors.aliyun.com/postgresql/repos/apt/ `lsb_release -cs`-pgdg main" |sudo tee  /etc/apt/sources.list.d/pgdg.list && \
sudo apt update && sudo apt install -y postgresql-16


# fish 命令
sudo apt update && \
sudo apt install -y curl gpg gnupg2 software-properties-common apt-transport-https lsb-release ca-certificates && \
curl -fsSL https://mirrors.aliyun.com/postgresql/repos/apt/ACCC4CF8.asc | sudo gpg --dearmor -o /etc/apt/trusted.gpg.d/postgresql.gpg && \
echo "deb https://mirrors.aliyun.com/postgresql/repos/apt/ $(lsb_release -cs)-pgdg main" |sudo tee  /etc/apt/sources.list.d/pgdg.list && \
sudo apt update && sudo apt install -y postgresql-16
```

