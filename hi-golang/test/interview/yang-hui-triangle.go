// 2019年2月17日 星期日
// 杨辉三角生成
// 2019.2 PHP工程师面试题
package main

import "fmt"

func YhTriangle(n int) [][]int {
	var quque [][]int = [][]int{}
	for i := 1; i <= n; i++ {
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

func main() {
	fmt.Println(YhTriangle(7))
	fmt.Println(YhTriangle(10))
}
