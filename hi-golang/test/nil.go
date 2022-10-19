package main

import "fmt"

func testNil() {
	var t1 any
	if t1 != nil {
		fmt.Println("t1 is not null")
	} else {
		fmt.Println("t1 is not init")
	}
}

func param(a ...any) {
	if a != nil {
		fmt.Println("input data not null. ")
		for i, v := range a {
			if v != nil {
				fmt.Println(i, "param is not null")
			}
		}
	}
}

func main() {
	//testNil()
	param()
	param(nil)
}
