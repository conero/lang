package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
 * @DATE        2019/5/21
 * @NAME        Joshua Conero
 * @DESCRIPIT   描述 descript
 **/

func cBufferToGo(i int, ch chan int)  {
	time.Sleep(1*time.Second)
	ch <- i
	fmt.Println(len(ch), cap(ch))
}

/**
c int 缓存数
 */
func cBuffer(det, c int)  {
	ch := make(chan int, c)

	for i := 0; i < det; i++ {
		go cBufferToGo(i, ch)
	}

	// 获取数据
	/*<-ch
	for c1 := range ch{
		fmt.Println(c1)
	}*/


	<- ch
	for {
		//x, ok := <- ch
		_, ok := <- ch
		if !ok{
			break
		}
		//fmt.Println(x)
	}
}

func buffer2()  {
	ch := make(chan int, 3)
	ch <- rand.Int()
	ch <- rand.Int()
	ch <- rand.Int()

	<- ch
	for i := range ch{
		fmt.Println(i)
	}
}

type strIntDick map[string]int
type mSt struct {
	mGs []chan strIntDick
	gsCtt int
}

func (m *mSt) Task(i int, mg chan strIntDick)  {
	dd := strIntDick{
		"id": i,
		"time": time.Now().Nanosecond(),
	}
	rd := rand.New(rand.NewSource(time.Now().Unix()))
	dd["rand"] = rd.Int()
	mg <- dd
}


// 数据结构
func dataStruct(n int)  {
	gsCtt := 50
	// gsCtt 数量必须和实际的个数一样，否则会出现 【fatal error: all goroutines are asleep - deadlock】
	mgs := make([]chan strIntDick, gsCtt)

	m := &mSt{
		mGs: mgs,
		gsCtt: gsCtt,
	}
	for i:=0; i < n; i++{
		mgs[i] = make(chan strIntDick)
		go m.Task(i, mgs[i])
	}

	// 协程自动处理
	for _, ch := range mgs{
		//<-ch

		dd, ok := <-ch
		if !ok{
			break
		}
		fmt.Println(dd)
	}
}

func main() {
	//cBuffer(20, 1024)
	//buffer2()

	dataStruct(50)
}
