/**
 * 2017年6月28日 星期三
 * Joshua Conero
 */
// main
package main

import (
	"fmt"

	//"../data"
	"../ds"
)

func main() {
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
