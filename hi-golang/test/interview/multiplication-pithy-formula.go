// 2019年2月17日 星期日
// 乘法口诀生成
// 2019.2 PHP工程师面试题

package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 9
	s := ""
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			s += strconv.Itoa(i) + "×" + strconv.Itoa(j) + "=" + strconv.Itoa(i*j) + "\t"
		}
		s += "\n"
	}
	fmt.Println(s)
}
