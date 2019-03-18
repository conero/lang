

# MySQL

> 2019年3月18日 星期一
>
> Joshua Conero





## 安装

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

    

