# sqler-go
- 20171023
- Joshua Conero

## v0.0.2/20171025
- 实现简单的 insert 语句生成
- func (sql *Sql) Field(fields ...interface{}) *Sql
    - field 参数从原来的string，扩展到 ...interface{}
- func (sql *Sql) Data(datas ... interface{}) *Sql
    - 设置参数实现(初步)
- func (sql *Sql) Insert() string
    - 可以生成简单 insert 语句

## v0.0.1/20171023
- 实现简单 select 语句生成