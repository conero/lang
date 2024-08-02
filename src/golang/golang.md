# Golang

*21世纪的C语言, Go语言已经成为云计算、云存储时代最重要的基础编程语言*

*简洁编程哲学*



> 特性

- expressive
- concise
- clean
- efficient
- with  garbage collection (**GC**)
-  run-time reflection



*拥有自动垃圾回收、一个包系统、函数作为一等公民、词法作用域、系统调用接口、只读的UTF8字符串等*

其他：

- 基于CSP的并发特性支持
- 轻量级线程 goroutine
- Go语言的代码通过**包**（package）组织



*__Go语言本身只有很少的特性__   没有隐式的数值转换，没有构造函数和析构函数，没有运算符重载，没有默认参数，也没有继承，没有泛型，没有异常，没有宏，没有函数修饰，更没有线程局部存储*





> make it easy to write programs that get the most out of multicore and networked machines



> *贝尔实验室的UNIX和C语言两大发明奠定了整个现代IT行业最重要的软件基础*

- 1970s  *Ken Thompson* 和 *Dennis M. Ritchie* 合力开发 **Unix** 操作系统

- *Dennis M. Ritchie* 为了解决UNIX系统的移植性问题而发明了C语言



*Ken Thompson、Rob Pike、Robert Griesemer(设计了V8引擎和HotSpot虚拟机) 为了解决在21世纪多核和网络化环境下越来越复杂的编程问题而发明了Go语言*



*Go语言由来自Google公司的Robert Griesemer，Rob Pike和Ken Thompson三位大牛于2007年9月开始设计和实现，然后于2009年的11月对外正式发布*



*Go语言尤其适合编写网络服务相关基础设施，同时也适合开发一些工具软件和系统软件。 但是Go语言确实是一个通用的编程语言，它也可以用在图形图像驱动编程、移动应用程序开发 和机器学习等诸多领域。目前Go语言已经成为受欢迎的作为无类型的脚本语言的替代者： 因为Go编写的程序通常比脚本语言运行的更快也更安全，而且很少会发生意外的类型错误。*



//@TODO [github 网站学习](<https://github.com/chai2010/advanced-go-programming-book>)

