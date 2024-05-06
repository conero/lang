package main

// @name 171. Excel表列序号
// @link: https://leetcode-cn.com/problems/excel-sheet-column-number/
// @data 2019年3月23日 星期六

import (
	"fmt"
	"math"
)

/**
给定一个Excel表格中的列名称，返回其相应的列序号。
例如，
    A -> 1
    B -> 2
    C -> 3
    ...
    Z -> 26
    AA -> 27
    AB -> 28
    ...

示例 1:
	输入: "A"
	输出: 1

示例 2:
	输入: "AB"
	输出: 28

示例 3:
	输入: "ZY"
	输出: 701
*/

// 解答
func titleToNumber(s string) int {
	// 65 - 90
	// A-Z
	chars := []byte(s)
	dec := 0
	charsLen := len(chars)
	for i := charsLen - 1; i >= 0; i-- {
		c := chars[i]
		cn := int(c)
		dec += (cn - 65 + 1) * int(math.Pow(float64(26), float64(charsLen-i-1)))
	}
	//fmt.Println(chars)
	return dec
}

// 测试实例
func main() {
	vfmt := "%t) [%v] VS [%v] <<= [%s]"
	var vin string
	var vout, vneed int
	// test 1
	//vin, vout, vneed = "Hello World", 5, lengthOfLastWord(vin)
	vin, vout = "A", 1
	vneed = titleToNumber(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))

	// test 2
	vin, vout = "AB", 28
	vneed = titleToNumber(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))

	// test 3
	vin, vout = "ZY", 701
	vneed = titleToNumber(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))
}
