/**
 * 2017年6月28日 星期三
 * Joshua Conero
 * 数据生成器
 */
package data

import (
	"math/rand"
	"time"
)

// 随机整数形
func GetRandInt() int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return rd.Int()
}

/*
// 获取数据字母：A-Z: 65-90; 97 - 122
func GetRandLetter() string {
	num := GetRandInt()
	nLen := len(num)
	i := 0
	letter := ""
	for i < nLen-1 {
		firstNum := nLen[i]
		if firstNum > 1 {
			letter += string(rune(nLen[i : i+1]))
			i = i + 1
		}
	}
	return letter
}
*/
