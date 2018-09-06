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

> 相关关键字

```go
if else	
for	break continue			// no do-while/while
switch
select
```

### if

```go
// c1
if x>y {
  return x  
}

// c2
data map[string]string{}
if v, has := data["name"]; has{
    return v
}
```



### for

```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }

//-------------------------------------------------
// _ 忽略字符
sum := 0
for _, value := range array {
    sum += value
}

// range 用于 string
for pos, char := range "中文\x80测试" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}

// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```



### switch

```go
// case condition 使用条件/等式
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}

// 多条件一致
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```



> 类型判断

```go
// 类型判断  =>  switch t := t.(type) 
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```



## Function /函数

> 与其他语言不同的是，go函数可以返回多值

```go
// 返回类型
var ConfData map[string]interface{} = map[string]interface{}{}

func HasKey(key string) (interface{}, bool){
    value, has := ConfData[key]
    return (value, has)
}

// 返回带值的参数
func (file *File) Write(b []byte) (n int, err error)

func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```



> defer

// @TODO ...

