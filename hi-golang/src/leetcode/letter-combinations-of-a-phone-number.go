package main

import (
	"fmt"
	"regexp"
	"strings"
)

/**
	@date 2018年10月30日 星期二
	@description 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
	@link https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
	@name letter-combinations-of-a-phone-number (lcpn)
 */



func main() {
	s := ""

	s = "5 "
	fmt.Printf("输入： %s, 输出：%s\n", s, letterCombinations(s))

	s = "2 3"
	s = "2 3"
	fmt.Printf("输入： %s, 输出：%s\n", s, letterCombinations(s))

	s = "2345"
	fmt.Printf("输入： %s, 输出：%s\n", s, letterCombinations(s))
}



// 二分法：每次只合并两排
func letterCombinations(digits string) []string {
	// 去掉空字符
	reg := regexp.MustCompile("\\s")
	digits = reg.ReplaceAllString(digits, "")

	dick := map[string]string{
		"1": "",
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
		"0": "",
	}
	queue := strings.Split(digits, "")
	vLen := len(queue)

	value := []string{}

	getValFn := func(s string) []string {
		if s1Tmp, has := dick[s]; has{
			s = s1Tmp
		}else{
			s = ""
		}
		return strings.Split(s, "")
	}
	c1 := ""
	for i := 0; i<vLen && c1 == ""; i++{
		c1 = queue[i]
		if vLen == 1{
			value = getValFn(c1)
			break
		}
		for j := i+1; j<vLen; j++{
			if j == 1{
				for _, s1 := range getValFn(c1) {
					for _, s2 := range getValFn(queue[j]) {
						// 两两合并
						value = append(value, s1+s2)
					}
				}
			}else{
				newQue := []string{}
				for _, cc1 := range value{
					for _, s2 := range getValFn(queue[j]) {
						// 两两合并
						newQue = append(newQue, cc1+s2)
					}
				}
				value = newQue
			}
		}
	}
	return value
}