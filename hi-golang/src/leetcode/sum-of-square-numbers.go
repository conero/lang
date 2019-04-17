package main

import (
	"fmt"
	"math"
)

/**
 * @DATE        2019/4/16
 * @NAME        Joshua Conero
 * @DESCRIPIT   633. 平方数之和
 * @LINK		https://leetcode-cn.com/problems/sum-of-square-numbers/
 **/

// 题目
/*
给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a2 + b2 = c。

示例1:
	输入: 5
	输出: True
	解释: 1 * 1 + 2 * 2 = 5

示例2:
	输入: 3
	输出: False
*/

// @TODO 超出时间限制 / StillNeedToDo
// 解答
func judgeSquareSum(c int) bool {
	var value bool
	for i := 0; i <= c; i++ {
		for j := i; j <= c; j++ {
			// 两个数大于时退出第二个数的寻找
			vn := math.Pow(float64(i), 2) + math.Pow(float64(j), 2)
			if int(vn) > c {
				break
			}
			if c == int(vn) {
				value = true
				break
			}
		}
		// 外层中断
		if value {
			break
		}
	}
	return value
}

// 控制台
func main() {
	var vin int
	var vout, vneed bool
	const format = "%v)  [%v] => [%v] VS [%v]\r\n"

	vin = 5
	vout, vneed = judgeSquareSum(vin), true
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

	vin = 3
	vout, vneed = judgeSquareSum(vin), false
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

	vin = 4
	vout, vneed = judgeSquareSum(vin), true
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

	vin = 2
	vout, vneed = judgeSquareSum(vin), true
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

	// 压力/性能测试
	vin = 10000000
	vout, vneed = judgeSquareSum(vin), true
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

	vin = 999999999
	vout, vneed = judgeSquareSum(vin), true
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

}
