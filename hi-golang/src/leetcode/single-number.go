package main

import "fmt"

// @name 136. 只出现一次的数字
// @link: https://leetcode-cn.com/problems/single-number/
// @data 2019年4月1日 星期一

// 题目
/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
	你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？

示例 1:
	输入: [2,2,1]
	输出: 1

示例 2:
	输入: [4,1,2,1,2]
	输出: 4
*/

// 解答
func singleNumber(nums []int) int {
	dick := map[int]int{}
	for _, v := range nums {
		if _, has := dick[v]; !has {
			dick[v] = 1
		} else {
			dick[v] += 1
		}
	}
	for n, c := range dick {
		if c == 1 {
			return n
		}
	}
	return 0
}

// 控制台
func main() {
	vfmt := "%t) [%v] VS [%v] <<= [%v]"
	var vin []int
	var vout, vneed int

	// test 1
	vin, vout = []int{2, 2, 1}, 1
	vneed = singleNumber(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))

	// test 2
	vin, vout = []int{4, 1, 2, 1, 2}, 4
	vneed = singleNumber(vin)
	fmt.Println(fmt.Sprintf(vfmt, vneed == vout, vout, vneed, vin))
}
