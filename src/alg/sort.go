/**
 * 2017年11月12日 星期日
 * 排序算法
 */
package alg

import (
	"math"
	"fmt"
)

// 排序算法
type Sort struct {
	Ergodic int		// 遍历次数
}

// 小到大还是大到小， 实现顺排
// 插入排序, 插入一个数据到有序队列中依然有效
// 整形排序实现
func (S *Sort) InertSort(data []int, numbers ...int) []int{
	S.Ergodic = 0
	// 遍历插入的数组
	for _,v := range numbers{
		dataLen := len(data)
		// 用于排序
		// 无效数组
		//if dataLen < 1{
		//	break
		//}
		sIt := 0
		lastInData := false			// 最末尾边界标记
		for{
			// 跳出循环
			if sIt >= dataLen{
				break
			}
			S.Ergodic += 1
			// 操作到第一个打的数组
			if data[sIt] > v {
				// 获取原来位置的数据
				targetN := data[sIt]
				data[sIt] = v
				nIdx := sIt
				if dataLen > 1 && nIdx < dataLen{
					sIt2 := sIt+1
					for{
						if sIt2 == dataLen{
							break
						}
						tmpN := data[sIt2]
						data[sIt2] = targetN
						targetN = tmpN
						sIt2 += 1
					}
				}
				data = append(data, targetN)
				lastInData = true
				break
			}
			sIt += 1
		}
		if lastInData == false{
			data = append(data, v)
		}
	}
	return data
}
// 数据
// 基于 InertSort 实现 插入排序
func (S *Sort) Inert(arr []int) []int{
	arr1 := S.InertSort([]int{}, arr...)
	return arr1
}

// 直接选择排序
func (S *Sort) SelectSort(array []int) []int{
	sortEdArr := []int{}
	idxMap := map[int]int{}		// 键值对匹配
	minV := -1
	minK := -1
	arrayLen := len(array)
	for{
		if len(idxMap) == arrayLen{
			break
		}
		xCount := 0
		for i,v := range array{
			// 已经遍历的直接跳过
			if _,has := idxMap[i]; has{
				continue
			}
			if xCount > 0{
				// 取小于等于的数据
				if v <= minV{
					minV = v
					minK = i
				}
			}else{
				minK = i
				minV = v
			}
			xCount += 1
		}
		idxMap[minK] = minV
		sortEdArr = append(sortEdArr, minV)
	}
	return sortEdArr
}

// 希尔排序
func (S *Sort) ShellSort(array []int) []int{
	aLen := len(array)
	vLen := aLen
	getBit := func(vlen int) int {
		return int(math.Ceil(float64(vlen/2)))
	}
	for vLen > 1{
		vLen = getBit(vLen)
		for i:=0; i<vLen; i++{
			n1, n2 := array[i], array[vLen+i+1]
			if n2 < n1{
				array[i] = n2
				array[vLen+i+1] = n1
			}
		}
		fmt.Println(array)
		println(vLen)
	}
	return array
}