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

// @TODO 程序算法设计存在问题
// @TODO StillNeedToDo
// 解答
func judgeSquareSum(c int) bool {
	var value bool
	for i := 1; i <= c; i++ {
		vn := math.Pow(float64(i-1), 2) + math.Pow(float64(i), 2)
		fmt.Println(i-1, i, vn, c)
		if c == int(vn) {
			value = true
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
}
