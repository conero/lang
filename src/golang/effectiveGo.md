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

*键值类型：integers, floating point and complex numbers, strings, pointers, interfaces, structs and arrays*

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



### Printing / 格式打印

> *fmt* 包



> **verb 词结构**

#### 通常/General

- `%v`   指向值的默认格式，任意类型；*struct* 时 `%+v` 包括 field names
- `%#v`  *go-syntax 语法式值输出*
- `%T`  *go-syntax 语法式值的类型*
- `%%`   *a literal percent sign; consumes no value*



#### Boolean

- `%t` *布尔结构： true/false*



#### Integer

- `%b`  *二进制(base 2)*
- `%c`  *Unicode code point*
- `%d`  *十进制*
- `%o`  *八进制1*
- `%q`  *go-syntax 单引号字符安全编码*
- `%x` *十六进制 (a-f)*
- `%X` *十六进制 (A-F)*
- `%U`  *Unicode 格式，Unicode format: U+1234; same as "U+%04X"*



#### Floating-point and complex constituents

> 浮点数和复数

- `%b` *指数为两幂的无十进制的科学符号, 如： -123456p-78*
- `%e` *科学计数法， 如： -1.234456e+78*
- `%E` *科学计数法， 如： -1.234456E+78*
- `%f` *浮点数，无指数。如： 123.456*
- `%F` *%f 同义*
- `%g` *大指数 %e，否则 %f*
- `%G` `%E~%F`



#### String and slice of bytes

- `%s` *字符串/切片片接字符*
- `%q` *go-syntax 带双引号的字符串*
- `%x` *十六进制 (a-f)，每两个字符一个位*
- `%X` *十六进制 (A-F)，每两个字符一个位*



#### Slice

- `%p`  *十六进制，地址表示形式, 带： 0x*



#### Pointer

- `%p` *十六进制，值表示形式, 带： 0x*



>  *来源/参照： go 源码 fmt/doc.go*

