// @Date：   2018年12月18日 星期二
// @Author:  Joshua Conero
// @Name:    库入口文件

// reference link
//    + https://www.cnblogs.com/golove/p/3286303.html
package main

import (
	"fmt"
)

func main() {
	// var datas interface{}
	// fmt.Scan(&datas)
	// fmt.Println(datas)

	fmt.Println("5555")

	// 数据匹配
	// var s string
	// fmt.Scanf("get the string: %s", &s)
	// fmt.Println(" usage: " + s)

	// 输出颜色
	// liunx/bash 中颜色正常
	// + https://studygolang.com/articles/2500
	fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, "testPrintColor", 0x1B)
}
