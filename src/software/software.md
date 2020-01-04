# 软件使用帮助

> 2018年11月14日 星期三



## 常用软件列表

| 软件名称                                                    | 商用类型  | 分类                    | 备注                                                     | 语言        |
| ----------------------------------------------------------- | --------- | ----------------------- | -------------------------------------------------------- | ----------- |
| PowerDesigner                                               | 商业      | 数据库工具              | 数据库模型设计工具                                       | SQL         |
| PhpStrom(PS)                                                | 商业      | JetBrains IDE           | PHP 代码编辑器                                           | PHP         |
| fiddler                                                     | 免费      | web 调试代理 (http)     | [查看官网](https://docs.telerik.com/fiddler)             | .NET(C#)    |
| *Wireshark*                                                 | 免费      | 网络抓包工具            | [https://www.wireshark.org/](https://www.wireshark.org/) |             |
| [*Typora*](https://typora.io/)                              | 免费/开源 | markdown 编辑/查看器    |                                                          | electron/Js |
| [Anaconda](https://www.anaconda.com/)                       | 含免费    | python 科学运算包       |                                                          | python      |
| Android Studio                                              | 免费      | Android 开发工具        | [中文官网](https://developer.android.google.cn/studio/)  |             |
| [Axure](https://www.axure.com/)                             | 商业      | 设计工具                | 类似 Visio                                               |             |
| [*Code*::*Blocks*](http://www.codeblocks.org/)              | 免费      | C/C++ 开发工具          | IDE                                                      |             |
| [Cygwin](https://www.cygwin.com/)                           | 免费      | windows 的liunx支持软件 | 操作系统插件，类 Linux件                                 |             |
| [exe4j](https://exe4j.apponic.com/)                         | 商业/免费 | Java 打包为exe的工具    | java 插件                                                | java        |
| [Lantern](https://github.com/getlantern/lantern)            | 商业/免费 | VSN 工具                |                                                          | go          |
| [LLVM](https://llvm.org/)                                   | 免费      | 编程语言编译器组件      | [llvm-mirror](https://github.com/llvm-mirror/llvm)       |             |
| [Postman](https://www.getpostman.com/)                      | 免费      | API 调试工具            |                                                          | electron/Js |
| [Q-Dir](https://q-dir.en.softonic.com/)                     | 免费      | 多窗口文件整理工具      |                                                          |             |
| [SQLiteExpert](http://www.sqliteexpert.com/)                | 商业/免费 | sqlite 可视化管理工具   |                                                          | sqlite      |
| [TDM-GCC](http://tdm-gcc.tdragon.net/)                      | 免费      | GCC 编译器套件          | 如windows 下为 cgo 提供gcc环境                           | gcc         |
| [Search and Replace](http://www.funduc.com/srshareware.htm) | 免费      | 文本查询/替换工具       |                                                          |             |
| [DBeaver](https://dbeaver.io/)                              | 商业/社区 | 通用数据库管理工具      |                                                          | Java/JVM    |





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



> 数据库模型通过SQL生成

`File > Reverser Engineer >  Database`



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



### DBeaver

*数据库通用连接工具*



#### MySQL 8.0 连接错误的解决方法

**错误信息**： `Public Key Retrieval is not allowed`

**解决**  将软件中的 `allowPublicKeyRetrieval` 设置为 `true` 即可。 [参考地址](https://www.jianshu.com/p/240ecb210222)





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



