# 数据库

> 2019年5月31日 星期五
>
> *Joshua Conero*





## sqlite

[官方网站](https://www.sqlite.org/)



*使用命令行工具：(Command Line Shell For SQLite)*

```shell
# 可创建数据库 dbname.db， 并进入交互的 shell 中
$ sqlite3 dbname.db

# 查看 sqlite 中的命令
sqlite> .help
```



### 数据类型

| 类型       | 描述       | 其他                            |
| ---------- | ---------- | ------------------------------- |
| NULL       | 空值       | 原生                            |
| INTEGER    | 整形       | 原生                            |
| REAL       | 浮点型     | 原生                            |
| TEXT       | 文本字符串 | 原生                            |
| BLOB       | blob 类型  | 原生                            |
| BOOLEAN    | 布尔类型   | integers 0 (false) and 1 (true) |
| NUMERIC    | 整形别名   | INTEGER                         |
| VARCHAR(n) | 字符串类型 | 字符串长度                      |
| CHAR       | 字符串类型 |                                 |



> 扩展类型

| Example Typenames From The CREATE TABLE Statement          or CAST Expression | Resulting Affinity | Rule Used To Determine Affinity |
| ------------------------------------------------------------ | ------------------ | ------------------------------- |
| INT    INTEGER    TINYINT    SMALLINT    MEDIUMINT    BIGINT    UNSIGNED BIG INT    INT2    INT8 | INTEGER            | 1                               |
| CHARACTER(20)    VARCHAR(255)    VARYING CHARACTER(255)    NCHAR(55)    NATIVE CHARACTER(70)    NVARCHAR(100)    TEXT    CLOB | TEXT               | 2                               |
| BLOB    *no datatype specified*                              | BLOB               | 3                               |
| REAL    DOUBLE    DOUBLE PRECISION    FLOAT                  | REAL               | 4                               |
| NUMERIC    DECIMAL(10,5)    BOOLEAN    DATE    DATETIME      | NUMERIC            | 5                               |



#### 字符串函数

```sqlite
-- replace 替换字符串
update website_resources set "path" = replace("path", 'https://swt.gz.gov.uk', '') where is_foreign = 0 and "path" like 'http%';
```





### 脚本示例

#### DDL

```sqlite
-- drop table AtGroup;
-- AtGroup  创建表自增
CREATE TABLE AtGroup(
    Gid            integer PRIMARY KEY autoincrement,
    GroupName       varchar(32) not null
);

-- 默认当前系统时间
CREATE TABLE client(
    id               integer PRIMARY KEY autoincrement,
    account          varchar(50) not null, 						-- comment '账号'
    password         varchar(50) not null, 						-- comment '密码'
    create_at        datetime not null default (datetime('now', 'localtime'))
);
```



##### update

数据更新

```sqlite
-- 使用 replace into 实现 update join. replace into 必须列出全部的列
-- insert or replace into  可是先新增或合并数据，根据表的约束条件
replace into website_elements("id", "path")
	select we.id, SUBSTRING(we."path", LENGTH(w.url)+1) from website_elements we
		join websites w on we.site_id = w.id
		where instr(we."path", w.url) > 0
; 
```





## SQLServer

微软(Microsoft)推出的关系型数据库，1988年



使用客户端 `sqlcmd` 进行数据库链接

```shell
# 链接本地数据库
sqlcmd -S 127.0.0.1 -U conero
```







## 达梦数据库

> dmdb  https://www.dameng.com/

当前版本号 V8



一个数据库，为其单独实例，类似于 oracle。



可使用软件自带的**DM数据迁移工具**，实现dm数据库与oracle、MySQL、pgsql的数据迁移；同时支持 sql、csv等文件迁移。



### linux 配置

安装数据库

```shell
# 下载安装包后安装程序
./DMInstall.bin -i

# 安装后需要 dminit 工具进行初始化
dminit

# 卸载代码软件
./uninstall.sh -i

# 服务状态查看
dmserver status
```



配置环境变量

```shell
# 在<~/.bash_profile>中配置 达梦数据库
# 使其生效
DM_HOME=/home/dmdb/dmdbms/data
export PATH=$PATH:$DM_HOME/bin
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:$DM_HOME/bin
```



服务器达梦数据配置文件：

```shell
/etc/dm_svc.conf
```





### 入门

使用 **达梦数据库配置助手** 工具创建数据库，与 oracle 相似，即每个数据库需要单独配置。

授权：默认授权只有1年，一年后开发者等无法使用必须使用

模式：数据库创建时其默认模式为与用于同名模式。



数据库导入时需要相同的文件名。



数据库相关版本信息查询

```sql
select * FROM v$instance;

-- 查询当前连接数据库信息
SELECT * FROM V$DATABASE;

-- 授权信息查看
select * FROM v$license;
```



### 部署

- 测试用户（前端）：`18798011264/Conero@2024`



正式站程序：

- https://zscq.gyamrszsj.cn/





本地部署（windows）：

```powershell
# 启动cgi服务
.\php-cgi.exe -b 0.0.0.0:9003

# 使用 wpg 启动
wpg web -r 'E:\conero\repository\conero-blockd\gybhzx2411\guiyanggch\' -R '^(.*)$ /index.php?s=/$1' -addr 19211 -c 'tcp://127.0.0.1:9003' -C gb18030


# admin  部署
wpg web -r 'E:\conero\repository\conero-blockd\gybhzx2411\guiyanggch\' -R '^(.*)$ /spzgadmin.php?s=/$1' -addr 19212 -c 'tcp://127.0.0.1:9003'
```





### 异常

- php 链接服务中文乱码 php74（pdo），修改`C:\windows\System32\dm_svc.conf` 加入行文件：

```shell
CHAR_CODE=(PG_UTF8)
```





## MongoDB

使用 C++开发的文档数据库



与常见关系数据库概念差异，其他的 database、index、primary key等类似

| SQL术语 | MongoDB    | 说明            |
| ------- | ---------- | --------------- |
| table   | collection | 数据表/集合     |
| row     | document   | 数据行/文档     |
| column  | field      | 数据字段/数据域 |



MongoDB的最小存储单位就是文档document对象。



开发驱动

- golang   [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)



其他基本概念

- 默认端口为 27017



### v4.2

使用 mongo shell 连接或查询MongoDB

> 连接数据库

```shell
# 连接数据库
mongo
# 指定端口号
mongo --port 28015
# 指定主机地址
mongo --host mongodb0.example.com:28015
# 使用密码连接
mongo "mongodb://alice@mongodb0.examples.com:28015/?authSource=admin"

# 指定
mongo -u root

# 指定 root 用户并以 admin 管理人登录
mongo -u root --authenticationDatabase admin
```



> 操作数据库

mongo shell 执行

```javascript
// 查看当前数据库
db

// 显示当前数据列表，展示类别及大小
show dbs
// 查看数据库列表，简单模式
db.getMongo().getDBNames()
// 


// 查看数据库方法
db.stats()

// 创建用户
// 使用 admin 登陆用户：  mongo --authenticationDatabase admin
db.createUser(
  {
    user: "root",
    pwd: "$1^34xiusG3k6GG+", // 使用强密码
    roles: [
      { role: "userAdminAnyDatabase", db: "admin" }, // 允许管理用户和角色
      "readWriteAnyDatabase" // 允许在所有数据库中读写
    ]
  }
)

// 创建用户并对其进行授权
db.createUser(
  {
    user: "idis_user",
    pwd: "T**o-88",
    roles: [
      { role: "readWrite", db: "idis_lhs" },
      { role: "readWrite", db: "handle" }
    ]
  }
)

// 修改用户权限
db.grantRolesToUser(
  "databaseManager",
  [
    { role: "readWrite", db: "specificDB" } // 给 specificDB 数据库授予权限
  ]
)

// 修改用户权限
db.grantRolesToUser(
  "databaseManager",
  [
    { role: "readWrite", db: "specificDB" } // 给 specificDB 数据库授予权限
  ]
)

// 删除用户权限
db.revokeRolesFromUser(
  "databaseManager",
  [
    { role: "readWrite", db: "specificDB" } // 移除 specificDB 数据库的权限
  ]
)
```



#### openEuler 安装

安装 MongoDB v4.2.25 数据库

```shell
#下载软件包
wget https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-rhel80-4.2.25.tgz

# 解压
tar -zxvf mongodb-linux-x86_64-rhel80-4.2.25.tgz

# 将包移动到 mongodb 中
mkdir ../mongodb
mv  mongodb-linux-x86_64-rhel80-4.2.25/* ../mongodb/

# 编辑 mongod 配置文件
vi /etc/mongod.conf

# 运行服务
/home/mongodb/bin/mongod -f /etc/mongod.conf
```





#### 备份及恢复

```shell
# 一次性导出 mongo 中全部数据库
mongodump --out /path/to/backup/folder/
# 指定用户
mongodump -u root --out /data/bak/250108-qy-new-openeuler

# 导出数据库 handle 
# 格式 --gzip
mongodump -uroot -p$pswd --authenticationDatabase admin --db handle -o ./250109-handle --gzip

# 恢复数据
mongorestore --db handle /home/bak/250109-handle --gzip

# 以 admin 角色登录数据库
mongo --authenticationDatabase admin
```



## neo4j

基于Java的图形数据库。neo4j是一个高性能的NoSQL图形数据库，它将结构化数据存储在网络（从数学角度称为图）上而不是传统的表中。‌ Neo4j是一个嵌入式的、基于磁盘的、具备完全事务特性的Java持久化引擎，特别适合处理具有复杂关系的数据‌。



- 官网：https://neo4j.com/
- github: https://github.com/neo4j/neo4j





### v3.3

#### windows 安装

- 安装参考   [Neo4j 安装、使用教程](https://blog.csdn.net/zeroheitao/article/details/122925845) ，依赖java 1.8 。 若安装Neo4j 4.x 版本，需使用 jdk11

```powershell
# ps 5.1 版本和管理原
cd D:\neo4j-community-3.5.5\bin

# 以服务形式安装
./neo4j install-service

# 运行服务器
./neo4j start

# 查看状态
./neo4j status
```



安装后默认数据密码：` neo4j/neo4j` ，浏览器访问地址：http://localhost:7474 接口修改密码及登录
