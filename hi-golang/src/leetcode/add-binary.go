package main

// @name 67. 二进制求和
// @link: https://leetcode-cn.com/problems/add-binary/
// @data 2019年3月26日 星期二

import (
	"fmt"
	"strconv"
)

/**
给定两个二进制字符串，返回他们的和（用二进制表示）。
输入为非空字符串且只包含数字 1 和 0。

示例 1:
	输入: a = "11", b = "1"
	输出: "100"

示例 2:
	输入: a = "1010", b = "1011"
	输出: "10101"
*/

// 解答
func addBinary(a string, b string) string {
	var s string
	aLen, bLen := len(a), len(b)
	//aLen := len(a)
	//bLen := len(b)
	//fmt.Println("LEN:", aLen, bLen, a ,b )
	lastBit := 0
	for i :=1; ;i++  {
		if i > aLen && i > bLen{
			break
		}
		aBit, bBit := 0, 0
		if i <= aLen{
			aBit, _ = strconv.Atoi(a[aLen-i:aLen-i+1])
		}
		if i <= bLen{
			bBit, _ = strconv.Atoi(b[bLen-i:bLen-i+1])
		}
		//fmt.Println(i, ":", aBit, bBit, lastBit, aBit + bBit + lastBit)
		aBit = aBit + bBit + lastBit
		lastBit = 0
		if aBit > 1{
			aBit -= 2
			lastBit = 1
		}
		s = strconv.Itoa(aBit) + s
	}
	if lastBit == 2{
		lastBit = 0
		s = "10" + s
	}
	if lastBit > 0{
		s = strconv.Itoa(lastBit) + s
	}
	return s
}

// 测试实例
func main() {
	vfmt := "%t) [%v] VS [%v] <<= [%v, %v]"
	var vinA, vinB, vout, vneed string

	// test 1
	vinA, vinB, vout = "11", "1", "100"
	vneed = addBinary(vinA, vinB)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vinA, vinB))

	// test 2
	vinA, vinB, vout = "1010", "1011", "10101"
	vneed = addBinary(vinA, vinB)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vinA, vinB))

	// test 2
	vinA, vinB, vout = "1111", "1111", "11110"
	vneed = addBinary(vinA, vinB)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vinA, vinB))
}
