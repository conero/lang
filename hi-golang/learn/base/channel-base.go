package main

import (
	"fmt"
	"io/ioutil"
)

/**
 * @DATE        2019/5/20
 * @NAME        Joshua Conero
 * @DESCRIPIT   Channel 测试
 **/

// 并发计算
func count10M() {
	num := 100000 // 计数
	det := 1000   // 跨度

	type aStruct map[string]int
	aStr := make(chan aStruct)

	i := 1
	times := 0
	for times < num {

		times0 := times
		times2 := i * det
		times = i * det

		i++

		astr := aStruct{
			"start": times0,
			"end":   times2,
			"id":    i,
		}
		// 开启线程
		go func() {
			//fmt.Println(i, times0, times2)
			for j := times0; j < times2; j++ {
			}
			aStr <- astr
		}()
		//times = i * det
		//i++
	}

	text := ""
	//接受参数
	<-aStr

	// 及时已经完成依然报错
	//METHOD 1
	//close(aStr)
	for va := range aStr {
		fmt.Println(va)
		text += fmt.Sprintf("%v \r\n", va)
	}

	//METHOD 2
	/*for {
		select {
		case x, ok := <-aStr:
			if !ok {
				break // channel was closed and drained
			}
			text += fmt.Sprintf("%v \r\n",x)
		}
	}*/
	//关闭 chan
	//close(aStr)

	ioutil.WriteFile("./_debug_output_learn-base-channel--base.log", []byte(text), 0777)
}

func main() {
	count10M()
}
