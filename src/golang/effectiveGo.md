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

Go's `defer` statement schedules a function call (the *deferred* function) to be run immediately before the function executing the `defer` returns。

Go的defer语句用来调度一个函数调用（被延期的函数），使其在执行defer的函数即将返回之前才被运行,被延期执行的函数，它的参数（包括接受者）实在defer执行的时候被求值的，而不是在调用执行的时候。也就是说被延期执行的函数的参数是按正常顺序被求值的。

- *defer 逆序执行*
- *用于在处理资源时，函数结束时执行*



```go
// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}
```



## Data/ 数据

*Go has two allocation primitives, the built-in functions `new` and `make`.*

go 使用内建函数 `new` 和 `make` 分配类型

### `new` 分配类型

- 与其他语言中的名称不同，它不会*初始化*内存，它只会将其*归零*。即 `new (T) => *T`, 它返回指向新分配的类型零值的指针 `T`

```go
// 两者等价
p := new(T)  // type *T
var v T      // type  T

// new eg1
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
// new eg2
// 使用“复合文字”
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    
    f := File{fd, name, nil, 0}
    return &f
}
```

>复合文字(composite literals)

```go
type T struct{
    Name string
    Age int
    Gender string    
}
// 函数
func Test(){
    // 复合文字的字段按顺序排列，并且必须全部存在。
    t := &T{"Joshua Conero", 28, 'M'}
    t2 := &T{
        Name: "Joshua Conero",
        Gender: 'M'
    }
}
```



### `make` 分配类型

函数 `make(T, `**args**`)` 

与 `new` 不同： It creates slices, maps, and channels only, and it returns an *initialized* (not *zeroed*) value of type `T` (not `*T`).  仅仅用于 `lices, maps, and channels`,其初始化非空的的值， 非*指针*



```go
// new 与 make 的主要区别
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic:
v := make([]int, 100)
```



### Arrays/ 数组

> 与 C 不同

- Arrays are values. Assigning one array to another copies all the elements. 
  - *值复制，非址引用*
- In particular, if you pass an array to a function, it will receive a *copy* of the array, not a pointer to it. 
  - *数组在函数中值传递为：值传递，非址传递*
- The size of an array is part of its type. The types `[10]int` and `[20]int` are distinct.



### Slices/ 切片

*切片保存对基础数组的引用，如果将一个切片分配给另一个切片，则两者都引用相同的数组。*





### Maps/ 字典

> *k-v* 数据结构

*其值访问与 array、slice相同*

```go
var M1 map[string]string = map[string]string{
    "a":"alt",
    "b":"ben",
}

// 获取值
ben := M1["b"]

// 存在与否
// string,bool := map[string]string
jan,has := M1["j"]
_,has := M1["j"]


```



// @TODO ...

