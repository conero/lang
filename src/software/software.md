# 软件使用帮助

> 2018年11月14日 星期三



## 常用软件列表

| 软件名称      | 商用类型 | 分类       | 备注               |
| ------------- | -------- | ---------- | ------------------ |
| PowerDesigner | 商业     | 数据库工具 | 数据库模型设计工具 |
|               |          |            |                    |
|               |          |            |                    |





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





### IDEs



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

