## postgresql

> 2024年5月24日 星期五
>
> Joshua Conero



先进的开源关系数据库



官网地址：[https://www.postgresql.org/](https://www.postgresql.org/)

中文手册：https://github.com/postgres-cn/pgdoc-cn





### psql

sql 查询工具

```shell
# 查看客户端编码格式
select name,setting,context from pg_settings where name like '%encoding%';

# 设置编码
\encoding UTF8

# 导入sql
\i D:/conero/xxx/test_data.sql

# 切换数据库
\c mydb
```

