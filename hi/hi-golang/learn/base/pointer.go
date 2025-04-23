/**
 * golang 指针测试
 * 2017年11月25日 星期六
 */
package main

import (
	"encoding/json"
	"fmt"
)

// 基础结构体指针
type pointerBaseStruct struct {
	Score  float64 `json:"score"`  // 分数
	IsPass bool    `json:"isPass"` // 是否通过
	Name   string  `json:"name"`   // 姓名
}

// 基本支持测试
func BasePointer() {
	tA := []int{}
	var tAPnt *[]int
	// 初始化
	tA = append(tA, []int{1, 8, 7, 9, 8, 0}...)
	fmt.Println(tA)
	fmt.Println(tAPnt)

	// 指针引用
	tAPnt = &tA
	fmt.Println("tAPnt 去数组 tA的地址以后", tAPnt)
	//直接计算
	//tAPnt = append(tAPnt, []int{1, 3, 1, 4, 5, 2, 1})			// Error: first argument to append must be slice; have *[]int
	*tAPnt = append(*tAPnt, []int{1, 3, 1, 4, 5, 2, 1}...) // Error: first argument to append must be slice; have *[]int

	// 去除值，并替换
	//tmpArr := *tAPnt
	//tmpArr = append(tmpArr, []int{1, 3, 1, 4, 5, 2, 1}...)
	//tAPnt = &tmpArr

	fmt.Println("tAPnt 去数组 tA的地址以后", tAPnt)
	fmt.Println("tAPnt 去数组 tA的地址以后", tA)
}

// 基础结构体测试
//
// 测试值更改
func demoStructBase(bp *pointerBaseStruct) {
	const vStr = `{"name":"张三","score":102.54,"isPass":true}`
	var tmpBp pointerBaseStruct

	err := json.Unmarshal([]byte(vStr), &tmpBp)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
		return
	}

	fmt.Printf("tmpBp(%p): %#v\n", &tmpBp, tmpBp)
	bp = &tmpBp
	bp.Name = "安红"
	fmt.Printf("bp-doing(%p): %#v\n", &bp, bp)
}

func demoString(vStr *string) {
	var tmpStr = "函数内部临时值"
	fmt.Println("函数内部值（临时）:" + tmpStr)
	// 此处可改变： 因为外部的 str 和函数内部的 vStr 指向同一个内存地址，所以修改会生效。
	// *vStr = "--..  函数内部赋值（step1）"
	//*vStr = tmpStr
	// 将 【vStr】副本指针指向【tmpStr】的内存地址
	vStr = &tmpStr
	fmt.Println("函数内部值2:" + *vStr)
	*vStr = "--..  函数内部赋值"
	fmt.Println("func内部值3:" + *vStr)
}

// 测试用例
func testDemoStructBase() {
	var bp pointerBaseStruct = pointerBaseStruct{
		Score:  35,
		IsPass: false,
		Name:   "李丽",
	}

	// 地址更变情况测试
	fmt.Printf("bp-start(%p): %#v\n", &bp, bp)
	fmt.Println()
	demoStructBase(&bp)
	fmt.Println()
	fmt.Printf("bp-end(%p): %#v\n", &bp, bp)
	fmt.Println()

	// 字符串处理
	var tmpStr = "函数外部值(起始)"
	fmt.Println("函数外部值:" + tmpStr)
	fmt.Println()
	demoString(&tmpStr)
	fmt.Println()
	fmt.Println("函数外部值:" + tmpStr)
}

func main() {
	// BasePointer()
	testDemoStructBase()
}
