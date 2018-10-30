package main

import "strings"

/**
	@date 2018年10月30日 星期二
	@description 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
	@link https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/
	@name letter-combinations-of-a-phone-number (lcpn)
 */

func main() {
	
}




func letterCombinations(digits string) []string {
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
	for i:=0; i<vLen ; i++{
		s1 := value[i]
		if s1Tmp, has := dick[s1]; has{
			s1 = s1Tmp
		}
	}
	return value
}