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
		Field("name", "sex", map[string]string{"test":"yang"}).
		Field("gender").
		Field("id").
		Select()
	fmt.Println(sql1)
	fmt.Println(sqler.Table("sys_user").Data(map[string]string{"name":"杨", "gender":"M", "othor":"awsome"}).Insert())

	fmt.Println(sqler.Table("sys_user").Update())
	fmt.Println(sqler.Table("sys_user").Delete())
	fmt.Println(sqler.Table("sys_user").Select())
	fmt.Println(sqler.Table("sys_user").Type("oracle").Select())
}
func main() {
	(&Test{}).Run()
}
