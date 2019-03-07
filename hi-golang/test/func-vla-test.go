/**
    2017年10月25日 星期三
	golang 不定长参数测试
*/
package main

import (
	"fmt"
	"time"
)

// 不定长参数求和
func Sum(ns ...int) (int, int) {
	t := 0
	for _, n := range ns {
		t = t + n
	}
	return len(ns), t
}

// 调用不定函数
func callSum(dv int, ns ...int) float32 {
	// ns 作为为不定函数参数的写法
	n, sum := Sum(ns...)
	t := (sum / n) * dv
	return float32(t)
}

// 函数作为参数
func callFnc(f func(ns ...int) int) (int, func() string) {
	sum := f(1, 2, 3, 5, 6, 7, 8, 9, 10)
	sum = sum / 8
	return sum, func() string {
		return "传入函数且返回函数"
	}
}

// 入库函数
func main() {
	fmt.Println(Sum(4, 8, 9, 41))
	fmt.Println(callSum(78, 7, 6, 9, 4, 14, 85, 47))
	var sumArgsNs []int
	sumArgsNs = []int{7, 4, 5, 7, 6, 3, 1, 5}
	//fmt.Println(sumArgsNs...)	// cannot use sumArgsNs (type []int) as type []interface {} in argument to fmt.Println
	fmt.Println(Sum(sumArgsNs...))

	// 匿名函数
	func() {
		fmt.Println("使用匿名函数")
		fmt.Println(time.Now())
	}()

	// 函数作为参数
	num, newF := callFnc(func(ns ...int) int {
		num := 0
		for i, v := range ns {
			num = num + (v / (i + 1))
		}
		return num
	})
	fmt.Println(num, newF()+": Jason maybe say fuck to you")
}
