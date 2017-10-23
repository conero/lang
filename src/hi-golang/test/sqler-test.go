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
		Field("name", "sex").
		Field("gender").
		Field("id").
		Select()
	fmt.Println(sql1)
	fmt.Println(sqler.Table("sys_user").Insert())
	fmt.Println(sqler.Table("sys_user").Update())
	fmt.Println(sqler.Table("sys_user").Delete())
	fmt.Println(sqler.Table("sys_user").Select())
}
func main() {
	(&Test{}).Run()
}
