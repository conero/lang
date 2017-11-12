/**
 * 排序测试
 * 2017年11月12日 星期日
 */
package main

import (
	"../alg"
	"fmt"
)
// 测试用例
type tCase struct {
}

// 运行台
func (t tCase)console()  {
	t.InsertSort()
}

// 插入排序测试
func (t tCase)InsertSort(){
	sort := alg.Sort{}
	baseArray := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("基本数组：", baseArray, ", 长度:", len(baseArray))
	fmt.Println("插入数字3=> ", sort.InertSort(baseArray, 3), sort.Ergodic )
	fmt.Println("插入数字0=> ", sort.InertSort([]int{1,2,3,4,5,6,7,8,9,10}, 0), sort.Ergodic )
	fmt.Println("插入数字300=> ", sort.InertSort([]int{1,2,3,4,5,6,7,8,9,10}, 300), sort.Ergodic )

	// baseArray 的值已经被改变
	fmt.Println("插入数字300=> ", sort.InertSort(baseArray, 300), sort.Ergodic, "， 长度：", len(baseArray))
	fmt.Println("插入数字300=> ", sort.InertSort(baseArray, 300, 4, 7, 45, 60, 14, 20), sort.Ergodic )

	// 空值排序
	fmt.Println("基准数组为空=> ", sort.InertSort([]int{}, 300, 4, 7, 45, 60, 14, 20, 47, 12, 74, 56, 104, 104, 27, 47444), sort.Ergodic)
	rsArr := []int{}
	fmt.Println("基准数组为空=> ", sort.InertSort(rsArr, 300, 4, 7, 45, 60, 14, 20, 47, 12, 74, 56, 104, 104, 27, 47444), sort.Ergodic)
	fmt.Println("空基准数组变化为 =>", rsArr)
	fmt.Println("空基准数组变化为 =>", sort.Inert([]int{300, 4, 7, 45, 60, 14, 20, 47, 12, 74, 56, 104, 104, 27, 47444}))
	}

func main() {
	(tCase{}).console()
}
