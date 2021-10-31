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



### 脚本示例

#### DDL

```sqlite
-- AtGroup  创建表自增
CREATE TABLE AtGroup(
    Gid            integer PRIMARY KEY autoincrement,
    GroupName       varchar(32) not null
);
```



