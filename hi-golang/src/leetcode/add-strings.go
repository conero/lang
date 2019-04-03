package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 解答
// @TODO StillNeedToDo
/*
	给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
	注意：
		num1 和num2 的长度都小于 5100.
		num1 和num2 都只包含数字 0-9.
		num1 和num2 都不包含任何前导零。
		你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
*/
func addStrings(num1 string, num2 string) string {
	ns := []string{}
	n1Len, n2Len := len(num1), len(num2)
	max := n1Len
	if n2Len > max {
		max = n2Len
	}

	last := 0
	for x := max; x > 0; x-- {
		x1, x2 := 0, 0
		if x < n1Len{
			x1, _ = strconv.Atoi(num1[x-1: x])
		}
		if x < n2Len{
			x2, _ = strconv.Atoi(num2[x-1: x])
		}
		x1 = x1 + x2 + last
		if x1 > 9{
			last = x1 - 10
		}else {
			last = 0
		}
		ns = append(ns, strconv.Itoa(x1))
	}

	return strings.Join(ns, "")
}

// 控制台
func main() {

}
