/**
 * 2017年6月28日 星期三
 * Joshua Conero
 */
// main
package main

import (
	"fmt"

	"../data"
	"../ds"
	"math/rand"
	"time"
)

func getRandInt() int{
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	return src.Int()
}
func getRandHasMax(max int) int{
	src := rand.New(rand.NewSource(time.Now().UnixNano()))
	return src.Intn(max)
}

// 数据正式测试
func RandIntTestCase()  {
	// 随机数生成器
	src := rand.New(rand.NewSource(99))
	fmt.Println(src.Int(), src.Intn(10))
	// 5577006791947779410
	fmt.Println(rand.Int())
	fmt.Println(getRandInt(), " ：时间戳作为种子", getRandHasMax(1000))
	fmt.Println(data.GetRandIntArray(50, 100))
	fmt.Println(data.GetRandArrBase(50, 100))
}

func main() {
	RandIntTestCase()
	fmt.Println("Hello World!")

	queue := new(ds.Queue).Init()
	i := 0
	for i < 10 {
		queue.In("like")
		i = i + 1
	}
	fmt.Println(queue.In("like"), queue.Front)
	fmt.Println(queue.Out())
	fmt.Println(queue.Out())
	fmt.Println(queue.Front)

	//fmt.Println(data.GetRandInt(), data.GetRandLetter())
	//fmt.Println(data.GetRandInt())

	//fmt.Println(string(rune(60)))	// ascii - int
}
