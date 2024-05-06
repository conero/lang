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

/*
给定一个单词列表，只返回可以使用在键盘同一行的字母打印出来的单词。键盘如下图所示。
American keyboard

示例：
	输入: ["Hello", "Alaska", "Dad", "Peace"]
	输出: ["Alaska", "Dad"]

注意：
	你可以重复使用键盘上同一字符。
	你可以假设输入的字符串将只包含字母。
*/

// 解答
func findWords(words []string) []string {
	dick := []string{
		"QWERTYUIOP",
		"ASDFGHJKL",
		"ZXCVBNM",
	}
	wordValue := []string{}
	for _, s := range words {
		var idx int = -1
		isBreak := false
		for _, w := range strings.Split(s, "") {
			w = strings.ToUpper(w)
			for i, comW := range dick {
				if strings.Index(comW, w) > -1 {
					if idx == -1 {
						idx = i
					} else if i != idx {
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
