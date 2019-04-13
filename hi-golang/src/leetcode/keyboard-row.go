package main

import (
	"fmt"
	"strings"
)

/**
 * @DATE        2019/4/13
 * @NAME        Joshua Conero
 * @DESCRIPIT   500. 键盘行
 * @LINK		https://leetcode-cn.com/problems/keyboard-row/
 **/

// @TODO 完善题目信息，算法已解决

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
		isBreak := false
		for _, w := range strings.Split(s, ""){
			w = strings.ToUpper(w)
			for i, comW := range dick{
				if strings.Index(comW, w) > -1{
					if idx == -1{
						idx = i
					}else if i != idx{
						isBreak = true
						break
					}
				}
			}
			if isBreak {
				break
			}
		}
		if isBreak {
			continue
		}
		wordValue = append(wordValue, s)
	}
	return wordValue

}


// 控制台
func main() {
	fmt.Println(findWords([]string{"Hello", "Alaska", "Dad", "Peace"}))
}


