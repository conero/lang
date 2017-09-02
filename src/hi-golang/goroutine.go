/**
 *	2017年9月1日 星期五
 *	go 协程与channel 测试
 */

package main 

import(
	"fmt"
	"time"
)




// 测试用例
type GoRountineTest struct{}



// 基本线程测试

func (g *GoRountineTest) BaseTest(){
	//
	t := func(vs []string, rChannel chan int) {
		vlen := 0
		for _, v := range vs{
			vlen += len(v)
		}
		rChannel <- vlen
	}

	RC := make(chan int)
	go t([]string{"yang", "su"}, RC)
	// 地址
	fmt.Println(RC)
	// 传入值
	fmt.Println(<-RC)


	go t([]string{"yang", "su", "love"}, RC)

	fmt.Println(<-RC)
}

// 阻塞式协程处理
func (g *GoRountineTest) ChokeTest()  {

	go (func() {
		fmt.Println(time.Now().Format("2006"))
	})()
	//
	ch := make(chan int)
	value := <-ch
	fmt.Println(value)
}

func main(){
	grt := &GoRountineTest{}
	grt.BaseTest()
	//grt.ChokeTest()
}

