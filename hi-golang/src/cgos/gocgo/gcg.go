package main

import "C"
import (
	"fmt"
	"time"
)
// 此处若写错，可能出错。 如 "//export 无空格"
//export DateString4Go
func DateString4Go() {
	fmt.Printf("From Go Func(go)::%s", time.Now().String())
}