```go
/*
	Package fmt implements formatted I/O with functions analogous
	to C's printf and scanf.  The format 'verbs' are derived from C's but
	are simpler.


	Printing

	The verbs:

	General:
		%v	the value in a default format
			when printing structs, the plus flag (%+v) adds field names
		%#v	a Go-syntax representation of the value
		%T	a Go-syntax representation of the type of the value
		%%	a literal percent sign; consumes no value

	Boolean:
		%t	the word true or false
	Integer:
		%b	base 2
		%c	the character represented by the corresponding Unicode code point
		%d	base 10
		%o	base 8
		%q	a single-quoted character literal safely escaped with Go syntax.
		%x	base 16, with lower-case letters for a-f
		%X	base 16, with upper-case letters for A-F
		%U	Unicode format: U+1234; same as "U+%04X"
	Floating-point and complex constituents:
		%b	decimalless scientific notation with exponent a power of two,
			in the manner of strconv.FormatFloat with the 'b' format,
			e.g. -123456p-78
		%e	scientific notation, e.g. -1.234456e+78
		%E	scientific notation, e.g. -1.234456E+78
		%f	decimal point but no exponent, e.g. 123.456
		%F	synonym for %f
		%g	%e for large exponents, %f otherwise. Precision is discussed below.
		%G	%E for large exponents, %F otherwise
	String and slice of bytes (treated equivalently with these verbs):
		%s	the uninterpreted bytes of the string or slice
		%q	a double-quoted string safely escaped with Go syntax
		%x	base 16, lower-case, two characters per byte
		%X	base 16, upper-case, two characters per byte
	Slice:
		%p	address of 0th element in base 16 notation, with leading 0x
	Pointer:
		%p	base 16 notation, with leading 0x
		The %b, %d, %o, %x and %X verbs also work with pointers,
		formatting the value exactly as if it were an integer.

	The default format for %v is:
		bool:                    %t
		int, int8 etc.:          %d
		uint, uint8 etc.:        %d, %#x if printed with %#v
		float32, complex64, etc: %g
		string:                  %s
		chan:                    %p
		pointer:                 %p
	For compound objects, the elements are printed using these rules, recursively,
	laid out like this:
		struct:             {field0 field1 ...}
		array, slice:       [elem0 elem1 ...]
		maps:               map[key1:value1 key2:value2 ...]
		pointer to above:   &{}, &[], &map[]

	Width is specified by an optional decimal number immediately preceding the verb.
	If absent, the width is whatever is necessary to represent the value.
	Precision is specified after the (optional) width by a period followed by a
	decimal number. If no period is present, a default precision is used.
	A period with no following number specifies a precision of zero.
	Examples:
		%f     default width, default precision
		%9f    width 9, default precision
		%.2f   default width, precision 2
		%9.2f  width 9, precision 2
		%9.f   width 9, precision 0

	Width and precision are measured in units of Unicode code points,
	that is, runes. (This differs from C's printf where the
	units are always measured in bytes.) Either or both of the flags
	may be replaced with the character '*', causing their values to be
	obtained from the next operand (preceding the one to format),
	which must be of type int.

	For most values, width is the minimum number of runes to output,
	padding the formatted form with spaces if necessary.

	For strings, byte slices and byte arrays, however, precision
	limits the length of the input to be formatted (not the size of
	the output), truncating if necessary. Normally it is measured in
	runes, but for these types when formatted with the %x or %X format
	it is measured in bytes.

	For floating-point values, width sets the minimum width of the field and
	precision sets the number of places after the decimal, if appropriate,
	except that for %g/%G precision sets the maximum number of significant
	digits (trailing zeros are removed). For example, given 12.345 the format
	%6.3f prints 12.345 while %.3g prints 12.3. The default precision for %e, %f
	and %#g is 6; for %g it is the smallest number of digits necessary to identify
	the value uniquely.

	For complex numbers, the width and precision apply to the two
	components independently and the result is parenthesized, so %f applied
	to 1.2+3.4i produces (1.200000+3.400000i).

	Other flags:
		+	always print a sign for numeric values;
			guarantee ASCII-only output for %q (%+q)
		-	pad with spaces on the right rather than the left (left-justify the field)
		#	alternate format: add leading 0 for octal (%#o), 0x for hex (%#x);
			0X for hex (%#X); suppress 0x for %p (%#p);
			for %q, print a raw (backquoted) string if strconv.CanBackquote
			returns true;
			always print a decimal point for %e, %E, %f, %F, %g and %G;
			do not remove trailing zeros for %g and %G;
			write e.g. U+0078 'x' if the character is printable for %U (%#U).
		' '	(space) leave a space for elided sign in numbers (% d);
			put spaces between bytes printing strings or slices in hex (% x, % X)
		0	pad with leading zeros rather than spaces;
			for numbers, this moves the padding after the sign

	Flags are ignored by verbs that do not expect them.
	For example there is no alternate decimal format, so %#d and %d
	behave identically.

	For each Printf-like function, there is also a Print function
	that takes no format and is equivalent to saying %v for every
	operand.  Another variant Println inserts blanks between
	operands and appends a newline.

	Regardless of the verb, if an operand is an interface value,
	the internal concrete value is used, not the interface itself.
	Thus:
		var i interface{} = 23
		fmt.Printf("%v\n", i)
	will print 23.

	Except when printed using the verbs %T and %p, special
	formatting considerations apply for operands that implement
	certain interfaces. In order of application:

	1. If the operand is a reflect.Value, the operand is replaced by the
	concrete value that it holds, and printing continues with the next rule.

	2. If an operand implements the Formatter interface, it will
	be invoked. Formatter provides fine control of formatting.

	3. If the %v verb is used with the # flag (%#v) and the operand
	implements the GoStringer interface, that will be invoked.

	If the format (which is implicitly %v for Println etc.) is valid
	for a string (%s %q %v %x %X), the following two rules apply:

	4. If an operand implements the error interface, the Error method
	will be invoked to convert the object to a string, which will then
	be formatted as required by the verb (if any).

	5. If an operand implements method String() string, that method
	will be invoked to convert the object to a string, which will then
	be formatted as required by the verb (if any).

	For compound operands such as slices and structs, the format
	applies to the elements of each operand, recursively, not to the
	operand as a whole. Thus %q will quote each element of a slice
	of strings, and %6.2f will control formatting for each element
	of a floating-point array.

	However, when printing a byte slice with a string-like verb
	(%s %q %x %X), it is treated identically to a string, as a single item.

	To avoid recursion in cases such as
		type X string
		func (x X) String() string { return Sprintf("<%s>", x) }
	convert the value before recurring:
		func (x X) String() string { return Sprintf("<%s>", string(x)) }
	Infinite recursion can also be triggered by self-referential data
	structures, such as a slice that contains itself as an element, if
	that type has a String method. Such pathologies are rare, however,
	and the package does not protect against them.

	When printing a struct, fmt cannot and therefore does not invoke
	formatting methods such as Error or String on unexported fields.

	Explicit argument indexes:

	In Printf, Sprintf, and Fprintf, the default behavior is for each
	formatting verb to format successive arguments passed in the call.
	However, the notation [n] immediately before the verb indicates that the
	nth one-indexed argument is to be formatted instead. The same notation
	before a '*' for a width or precision selects the argument index holding
	the value. After processing a bracketed expression [n], subsequent verbs
	will use arguments n+1, n+2, etc. unless otherwise directed.

	For example,
		fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	will yield "22 11", while
		fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)
	equivalent to
		fmt.Sprintf("%6.2f", 12.0)
	will yield " 12.00". Because an explicit index affects subsequent verbs,
	this notation can be used to print the same values multiple times
	by resetting the index for the first argument to be repeated:
		fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)
	will yield "16 17 0x10 0x11".

	Format errors:

	If an invalid argument is given for a verb, such as providing
	a string to %d, the generated string will contain a
	description of the problem, as in these examples:

		Wrong type or unknown verb: %!verb(type=value)
			Printf("%d", hi):          %!d(string=hi)
		Too many arguments: %!(EXTRA type=value)
			Printf("hi", "guys"):      hi%!(EXTRA string=guys)
		Too few arguments: %!verb(MISSING)
			Printf("hi%d"):            hi%!d(MISSING)
		Non-int for width or precision: %!(BADWIDTH) or %!(BADPREC)
			Printf("%*s", 4.5, "hi"):  %!(BADWIDTH)hi
			Printf("%.*s", 4.5, "hi"): %!(BADPREC)hi
		Invalid or invalid use of argument index: %!(BADINDEX)
			Printf("%*[2]d", 7):       %!d(BADINDEX)
			Printf("%.[2]d", 7):       %!d(BADINDEX)

	All errors begin with the string "%!" followed sometimes
	by a single character (the verb) and end with a parenthesized
	description.

	If an Error or String method triggers a panic when called by a
	print routine, the fmt package reformats the error message
	from the panic, decorating it with an indication that it came
	through the fmt package.  For example, if a String method
	calls panic("bad"), the resulting formatted message will look
	like
		%!s(PANIC=bad)

	The %!s just shows the print verb in use when the failure
	occurred. If the panic is caused by a nil receiver to an Error
	or String method, however, the output is the undecorated
	string, "<nil>".

	Scanning

	An analogous set of functions scans formatted text to yield
	values.  Scan, Scanf and Scanln read from os.Stdin; Fscan,
	Fscanf and Fscanln read from a specified io.Reader; Sscan,
	Sscanf and Sscanln read from an argument string.

	Scan, Fscan, Sscan treat newlines in the input as spaces.

	Scanln, Fscanln and Sscanln stop scanning at a newline and
	require that the items be followed by a newline or EOF.

	Scanf, Fscanf, and Sscanf parse the arguments according to a
	format string, analogous to that of Printf. In the text that
	follows, 'space' means any Unicode whitespace character
	except newline.

	In the format string, a verb introduced by the % character
	consumes and parses input; these verbs are described in more
	detail below. A character other than %, space, or newline in
	the format consumes exactly that input character, which must
	be present. A newline with zero or more spaces before it in
	the format string consumes zero or more spaces in the input
	followed by a single newline or the end of the input. A space
	following a newline in the format string consumes zero or more
	spaces in the input. Otherwise, any run of one or more spaces
	in the format string consumes as many spaces as possible in
	the input. Unless the run of spaces in the format string
	appears adjacent to a newline, the run must consume at least
	one space from the input or find the end of the input.

	The handling of spaces and newlines differs from that of C's
	scanf family: in C, newlines are treated as any other space,
	and it is never an error when a run of spaces in the format
	string finds no spaces to consume in the input.

	The verbs behave analogously to those of Printf.
	For example, %x will scan an integer as a hexadecimal number,
	and %v will scan the default representation format for the value.
	The Printf verbs %p and %T and the flags # and + are not implemented.
	The verbs %e %E %f %F %g and %G are all equivalent and scan any
	floating-point or complex value. For float and complex literals in
	scientific notation, both the decimal (e) and binary (p) exponent
	formats are supported (for example: "2.3e+7" and "4.5p-8").

	Input processed by verbs is implicitly space-delimited: the
	implementation of every verb except %c starts by discarding
	leading spaces from the remaining input, and the %s verb
	(and %v reading into a string) stops consuming input at the first
	space or newline character.

	The familiar base-setting prefixes 0 (octal) and 0x
	(hexadecimal) are accepted when scanning integers without
	a format or with the %v verb.

	Width is interpreted in the input text but there is no
	syntax for scanning with a precision (no %5.2f, just %5f).
	If width is provided, it applies after leading spaces are
	trimmed and specifies the maximum number of runes to read
	to satisfy the verb. For example,
	   Sscanf(" 1234567 ", "%5s%d", &s, &i)
	will set s to "12345" and i to 67 while
	   Sscanf(" 12 34 567 ", "%5s%d", &s, &i)
	will set s to "12" and i to 34.

	In all the scanning functions, a carriage return followed
	immediately by a newline is treated as a plain newline
	(\r\n means the same as \n).

	In all the scanning functions, if an operand implements method
	Scan (that is, it implements the Scanner interface) that
	method will be used to scan the text for that operand.  Also,
	if the number of arguments scanned is less than the number of
	arguments provided, an error is returned.

	All arguments to be scanned must be either pointers to basic
	types or implementations of the Scanner interface.

	Like Scanf and Fscanf, Sscanf need not consume its entire input.
	There is no way to recover how much of the input string Sscanf used.

	Note: Fscan etc. can read one character (rune) past the input
	they return, which means that a loop calling a scan routine
	may skip some of the input.  This is usually a problem only
	when there is no space between input values.  If the reader
	provided to Fscan implements ReadRune, that method will be used
	to read characters.  If the reader also implements UnreadRune,
	that method will be used to save the character and successive
	calls will not lose data.  To attach ReadRune and UnreadRune
	methods to a reader without that capability, use
	bufio.NewReader.
*/
```



### Append

*切片添加元素*

```go
func append(slice []T, elements ...T) []T
```



## Initialization/ 初始化

// @TODO ...

