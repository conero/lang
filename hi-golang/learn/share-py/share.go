package main

import "C"
import "fmt"

//export joshua_conero
func joshua_conero() {
    fmt.Println("Joshua Conero demo it.")
}

//export cal
func cal(x, y int) int {
	return  x*x - y
}

//export cal_about
func cal_about() string {
    return "x*x - y"
}

func main()  {
	// Need a main function to make CGO compile package as C shared library
}