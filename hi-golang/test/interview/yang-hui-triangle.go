// 2019年2月17日 星期日
// 杨辉三角生成
// 2019.2 PHP工程师面试题

package main

import "fmt"

func YhTriangle(n int) string {
	var s string

	var queque [][]int = [][]int{}
	for i := 1; i <= n; i++ {
		if len(queque) == 0 {
			queque = append(queque, []int{1})
			continue
		} else {
			queque = append(queque, []int{})
		}
		j := i - 2
		oLine := queque[j]
		fmt.Println(oLine, i, j, queque)
		vlen := len(oLine)
		line := []int{}
		for m := 0; m < vlen; m++ {
			l, r := oLine[m], 0
			if m < vlen-1 {
				r = oLine[m+1]
			}
			line = append(line, l+r)
		}
		line = append([]int{1}, line...)
		line = append(line, 1)
		queque[j] = line
	}

	fmt.Println(queque)
	return s
}

func main() {
	fmt.Println(YhTriangle(7))
}
