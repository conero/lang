package main

import (
	"fmt"
	"strings"
)

/*
	@name 191. 位1的个数
	@Link https://leetcode-cn.com/problems/number-of-1-bits/
	@date 2019年3月15日 星期五
*/

//题目
/*
	编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）。

	示例 1：
		输入：00000000000000000000000000001011
		输出：3
	解释：输入的二进制串 00000000000000000000000000001011 中，共有三位为 '1'。

	示例 2：
		输入：00000000000000000000000010000000
		输出：1
	解释：输入的二进制串 00000000000000000000000010000000 中，共有一位为 '1'。

	示例 3：
		输入：11111111111111111111111111111101
		输出：31
	解释：输入的二进制串 11111111111111111111111111111101 中，共有 31 位为 '1'。

	提示：
		请注意，在某些语言（如 Java）中，没有无符号整数类型。在这种情况下，输入和输出都将被指定为有符号整数类型，并且不应影响您的实现，因为无论整数是有符号的还是无符号的，其内部的二进制表示形式都是相同的。
		在 Java 中，编译器使用二进制补码记法来表示有符号整数。因此，在上面的 示例 3 中，输入表示有符号整数 -3。

	进阶:
		如果多次调用这个函数，你将如何优化你的算法？
*/

// 编译通过，但是感觉做法有点凑巧
// @TODO 可以继续优化
// 解答
func hammingWeight(num uint32) int {
	bits := 0
	s := fmt.Sprintf("%b", num)
	for _, sn := range strings.Split(s, "") {
		if sn == "1" {
			bits += 1
		}
	}
	//fmt.Println(s, num)
	return bits
}

// 控制台
func main() {
	var num uint32
	var out, vsout int
	var fmtStr = "%t)  [%d] => [%v] VS [%v]\n"

	// 测试用例
	num = 00000000000000000000000000001011
	out, vsout = hammingWeight(num), 3
	fmt.Printf(fmtStr, out == vsout, num, out, vsout)

	//为什么此处为的二进制打印为 521
	//fmt.Println(521 == num, 00000000000000000000000010000000)

	// case
	num = 00000000000000000000000010000000
	out, vsout = hammingWeight(num), 1
	fmt.Printf(fmtStr, out == vsout, num, out, vsout)

	//Error： constant 11111111111111111111111111111101 overflows uint32
	// case
	//num = 11111111111111111111111111111101
	//out, vsout = hammingWeight(num), 31
	//fmt.Printf(fmtStr, out == vsout, num, out, vsout)
}
