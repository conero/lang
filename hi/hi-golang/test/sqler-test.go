// sqler-test
package main

import (
	"fmt"

	sqler "../src/sqler"
)

type Test struct{}

func (t *Test) Run() {
	fmt.Println(sqler.Version)
	t.selectTest()
}
func (t *Test) selectTest() {
	sql1 := sqler.
		Table("sys_user").
		Field("name", "sex", map[string]string{"test": "yang"}).
		Field("gender").
		Field("id").
		Where(map[string]int{"age": 58}, map[string]string{"company": "conero@cn"}).
		Select()
	fmt.Println(sql1)
	fmt.Println(sqler.Table("sys_user").Data(map[string]string{"name": "杨", "gender": "M", "othor": "awsome"}).Insert())
	// 更新数据查询
	fmt.Println(sqler.Table("sys_user").
		Data(map[string]int{"age": 20, "stru_id": 5}, map[string]string{"nick": "Joshua Conero"}).
		Where("name is null").
		Update())
	// 删除数据
	fmt.Println(sqler.Table("sys_user").
		Where(map[string]string{"name": "Susie", "gender": "F"}).
		Delete())
	fmt.Println(sqler.Table("sys_user").Select())
	// oracle 查询数据
	fmt.Println(sqler.Table("sys_user").
		Type("oracle").
		Where(`"name" like '%中%'`).
		Where(map[string]string{"gender": "M"}).
		Select())
}
func main() {
	(&Test{}).Run()
}
