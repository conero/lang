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

