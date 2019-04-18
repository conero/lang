package main

import (
	"fmt"
	"math"
)

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

// @TODO 超出时间限制 / StillNeedToDo
// 解答
// 求指出最大三个数
func maximumProduct(nums []int) int {
	var value, vlen int = 0, len(nums)
	var valueNil bool = true
	for i := 0; i < vlen; i++ {
		for j := i + 1; j < vlen; j++ {
			for k := j+ 1; k < vlen; k++{
				if valueNil{
					value = nums[i]*nums[j]*nums[k]
					valueNil = false
					//fmt.Println(nums[i], nums[j], nums[k], value)
					continue
				}
				value = int(math.Max(float64(value), float64(nums[i]*nums[j]*nums[k])))
				//fmt.Println(nums[i], nums[j], nums[k], value)
			}
		}
	}
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
