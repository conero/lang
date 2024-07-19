# KingbaseES

> 人大金仓数据库
>
> Joshua Conero
>
> 2024年4月26日
>
> 文档基于 V8R6 等



KingbaseES是人大金仓自主研发的企业级大型通用数据库管理系统，提供Oracle、、pgsql、MySQL和SQLServer兼容模式，在应用不改、性能不降、习惯不变的情况下，实现国外数据库的迁移替代。



工具

- ksql        命令行SQL连接工具
- Kstudio  SQL连接客户端（基于[Dbeaver](https://github.com/dbeaver/dbeaver)开发）
- [KDts](https://help.kingbase.com.cn/v8/development/develop-transfer/kdts-plus/index.html)       oracle，pgsql，mysql等数据迁移工具 



**个人感受**

对开发者很不友好，在没有授权的情况下难以支持开发，各种功能限制。如果不是“信创”鬼才用呢，难用的很。



### 安装

**Windows 是服务启动错误排除**

- 服务无法正常启动是错误排查

```powershell
# 使用 sys_ctrl 进行排查
.\sys_ctl.exe -D 'D:\Program Files\Kingbase\ES\V8\data' start
```



数据库模式切换

```shell
# 使用 initdb 切换数据库模式，再不重装软件的情况下
initdb -U用户 -m pg -D data目录
```



### 概念

#### Schema 模式

属于一个逻辑分割，命名空间。一个数据库可以包含一个或多个命名的模式。

- public                   新建用户默认继承public角色权限，在任何数据库的public 模式下有usage和create的权限（默认）。





### 用户

用户体系创建三权分立的约束，数据库管理员(system)，安全管理员(sso)，审计管理员(sao)各自维护自己权限许可范围内的用户。



创建用户及数据库使外界可访问

```sql
-- 创建用户以及带密码
CREATE USER jc_rdzcgl WITH PASSWORD 'password';

-- 创建测试数据库
create database jc_test;
-- 将数据库授权给用户
GRANT ALL PRIVILEGES ON database jc_test TO jc_rdzcgl;
```



### SQL

#### 数据类型

版本号 *V008R006C008B0014*

- 整形
  - smallint                 ，tinyint 不存在



#### select 查询

```sql
# 分页查询，页码（偏移量）20，从10行开始 
SELECT * FROM "ed_assets" LIMIT 20 OFFSET 10;	
```





##### like

整形不支持 like 操作，需使用 cast 函数将整形转化为字符型。

```sql
SELECT * FROM "ed_assets" WHERE "code" LIKE '%6%' AND cast(id AS varchar) LIKE '%8%';	
```





#### 系统信息

```sql
-- 查看全部表信息
SELECT table_name, table_schema FROM information_schema.tables;
-- 查看字段信息
SELECT table_name, table_schema FROM information_schema.COLUMNS WHERE table_name = 'tbname';

-- 查看指定模式下表信息
SELECT * FROM "information_schema"."tables"
	WHERE "table_schema" = 'public' AND "table_catalog" = 'ksj_zcglxt_24'
;
```



数据库模式类型查询

```sql
SELECT current_setting('database_mode');

-- 查询数据库的兼容模式
SELECT current_setting('compatibility');

SELECT * FROM pg_settings WHERE name LIKE 'database_mode';

# 查询数据最大的连接数
show max_connections;
```



#### license

```sql
# 查看 license 天数（有限期）
select get_license_validdays ();

# 查看 license 全部信息
select get_license_info();
```





### Ksql

Linux环境下，使用工具登陆数据库

```shell
# 切换至，kingbase用户下
su - kingbase

# 使用 system 用户登陆（服务本地免密码情况下）
ksql test system

# 使用 ksql 链接数据
ksql -h 192.168.3.111 -U SYSTEM -p 54321 -W test
ksql test -h 192.168.3.111 -U SYSTEM -p 54321 -W

# 指定数据库名
ksql test -h 192.168.3.111 -U SYSTEM -p 54321 -W -d dbname

# 显示用户下所有数据库
ksql -U system -l

# 版本信息查询
ksql -V
```



Ksql 下SQL 相关查询

```sql
# 查看当前系统用户
select * from sys_user;

# 系统级别的所有用户
SELECT * FROM SYS.DBA_USERS;

# 登录成功后
# 查看当前用户信息
select user;
```





Ksql 下命令话查询

```sql
#
# 命令行查询
# 查看当前用户及信息
\du

# 显示数据库列表
\l
```





### 常见问题

- V008R006C008B0014 tinyint 不存在；可以将其转换为 smallint
- 类型限制：如int不能使用 like 字符字符串，（数字）字符串不能等于整形。提示：*HINT:  没有匹配指定名称和参数类型的操作符. 您也许需要增加明确的类型转换.* 。需要手动使用 **CAST(vchar as int) = 0**



