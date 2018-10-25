package main

import (
	"fmt"
	"strings"
)

/**
	@date 2018年10月25日 星期四
	@description 给定一个字符串，找出不含有重复字符的最长子串的长度。
	@link https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
	@name lengthOfLongestSubstring (lols)
 */




/**
		示例 1:

		输入: "abcabcbb"
		输出: 3
		解释: 无重复字符的最长子串是 "abc"，其长度为 3。
		示例 2:

		输入: "bbbbb"
		输出: 1
		解释: 无重复字符的最长子串是 "b"，其长度为 1。
		示例 3:

		输入: "pwwkew"
		输出: 3
		解释: 无重复字符的最长子串是 "wke"，其长度为 3。
			 请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串。
 */
func case1()  {

	var testFn = func(s string) {
		var max int
		var sub string
		sub, max = LolsA(s)
		fmt.Println(s, sub, max)
	}
	var testFnArr = func(ss []string) {
		for _, s := range ss{
			testFn(s)
		}
	}


	testFnArr([]string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"a",
		"ab",
		"joshua",
		" ",
		"uqinntq",
	})
}
func main() {
	case1()
}

/**
	Lols algorithm
	暴力遍历法
 */
func LolsA(s string) (string, int)  {
	// 收尾指针(指向字符串的为)
	//s = strings.TrimSpace(s)
	pe := len(s)
	sLen := pe 					// 长度
	maxLen := 0			// 最大长度
	var maxLastStr string		// 最大的最后遍历的字符串

	// 遍历值
	for ps := 0;  ps < sLen; ps ++ {
		tmpC := s[ps:ps+1]
		tmpIdx := strings.Index(s[ps+1:], tmpC)
		// 求出最大区间段
		switch tmpIdx {
		case -1:
			pe = sLen - 1
		default:
			pe = tmpIdx + ps
		}
		//fmt.Println(ps, pe, "<=")
		// 子字符串处理
		for j := ps + 1; j < pe; j++ {
			tmpS2 := s[j: j +1]
			k := j + 1
			//for ; k < pe && tmpS2 != s[k: k +1]; k++{
			for ; k < pe && tmpS2 != s[k: k +1]; k++{
				//pe = k - 1
			}
			//fmt.Println(tmpC, tmpS2, k, "5", s[k-1: k])
			if tmpS2 == s[k: k +1]{
				//if tmpS2 == s[k-1: k]{
				pe = k - 1
				//break
			}else {
				//fmt.Println(tmpC, tmpS2, k, "5", s[k-1: k])
				//pe = k
			}

			//fmt.Println(tmpC, tmpS2, k, "5", s[k-1: k])
			//if tmpS2 == s[k: k+1]{
			//	pe = k - 1
			//}
			//if k == sLen{
			//	k -= 1
			//}
			//if tmpS2 == s[k: k +1]{
			//	pe = k
			//}
			//fmt.Println(tmpS2, s[k: k +1], ": s2")
			//pe = k + 1
		}
		// !!
		// 边界检测
		//if pe < sLen - 1{
		//	pe += 1
		//}

		pe += 1
		tmpS := s[ps: pe]
		//if ps == pe{
		//	if ps > 0{
		//		tmpS = s[ps - 1: pe]
		//	}
		//}

		// 去中间空字符
		//reg := regexp.MustCompile("\\s")
		//tmpS = reg.ReplaceAllString(tmpS, "")

		tmpLen := len(tmpS)
		if tmpLen >= maxLen{
			maxLen = tmpLen
			maxLastStr = tmpS
		}
	}

	return maxLastStr, maxLen
}

/**
	标准函数
 */
func lengthOfLongestSubstring(s string) int {
	_, max := LolsA(s)
	return max
}


