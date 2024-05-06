package main

// #include <stdio.h>
// #include "base.h"
// #include <stdlib.h>
import "C"

import (
	"fmt"
)

// golang的cgo是调用gcc编译c代码的，gcc工具链在linux上很方便，但windows上是没有的。而windows上一般用的vc，golang是不支持的。
func main() {
	fmt.Println(C.count)
	fmt.Println(C.add())
}
