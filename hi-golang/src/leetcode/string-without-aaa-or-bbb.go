package main

import (
	"fmt"
	"strings"
)

/**
   @date 2018年10月26日 星期五
   @description 不含 AAA 或 BBB 的字符串
   @link https://leetcode-cn.com/problems/string-without-aaa-or-bbb/
   @name 不含 AAA 或 BBB 的字符串


	给定两个整数 A 和 B，返回任意字符串 S，要求满足：
		S 的长度为 A + B，且正好包含 A 个 'a' 字母与 B 个 'b' 字母；
		子串 'aaa' 没有出现在 S 中；
		子串 'bbb' 没有出现在 S 中。


	示例 1：
		输入：A = 1, B = 2
		输出："abb"
		解释："abb", "bab" 和 "bba" 都是正确答案。

	示例 2：
		输入：A = 4, B = 1
		输出："aabaa"


	提示：
		0 <= A <= 100
		0 <= B <= 100
		对于给定的 A 和 B，保证存在满足要求的 S。
*/
// @TODO 未解决 - NeedToDoStill.
// 解题
func strWithout3a3b(A int, B int) string {
	sArr := []string{}
	for A > 0 || B > 0 {
		firstA := A > B
		sArrLen := len(sArr)
		if sArrLen >= 2 && sArr[sArrLen-1] == sArr[sArrLen-2] {
			firstA = sArr[sArrLen-1] == "b"
		}
		if firstA {
			sArr = append(sArr, "a")
			A -= 1
		} else {
			sArr = append(sArr, "b")
			B -= 1
		}
	}
	return strings.Join(sArr, "")
}

// d1
func strWithout3a3b_draft1(A int, B int) string {
	var s string = ""
	//ab := A + B
	for {
		// a
		if A < 3 {
			s += strings.Repeat("a", A)
			A = 0
		} else {
			A -= 2
			s += strings.Repeat("a", 2)
		}
		// b
		if B < 3 {
			s += strings.Repeat("b", B)
			B = 0
		} else {
			B -= 2
			s += strings.Repeat("b", 2)
		}

		if A == 0 && B == 0 {
			break
		}
	}
	return s
}

// d2
func strWithout3a3b_draft2(A int, B int) string {
	var s string = ""
	//ab := A + B
	// 字典
	dick := map[string]string{
		"A": "a",
		"B": "b",
	}
	var max, min int
	var maxK, minK string
	if A-B > 0 {
		max, min = A, B
		maxK, minK = "A", "B"
	} else {
		max, min = B, A
		maxK, minK = "B", "A"
	}
	//fmt.Println([]interface{}{max, min, maxK, minK})
	//fmt.Println(dick)
	for {
		maxCt, minCt := max-2, min-2
		max -= 2
		min -= 2
		if maxCt > 0 {
			s += strings.Repeat(dick[maxK], maxCt)
		}
		if minCt > 0 {
			s += strings.Repeat(dick[minK], minCt)
		}
		if max <= 0 && min <= 0 {
			break
		}
	}
	return s
}

// 运行
func main() {
	fmt.Println(strWithout3a3b(5, 7))
	fmt.Println(strWithout3a3b(1, 2))
	fmt.Println(strWithout3a3b(4, 1))
	fmt.Println(strWithout3a3b(1, 3)) //bbab
}