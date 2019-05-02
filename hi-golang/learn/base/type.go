/**
2017年11月2日 星期四
Joshua Conero
golang 类型测试
*/

package main

import (
	"fmt"
)

type BaseTypeCaseTest struct {
}

// 运行台
func (btct *BaseTypeCaseTest) Console() {
	// btct.baseType()
	btct.childArray()
}

func (btct BaseTypeCaseTest) baseType() {
	var str string
	var mapType map[string]interface{}
	// 未初始化的对象
	fmt.Println(mapType, mapType == nil, str)
	// 直接赋值
	//mapType["test"] = 5		// ERROR:  panic: assignment to entry in nil map

	fmt.Println(mapType)

	// 类型别名
	var number1 int = 8
	type nInt int
	type nInt2 = int // 类型别名
	var number2 nInt = 47
	var number3 nInt2 = 47
	fmt.Println(number1, number2)
	//number1 = number2			// ERROR:   cannot use number2 (type nInt) as type int in assignment
	number2 = nInt(number1)
	number1 = number3
	fmt.Println(number1, number2, number3)
}

func (btct BaseTypeCaseTest) childArray() {
	a1 := []int{1, 2, 4, 5, 14, 74, 100, 104}
	fmt.Println(a1, len(a1))
	a1 = a1[0:2]
	fmt.Println(a1, len(a1))
	fmt.Println(a1[0:3], len(a1[0:3]))
}

func main() {
	(&BaseTypeCaseTest{}).Console()
}
