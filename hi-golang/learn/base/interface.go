/**
2017年11月4日 星期六
Joshua Conero
golang interface{} 类型测试
*/

package main

import (
	"fmt"
)

func anyparam(any interface{}) interface{} {
	return any
}

// 基本测试
func baseInterface() {
	// ?? fmt.Println 与 println 的区别
	var any interface{} // 地址
	any = 5
	fmt.Println(any, "int", anyparam(8*9) == 72, any == 5)
	println(any, any.(int), any == 5)

	// 数组
	any = []interface{}{
		8,
		74,
		"test",
		0.245,
	}
	fmt.Println(any, "array", any.([]interface{})[0] == 8)
	println(any, any.([]interface{}), anyparam(8*9) == 72, any.([]interface{})[1] == 8)
}

func main() {
	baseInterface()
}