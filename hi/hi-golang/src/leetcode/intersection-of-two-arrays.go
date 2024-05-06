package main

import "fmt"

/**
 * @DATE        2019/5/13
 * @NAME        Joshua Conero
 * @DESCRIPIT   349. 两个数组的交集
 * @LINK 		https://leetcode-cn.com/problems/intersection-of-two-arrays/
 **/

// 题目
/*
给定两个数组，编写一个函数来计算它们的交集。

示例 1:
	输入: nums1 = [1,2,2,1], nums2 = [2,2]
	输出: [2]

示例 2:
	输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
	输出: [9,4]

说明:
	输出结果中的每个元素一定是唯一的。
	我们可以不考虑输出结果的顺序。
*/

// 解题
func intersection(nums1 []int, nums2 []int) []int {
	var data []int = []int{}
	var min, max []int
	if len(nums1) > len(nums2) {
		min = nums2
		max = nums1
	} else {
		min = nums1
		max = nums2
	}
	var dickRaw map[int]int = map[int]int{}  // 原料字典
	var dickCurr map[int]int = map[int]int{} // 当前数据字典
	for _, v := range max {
		dickRaw[v] = 1
	}
	// 遍历
	for _, v := range min {
		if _, has := dickRaw[v]; has {
			if _, has := dickCurr[v]; !has {
				dickCurr[v] = 1
				data = append(data, v)
			}

		}
	}
	return data
}

// 控制台
func main() {
	const format = "{%v, %v} => {%v} VS {%v}\r\n"
	var nums1, nums2, vout, vneed []int

	// case
	nums1, nums2, vneed = []int{1, 2, 2, 1}, []int{2, 2}, []int{2}
	vout = intersection(nums1, nums2)
	fmt.Printf(format, nums1, nums2, vout, vneed)

	// case
	nums1, nums2, vneed = []int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}
	vout = intersection(nums1, nums2)
	fmt.Printf(format, nums1, nums2, vout, vneed)
}
