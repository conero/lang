

# MySQL

> 2019年3月18日 星期一
>
> Joshua Conero





## 安装

### 配置

数据库表：*performance_schema*

```mysql
# MySQL ，通过 SQL 修改 数据库的配置信息
# performance_schema.global_variables 全局配置
set global max_allowed_packet = 200*1024*1024;
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

# 选择数据库
use dbname;
# 使用 source 导入脚本
source d:/.../name.sql
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