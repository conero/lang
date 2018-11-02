package main

import "C"
import "fmt"

// https://colobu.com/2018/08/28/c-and-go-calling-interaction/
// https://studygolang.com/articles/9883
// https://www.jianshu.com/p/7d67068848a4

//export HelloFromGo
func HelloFromGo() {
	fmt.Printf("Hello from Go!\n")
}
