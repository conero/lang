package main

import (
	"fmt"
	"math"
)

/**
 * @DATE        2019/4/18
 * @NAME        Joshua Conero
 * @DESCRIPIT   268. 缺失数字
 * @LINK   		https://leetcode-cn.com/problems/missing-number/
 **/

// 解答
func missingNumber(nums []int) int {
	var value int
	var max int
	var hasValueMk bool = false
	var numDick map[int]int = map[int]int{}

	// 清楚最大值
	for _, v := range nums {
		max = int(math.Max(float64(v), float64(max)))
		numDick[v] = 1
	}
	// 最大值查看
	//fmt.Println(numDick[2], numDick)
	for i := 0; i <= max; i++ {
		if _, has := numDick[i]; !has {
			value = i
			hasValueMk = true
			break
		}
	}
	if !hasValueMk {
		value = max + 1
	}
	return value
}

// 控制台
func main() {
	var vin []int
	var vout, vneed int
	const format = "%v) [%v] => [%v] VS [%v]\n"

	vin = []int{3, 0, 1}
	vout, vneed = missingNumber(vin), 2
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

	vin = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	vout, vneed = missingNumber(vin), 8
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

	vin = []int{0}
	vout, vneed = missingNumber(vin), 1
	fmt.Printf(format, vout == vneed, vin, vout, vneed)

	vin = []int{0, 1}
	vout, vneed = missingNumber(vin), 2
	fmt.Printf(format, vout == vneed, vin, vout, vneed)
}
