package main

import (
	"fmt"
	"strconv"
)

// @name 258. 各位相加
// @link: https://leetcode-cn.com/problems/add-digits/
// @data 2019年4月6日 星期六

/*
给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。

示例:
	输入: 38
	输出: 2

解释: 各位相加的过程为：3 + 8 = 11, 1 + 1 = 2。 由于 2 是一位数，所以返回 2。

进阶:
	你可以不使用循环或者递归，且在 O(1) 时间复杂度内解决这个问题吗？
*/

// @TODO CanBeBetter
// 解答
/*
	1. 实现开根预算实现位上数字解析
*/
func addDigits(num int) int {
	n := num
	for n > 9 {
		//math.Sqrt()
		n2s := strconv.Itoa(n)
		s1, s2 := n2s[0:1], n2s[1:]
		sn1, _ := strconv.Atoi(s1)
		sn2, _ := strconv.Atoi(s2)
		//fmt.Println(n, sn1, sn2)
		n = sn1 + sn2
	}
	return n
}

// 控制器
func main() {
	vfmt := "%t) [%v] VS [%v] <<= [%v]"
	var vin, vout, vneed int

	// test 1
	vin, vneed = 38, 2
	vout = addDigits(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))

	// test 1
	vin, vneed = 10, 1
	vout = addDigits(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))

}
