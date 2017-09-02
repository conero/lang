/**
 * 2017年7月26日 星期三
 * 闭包函数测试
 */
 package main

 import(
	 "fmt"
 )

 func main(){	 
	 // 匿名行数
	 (func (){
		 fmt.Println(8*9)
	 })()
 }