/**
 * golang 指针测试
 * 2017年11月25日 星期六
 */
package main

import (
	"fmt"
)

// 基本支持测试
func BasePointer() {
	tA := []int{}
	var tAPnt *[]int
	// 初始化
	tA = append(tA, []int{1, 8, 7, 9, 8, 0}...)
	fmt.Println(tA)
	fmt.Println(tAPnt)

	// 指针引用
	tAPnt = &tA
	fmt.Println("tAPnt 去数组 tA的地址以后", tAPnt)
	//直接计算
	//tAPnt = append(tAPnt, []int{1, 3, 1, 4, 5, 2, 1})			// Error: first argument to append must be slice; have *[]int
	*tAPnt = append(*tAPnt, []int{1, 3, 1, 4, 5, 2, 1}...) // Error: first argument to append must be slice; have *[]int

	// 去除值，并替换
	//tmpArr := *tAPnt
	//tmpArr = append(tmpArr, []int{1, 3, 1, 4, 5, 2, 1}...)
	//tAPnt = &tmpArr

	fmt.Println("tAPnt 去数组 tA的地址以后", tAPnt)
	fmt.Println("tAPnt 去数组 tA的地址以后", tA)
}

func main() {
	BasePointer()
}
