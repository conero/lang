# **Golang** 



## *语法学习* 

### cgo

> go 语言中调用 *c/c++* 

golang的cgo是调用gcc编译c代码的，gcc工具链在linux上很方便，但windows上是没有的。而windows上一般用的vc，golang是不支持的。

可用 *MinGW*



### 概述 

主要特性
- 自动垃圾回收(gc)
- 更丰富的内置类型
- 多函数多返回值: python
- 错误处理
- 匿名函数和闭包
- 类型和接口
- 并发编程: 轻量级协程支持， goroutine/channel， Erlang 风格(Erlang风格的并发模型, 大部分语言采用共享内存模型.)
- 反射
- 语言交互性: cgo



值类型， 应用类型



> 类型

- 数字型
  - int/uint 
  - int8/int16 
  - uint/uint8/unit16/unit32/uint64
  - float32/float64
- 字符串
  - string
- 数组/切片 
  - []type
- 字典
  - map[type]type
- 结构体
  - struct
- 接口
  - interface



### 基本语法

> 字符串, "" 以及 `` 表示

```go
func main(){
    var str1 = "字符串"
	str2 := `多行字符串`
    fmt.Println(str1, str2)
}

```

> 域源代码为 *utf-8* 编码格式

> 大小写敏感

> 注释， // 单行， /*多行*/

```go
func init(){
	// 单行注释
    t1, n1 := `ffff`, 5
    /*
    多行注释
    */
    fmt.Println(t1, n1)
}
```

> tokens : *标识符*，*关键字*，*运算符和标点符号*以及*文字* 

> 标识符中的第一个字符必须是字母。以字母开头

> **关键字**

```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

> **运算符和标点符号**

```go
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=
```

####  数字 / number

> 整形

支持十进制、八进制、十六进制

> float 浮点数

```go
n10, n8, n16 := 0, o5, 0x100
f1, f2 := 072.47, 3.45e5
// 1e-3 = 0.001
// 1e3 = 1000
```



```go
// utype = 2^n - 1
uint8       the set of all unsigned  8-bit integers (0 to 255), 2百
uint16      the set of all unsigned 16-bit integers (0 to 65535), 6万
uint32      the set of all unsigned 32-bit integers (0 to 4294967295) 42亿
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

uint        either 32 or 64 bits
int         same size as uint
uintptr     an unsigned integer large enough to store the uninterpreted bits of a pointer value
```



#### string/字符串

```go
string_lit             = raw_string_lit | interpreted_string_lit
// 原始字符串
raw_string_lit         = "`" { unicode_char | newline } "`"
// 解析字符串
interpreted_string_lit = `"` { unicode_value | byte_value } `"`

n := len(raw_string_lit)	// 0 ... len(s) -1
```



#### 常量 / constant

> 类型

- boolean  布尔
- rune  符文类型
- integer 整形
- floating-point 浮点数
- complex 符合类型
- string 字符串

Rune, integer, floating-point, and complex constants are collectively called *numeric constants*. 



#### 变量/variable

>  结构化的变量: array, slice, map



#### 类型/ type

> 复合类型 *Composite types* 

array, struct, pointer, function, interface, slice, map, and channel

> 基础类型/ *underlying type* 

 boolean, numeric, or string 



### 接口/interface

> go 语言采用 *组合继承* ；接口实现的方法，一个 ```struct``` 实现 ``interface`` 所有方法则实现了接口

```go
package main

import "fmt"

type A interface {
	Name() string
	Cal(params ...int) int
}

type A1 struct {
}

func (a1 *A1) Name() string {
	return "Add"
}

func (a1 *A1) Cal(params ...int) int {
	all := 0
	for _, n := range params {
		all += n
	}
	return all
}

type A2 struct {
}

func (a2 *A2) Name() string {
	return "Reduce"
}

func (a2 *A2) Cal(params ...int) int {
	all := 0
	for _, n := range params {
		all -= n
	}
	return all
}

func GetA(vtype string) A {
	var a A = nil
	switch vtype {
	case "A1":
		//a = A1{}
		// ERROR: cannot use A1 literal (type A1) as type A in assignment:
		a = new(A1)
	case "A2":
		a = new(A2)
	}
	return a
}

func main() {
	param := []int{100, 10, 1, 0}
	var a1, a2 A
	a1 = GetA("A1")
	a2 = GetA("A2")
	fmt.Println(a1.Name(), a1.Cal(param...))
	fmt.Println(a2.Name(), a2.Cal(param...))
}

```







## import

### Golang 环境设置



通过>go env 可以查看go安装的配置信息
　　Goroot: go的安装路径,官方包路径根据这个设置自动匹配。如同golang的安装一致。
　　Gopath: 工作路径，如果设置错误可能是系统运行有问题。尤其是在这种情况下安装的第三方包，在引用是也会出错。我的经验是这种情况下，应该卸载掉当前的语言环境，重新安装。



绝对路径/相对路径 加载文件/包

1.   相对路径
  import “./model” //当前文件同一目录的model目录，但是不建议这种方式来import
2.   绝对路径
  import “shorturl/model”//加载gopath/src/shorturl/model模块

面展示了一些import常用的几种方式，但是还有一些特殊的import，让很多新手很费解，下面我们来一一讲解一下到底是怎么一回事



### 点操作

1.    点操作
  我们有时候会看到如下的方式导入包
  import(
    . "fmt"
  )
  这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调用的fmt.Println("helloworld")可以省略的写成Println("helloworld")
2.   别名操作
  别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字
  import(
    f "fmt"
  )
  别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("helloworld")
3.   _ 操作
  这个操作经常是让很多人费解的一个操作符，请看下面这个import
  import (
  "database/sql"
    _ "github.com/ziutek/mymysql/godrv"
  )
  _操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数
  注意：
  同一个目录下不能定义不同的package
  应该分为两个不同的目录才可以。
  go，模块中要导出的函数，必须首字母大写。也即是 大写开头方法为共有方法，否则包内私有方法

### Golang 命令

go get = git clone + go install

	go build 加上可以编译的go源文件可以得到一个可执行文件。
	go install 在编译源代码之后还安装到指定的目录
	go get 从指定源上面下载或者更新指定的代码和依赖，并对他们进行编译和安装


## 内建函数或GO基础包

	go/src/builtin/builtin.go   内建函数文档包，可直接使用而不用引入包
	go调用C/C++，或dll文件



### 保留函数

	无参数、返回值
	main	只能应用于package main
	init		能够应用于所有的package	可选
	

	C/C++中接收命令行参数是通过~main函数传入，而go可以使用 os ,flag 库。	os.Args 命令行



### 指针

- 地址

  ```go
  type T1 struct{
  }
  t1 := &T1{}				// T1 的地址
  t2 := *t1				// t1 的值
  ```


### 时间

> 格式化

```
范本： 2006-01-02 15:04:05

如： 获取当前的时间
	time.Now().Format("2006-01-02 15:04:05")  
```





## 参考资料

- 《Go 语言编程》(2012.8)     - 徐式伟 