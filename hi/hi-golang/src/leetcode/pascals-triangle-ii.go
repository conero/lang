package main

import "fmt"

/**
@date 2019年2月18日 星期一
@description 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
@link https://leetcode-cn.com/problems/pascals-triangle-ii/
@name pascals-triangle-ii  杨辉三角 II


示例:

	输入: 3
	输出: [1,3,3,1]
进阶：
	你可以优化你的算法到 O(k) 空间复杂度吗？

*/

func getRow(rowIndex int) []int {
	d := []int{}
	var quque [][]int = [][]int{}
	rowIndex += 1
	for i := 1; i <= rowIndex; i++ {
		if len(quque) == 0 {
			quque = append(quque, []int{1})
			continue
		} else {
			quque = append(quque, []int{})
		}
		j := i - 2
		oLine := quque[j]
		vlen := len(oLine)
		line := []int{}
		for m := 0; m < vlen-1; m++ {
			l, r := oLine[m], oLine[m+1]
			line = append(line, l+r)
		}
		line = append([]int{1}, line...)
		line = append(line, 1)
		quque[i-1] = line
	}
	if len(quque) > 0 {
		d = quque[len(quque)-1]
	}
	return d
}

func main() {
	// 3 => [1,3,3,1]
	fmt.Println(getRow(3))
}