- [第二章　程序结构](<https://docs.hacknode.org/gopl-zh/ch2/ch2.html>)



//@TODO PART II    <https://chai2010.cn/advanced-go-programming-book/>

## cmd

> 命令行

```shell
# 安装包
go get <path>
# 安装且更新为最新的包
go get -u <path>
# 升级全部依赖
go get -u all

# go fmt 
# go 脚本格式化; 全目录，like IDE
go fmt ./...

# 列表当前项目包或模块
go list -m -json all
```



### tool

**nm**  输出在给定的可执行文件、对象文件或存档中涉及的符号。

```shell
# 打印中 yangsu.exe 的符号
#
# 如果使用 <-ldflags "-s -w"> 清除符号后，无法获取符号
go tool nm ./yangsu.exe
```





## 关于

Go 语言族

![](../image/golang/lang-family.png)



### 代码

> workshop    **GOPATH** 用于可指明其目录

- Go programmers typically keep all their Go code in a single *workspace*.    
  - go 程序员通常将程序代码置于单个``工作区``中
- A workspace contains many version control *repositories* (managed by Git, for example).
  - 工作区含多个版本控制
- Each repository contains one or more *packages*.
  - 每个仓库包含单/多个``包``
- Each package consists of one or more Go source files in a single directory.
  - 每个包含有多个文件再单一目录下(*一个目录一个包*)。
- The path to a package's directory determines its *import path*.
  - 包的路径与 *import* 路径一致



> 包

```go
// path = ../name 
package name
```

*单个包中包名应该一致。*



### Testing

> go 拥有一个轻量级的测试框架，通过 ``go test`` 命令进行测试

- 测试文件名 ``{name}_test.go``

- 测试方法入口  ``Test{name}``

```go
func TestName (t *testing.T)
// t.Error/t.Fail   意识着测试失败
```



执行测试命令

```powershell
# 指定文件夹
go test ./str/

# 指定文件
go test -v ./str/calc_test.go ./str/calc.go


# 指定文件函数
go test -v ./str/ '-test.run' TestFloatSimple

# 指定包以及测试函数
go test -v ./ysu/tools '-test.run' '^\QTestNewXRsa\E$'
```





**Benchmark**  golang 自带的基准测试：

```go
func BenchmarkStructToMapLStyle(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		StructToMapLStyle(tom)
	}
}


// goos: windows
// goarch: amd64
// pkg: gitee.com/conero/uymas/util
// cpu: xxx(R) Core(TM) i5-10200H CPU @ 2.40GHz
// BenchmarkStructToMapLStyle
// BenchmarkStructToMapLStyle-8   	   90331	     12899 ns/op
// 标识循环：90331 次, 每次循环花费的时间: 12899 ns.
```





### fetch

> $> go fetch {url}   获取并安装远程包






## *语法学习* 

### cgo

> go 语言中调用 *c/c++* 

golang的cgo是调用gcc编译c代码的，gcc工具链在linux上很方便，但windows上是没有的。而windows上一般用的vc，golang是不支持的。

可用 *MinGW*



> 环境安装/window下 MinGW

*window 需要环境  64位 `gcc`编译器，如:  [mingw-64 网站](http://mingw-w64.org) /  [下载地址](<https://sourceforge.net/projects/mingw-w64/>)；32 位下 __cgo__ 不可用。*

文件下载历史：https://sourceforge.net/projects/mingw-w64/files/mingw-w64/



<!--[@todo https://blog.csdn.net/RA681t58CJxsgCkJ31/article/details/80504482]-->



> go 编译为 dll 文件

```shell
# 编译单文件
go build -buildmode=c-shared -o exportgo.dll exportgo.go

# 编译目录
go build -buildmode=c-shared -o shared.dll .
```





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

转义字符
\x0c   清屏并输出
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





go 组合与继承的区别:

- 如果一个struct嵌套了另一个有名结构体，那么这个模式就叫组合。 

- 如果一个struct嵌套了另一个匿名结构体（只有类型没有名字），那么这个结构可以直接访问匿名结构体的方法，从而实现了继承。 

- 如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现了多重继承。



## import

### Golang 环境设置



通过 `$ go env` 可以查看go安装的配置信息
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
    ```go
    import(
    	. "fmt"
    )
    ```
    这个点操作的含义就是这个包导入之后在你调用这个包的函数时，你可以省略前缀的包名，也就是前面你调用的`fmt.Println("helloworld")`可以省略的写成`Println("helloworld")`
    
2.    别名操作
      别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字
      ```go
      import(
      	f "fmt"
      )
      ```
      别名操作的话调用包函数时前缀变成了我们的前缀，即 `f.Println("helloworld")`
      
3.    _ 操作
      这个操作经常是让很多人费解的一个操作符，请看下面这个import
      ```go
      import (
          "database/sql"
          _ "github.com/ziutek/mymysql/godrv"
      )
      ```
      `_`操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的`init`函数
      
      

注意：
      	同一个目录下不能定义不同的package
      	应该分为两个不同的目录才可以。
      go，模块中要导出的函数，必须首字母大写。也即是 大写开头方法为共有方法，否则包内私有方法




## 并发编程/goroutine

> *CSP* 线程并发模型



### channel

*分为两种： `无缓存通道`、`带缓存通道`。单数据传递使用无缓存通道可行，大量数据需要带缓存的通道。*

```go
// 无缓存的数据通道
ch := make(chan int)

// 带缓存的数据通道
ch2:= make(chan int, 1024)
```





## godoc

> *golang 注释文档格式*

- *注释符`//`后面要加空格, 例如: `// xxx`*
- *在`package, const, type, func`等`关键字`上面并且紧邻关键字的注释才会被展示*

```go
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// 注: 上面注释不会显示为文档；单独的 “//” 可用换行

// Package path implements utility routines for manipulating slash-separated
// paths.
//
// The path package should only be used for paths separated by forward
// slashes, such as the paths in URLs. This package does not deal with
// Windows paths with drive letters or backslashes; to manipulate
// operating system paths, use the path/filepath package.
package path
```

- *`Package`的注释如果行数过多可以考虑用 `doc.go` 文档代替*

- _**如果当前包目录下包含多个Package注释的go文件(包括doc.go), 那么按照文件名的字母数序优先显示**_

- _`Package`的注释会出现在godoc的[包列表](#)中, 但只能展示大约523字节的长度_

- *在无效注释中以`BUG(who)`开头的注释, 将被识别为已知bug, 显示在`bugs`区域*

  ```go
  // BUG(who): 我是bug说明
  
  // Package banana ...
  package banana
  
  ```

- *如果`bug注释`和`关键字注释`中间无换行, 那么`混合的注释`将被显示在`bugs`和`godoc列表`两个区域内*

  ```go
  // BUG(who): 我是bug注释
  // Package banana ...也是pkg注释
  package banana
  
  ```

  

- *URL将被转化为HTML链接*



> **Example**

- *文件必须放在当前包下*
- *文件名以`example`开头, `_`连接, `test`结尾, 如:`example_xxx_test.go`*
- *包名是`当前包名` + `_test`, 如: `strings_test`*
- *函数名称的格式`func Example[FuncName][_tag]()`*
- *函数注释会展示在页面上*
- *函数结尾加上`// Output:`注释, 说明函数返回的值*



> 命令行

```shell
# 开启一个godoc小型server, -play可以使用playground运行Example代码
godoc -http=:6060 -play
```



### 标签

```go
// Deprecated: 遗弃代码，IDE 高亮显示
```





## go-line command

>go 编译器命令行参考



**查看帮助**

```powershell
go help
```



~~go get = git clone + go install~~

```shell
# 代码编译
go build 加上可以编译的go源文件可以得到一个可执行文件。
# 批量[cmd]下所有二级制程序编译
# 注意：cmd 目录有效，其他目录经测试无法通过
go build -o ./bin ./cmd/...

# 使用 ldflags 简化过编译文件大小
# 优化文件大小
# -s 忽略调试信息
# -w 忽略DWARF表符号
# -trimpath  从输出的二进制文件中移除路径信息
go build -o ./bin -ldflags "-s -w" ./cmd/...
# 安装程序
go install -ldflags "-w -s" -trimpath .\cmd\umsp\

# go 代码生成，当前下的所有代码
go generate ./...

# 本地项目安装
go install 在编译源代码之后还安装到指定的目录

# 依赖下载，在mod开启的情况下，其实下载到本地项目
go get 从指定源上面下载或者更新指定的代码和依赖，并对他们进行编译和安装

#
#
# 项目模块
go mod 项目模块系统

go mod init 	项目初始化
# 项目不在 【GOPATH】中时，需要指定模块名称，如
go mod init [module] 

go mod vendor 	模块自动拉取，并生产vendor目录
```





> go 自定义*tag*编译

- 编写脚本

```go
// +build embed_mode
// 自定义编译模式
```



脚本写法2

```go
//go:build embed
// 自定义编译模式
```





- 编译程序

```shell
# 使用自定义模式进行编译即可
go build -tags embed_mode

# 支持多个tags 如t1/t2/t3
go build -tags "t1 t2 t3"

# 不使用 tags
go build 
```



### gopath、go module

*go 1.11 新增 对 go module 的支持，提供 vendor 库支持；而 gopath 几何与 go 一起诞生，后者必须在 gopath 进行项目开发。*



**实际测试 IntellJ + go plugin**  时间: 2019年10月25日 星期五， 无法友好的提示 vendor 目录下的依赖，虽然可运行带不够友好。  



**go module** 使用 replace 指令

```go
module conero.cn/firefly/secret

go 1.21.3

replace conero.cn/firefly/ffutil => ../ffutil

require (
    conero.cn/firefly/ffutil v0.0.0 // 可以是任意占位符版本号
    gitee.com/conero/uymas v1.4.0-alpha.1.0.20240116060107-6d9af8acb109
)
```

使用 `go mod vendor` 可更新，replace下的最新代码。



### 条件编译

编译常量，也称编译标签。其在文件的首行:

```go
// +build

// 示例
// +build linux darwin
// +build amd64
// +build linux,386 darwin,!cgo
// +build linux,cgo darwin,cgo
// +build !linux,!darwin !cgo

// 操心系统来源常量: runtime.GOOS
// 操心系统架构来源常量: runtime.GOARCH
```



条件编译也可在文件命名是指定，指定规则：

```shell
*_GOOS
*_GOARCH
*_GOOS_GOARCH

# 如: 
#  source_windows_amd64.go
```



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



## 资源/库

### linting 工具

- golint    代码质量检查工具，它提供了一套标准的代码风格和潜在问题检查。https://github.com/golang/lint （archived） 已停止开发
- gosec   功能更加强大的静态分析工具，专注于检测Go代码中的安全漏洞
- golangci-lint     https://github.com/golangci/golangci-lint
- staticcheck     https://github.com/dominikh/go-tools



```shell
# golangci-lint
golangci-lint run

# 运行所有
golangci-lint run ./...

# 指定目录
golangci-lint run dir1 dir2/...
golangci-lint run file1.go
```



### 图像/GUI

- [ebiten](https://github.com/hajimehoshi/ebiten)              2D 游戏引擎





#### fyne

*跨平台的 GUI 库，[GitHub代码库](<https://github.com/fyne-io/fyne>)。*

> window GUI APP 默认含 cmd 窗口，可使用参数选项来消除

```powershell
go build -ldflags -H=windowsgui <filename>
```





### Websocket

要使用 Golang 开发 WebSocket，选择基本就在 [x/net/websocket](http://godoc.org/golang.org/x/net/websocket) 和 [gorilla/websocket](http://godoc.org/github.com/gorilla/websocket) 之间。《Go Web 编程》一书中的例子使用了 `x/net/websocket` 作为开发包，而且貌似它也更加官方且正式。而实际根据我在网上查询得到的反馈看来，并非如此。`x/net/websocket` 貌似 Bug 较多，且较为不稳定，问题解决也并不及时。相比之下，`gorilla/websocket` 则更加优秀。



### 规范

#### 项目目录规范

参考：[golang-standards/project-layout](https://github.com/golang-standards/project-layout)

主要目录结构

- cmd                            项目应用程序，如二进制包
- internal                      私有包，不可为外界引用，**Go强制执行**的。
- pkg                             外部程序可应用的库，当顶级目录有很多非go代码是可使用该目录
- vendor                       go modules 安装库
- api                              如Swagger相关的各种 json 文件
- web                            静态web资源，或者服务端模板或者SPA文件
- configs/config           配置文件模板或者默认配置
- init                              系统初始化，进程管理等配置
- scripts                        项目shell等相关脚本
- test                             代码测试目录
- docs                           项目相关文件，设计或计划文档等
- tools                          项目支持的工具
- examples                  应用程序或公共实例代码
- assets                        静态资源，如图形
- website                      网站数据





### 交叉编译

go 对windows 版本支持介绍如：

- 1.10   最后支持 Windows XP（2001）、Windows Vista 的系统
- 1.20   最后支持 win7(2009)、win8(2012)的系统



#### Windows for Linux/max

*386 编译的程序为： 32位。*

```shell
 # 设置环境变量
 # linux, darwin, windows, netbsd
 $env:GOOS="linux"
 
 # 系统架构
 # amd64, 386, arm, ppc64
 $env:GOARCH = "amd64"
 
 # 运行，即使对应操作系统的脚本
 go build
```



| 操作系统 GOOS | 支持架构 GOARCH                                              |      |
| ------------- | ------------------------------------------------------------ | ---- |
| windows       | 386, amd64                                                   |      |
| linux         | 386,amd64,arm,arm64,ppc64,ppc64le,mips,mipsle,mips64,mips64le,s390x |      |
| darwin        | amd64,arm64                                                  |      |
| android       | 386,amd64,arm,arm64                                          |      |
| aix           | ppc64                                                        |      |
| dragonfly     | amd64                                                        |      |
| freebsd       | 386,amd64,arm                                                |      |
| illumos       | amd64                                                        |      |
| js            | wasm                                                         |      |
| netbsd        | 386,amd64,arm                                                |      |
| openbsd       | 386,amd64,arm,arm64                                          |      |
| plan9         | 386,amd64,arm                                                |      |
| solaris       | amd64                                                        |      |



## 附录

学习资源

- zh book [Go 语言原本](https://golang.design/under-the-hood/)        go语言学习资料
- zh book [Go 程序员面试笔试宝典](https://golang.design/go-questions/)
- 





### 参考资料

- 《Go 语言编程》(2012.8)     - 徐式伟 
- [《go语言圣经》](<https://docs.hacknode.org/gopl-zh/>)（中文版）



### todos

- //@TODO  [godoc 介绍以及 Golang 注释规范 ](https://blog.cyeam.com/golang/2018/09/03/godoc#%E6%A0%87%E9%A2%98%E5%92%8C%E6%AE%B5%E8%90%BD)

