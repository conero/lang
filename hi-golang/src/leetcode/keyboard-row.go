package main

import "strings"

/**
 * @DATE        2019/4/13
 * @NAME        Joshua Conero
 * @DESCRIPIT   500. 键盘行
 * @LINK		https://leetcode-cn.com/problems/keyboard-row/
 **/

// 解答
func findWords(words []string) []string {
	dick := []string{
		"QWERTYUIOP",
		"ASDFGHJKL",
		"ZXCVBNM",
	}
	wordValue := []string{}
	for _, s := range words{
		var idx int = -1
		for _, w := range strings.Split(s, ""){
			w = strings.ToUpper(w)

		}
	}
	return wordValue

}


// 控制台
func main() {
}


