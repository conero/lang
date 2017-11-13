/**
 * 2017年6月28日 星期三
 * Joshua Conero
 * 数据生成器
 */
package data

import (
	"math/rand"
	"time"
)

// 基数
var baseRandInt int = 1
var baseRandIntMax int = 200
// 后去基数
func getBaseInt() int {
	if baseRandInt >= baseRandIntMax{
		baseRandInt = 0
	}
	baseRandInt += 1
	return baseRandInt
}

// 随机整数形
func GetRandInt() int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano() * int64(getBaseInt())))
	return rd.Int()
}

// 随机整数形规定最大值
func GetRandIntHasMax(max int) int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano() * int64(getBaseInt())))
	return rd.Intn(max)
}

// 获取数组， (bit, max)
func GetRandIntArray(params ...int) []int{
	bit := 10			// 数组位数
	max := 1000					// 随机数的最大值
	data := []int{}
	if len(params) > 0{
		bit = params[0]
	}
	if len(params) > 1{
		max = params[1]
	}
	// 内部随机速生成器
	randnInt := func(n0 int) int {
		rd := rand.New(rand.NewSource(time.Now().UnixNano()+int64(n0)))
		return rd.Intn(max)
	}
	i := 1
	for  {
		if i > bit{
			break
		}
		// 由于程序运行中单纯靠时间错，可能依然有很多项目的随机水
		// 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 442 442 442 442 442 442 442 442 442 442 442 4
		//data = append(data, GetRandIntHasMax(max))
		data = append(data, randnInt(i))
		i += 1
	}
	return data
}

// 获取数组， (bit, max)
// 纯时间戳驱动
func GetRandArrBase(params ...int) []int{
	bit := 10			// 数组位数
	max := 1000					// 随机数的最大值
	data := []int{}
	if len(params) > 0{
		bit = params[0]
	}
	if len(params) > 1{
		max = params[1]
	}
	i := 1
	for  {
		if i > bit{
			break
		}
		// 由于程序运行中单纯靠时间错，可能依然有很多项目的随机数
		// 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 273 442 442 442 442 442 442 442 442 442 442 442 4
		data = append(data, GetRandIntHasMax(max))
		i += 1
	}
	return data
}

/*
// 获取数据字母：A-Z: 65-90; 97 - 122
func GetRandLetter() string {
	num := GetRandInt()
	nLen := len(num)
	i := 0
	letter := ""
	for i < nLen-1 {
		firstNum := nLen[i]
		if firstNum > 1 {
			letter += string(rune(nLen[i : i+1]))
			i = i + 1
		}
	}
	return letter
}
*/
