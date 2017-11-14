/**
 * 排序测试
 * 2017年11月12日 星期日
 */
package main

import (
	"../alg"
	"../data"
	"fmt"
	"time"
)

// 测试用例
type tCase struct {
}

// 运行台
func (t tCase) console() {
	//t.InsertSort()	// 插入排序
	//t.SelectSort()	// 选择排序
	//t.ShellSort()		// 希尔排序
	t.SortVsRuntimes()
}

// 插入排序测试
func (t tCase) InsertSort() {
	sort := alg.Sort{}
	baseArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("基本数组：", baseArray, ", 长度:", len(baseArray))
	fmt.Println("插入数字3=> ", sort.InertSort(baseArray, 3), sort.Ergodic)
	fmt.Println("插入数字0=> ", sort.InertSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0), sort.Ergodic)
	fmt.Println("插入数字300=> ", sort.InertSort([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 300), sort.Ergodic)

	// baseArray 的值已经被改变
	fmt.Println("插入数字300=> ", sort.InertSort(baseArray, 300), sort.Ergodic, "， 长度：", len(baseArray))
	fmt.Println("插入数字300=> ", sort.InertSort(baseArray, 300, 4, 7, 45, 60, 14, 20), sort.Ergodic)

	// 空值排序
	fmt.Println("基准数组为空=> ", sort.InertSort([]int{}, 300, 4, 7, 45, 60, 14, 20, 47, 12, 74, 56, 104, 104, 27, 47444), sort.Ergodic)
	rsArr := []int{}
	fmt.Println("基准数组为空=> ", sort.InertSort(rsArr, 300, 4, 7, 45, 60, 14, 20, 47, 12, 74, 56, 104, 104, 27, 47444), sort.Ergodic)
	fmt.Println("空基准数组变化为 =>", rsArr)
	fmt.Println("空基准数组变化为 =>", sort.Inert([]int{300, 4, 7, 45, 60, 14, 20, 47, 12, 74, 56, 104, 104, 27, 47444}))

}

func (t tCase) SelectSort()  {
	baseArr := data.GetRandIntArray(10000*1, 10000)
	aSort := &alg.Sort{}
	fmt.Println("基准随机数组： ", baseArr)
	t1 := getRunTimes()
	fmt.Println("选择排序：", aSort.SelectSort(baseArr))
	fmt.Println("  ", t1())
	//fmt.Println("基准随机数组： ", baseArr)
	t2 := getRunTimes()
	fmt.Println("插入排序: ", aSort.Inert(baseArr))
	fmt.Println("  ", t2())
}

// 希尔排序
func (t tCase) ShellSort(){
	baseArr := data.GetRandIntArray(10)
	aSort := &alg.Sort{}
	fmt.Println("基准随机数组： ", baseArr)
	fmt.Println("希尔排序处理：", aSort.ShellSort(baseArr))
	fmt.Println("希尔排序处理2：", aSort.ShellSort2(baseArr))
}

// 排序比较
func (t tCase) SortVsRuntimes(){
	baseArr := data.GetRandIntArray(1000*5, 10000)
	aSort := &alg.Sort{}
	fmt.Println("基准随机数组： ", baseArr)
	fmt.Println("基准数组长度： ", len(baseArr))
	t1 := getRunTimes()
	fmt.Println("A. 选择排序：", aSort.SelectSort(baseArr))
	fmt.Println("  1).", t1())
	//fmt.Println("基准随机数组： ", baseArr)
	t2 := getRunTimes()
	fmt.Println("B. 插入排序: ", aSort.Inert(baseArr))
	fmt.Println("  1).", t2())
	t3 := getRunTimes()
	fmt.Println("B2. 希尔排序: ", aSort.ShellSort2(baseArr))
	fmt.Println("  1).", t3())
}

// 获取运行时间
func getRunTimes()  func() string {
	t1 := time.Now().UnixNano()
	getTimes := func() string{
		DetSec := (float64(time.Now().UnixNano()) - float64(t1)) / float64(1E9)
		return fmt.Sprintf("用时 %0.4f 秒", DetSec)
	}
	return getTimes
}
func main() {
	(tCase{}).console()
}
