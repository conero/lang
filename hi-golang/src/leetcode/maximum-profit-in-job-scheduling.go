package main

import (
	"fmt"
)

/**
 * @DATE        2019年10月24日 星期四
 * @NAME        Joshua Conero
 * @DESCRIPIT   1235. 规划兼职工作
 * @LINK 		https://leetcode-cn.com/problems/maximum-profit-in-job-scheduling/
 **/

// @TODO StillNeedToDoThat
// 题目

/*
你打算利用空闲时间来做兼职工作赚些零花钱。
这里有 n 份兼职工作，每份工作预计从 startTime[i] 开始到 endTime[i] 结束，报酬为 profit[i]。
给你一份兼职工作表，包含开始时间 startTime，结束时间 endTime 和预计报酬 profit 三个数组，请你计算并返回可以获得的最大报酬。
注意，时间上出现重叠的 2 份工作不能同时进行。
如果你选择的工作在时间 X 结束，那么你可以立刻进行在时间 X 开始的下一份工作。

示例 1：
	输入：startTime = [1,2,3,3], endTime = [3,4,5,6], profit = [50,10,40,70]
	输出：120
	解释：
		我们选出第 1 份和第 4 份工作，
		时间范围是 [1-3]+[3-6]，共获得报酬 120 = 50 + 70。

示例 2：
	输入：startTime = [1,2,3,4,6], endTime = [3,5,10,6,9], profit = [20,20,100,70,60]
	输出：150
	解释：
		我们选择第 1，4，5 份工作。
		共获得报酬 150 = 20 + 70 + 60。

示例 3：
	输入：startTime = [1,1,1], endTime = [2,3,4], profit = [5,6,4]
	输出：6

提示：
	1 <= startTime.length == endTime.length == profit.length <= 5 * 10^4
	1 <= startTime[i] < endTime[i] <= 10^9
	1 <= profit[i] <= 10^4

*/

// 解答
// 规则检测法则
func jobScheduling(startTime []int, endTime []int, profit []int) int {
	var maxProfit int
	// 遍历时间
	vlen := len(startTime)
	for i := 0; i < vlen; i++{
		st, et, pt := startTime[i], endTime[i], profit[i]
		_ = st	// 忽略数字
		tmpProfit := 0
		sEQue := [][]int{}
		for j := i+1; j<vlen; j++{
			st2, et2, pt2 := startTime[j], endTime[j], profit[j]
			if len(sEQue) == 0{
				if st2 < et{
					if j == vlen -1{
						sEQue = [][]int{}
					}
					continue
				}
				sEQue = append(sEQue, []int{st2, et2, pt + pt2})
				continue
			}
			last := sEQue[len(sEQue) -1]
			if st2 < last[1]{
				if j == vlen -1{
					sEQue = [][]int{}
				}
				continue
			}
			sEQue = append(sEQue, []int{st2, et2, last[2] + pt2})
		}
		fmt.Println(sEQue)
		if tmpProfit > maxProfit{
			maxProfit = tmpProfit
		}
	}

	return maxProfit
}

// 控制台
func main() {
	const format = "%v)  {%v, %v, %v} => %v  VS %v\r\n"
	var stime, etime, profit []int
	var vout, want int

	// case
	stime, etime, profit = []int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}
	want = 120
	vout = jobScheduling(stime, etime, profit)
	fmt.Printf(format, vout == want, stime, etime, profit, vout, want)

	// case
	stime, etime, profit = []int{1, 2, 3, 4, 6}, []int{3, 5, 10, 6, 9}, []int{20, 20, 100, 70, 60}
	want = 150
	vout = jobScheduling(stime, etime, profit)
	fmt.Printf(format, vout == want, stime, etime, profit, vout, want)

	// case
	stime, etime, profit = []int{1, 1, 1}, []int{2, 3, 4}, []int{5, 6, 4}
	want = 6
	vout = jobScheduling(stime, etime, profit)
	fmt.Printf(format, vout == want, stime, etime, profit, vout, want)
}
