# effective Go

> Joshua Conero
>
> 2018年9月3日 星期一



version-go 1

> [https://golang.google.cn/doc/effective_go.html](https://golang.google.cn/doc/effective_go.html) 学习笔记/以及翻译





> ``go fmt`` 格式化 go 语言格式

> 分号

*go 代码不以分号结尾，但在程序中依然可能出现分号*

```go
for i=0; i<100; i++{
    fmt.Println(i)
}
```





## 注释

> C-style 		多行注释    /**/
>
> C++-Style   	单行注释

```go
package main
// 包引入
import "fmt"

func main(){
    /**
    *      多行注释
    **/
    fmt.Println(5+9*14, "Hi-Man")
}
```



## 命名

> 包级可见命名 ``必须以大写字母开头`` (公用)

```go
var A struct{
    PubName string	// public 属性
    priName string	// private 属性
}
```



> By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to construct an agent noun: `Reader`, `Writer`, `Formatter`, `CloseNotifier` etc.

*通常单方法接口命名以 -er 结尾， 如: Reader, Writer*



> Getter

*If you have a field called `owner` (lower case, unexported), the getter method should be called `Owner`(upper case, exported), not `GetOwner`.*



## 控制语句

// @TODO ...

