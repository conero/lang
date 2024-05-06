package main

// @name 824. 山羊拉丁文
// @link: https://leecode-cn.com/problems/goat-latin/
// @data 2019年3月8日 星期五

import "fmt"

/**
@date 2019年2月18日 星期一
@description 给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
@link https://leetcode-cn.com/problems/pascals-triangle/
@name pascals-triangle


示例:
输入: 5
输出:
[
	 [1],
	[1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

*/

/*
todos 例子
	@TODO 实现失败：  超出时间限制
	@TODO 未解决 - NeedToDoStill.
*/

// 解答
func generateXXXX(numRows int) [][]int {
	var quque [][]int = [][]int{}
	for i := 1; i <= numRows; i++ {
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
	return quque
}

// 测试实例
func main() {
	fmt.Println(generateXXXX(5))
}
