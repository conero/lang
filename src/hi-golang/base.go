/**
* 2017年9月2日 星期六
* golang 基础代码示例
*/

package main

import (
	"fmt"
)

type BaseLearn struct {}

// 基本变量测试
func (b *BaseLearn) BaseVariable()  {
	var i, j int
	i = 5
	j = 9
	fmt.Println(i, j)

	// 值交换
	i,j = j,i
	fmt.Println(i, j)

	// 类型强制转换
	t := 1994.25
	fmt.Println(t, int32(t), int64(t))
}

// 类型测试
func (b *BaseLearn) TypeTest(){
	msg := `
		BaseLearn.TypeTest()
	`
	fmt.Print(msg)
}


// go1.9 类型别名
type AliasBaseLearn = BaseLearn
func (abl *AliasBaseLearn) ATypeTest(){
	abl.TypeTest()
}

// 结构体继承
type ExtendBaseLearn struct{
	BaseLearn
	Name string 
	name string
	privateSatck map[string]string
}
// 链式指针
func (ebl *ExtendBaseLearn) Link() (*ExtendBaseLearn){
	return ebl
}
// 链式方法
func (ebl *ExtendBaseLearn) LinkCallBack(CallBack func(cEbl *ExtendBaseLearn)) (*ExtendBaseLearn){
	CallBack(ebl)
	return ebl
}

func (ebl *ExtendBaseLearn) Set(k ,v string) bool{
	// fmt.Println(ebl.privateSatck, "K-V", nil)
	if ebl.privateSatck == nil{
		ebl.privateSatck = map[string]string{}
	}
	ebl.privateSatck[k] = v
	return true
}
func (ebl *ExtendBaseLearn) GET(k string) string{
	return ebl.privateSatck[k]
}

// 基本接口
type BaseInterface interface{
	Set(k, v string) bool
	Get(k string) string
}
type BITest struct{
	privateSatck map[string]string
}
func (bit *BITest) Set(k, v string) bool {
	// 初始化判断
	if bit.privateSatck == nil{
		bit.privateSatck = map[string]string{}
	}
	bit.privateSatck[k] = v
	return true
}
func (bit *BITest) Get(k string) string {
	return bit.privateSatck[k]
}

func main()  {
	// 结果题调用1
	var bl *BaseLearn = &BaseLearn{}
	bl.BaseVariable()

	// 简洁命名方法
	(&AliasBaseLearn{}).ATypeTest()
	//(&BaseLearn{}).ATypeTest()

	// 组合调用
	ebl := &ExtendBaseLearn{}
	//ebl.Link().ATypeTest()
	fmt.Println(ebl.LinkCallBack(func(cEbl *ExtendBaseLearn){
		cEbl.name = "Susanna"
		cEbl.Name = "Emma"
	}).name, ebl.Name)
	
	// 接口的实现
	var bit BaseInterface = new(BITest)	
	bit.Set("name", "SuGoddess")
	fmt.Println(bit.Get("name"))

	// 接口赋值
	//var bit2 BaseInterface = ebl
	//var bit2 BaseInterface = new(ExtendBaseLearn)

	// var bit2 BaseInterface = &ExtendBaseLearn{}
	// // Error cannot use ExtendBaseLearn literal (type *ExtendBaseLearn) as type BaseInterface in assignment:
	// bit.Set("name", "SuGoddess is My Lover")
	// fmt.Println(bit.Get("name"))
}