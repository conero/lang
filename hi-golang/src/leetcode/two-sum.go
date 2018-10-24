package main

import "fmt"

func twoSum(nums []int, target int) []int {
	vlen := len(nums)
	var value []int
	for i, n := range nums{
		success := false
		for j := i +1; j<vlen; j++{
			if (target - n) == nums[j]{
				success = true
				value = []int{i, j}
				break
			}
		}
		if success{
			break
		}
	}
	return value
}

// 示例:
//
//	给定 nums = [2, 7, 11, 15], target = 9
//	因为 nums[0] + nums[1] = 2 + 7 = 9
//	所以返回 [0, 1]
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))
}
