package main

import "fmt"

/**
 * @DATE        2019/4/16
 * @NAME        Joshua Conero
 * @DESCRIPIT   628. 三个数的最大乘积
 * @LINK		https://leetcode-cn.com/problems/maximum-product-of-three-numbers/
 **/

// 题目
/*
给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

示例 1:
	输入: [1,2,3]
	输出: 6

示例 2:
	输入: [1,2,3,4]
	输出: 24

注意:
	给定的整型数组长度范围是[3,10^4]，数组中所有的元素范围是[-1000, 1000]。
	输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。
 */


// @TODO 程序算法设计存在问题
// @TODO StillNeedToDo
// 解答
// 求指出最大三个数
func maximumProduct(nums []int) int {
	var value int
	var n1, n2, n3 int = 0, 0, 0
	var n10, n20, n30 bool = false, false, false
	for _, v := range nums {
		if !n10 {
			n1 = v
			n10 = true
			continue
		} else if !n20 {
			n2 = v
			n20 = true
			continue
		} else if !n30 {
			n3 = v
			n30 = true
			continue
		}

		// 最大的数字
		if v > n1 {
			n1, v = v, n1
		}
		//第二大的数字
		if v > n2 {
			n2, v = v, n2
		}
		// 第三大的数字
		if v > n3 {
			n3 = v
		}
	}
	//fmt.Println(n1, n2, n3)
	value = n1 * n2 * n3
	return value
}

// 控制台
func main() {
	var vin []int
	var vout, need int
	var format string = "%v)  [%v] => [%v] VS [%v]\r\n"

	vin = []int{1, 2, 3}
	need, vout = 6, maximumProduct(vin)
	fmt.Printf(format, vout == need, vin, vout, need)

	vin = []int{1, 2, 3, 4}
	need, vout = 24, maximumProduct(vin)
	fmt.Printf(format, vout == need, vin, vout, need)

	vin = []int{-1, -2, -3}
	need, vout = -6, maximumProduct(vin)
	fmt.Printf(format, vout == need, vin, vout, need)

	vin = []int{-4, -3, -2, -1, 60}
	need, vout = 720, maximumProduct(vin)
	fmt.Printf(format, vout == need, vin, vout, need)
}
