# 软件使用帮助

> 2018年11月14日 星期三



## 常用软件列表

| 软件名称      | 商用类型 | 分类                | 备注                                                     | 语言     |
| ------------- | -------- | ------------------- | -------------------------------------------------------- | -------- |
| PowerDesigner | 商业     | 数据库工具          | 数据库模型设计工具                                       | SQL      |
| PhpStrom(PS)  | 商业     | JetBrains IDE       | PHP 代码编辑器                                           | PHP      |
| fiddler       | 免费     | web 调试代理 (http) | [查看官网](https://docs.telerik.com/fiddler)             | .NET(C#) |
| *Wireshark*   | 免费     | 网络抓包工具        | [https://www.wireshark.org/](https://www.wireshark.org/) |          |





> 使用软件推荐

| 名称                                                         | 语言      | 商用 | 其他                                                         |
| ------------------------------------------------------------ | --------- | ---- | ------------------------------------------------------------ |
| [Oracle SQL Developer](https://www.oracle.com/database/technologies/appdev/sql-developer.html) | JAVA(JVM) | 免费 | 查询分析器,[文档](https://docs.oracle.com/en/database/oracle/sql-developer/) |
|                                                              |           |      |                                                              |
|                                                              |           |      |                                                              |





## 数据库及相关

### PowerDesigner

> 修改`模块生成规则` 的方法(*v 16.5.0.3982*)

```ini
# 1
Database/Edit Current DBMS…


# 2 如常量
找到 Scipt/Objects/Reference/ConstName
# 默认为： FK_%.U8:CHILD%_%.U9:REFR%_%.U8:PARENT%
# 修改为新的
FK_%.U40:REFR%
```





### Oracle SQL Developer

- [文档 https://docs.oracle.com/en/database/oracle/sql-developer/](https://docs.oracle.com/en/database/oracle/sql-developer/)



> **显示行数**

*版本 18.2.0.183* `首选项>>代码编辑器>>行装订线>>显示行数`



> **结合oracle 数据库和 SQL Developer 做数据处理**

- *场景： 一堆数据 Excel 数据表，通过对其的逻辑处理得到另一个结果集。将 Excel 数据导入 Oracle，然后借助 SQL 特性处理*
  1. *整体处理 Excel 数据表，使首行为同数据表对应的列字段集合*
  2. *将 Excel 格式另保存为 `csv` 格式*
  3. *创建对应的数据表*
  4. *使用 SQL Developer 导入，`csv` 数据*
  5. *通过 SQL 处理数据*



## IDEs



### JetBrains

> [JetBrains 系列IDE](https://www.jetbrains.com/)



*常用快捷键*

- `Ctrl+F4`  关闭当前的窗口
- `Ctrl+D`    *复制行（double）*
- `Ctrl+X`     *删除行*
- `Ctrl+E`     *最近打开的文件*
- `Ctrl+[/]`  *跳到大括号的开头结尾*
- `Ctrl＋F12`  *可以显示当前文件的结构*
- `Ctrl＋N`       *快速打开类*



#### PhpStrom(PS)

> *技巧 (Windows)*
>
> [官网文档](https://www.jetbrains.com/help/phpstorm/meet-phpstorm.html)

- `Shift`    两次，全局文件搜索
- `Ctrl+B, Ctrl+Click`  调整至源代码



