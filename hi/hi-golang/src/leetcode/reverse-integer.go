package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
   @date 2018年10月25日 星期四
   @description 给定一个 32 位有符号整数，将整数中的数字进行反转。
   @link https://leetcode-cn.com/problems/reverse-integer/
   @name 反转整数


	给定一个 32 位有符号整数，将整数中的数字进行反转。
	示例 1:

	输入: 123
	输出: 321
	 示例 2:

	输入: -123
	输出: -321
	示例 3:

	输入: 120
	输出: 21
	注意:

	假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。根据这个假设，如果反转后的整数溢出，则返回 0。
*/

func main() {
	fmt.Println(reverse(123), reverse(-123), reverse(120), reverse(1534236469), reverse(1463847412))

	fmt.Println(int(math.Pow(-2, 31)), math.Pow(-2, 31-1))
	// 32 bit(
	var i32 int32 = 1534236469
	fmt.Println(i32)
}

/*
*
基于字符串的字符串处理
*/
func baseStrHandle(x int) int {
	isNig := false
	if x < 0 {
		isNig = true
	}
	x = int(math.Abs(float64(x)))
	str := strconv.Itoa(x)
	sLen := len(str)
	strQue := []string{}
	for i := sLen; i > 0; i-- {
		strQue = append(strQue, str[i-1:i])
	}

	value, err := strconv.Atoi(strings.Join(strQue, ""))
	if err != nil {
		value = 0
	} else {
		if isNig && value != 0 {
			value = value * -1
		}
	}

	// 32 位整形判断   [−2^31,  2^(31 − 1)]
	s32 := int(math.Pow(-2, 31))
	e32 := int(math.Pow(2, 31)) - 1
	if value < s32 || value > e32 {
		value = 0
	}
	return value
}

func reverse(x int) int {
	return baseStrHandle(x)
}
