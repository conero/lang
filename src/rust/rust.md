# Rust

> A safe, concurrent, practical language(安全、并行、实用)



*参考地址*：

- [https://github.com/rust-lang/rust](https://github.com/rust-lang/rust)
- [[https://www.rust-lang.org](https://www.rust-lang.org/)]



> **Rust** 是一种系统编程语言。 它有着惊人的运行速度，能够防止段错误，并保证线程安全。 



> 特点

- 零开销抽象
- 转移语义
- 保证内存安全
- 线程无数据竞争
- 基于 trait 的泛型
- 模式匹配
- 类型推断
- 极小运行时
- 高效 C 绑定 



*在 Rust 中引入并行是相对低风险的操作：编译器会为你捕获经典的错误。*

*在 Rust 中，变量默认是不可变的*



> rustc uses LLVM as its back end. 使用 **LLVM** 作为后端



> 家谱图

![](../image/rust-familys.png)



> 方向

- *底层系统编程*
- *CLI Application*
- *Web Server*



> Rust 工具

- *cargo*                 依赖管理器和构建工具
- *rustfmt*              代码格式化
- *Rust Language Server*
- *rustup*               rust 安装工具



> window 使用 **rustup/rust**

```powershell
# 更新 rust
$ rustup update 
# 卸载 rust/rustup
$ rust self uninstall

# 查看版本
$ rustc --version
$ cargo --version
```





## 基础学习

> HelloWorld

```rust
// main.rs
// Rust 的缩进风格使用 4 个空格，而不是 1 个制表符（tab）
// 运行: >rustc main.rs
fn main(){
    println!("Hello, world!");
}
```




> 变量/常量

*变量默认是不可改变的（immutable）; 常量（constants）不光默认不能变，它总是不能变。*

*常量可以在任何作用域中声明，包括全局作用域*

*定义一个与之前变量同名的新变量，而新变量会 **隐藏** 之前的变量。Rustacean 们称之为第一个变量被第二个 **隐藏** 了，这意味着使用这个变量时会看到第二个值。*

*`mut` 与隐藏的一个区别是，当再次使用 `let` 时，实际上创建了一个新变量，我们可以改变值的类型，但复用这个名字。*

```rust
// immutable
let x = 5;
// mutable
let mut y = 8;
// constant
const MAX_POINTS: u32 = 100_000;

//  隐藏/ shadowing; x = 25
let x = x * x;
```



### 数据类型

> date type

- 数据类型
  - 标量（scalar）
    - 整型 (int)
      - i32      *默认*
    - 浮点型 (float)
      - f32
      - f64      *默认*
    - 布尔类型(bool)
      - `bool: true/false`
    - 字符类型(char)
  - 复合（compound)
    - 元组 (tuple)
    - 数组 (array)



*字符类型: `char` 由单引号指定，不同于字符串使用双引号。*

*元组是一个将多个其他类型的值组合进一个复合类型的主要方式。使用 `.` 访问对应的索引、模式匹配解构元组值*

*Rust 中的数组是固定长度的：一旦声明，它们的长度不能增长或缩小。元组不同，数组中的每个元素的类型必须相同。*




> 整型 （8,16,32,64）

| 长度   | 有符号  | 无符号  |
| ------ | ------- | ------- |
| 8-bit  | `i8`    | `u8`    |
| 16-bit | `i16`   | `u16`   |
| 32-bit | `i32`   | `u32`   |
| 64-bit | `i64`   | `u64`   |
| arch   | `isize` | `usize` |



> 数字字面值

| 数字字面值       | 例子          |
| ---------------- | ------------- |
| Decimal          | `98_222`      |
| Hex              | `0xff`        |
| Octal            | `0o77`        |
| Binary           | `0b1111_0000` |
| Byte (`u8` only) | `b'A'`        |



> 数值运算

`+ 加, - 减, * 乘, / 除, % 取余`



### 注释

- 普通注释

- *documentation comments/文档注释*



`// 单行注释 ` 单行注释

`/// 文档注释 `    *它们会生成 HTML文档，支持 Markdown 注解来格式化文本。`$ cargo doc` 可生成文档， `$ cargo doc --open` 会构建当前 crate 文档（同时还有所有 crate 依赖的文档）的 HTML 并在浏览器中打开*





### 函数(fn)

*`main` 函数，它是很多程序的入口点。你也见过 `fn` 关键字，它用来声明新函数。*

*Rust 代码中的函数和变量名使用 __snake case__ 规范风格。在 __snake case__ 中，所有字母都是小写并使用下划线分隔单词。*

*Rust 是一门基于表达式（expression-based）的语言，这是一个需要理解的（不同于其他语言）重要区别。*

_**语句**（*Statements*）是执行一些操作但不返回值的指令。表达式（*Expressions*）计算并产生一个值。_

*我们用来创建新作用域的大括号（代码块），`{}`，也是一个表达式*



> 参数

*参数是特殊变量，是函数签名的一部分。当函数拥有参数（形参）时，可以为这些参数提供具体的值（实参）*



### 控制流

*Rust 代码中最常见的用来控制执行流的结构是 `if` 表达式和循环。*

_代码中的条件 **必须** 是 `bool` 值。如果条件不是 `bool` 值，我们将得到一个错误。_



**条件控制**

> `if - else if - else`    多重条件



**循环控制**

*Rust 有三种循环：`loop`、`while` 和 `for`。*

- `loop` 重复执行
- `while` 条件循环
- `for` 遍历循环



//@TODO   *[所有权](https://kaisery.github.io/trpl-zh-cn/ch04-00-understanding-ownership.html)*

## cargo/ 包管理器 

> *Rust 的构建系统和包管理器。 采用 Toml 格式的配置文件*

```shell
$ cargo --version
# 利用 cargo 创建项目
$ cargo new <project>

# 构建项目
$ cargo build
# 优化编译项目/发布时
$ cargo build --release

# 编译并运行生成的可执行文件
$ cargo run

# 快速检查代码确保其可以编译， 比 build 快
$ cargo check
```





## rustdoc/ 文档生成器





## rustup 工具链

> 升级 rust命令

``rustup update``



## 历史

- 2006      *Graydon Hoare*     ``personal project``
  - The language grew out of a personal project started in 2006 by Mozilla employee Graydon Hoare
- 2009    **Mozilla** *began sponsoring the project* 
- 2010    **Mozilla** announced it
  - *work shifted from the initial compiler (written in ``OCaml``) to the self-hosting compiler written in Rust. Named rustc, it successfully compiled itself in 2011*
- *2012.01.21  release v0.1*
- 2015.05.16   release **v1.0**





### Rust 2015



### Rust 2018

> Rust 1.29-1.31



> Rust 2015



## 其他语言

### GoLang

**golang**

>  定位: 位解决分布式系统、服务器应用开发
>
> vs: Java/Python/Ruby



**rust**

>  定位: 解决单机安全问题、高性能场景偏系统底层开发
>
> vs: C/C++





## 附录

### 参考

- [Rust 程序设计语言-简体中文](https://kaisery.github.io/trpl-zh-cn/)
- [cargo-doc](https://doc.rust-lang.org/cargo/)