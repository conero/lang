package main

// @name 69. x 的平方根
// @link: https://leetcode-cn.com/problems/sqrtx/
// @data 2019年3月26日 星期二

import (
	"fmt"
	"math"
)

/**
实现 int sqrt(int x) 函数。
计算并返回 x 的平方根，其中 x 是非负整数。
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:
	输入: 4
	输出: 2

示例 2:
	输入: 8
	输出: 2
	说明: 8 的平方根是 2.82842...,
		 由于返回类型是整数，小数部分将被舍去。
*/

// @TODO 可以改善，使用真正的算法解决
// 解答
func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

// 测试实例
func main() {
	vfmt := "%t) [%v] VS [%v] <<= [%v]"
	var vin, vout, vneed int

	// test 1
	vin, vout = 4, 2
	vneed = mySqrt(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))

	// test 2
	vin, vout = 8, 2
	vneed = mySqrt(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))

}
