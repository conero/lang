# sqler-go
- 20171023
- Joshua Conero

## V0.1.0/20171027
- 实现基本的数据库语法生成
- 出口以后会自动，清理运行时的缓存；用于准备下一次调用(即单出口多出口)
```go
    // 数据库查询实例
	fmt.Println(sqler.
		Table("sys_user").
		Field("name", "sex", map[string]string{"test": "yang"}).
		Field("gender").
		Field("id").
		Where(map[string]int{"age":58}, map[string]string{"company":"conero@cn"}).
        Select())
    // SELECT `name`,`sex`,`test` AS `yang`,`gender`,`id` FROM `sys_user` WHERE `age`=58 AND `company`='conero@cn'

    // 数据新增
	fmt.Println(sqler.
        Table("sys_user").
        Data(map[string]string{"name": "杨", "gender": "M", "othor": "awsome"}).
        Insert())
    // INSERT INTO `sys_user`(`gender`,`othor`,`name`) VALUES ('M','awsome','杨')

	// 更新数据查询
	fmt.Println(sqler.Table("sys_user").
		Data(map[string]int{"age": 20, "stru_id": 5}, map[string]string{"nick": "Joshua Conero"}).
		Where("name is null").
        Update())
    // UPDATE `sys_user` SET `nick`='Joshua Conero',`stru_id`=5,`age`=20 WHERE name is null

	// 删除数据
	fmt.Println(sqler.Table("sys_user").
		Where(map[string]string{"name": "Susie", "gender": "F"}).
        Delete())
    //  DELETE FROM `sys_user` WHERE `name`='Susie' AND `gender`='F'
    fmt.Println(sqler.Table("sys_user").Select())
    //  SELECT * FROM `sys_user`

	// oracle 查询数据
	fmt.Println(sqler.Table("sys_user").
		Type("oracle").
		Where(`"name" like '%中%'`).
		Where(map[string]string{"gender": "M"}).
        Select())
    //  SELECT * FROM "sys_user" WHERE "name" like '%中%' AND "gender"='M'
```

- 支持数据库: oracle, mysql
- 入口方法：
    - func Table(name string) *Sql  通过设置数据库表名入口 创建sqler
    - func Type(vType string) *Sql  通过设置数据库类型入口 创建sqler
- 中间件：
    - func (sql *Sql) Type(vtype string) *Sql
    - func (sql *Sql) Table(tableParam ...string) *Sql  设置数据表名， 参数 @param table[, alias]
    - func (sql *Sql) Field(fields ...interface{}) *Sql 设置数据列名，用于 Select
    - func (sql *Sql) Data(datas ...interface{}) *Sql   设置数据， 用于 Insert/Update
    - func (sql *Sql) Where(wheres ...interface{}) *Sql 设置查询条件， 用于 Select/Update/Delete
- 出口
    - func (sql *Sql) Insert() string   新增数据
    - func (sql *Sql) Delete() string   删除数据
    - func (sql *Sql) Update() string   更新数据
    - func (sql *Sql) Select() string   查询数据

## v0.0.3/20171026
- db_adapter 数据库适配器类型新增“正则”处理数据库列名被重复“标识化”
    - 如: mysql "noun" => ""noun""
    
- 修复 sqler.table 被重复“标识化”
    - 如: mysql -> oracle 
        - \`table_name\` => "\`table_name\`"    
    
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