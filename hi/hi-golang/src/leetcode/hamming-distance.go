package main

import (
	"fmt"
	"math"
)

// @name 461. 汉明距离
// @link https://leetcode-cn.com/problems/hamming-distance/
// @date 2019年3月15日 星期五

// 题目
/*
	两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
	给出两个整数 x 和 y，计算它们之间的汉明距离。

	注意：
		0 ≤ x, y < 231.

	示例:
		输入: x = 1, y = 4
		输出: 2

	解释:
	1   (0 0 0 1)
	4   (0 1 0 0)
		   ↑   ↑

	上面的箭头指出了对应二进制位不同的位置。
*/

// @TODO 解答时遇到问题
// 解
func hammingDistance(x int, y int) int {
	xy := 0
	// 求出 BQ
	getBQ := func(n int) []int {
		nn := []int{}
		for {
			//fmt.Println(n)
			if n > 1 && n%2 == 0 {
				nn = append(nn, 0)
				n = n / 2
			} else if n > 1 {
				nn = append(nn, 1)
				n = (n - 1) / 2
			}
			if n == 1 {
				nn = append(nn, 1)
				break
			} else if n < 1 {
				nn = append(nn, 0)
				break
			}
		}
		return nn
	}
	//fmt.Println(getBQ(1), getBQ(4), getBQ(5))

	xbq, ybq := getBQ(x), getBQ(y)
	xct, yct := len(xbq), len(ybq)
	max := int(math.Max(float64(xct), float64(yct)))
	fmt.Println(xbq, ybq, xct, yct, max)
	// 传入参数与常量的结果不一致
	fmt.Println(x, y, getBQ(x), getBQ(y), getBQ(1), getBQ(4))

	// 比较
	for i := 1; i <= max; i++ {
		xb, yb := 0, 0
		if i <= xct {
			xb = xbq[xct-i]
		}
		if i <= yct {
			yb = ybq[yct-i]
		}
		if xb != yb {
			xy += 1
		}
	}

	return xy
}

// 控制台
func main() {
	var sfmt string = "%t) [%v, %v] => [%v] VS [%v]\n"
	var x, y, out, vsout int

	// case
	x, y, out, vsout = 1, 4, hammingDistance(x, y), 2
	fmt.Printf(sfmt, out == vsout, x, y, out, vsout)
}
