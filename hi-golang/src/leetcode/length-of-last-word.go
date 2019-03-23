package main

// @name 58. 最后一个单词的长度
// @link: https://leetcode-cn.com/problems/length-of-last-word/
// @data 2019年3月23日 星期六

import (
	"fmt"
	"strings"
)

/**

给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
如果不存在最后一个单词，请返回 0 。
说明：一个单词是指由字母组成，但不包含任何空格的字符串。

示例:
	输入: "Hello World"
	输出: 5

*/


// 解答
func lengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	vlen, slen := 0, len(s)
	if slen != 0 && " " == s[slen-1:] {
		fmt.Println(slen, s)
		return vlen
	}
	ss := strings.Split(s, " ")
	lastLen := len(ss)
	if lastLen > 0 {
		vlen = len(ss[lastLen-1])
	}
	return vlen
}

// 测试实例
func main() {
	vfmt := "%t) [%v] VS [%v] <<= [%s]"
	var vin string
	var vout, vneed int
	// test 1
	//vin, vout, vneed = "Hello World", 5, lengthOfLastWord(vin)
	vin, vout = "Hello World", 5
	vneed = lengthOfLastWord(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))

	// test 2
	vin, vout = "a ", 1
	vneed = lengthOfLastWord(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))
}
