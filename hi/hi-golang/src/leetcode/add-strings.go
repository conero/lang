package main

// @name 415. 字符串相加
// @link: https://leetcode-cn.com/problems/add-strings/
// @data 2019年4月3日

import (
	"fmt"
	"strconv"
	"strings"
)

// 解答
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
	for x := 1; x <= max; x++ {
		x1, x2 := 0, 0
		tmpInt := 0
		if x <= n1Len {
			tmpInt = n1Len - x
			x1, _ = strconv.Atoi(num1[tmpInt : tmpInt+1])
		}
		if x <= n2Len {
			tmpInt = n2Len - x
			x2, _ = strconv.Atoi(num2[tmpInt : tmpInt+1])
		}
		//fmt.Println(x, x1, x2, last, x1+x2+last)
		x1 = x1 + x2 + last
		bitN := x1
		if x1 > 9 {
			bitN = x1 - 10
			last = (x1 - bitN) / 10
		} else {
			last = 0
		}
		ns = append(ns, strconv.Itoa(bitN))
	}
	if last > 0 {
		ns = append(ns, strconv.Itoa(last))
	}
	ns1 := []string{}
	for j := len(ns) - 1; j > -1; j-- {
		ns1 = append(ns1, ns[j])
	}
	return strings.Join(ns1, "")
}

// 控制台
func main() {
	vfmt := "%t) [%v] VS [%v] <<= [%v, %v]"
	var vinA, vinB, vout, vneed string

	// test 1
	// 4-6-9
	vinA, vinB, vout = "9999", "999999", "1009998"
	vneed = addStrings(vinA, vinB)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vinA, vinB))
}
