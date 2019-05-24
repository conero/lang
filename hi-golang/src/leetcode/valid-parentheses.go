package main

import (
	"fmt"
	"regexp"
)

/**
 * @DATE        2019/5/24
 * @NAME        Joshua Conero
 * @DESCRIPIT   20. 有效的括号
 * @LINK 		https://leetcode-cn.com/problems/valid-parentheses/
 **/

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
	输入: "()"
	输出: true

示例 2:
	输入: "()[]{}"
	输出: true

示例 3:
	输入: "(]"
	输出: false

示例 4:
	输入: "([)]"
	输出: false

示例 5:
	输入: "{[]}"
	输出: true
*/

// 解答
// 规则检测法则
func isValid(s string) bool {
	isV := false
	/*sign1 := regexp.MustCompile(`\([^\(\)]*\)`)
	sign2 := regexp.MustCompile(`\{[^\{\}]*\}`)
	sign3 := regexp.MustCompile(`\[[^\[\]]*\]`)*/

	sign1 := regexp.MustCompile(`\(\)`)
	sign2 := regexp.MustCompile(`\{\}`)
	sign3 := regexp.MustCompile(`\[\]`)

	for {
		oneOf := false
		if sign1.MatchString(s) {
			oneOf = true
			s = sign1.ReplaceAllString(s, "")
		}

		if sign2.MatchString(s) {
			oneOf = true
			s = sign2.ReplaceAllString(s, "")
		}
		//fmt.Println(sign3.MatchString(s), oneOf, s)
		if sign3.MatchString(s) {
			oneOf = true
			s = sign3.ReplaceAllString(s, "")
		}
		//fmt.Println(sign3.MatchString(s), oneOf, s)
		if s == "" {
			isV = true
			break
		}
		if !oneOf {
			break
		}
	}
	return isV
}

// 控制台
func main() {
	const format = "%v)  %v => %v  VS %v\r\n"
	var vin string
	var vout, want bool

	// case
	vin, want = "()", true
	vout = isValid(vin)
	fmt.Printf(format, vout == want, vin, vout, want)

	// case
	vin, want = "()[]{}", true
	vout = isValid(vin)
	fmt.Printf(format, vout == want, vin, vout, want)

	// case
	vin, want = "(]", false
	vout = isValid(vin)
	fmt.Printf(format, vout == want, vin, vout, want)

	// case
	vin, want = "([)]", false
	vout = isValid(vin)
	fmt.Printf(format, vout == want, vin, vout, want)

	// case
	vin, want = "{[]}", true
	vout = isValid(vin)
	fmt.Printf(format, vout == want, vin, vout, want)
}
