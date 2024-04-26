# KingbaseES

> 人大金仓数据库
>
> Joshua Conero
>
> 2024年4月26日



工具

- ksql        命令行SQL连接工具
- Kstudio  SQL连接客户端





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

#### 系统信息

```sql
-- 查看全部表信息
SELECT table_name, table_schema FROM information_schema.tables;
SELECT table_name, table_schema FROM information_schema.tables WHERE table_name = 'tbname';
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









