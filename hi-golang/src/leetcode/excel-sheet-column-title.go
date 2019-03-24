package main

// @name 168. Excel表列名称
// @link: https://leetcode-cn.com/problems/excel-sheet-column-title/
// @data 2019年3月24日 星期日

import (
	"fmt"
)

/**
给定一个正整数，返回它在 Excel 表中相对应的列名称。
例如，

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...

示例 1:
	输入: 1
	输出: "A"

示例 2:
	输入: 28
	输出: "AB"

示例 3:
	输入: 701
	输出: "ZY"
*/

// 解答
// 26 进制转10进制
func convertToTitle(n int) string {
	var s string
	const x26  = 26
	for {
		c26 := 0
		if n > x26{
			yushu := n % x26
		}else {
			c26 = n
			n = -1
		}
		if n < 0{
			break
		}
	}
	return s
}

// 测试实例
func main() {
	vfmt := "%t) [%v] VS [%v] <<= [%v]"
	var vin int
	var vout, vneed string
	// test 1
	//vin, vout, vneed = "Hello World", 5, lengthOfLastWord(vin)
	vin, vout = 1, "A"
	vneed = convertToTitle(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))

	// test 2
	vin, vout = 28, "AB"
	vneed = convertToTitle(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))

	// test 3
	vin, vout = 701, "ZY"
	vneed = convertToTitle(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))
}
