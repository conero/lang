## Zig

> Joshua Conero
>
> 2025年1月21日





ziglang 始于2016由  Andrew Kelley 开发，基于 LLVM 后端



git 仓库：https://codeberg.org/ziglang/zig   （2025/11/26 由 https://github.com/ziglang/zig 迁入）



### 命令

```shell
# zig 项目初始化
zig init

# 查看zig标准库文档
zig std

# zig 代码格式化
zig fmt ..
zig fmt ./..
```



zig 交叉编译

```powershell
# linux
zig build '-Dtarget=x86_64-linux-musl' --release=safe
zig build '-Dtarget=x86_64-linux-gnu'
zig build '-Dtarget=x86_64-linux'
zig build '-Dtarget=aarch64-linux'

# macos
zig build '-Dtarget=x86_64-macos'

# windows
zig build '-Dtarget=x86_64-windows'
zig build '-Dtarget=x86_64-windows-gnu'

# wasi 编码
zig build '-Dtarget=wasm32-wasi-musl'
```





Windows 下命令行乱码的处理方法

```powershell
# powershell
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

# 查看命令行编码
chcp
# 设置名行编码
# utf-8
chcp 65001
# gbk
chcp 936
```





### 语法

> **特性**

- 支持泛型，具有强大的编译时元编程功能
- 以 `@` 开头的函数是内置函数。它们是由编译器提供的，而不是标准库提供的
- Zig 支持任意位宽度的整数
- 函数参数是常量
- 数组：使用 `_` 让编译器推导长度
- 字符串是字节（u8）的序列（即数组或切片）
  - 字符串字面形式类型: `[]const u8`
- 自动进行类型转化
- 类型是 zig 的一等公民



> **注释**

- Zig 没有像 C 语言中类似 `/* ... */` 的多行注释。
- `//!` 被称为顶级文档注释，可以放在文件的顶部。
- 三斜线注释 (`///`) 被称为文档注释，可以放在特定位置，如声明之前。





#### 类型

```c
//
var variable:type = value;

// 常量
const variable = value;

//解构赋值
var z: u32 = undefined;
// var z: u32 = undefined;
const x, var y, z = [3]u32{ 1, 2, 3 }; // 数组


// 字符串的写法
const message_1 = "hello";
const message_2 = [_]u8{ 'h', 'e', 'l', 'l', 'o' };
const message_3: []const u8 = &.{ 'h', 'e', 'l', 'l', 'o' };
```



- 基本类型
  - 数字
  - 布尔 - boolean
  - 字符
- 复合类型
  - 字符串 u8
  - 元组（tuple） .{}
  - 结构体（struct） strcut
  - 联合  union
  - 数组 [_]T
  - 切片 []T
  - 枚举  enum
- 其他
  - 向量
  - 指针
  - 零位类型    void






使用 `undefined` 使变量保持未初始化状态。

**字符串字面量**既可以隐式转换为 u8 切片，也可以隐式转换为以 0 结尾的指针。

实际上枚举仅仅是一种命名空间



**struct**

结构体允许使用默认值





#### 语句

- 判断
  - if else
  - switch
- 循环，支持 continue/break
  - for
  - while
- 内联   inline
- defer  延期执行，与 go 语言中类似
- 可选类型   ?T   表 null or T







**三元表达式**

```c
const a: u32 = 5;
const b: u32 = 4;
// 下方 result 的值应该是47
const result = if (a != b) 47 else 3089;
```





#### 函数

- 函数的参数是不可变的（函数参数是常量）。

- 原始类型（整型、布尔这种）传递完全使用值传递，像复合类型（结构体、联合、数组等），这些传递均是由编译器来决定究竟是使用“值传递”还是“引用传递”。

- 不支持函数重载



**内建函数** 由编译器提供而非标准库，并以 @ 为前缀。

**编译器** 运行，使用 comptime 修饰，其在编译器是必须是已知的。



#### 指针

指针是指向一块内存区域地址的变量，它存储了一个地址。

```shell
# 指向 x 的指针
*x
# 指向常量 x 的指针，用来读取x
* const x
# 从 *X 到 *const X 的转换是隐式的，因为这是安全的。
# 从 *const X 到 *X 的转换需要显式执行，并且应当谨慎对待，以避免意外地修改原本应保持不变的数据。
```





### 内存管理

采取了类似 C 的方案，完全由程序员管理内存。





### 基础库

#### 格式化输出

数据格式化输出，如下：

```c
const std = @import("std");

std.debug.print("{s}", "中文")
```

格式化参数规则：`{[argument][specifier]:[fill][alignment][width].[precision]}`

参数规则，

- *fill*（填充字符）是单个 Unicode 码点，用于填充格式化后的文本
- *alignment*（对齐方式）是 `<`、`^`、`>` 三个字符之一，分别表示文本左对齐、居中对齐、右对齐
- *width*（宽度）是字段以 Unicode 码点计的总宽度
- *precision*（精度）指定格式化数字应保留的小数位数



- 类型相关的格式化选项
  - `x` 和 `X`  十六进制是输出格式（HEX）
  - `s`  字符串输出
  - `e`  以科学计数法输出浮点数值
  - `d` 以十进制表示法输出数值
  - `b` 以二进制表示法输出整数值
  - `o` 以八进制表示法输出整数值
  - `c` 将整数作为 ASCII 字符输出。整数类型最多只能有 8 位
  - `u` 将整数作为 UTF-8 序列输出。整数类型最多只能有 21 位。
  - `?`：输出可选值，要么是解包后的值，要么是 `null`；后面可以跟底层值的格式说明符。
  - `!`：输出错误联合值，要么是解包后的值，要么是格式化后的错误值；后面可以跟底层值的格式说明符。
  - `*`：输出值的地址，而不是值本身。
  - `any`：使用默认格式输出任意类型的值。







### 附录

- [Zig 语言圣经](https://course.ziglang.cc/)
- [Zig 语言中文社区](https://ziglang.cc/)
- [Zig语言中文手册](https://sxwangzhiwen.github.io/zigcndoc/zigcndoc.html)
- [pedropark99/zig-book](https://github.com/pedropark99/zig-book) 英文版zig入门手册，https://pedropark99.github.io/zig-book/



