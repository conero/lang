package main

import (
	"fmt"
	"math"
)

/**
 * @DATE        2019/4/14
 * @NAME        Joshua Conero
 * @DESCRIPIT   描述 descript
 **/

// @TODO StillNeedToDo 算法解法1
// 解答
func prefixesDivBy5(A []int) []bool {
	var value []bool = []bool{}
	var dec int = 0
	for i, v := range A {
		dec += v * int(float64(math.Pow(float64(2), float64(i))))
		fmt.Println(i, v, dec)
		if dec%5 == 0 {
			value = append(value, true)
		} else {
			value = append(value, false)
		}
	}
	return value
}

// 控制台
func binaryPDBy5(vin, vneed []bool) bool {
	if len(vin) != len(vneed) {
		return false
	}
	for i, v := range vin {
		if v != vneed[i] {
			return false
		}
	}
	return true
}

// 控制台
func main() {
	var vin []int
	var vout, vneed []bool
	var vformat string = "%v) %v => %v / %v\r\n"

	//// test
	//vin = []int{0, 1, 1}
	//vneed = []bool{true, false, false}
	//vout = prefixesDivBy5(vin)
	//fmt.Printf(vformat, binaryPDBy5(vout, vneed), vin, vout, vneed)
	//
	//// test
	//vin = []int{1, 1, 1}
	//vneed = []bool{false, false, false}
	//vout = prefixesDivBy5(vin)
	//fmt.Printf(vformat, binaryPDBy5(vout, vneed), vin, vout, vneed)
	//
	//// test
	//vin = []int{0, 1, 1, 1, 1, 1}
	//vneed = []bool{true, false, false, false, true, false}
	//vout = prefixesDivBy5(vin)
	//fmt.Printf(vformat, binaryPDBy5(vout, vneed), vin, vout, vneed)
	//
	//// test
	//vin = []int{1, 1, 1, 0, 1}
	//vneed = []bool{false, false, false, false, false}
	//vout = prefixesDivBy5(vin)
	//fmt.Printf(vformat, binaryPDBy5(vout, vneed), vin, vout, vneed)

	// test
	vin = []int{1, 1, 0, 0, 0, 1, 0, 0, 1}
	vneed = []bool{false, false, false, false, false, false, false, false, false}
	vout = prefixesDivBy5(vin)
	fmt.Printf(vformat, binaryPDBy5(vout, vneed), vin, vout, vneed)
}
