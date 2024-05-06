package main

// @name 88. 合并两个有序数组
// @link: https://leetcode-cn.com/problems/merge-sorted-array/
// @data 2019年3月26日 星期二

import "fmt"

// 题目
/*
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
说明:
	初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
	你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

示例:
	输入:
		nums1 = [1,2,3,0,0,0], m = 3
		nums2 = [2,5,6],       n = 3

	输出: [1,2,2,3,5,6]
*/

// @TODO StillNeedToDoYet
// 解答
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	mv := []int{}
	for i := 1; ; i++ {
		if i > m && i > n {
			break
		}
		hasM, hasN := false, false
		if i <= m {
			hasM = true
		}
		if i <= n {
			hasN = true
		}
		insertM := false
		for j, mTmp := range mv {
			if !hasM {
				break
			}
			if nums1[i-1] < mTmp {
				mv[j] = nums1[i-1]
				lastMv := mv[j+1:]
				lastMv = append([]int{mTmp}, lastMv...)
				mv = append(mv[0:i], lastMv...)
				insertM = true
				break
			}
		}
		if !insertM && hasM {
			mv = append(mv, nums1[i-1])
		}
		//----------------------[ N ]
		insertN := false
		for j, mTmp := range mv {
			if !hasN {
				break
			}
			if nums2[i-1] < mTmp {
				mv[j] = nums2[i-1]
				lastMv := mv[j+1:]
				lastMv = append([]int{mTmp}, lastMv...)
				mv = append(mv[0:i], lastMv...)
				insertN = true
				break
			}
		}
		if !insertN && hasN {
			mv = append(mv, nums2[i-1])
		}
	}
	return mv
}

func main() {
	nums1, nums2 := []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}
	m, n := 3, 3
	fmt.Println([]int{1, 2, 2, 3, 5, 6}, merge(nums1, m, nums2, n))
}
