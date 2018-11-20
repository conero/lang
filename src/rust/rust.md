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

- __不能在相同作用域中同时存在可变和不可变引用的规则。__
- __Rust 倾向于确保暴露出可能的错误__
- __字符串字面值不可变__
  - 字符串值(str)被硬编码进程序里。字符串字面值是很方便的，不过他们并不适合使用文本的每一种场景。原因之一就是他们是不可变的
- __Rust 永远也不会自动创建数据的 “深拷贝”。因此，任何 **自动** 的复制可以被认为对运行时性能影响较小。__



*实现变量时：注意自始到终的顺序*



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



> 使用元组来返回多个值

```rust
fn calculate_length(s: String) -> (String, usize) {
    let length = s.len(); // len() 返回字符串的长度

    (s, length)
}
```



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



### 所有权



*所有权（系统）是 Rust 最独特的功能，其令 Rust 无需垃圾回收（garbage collector）即可保障内存安全。Rust 的核心功能（之一）是 **所有权**（*ownership*）*



*所有运行的程序都必须管理其使用计算机内存的方式。*

- *`垃圾回收机制`       在程序运行时不断地寻找不再使用的内存: Java/Go等*
- *`程序员必须亲自分配和释放内存`:  C/C++等*
- *`通过所有权系统管理内存，编译器在编译时会根据一系列的规则进行检查`:  Rust等*



*内存与分配： Rust 采取了一个不同的策略：内存在拥有它的变量离开作用域后就被自动释放。*

```rust
// 创建可变字符串
let mut s = String::from("Joshua Conero is");

s.push_str(" learning Rust.");

// Joshua Conero is learning Rust.
println!("{}", s);

```





> 所有权规则

1. Rust 中的每一个值都有一个被称为其 **所有者**（*owner*）的变量。
2. 值有且只有一个所有者。
3. 当所有者（变量）离开作用域，这个值将被丢弃。



```rust
{
    let s = String::from("hello"); // 从此处起，s 是有效的
    // 使用 s
}                                  // 此作用域已结束，
                                   // s 不再有效
// 说明：
//  这是一个将 String 需要的内存返回给操作系统的很自然的位置：当 s 离开作用域的时候。当变量离开作用域，Rust 为我们调用一个特殊的函数。这个函数叫做 drop，在这里 String 的作者可以放置释放内存的代码。Rust 在结尾的 } 处自动调用 drop。
```



> 变量与数据交互的方式:

- 移动
- 克隆

```rust
let x = 5;
let y = x;

//  s1 被 移动 到了 s2 中
let s1 = String::from("hello");
let s2 = s1;

println!("{}, world!", s1);		// error[E0382]: use of moved value: `s1`

// 当我们将 s1 赋值给 s2，String 的数据被复制了，这意味着我们从栈上拷贝了它的指针、长度和容量。我们并没有复制指针指向的堆上数据。
// 这就有了一个问题：当 s2 和 s1 离开作用域，他们都会尝试释放相同的内存。这是一个叫做 二次释放（double free）的错误，也是之前提到过的内存安全性 bug 之一。两次释放（相同）内存会导致内存污染，它可能会导致潜在的安全漏洞。
// 与其尝试拷贝被分配的内存，Rust 则认为 s1 不再有效，因此 Rust 不需要在 s1 离开作用域后清理任何东西
```

> String 克隆

```rust
let s1 = String::from("hello");
let s2 = s1.clone();

println!("s1 = {}, s2 = {}", s1, s2);
```

> 只在栈上的数据： 拷贝

如类型:  整形、布尔、浮点型、字符型(char)、元组



> 所有权与函数

*将值传递给函数在语义上与给变量赋值相似。向函数传递值可能会移动或者复制，就像赋值语句一样。*

```rust
fn main() {
    let s = String::from("hello");  // s 进入作用域

    takes_ownership(s);             // s 的值移动到函数里...
                                    // ... 所以到这里不再有效

    let x = 5;                      // x 进入作用域

    makes_copy(x);                  // x 应该移动函数里，
                                    // 但 i32 是 Copy 的，所以在后面可继续使用 x

} // 这里, x 先移出了作用域，然后是 s。但因为 s 的值已被移走，
  // 所以不会有特殊操作

fn takes_ownership(some_string: String) { // some_string 进入作用域
    println!("{}", some_string);
} // 这里，some_string 移出作用域并调用 `drop` 方法。占用的内存被释放

fn makes_copy(some_integer: i32) { // some_integer 进入作用域
    println!("{}", some_integer);
} // 这里，some_integer 移出作用域。不会有特殊操作
```



> 返回值与作用域

*返回值也可以转移所有权。*

```rust
fn main() {
    let s1 = gives_ownership();         // gives_ownership 将返回值
                                        // 移给 s1

    let s2 = String::from("hello");     // s2 进入作用域

    let s3 = takes_and_gives_back(s2);  // s2 被移动到
                                        // takes_and_gives_back 中, 
                                        // 它也将返回值移给 s3
} // 这里, s3 移出作用域并被丢弃。s2 也移出作用域，但已被移走，
  // 所以什么也不会发生。s1 移出作用域并被丢弃

fn gives_ownership() -> String {             // gives_ownership 将返回值移动给
                                             // 调用它的函数

    let some_string = String::from("hello"); // some_string 进入作用域.

    some_string                              // 返回 some_string 并移出给调用的函数
}

// takes_and_gives_back 将传入字符串并返回该值
fn takes_and_gives_back(a_string: String) -> String { // a_string 进入作用域

    a_string  // 返回 a_string 并移出给调用的函数
}
```



*__变量的所有权总是遵循相同的模式：将值赋给另一个变量时移动它。当持有堆中数据值的变量离开作用域时，其值将通过 `drop` 被清理掉，除非数据被移动为另一个变量所有。__*



#### 引用与借用

_`&` 符号就是 **引用**，它们允许你使用值但不获取其所有权。与使用 `&` 引用相反的操作是 **解引用**（*dereferencing*），它使用解引用运算符，`*`。_



_函数签名使用 `&` 来表明参数 `s` 的类型是一个引用。让我们增加一些解释性的注释：_

```rust
fn calculate_length(s: &String) -> usize { // s 是对 String 的引用
    s.len()
} // 这里，s 离开了作用域。但因为它并不拥有引用值的所有权，
  // 所以什么也不会发生
```

_变量 `s` 有效的作用域与函数参数的作用域一样，不过当引用离开作用域后并不丢弃它指向的数据，因为我们没有所有权。当函数使用引用而不是实际值作为参数，无需返回值来交还所有权，因为就不曾拥有所有权。_

_我们将获取引用作为函数参数称为 **借用**（*borrowing*）。正如现实生活中，如果一个人拥有某样东西，你可以从他那里借来。当你使用完毕，必须还回去。_



> 可变引用

*可变引用有一个很大的限制：在特定作用域中的特定数据有且只有一个可变引用。*



> 引用的规则

- 在任意给定时间，**要么** 只能有一个可变引用，**要么** 只能有多个不可变引用。
- 引用必须总是有效



#### slice

_另一个没有所有权的数据类型是 *slice*。slice 允许你引用集合中一段连续的元素序列，而不用引用整个集合。_



> `slice 字符串`

```rust
let s = String::from("hello world");

let hello = &s[0..5];
// 如果想要从第一个索引（0）开始，可以不写两个点号之前的值。
let hello = &s[..5];
let world = &s[6..11];

let len = s.len();
// 如果 slice 包含 String 的最后一个字节，也可以舍弃尾部的数字。
let slice = &s[3..len];
let slice = &s[3..];

// 也可以同时舍弃这两个值来获取整个字符串的 slice。
let slice = &s[..];
```

_`start..end` 语法代表一个以 `start` 开头并一直持续到但不包含 `end` 的 range。_



*字符串字面值就是 slice*

```rust
let s = "Joshua Conero";
```

这里 `s` 的类型是 `&str`：它是一个指向二进制程序特定位置的 slice。这也就是为什么字符串字面值是不可变的；`&str` 是一个不可变引用。



> 其他类型的 slice

```rust
let a = [1, 2, 3, 4, 5];
let slice = &a[1..3];
```





#### 堆 stack/栈 heap

*栈和堆都是代码在运行时可供使用的内存，但是它们的结构不同。*



*栈 __后进先出__(last in, first out), 有 __进栈__ 和 __出栈__ 操作。*

_栈的操作是十分快速的，这主要是得益于它存取数据的方式：因为数据存取的位置总是在栈顶而不需要寻找一个位置存放或读取数据。另一个让操作栈快速的属性是，栈中的所有数据都必须占用已知且固定的大小。_



_`堆` 在编译时大小未知或大小可能变化的数据，要改为存储在堆上。堆是缺乏组织的：当向堆放入数据时，你要请求一定大小的空间。操作系统在堆的某处找到一块足够大的空位，把它标记为已使用，并返回一个表示该位置地址的 **指针**（*pointer*）。这个过程称作 **在堆上分配内存**（*allocating on the heap*），有时简称为 “分配”（allocating）。将数据推入栈中并不被认为是分配。因为指针的大小是已知并且固定的，你可以将指针存储在栈上，不过当需要实际数据时，必须访问指针。_



### 结构体

*struct*，或者 *structure*，是一个自定义数据类型，允许你命名和包装多个相关的值，从而形成一个有意义的组合。

*和元组一样，结构体的每一部分可以是不同类型。但不同于元组，结构体需要命名各部分数据以便能清楚的表明其值的意义。由于有了这些名字，结构体比元组更灵活：不需要依赖顺序来指定或访问实例中的值。*

```rust
// 创建结构体
struct User{
    name: String,    
    sex: String,
    age: i32,
    logined: bool,
};

// 定义 User 结构体
let jc = User {
    name: String::from("Joshua Conero")，
    sex: String::from("M"),
    age: 22,
    logined: true,
};

println!("age: {}", jc.age);


// 定义 User 结构体
let jc = User {
    name: String::from("Joshua Conero")，
    sex: String::from("M"),
    age: 22,
    logined: true,
};

// 变量与字段同名时的字段初始化简写语法
fn build_user(name: String, sex: String) -> User {
    User {
        name,
        sex,
        age: 0,
    	logined: true,
    }
}

// 使用结构体更新语法从其他实例创建实例
let mut sr = User {
    name: String::from("Susanna Jahn"),
    sex: "F",
    age: jc.age,
    ..jc
}
// 打印 sr 变量
println!("sr is {:?}", sr);
```



> 元组结构体

_也可以定义与元组类似的结构体，称为 **元组结构体**（*tuple structs*）。元组结构体有着结构体名称提供的含义，但没有具体的字段名，只有字段的类型。当你想给整个元组取一个名字，并使元组成为与其他元组不同的类型时，元组结构体是很有用的，这时像常规结构体那样为每个字段命名就显得多余和形式化了。_

```rust
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);

let black = Color(0, 0, 0);
let origin = Point(0, 0, 0);
```



> 方法

_**方法** 与函数类似：它们使用 `fn` 关键字和名称声明，可以拥有参数和返回值，同时包含在某处调用该方法时会执行的代码。不过方法与函数是不同的，因为它们在结构体的上下文中被定义（或者是枚举或 trait 对象的上下文），并且它们第一个参数总是 `self`，它代表调用该方法的结构体实例。_



_**关联函数**             `impl` 块的另一个有用的功能是：允许在 `impl` 块中定义 **不** 以 `self` 作为参数的函数。这被称为 **关联函数**（*associated functions*），因为它们与结构体相关联。它们仍是函数而不是方法，因为它们并不作用于一个结构体的实例。_



```rust
// 创建结构体
struct User{
    name: String,    
    sex: String,
    age: i32,
    logined: bool,
};

//  impl 块（impl 是 implementation 的缩写）
impl User{
    // 定义方法
    // 第一个参数为 self， 与 python 相似
    fn toString(self){
        println!("The guy, {}. sex: {}, age {}. logined status like {}", 
        	self.name, self.sex, self.age. self.logined
        );
    }
    // &self 方法可以选择获取 self 的所有权
    fn addAge(&self, age i32){
        self.age = age;
    }
    // 关联函数经
    // 关联函数经常被用作返回一个结构体新实例的构造函数
    fn male(name: String) -> User{
        User {
            name,
            sex: String::from("M"),
            age: 0,
            logined: false,
        }
    }
}

fn main(){
    let jc = User {
        name: String::from("Joshua Conero")，
        sex: String::from("M"),
        age: 22,
        logined: true,
    };
    // 调用方法
    jc.addAge(1);
    jc.toString();
    
    // 调用关联函数
    let mut jc2 = User::male(String::from("Joshua Conero"));
    jc2.age = 22;
    jc2.toString();
}
```



### 枚举

**枚举**（*enumerations*），也被称作 *enums*。枚举允许你通过列举可能的值来定义一个类型。

枚举是一个很多语言都有的功能，不过不同语言中其功能各不相同。Rust 的枚举与 F#、OCaml 和 Haskell 这样的函数式编程语言中的 **代数数据类型**（*algebraic data types*）最为相似。



```rust
// 方向
// top 等称为枚举的成员
enum Direction{
    top, bottom,right,left,
}

// 访问其成员
Direction::top;
// 标题栏
struct titleBar{
    dct: Direction,
    text: String,
}

let topTitleBar = titleBar {
    dct: Direction::top,
    text: String::from("welcome to the Rust World."),
}


// 性别
enum Gender{
    M(String),
    F(String)
}
let man = Gender::M(String::from("man"));
let woman = Gender::F(String::from("woman"));

// 枚举其他
enum IpAddr {
    //name(type)
    V4(u8, u8, u8, u8),
    V6(String),
}

let home = IpAddr::V4(127, 0, 0, 1);
let loopback = IpAddr::V6(String::from("::1"));
```



#### option-枚举

_`Option` 是标准库定义的另一个枚举。`Option` 类型应用广泛因为它编码了一个非常普遍的场景，即一个值要么有值要么没值。_



_关于**空值 / null**。 问题不在于概念而在于具体的实现。为此，Rust 并没有空值，不过它确实拥有一个可以编码存在或不存在概念的枚举。这个枚举是 `Option<T>`，而且它[定义于标准库中](https://doc.rust-lang.org/std/option/enum.Option.html)，如下:_

```rust
enum Option<T> {
    Some(T),
    None,
}
```



#### match-控制流运算符

_Rust 有一个叫做 `match` 的极为强大的控制流运算符，它允许我们将一个值与一系列的模式相比较并根据相匹配的模式执行相应代码。模式可由字面值、变量、通配符和许多其他内容构成。_

```rust
// 使用 枚举 做模式匹配
enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter,
}

fn value_in_cents(coin: Coin) -> u32 {
    match coin {
        Coin::Penny => {
            println!("Lucky penny!");
            1
        },
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter => 25,
    }
}
```



> 绑定值的模式

*匹配分支的另一个有用的功能是可以绑定匹配的模式的部分值。这也就是如何从枚举成员中提取值的。*

```rust
enum UsState {
    Alabama,
    Alaska,
    // ... etc
}

enum Coin {
    Penny,
    Nickel,
    Dime,
    Quarter(UsState),
}

fn value_in_cents(coin: Coin) -> u32 {
    match coin {
        Coin::Penny => 1,
        Coin::Nickel => 5,
        Coin::Dime => 10,
        Coin::Quarter(state) => {
            println!("State quarter from {:?}!", state);
            25
        },
    }
}
```



> 匹配 Option<T>

**匹配是穷尽的**

```rust
// 空值配置
fn plus_one(x: Option<i32>) -> Option<i32> {
    match x {
        // 匹配空
        None => None,
        // 非空
        Some(i) => Some(i + 1),
    }
}

let five = Some(5);
let six = plus_one(five);
let none = plus_one(None);


// _通配符
let some_u8_value = 0u8;
match some_u8_value {
    1 => println!("one"),
    3 => println!("three"),
    5 => println!("five"),
    7 => println!("seven"),
    // 通配符
    _ => (),
}
```



#### if-let-简单控制流

_`if let` 获取通过 `=` 分隔的一个模式和一个表达式。它的工作方式与 `match` 相同，这里的表达式对应 `match` 而模式则对应第一个分支。换句话说，可以认为 `if let` 是 `match` 的一个语法糖，它当值匹配某一模式时执行代码而忽略所有其他值。_

```rust
let some_u8_value = Some(0u8);
match some_u8_value {
    Some(3) => println!("three"),
    _ => (),
}

// 前后等价
if let Some(3) = some_u8_value {
    println!("three");
}
```



### 模块/module

**模块**（*module*）是一个包含函数或类型定义的命名空间，你可以选择这些定义能（公有）或不能（私有）在其模块外可见。下面是一个模块如何工作的梗概：

- 使用 `mod` 关键字声明新模块。此模块中的代码要么直接位于声明之后的大括号中，要么位于另一个文件。
- 函数、类型、常量和模块默认都是私有的。可以使用 `pub` 关键字将其变成公有并在其命名空间之外可见。
- `use` 关键字将模块或模块中的定义引入到作用域中以便于引用它们。



_因为枚举也像模块一样组成了某种命名空间，也可以使用 `use` 来导入枚举的成员。对于任何类型的 `use`语句，如果从一个命名空间导入多个项，可以在最后使用大括号和逗号来列举它们_

_为了一次将某个命名空间下的所有名称都引入作用域，可以使用 `*` 语法，这称为 **glob 运算符**（*glob operator*）。_

_使用-super-访问父模块_

```rust
pub mod A{
    pub fn test(){}    
    
    pub mod B{
        pub fn test(){
            // 访问父模块
            super::test();
        }
    }
    
    // 私有模式
    mod inner{}
    
    // 外部可访问的方法
    pub fn name(){}
    pub mod F{
        pub mod FF{
            //
        }
    }
    pub mod I{}
}

use A::B;
use A::test;
use A::{F, I};  // 如果从一个命名空间导入多个项，可以在最后使用大括号和逗号来列举它们
use A::F::*;    // 为了一次将某个命名空间下的所有名称都引入作用域，可以使用 * 语法，这称为 glob 运算符（glob operator）。
fn main(){
    A::test();
    A::B::test();
    B::test();
    test();
}
```



> mod-和文件系统

__Rust 默认只知道 *src/lib.rs* 中的内容。如果想要对项目加入更多文件，我们需要在 *src/lib.rs* 中告诉 Rust 去寻找其他文件。__



**模块文件系统规则：**

- 如果一个叫做 `foo` 的模块没有子模块，应该将 `foo` 的声明放入叫做 *foo.rs* 的文件中。
- 如果一个叫做 `foo` 的模块有子模块，应该将 `foo` 的声明放入叫做 *foo/mod.rs* 的文件中。



*读取模块信息： `Rust -> src/lib.rs -> ~/mod.rs`*



**私有性规则**

1. 如果一个项是公有的，它能被任何父模块访问
2. 如果一个项是私有的，它能被其直接父模块及其任何子模块访问



### 通用集合类型

> 查看      `std::collections`

_Rust 标准库中包含一系列被称为 **集合**（*collections*）的非常有用的数据结构。大部分其他数据类型都代表一个特定的值，不过集合可以包含多个值。不同于内建的数组和元组类型，这些集合指向的数据是储存在堆上的，这意味着数据的数量不必在编译时就已知并且可以随着程序的运行增长或缩小。_

*集合存于__堆(heap)__上*



*常用的集合类型：*

- Vector                                     *允许我们一个挨着一个地储存一系列数量可变的值*
- 字符串 (string)                        *是字符的集合*
- 哈希 Map (hash map)           *k-v 数据键值对*



#### Vector

`Vec<T>`   *vector 允许我们在一个单独的数据结构中储存多于一个的值，它在内存中彼此相邻地排列所有的值。vector 只能储存相同类型的值。vector 是用泛型实现的。*



```rust
// 新建 Vector
// 注意这里我们增加了一个类型注解。因为没有向这个 vector 中插入任何值，Rust 并不知道我们想要储存什么类型的元素。这是一个非常重要的点。
let x: Vec<i32> = Vec::new();

// 使用宏定义 Vector；使用初始值创建宏
let y = vec![1, 2, 3];


// 更新 Vector
// 1. 使用 push 方法向 vector 增加值
x.push(5);
x.push(6);
y.push(99);

// 读取值
let x1: &i32 = &x[0];
let y1: Option<&i32> = y.get(0);

// 遍历
for i in &y{
    println!("{}", i);
}
```



#### 字符串

- `str` 字符串字面值，硬编码的方式，不可变。

- _`String` 类型，为了支持一个可变，可增长的文本片段，需要在堆上分配一块在编译时未知大小的内存来存放内容。_

_Rust 的核心语言中只有一种字符串类型：`str`，字符串 slice，它通常以被借用的形式出现，`&str`。_



```rust
// 新建字符串(空字符)
let mut s = String::new();

// 带初始值字符串
let data = "initial contents";
let s = data.to_string();
let s = String::from("initial contents");

// the method also works on a literal directly:
let s = "initial contents".to_string();
```



> 更新字符串

- 使用-push-附加字符串
- 使用--运算符或-format-宏连接字符串



_`String` 的大小可以增长其内容也可以改变，就像可以放入更多数据来改变 `Vec` 的内容一样。另外，`String` 实现了 `+` 运算符作为连接运算符以便于使用。_

```rust
let mut s = String::from("foo");
s.push_str("bar");
// s => "foo bar"

let mut s1 = String::from("Joshu");
s1.push('a');
// s1 -> "Joshua"
```



使用`+` 符号更新字符串

```rust
let s1 = String::from("Hello");
let s2 = String::from("world!");

let s3 = s1 + " " + &s2; // note s1 has been moved here and can no longer be used
let s3 = format!("{} {}", s1, s2);

// 遍历字符串
for b in s1.bytes(){
    println!("{}", b);
}
```



#### 哈希 map

_`HashMap<K, V>` 类型储存了一个键类型 `K` 对应一个值类型 `V` 的映射。它通过一个 **哈希函数**（*hashing function*）来实现映射，决定如何将键和值放入内存中。_

_同样类似于 vector，哈希 map 是同质的：所有的键必须是相同类型，值也必须都是相同类型。_



*新建哈希 map,通过 insert 更新值。可以通过 `get` 方法并提供对应的键来从哈希 map 中获取值。*

```rust
use std::collections::HashMap;
// method 1
let mut scores = HashMap::new();

scores.insert(String::from("Blue"), 10);
scores.insert(String::from("Yellow"), 50);
// 读取值
let key = "Blue".to_string();
println!("{}", scores.get(&key));	// => 10
// 循环, map
for (k, v) in &scores {
    println!("{}: {}", k, v);
}
println!("{:?}", scores);
// 检测 scores 是否存在 "Red" 键值，不存在就新增子 => Red ~ 50
scores.entry(String::from("Red")).or_insert(50);



// method 2
let teams  = vec![String::from("Blue"), String::from("Yellow")];
let initial_scores = vec![10, 50];

let scores: HashMap<_, _> = teams.iter().zip(initial_scores.iter()).collect();
```



### 错误处理

_Rust 将错误组合成两个主要类别：**可恢复错误**（*recoverable*）和 **不可恢复错误**（*unrecoverable*）。可恢复错误通常代表向用户报告错误和重试操作是合理的情况，比如未找到文件。不可恢复错误通常是 bug 的同义词，比如尝试访问超过数组结尾的位置。_



>  **Panic 中的栈展开与终止**

当出现 `panic!` 时，程序默认会开始 **展开**（*unwinding*），这意味着 Rust 会回溯栈并清理它遇到的每一个函数的数据，不过这个回溯并清理的过程有很多工作。另一种选择是直接 **终止**（*abort*），这会不清理数据就退出程序。那么程序所使用的内存需要由操作系统来清理。如果你需要项目的最终二进制文件越小越好，panic 时通过在 *Cargo.toml* 的 `[profile]` 部分增加 `panic = 'abort'`，可以由展开切换为终止。例如，如果你想要在release模式中 panic 时直接终止：

```toml
[profile.release]
panic = 'abort'
```



> result-与可恢复的错误

```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```



> **Result 例子**

```rust
use std::fs::File;

fn main() {
    let f = File::open("hello.txt");
	// 使用 match 处理 Result
    let f = match f {
        Ok(file) => file,
        Err(error) => {
            panic!("There was a problem opening the file: {:?}", error)
        },
    };
}
```



*匹配不同的错误*

```rust
use std::fs::File;
use std::io::ErrorKind;

fn main() {
    let f = File::open("hello.txt");

    let f = match f {
        Ok(file) => file,
        Err(ref error) if error.kind() == ErrorKind::NotFound => {
            match File::create("hello.txt") {
                Ok(fc) => fc,
                Err(e) => {
                    panic!(
                        "Tried to create file but there was a problem: {:?}",
                        e
                    )
                },
            }
        },
        Err(error) => {
            panic!(
                "There was a problem opening the file: {:?}",
                error
            )
        },
    };
}
```



*失败时-panic-的简写unwrap-和-expect*

- 如果 `Result` 值是成员 `Ok`，`unwrap` 会返回 `Ok` 中的值。如果 `Result` 是成员 `Err`，`unwrap` 会为我们调用 `panic!`。
- 另一个类似于 `unwrap` 的方法它还允许我们选择 `panic!` 的错误信息：`expect`。使用 `expect` 而不是 `unwrap` 并提供一个好的错误信息可以表明你的意图并更易于追踪 panic 的根源。

`expect` 与 `unwrap` 的使用方式一样：返回文件句柄或调用 `panic!` 宏。`expect` 用来调用 `panic!` 的错误信息将会作为参数传递给 `expect` ，而不像`unwrap` 那样使用默认的 `panic!` 信息。



*回调处理 Result*

```rust
use std::io;
use std::io::Read;
use std::fs::File;

fn read_username_from_file() -> Result<String, io::Error> {
    let f = File::open("hello.txt");

    let mut f = match f {
        Ok(file) => file,
        Err(e) => return Err(e),
    };

    let mut s = String::new();

    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    }
}
```



*传播错误的简写 `?`*

```rust
use std::io;
use std::io::Read;
use std::fs::File;

fn read_username_from_file() -> Result<String, io::Error> {
    let mut f = File::open("hello.txt")?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}
```



*错误处理指导原则*

- 有害状态并不包含 **预期** 会偶尔发生的错误
- 之后的代码的运行依赖于处于这种有害状态
- 当没有可行的手段来将有害状态信息编码进所使用的类型中的情况



### 泛型trait-和生命周期

*泛型是具体类型或其他属性的抽象替代。泛型用于为函数签名或结构体等创建定义，允许我们创建许多不同的具体数据类型。*

*定义函数时可以在函数签名的参数数据类型和返回值中使用泛型。以这种方式编写的代码将更灵活并能向函数调用者提供更多功能，同时不引入重复代码。*



> *函数使用范类*

```rust
fn largest<T>(list: &[T]) -> T {
    let mut largest = list[0];

    for &item in list.iter() {
        if item > largest {
            largest = item;
        }
    }

    largest
}

fn main() {
    let number_list = vec![34, 50, 25, 100, 65];

    let result = largest(&number_list);
    println!("The largest number is {}", result);

    let char_list = vec!['y', 'm', 'a', 'q'];

    let result = largest(&char_list);
    println!("The largest char is {}", result);
}
```



> 结构体定义中的泛型

```rust
struct Area<T, U>{
    w: T,
    h: T,
    a: T,
    s: U,
}
```



#### trait定义共享的行为(接口)

> **默认实现**

_有时为 trait 中的某些或全部方法提供默认的行为，而不是在每个类型的每个实现中都定义自己的行为是很有用的。这样当为某个特定类型实现 trait 时，可以选择保留或重载每个方法的默认行为。_



*定义 trait*

```rust
// 定义 trait
trait A{
    // 定义方法，但是不具体实现它
    fn name(&self) -> String;
    fn name2(&self);    
    
    // 默认实现
    fn about(&self) -> String {
        String::from("Joshua Conero. Use A.")
    }
}

// AStr 结构体
struct AStr{
    _name: String,
}

// AStr 实现 A trait
impl A for AStr{
    // 必须实现 name 相关方法
}
```



> **trait-bounds**

_我们可以限制泛型不再适用于任何类型，编译器会确保其被限制为那些实现了特定 trait 的类型，由此泛型就会拥有我们希望其类型所拥有的功能。这被称为指定泛型的 *trait bounds*。_



#### 生命周期与引用有效性

_Rust 中的每一个引用都有其 **生命周期**（*lifetime*），也就是引用保持有效的作用域。大部分时候生命周期是隐含并可以推断的，正如大部分时候类型也是可以推断的一样。类似于当因为有多种可能类型的时候必须注明类型，也会出现引用的生命周期以一些不同方式相关联的情况，所以 Rust 需要我们使用泛型生命周期参数来注明他们的关系，这样就能确保运行时实际使用的引用绝对是有效的。_



>  **生命周期注解语法**

```rust
&i32        // a reference
// 后两者存在的生命周期一样长
&'a i32     // a reference with an explicit lifetime
&'a mut i32 // a mutable reference with an explicit lifetime
```



*限制函数参数具有相同的生命周期*

```rust
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() {
        x
    } else {
        y
    }
}

fn main() {
    let string1 = String::from("abcd");
    let string2 = "xyz";

    let result = longest(string1.as_str(), string2);
    println!("The longest string is {}", result);
}
```



_当在函数中使用生命周期注解时，这些注解出现在函数签名中，而不存在于函数体中的任何代码中。这是因为 Rust 能够分析函数中代码而不需要任何协助，不过当函数引用或被函数之外的代码引用时，参数或返回值的生命周期可能在每次函数被调用时都不同。这可能会产生惊人的消耗并且对于 Rust 来说通常是不可能分析的。在这种情况下，我们需要自己标注生命周期。_



```rust
// 如下可编译， y 与函数体无任何关联
fn longest<'a>(x: &'a str, y: &str) -> &'a str {
    x
}
```



*结构体*

```rust
struct ImportantExcerpt<'a> {
    part: &'a str,
}

fn main() {
    let novel = String::from("Call me Ishmael. Some years ago...");
    let first_sentence = novel.split('.')
        .next()
        .expect("Could not find a '.'");
    let i = ImportantExcerpt { part: first_sentence };
}
```



>  **生命周期省略规则**

_被编码进 Rust 引用分析的模式被称为 **生命周期省略规则**（*lifetime elision rules*）。这并不是需要程序员遵守的规则；这些规则是一系列特定的场景，此时编译器会考虑，如果代码符合这些场景，就无需明确指定生命周期。_

_函数或方法的参数的生命周期被称为 **输入生命周期**（*input lifetimes*），而返回值的生命周期被称为 **输出生命周期**（*output lifetimes*）。_



_生命周期省略规则。第一条规则适用于输入生命周期，后两条规则适用于输出生命周期。_

1. 每一个是引用的参数都有它自己的生命周期参数。换句话说就是，有一个引用参数的函数有一个生命周期参数：`fn foo<'a>(x: &'a i32)`，有两个引用参数的函数有两个不同的生命周期参数，`fn foo<'a, 'b>(x: &'a i32, y: &'b i32)`，依此类推。
2. 如果只有一个输入生命周期参数，那么它被赋予所有输出生命周期参数：`fn foo<'a>(x: &'a i32) -> &'a i32`。
3. 如果方法有多个输入生命周期参数，不过其中之一因为方法的缘故为 `&self` 或 `&mut self`，那么 `self` 的生命周期被赋给所有输出生命周期参数。这使得方法编写起来更简洁。



> **静态生命周期**

_`'static` 生命周期存活于整个程序期间。所有的字符串字面值都拥有 `'static` 生命周期_

```rust
let s: &'static str = "I have a static lifetime.";
```





### 测试



Rust 中的测试函数是用来验证非测试代码是否按照期望的方式运行的。测试函数体通常执行如下三种操作：

1. 设置任何所需的数据或状态
2. 运行需要测试的代码
3. 断言其结果是我们所期望的



//@TODO   *[测试-编写测试](https://kaisery.github.io/trpl-zh-cn/ch11-01-writing-tests.html)*



### rust-的面向对象特性

> OOP(*Object-Oriented Programming*)

_面向对象的程序是由对象组成的。一个 **对象** 包含数据和操作这些数据的过程。这些过程通常被称为 **方法** 或 **操作**_

_Rust 是面向对象的：结构体和枚举包含数据而 impl 块提供了在结构体和枚举之上的方法。虽然带有方法的结构体和枚举并不被 **称为** 对象，但是他们提供了与对象相同的功能_



_Rust 的结构体自身被标记为 `pub`，这样其他代码就可以使用这个结构体，但是在结构体内部的字段仍然是私有的。_

```rust
// 对象
pub struct A{}
// 接口
pub trait A2A{
    pub mode(&self){
    }
}

// 方法等
// A 实现 A2A； 类似 A 继承于 A2A
impl A2A for A{
    // 公有方法 - n1
    pub fn n1(&self){}
    // 私有方法
    fn priv_n1(self){}
}

```



_继承： Rust 代码可以使用默认 trait 方法实现来进行共享。_





### 宏

_从根本上来说，宏是一种为写其他代码而写代码的方式，即所谓的*元编程*。_

_元编程对于减少大量编写和维护的代码是非常有用的，它也扮演了函数的角色。但宏有一些函数所没有的附加能力。_



> 与函数的区别

- 一个函数标签必须声明函数参数个数和类型。而宏只接受一个可变参数
- 宏可以在编译器翻译代码前展开，例如，宏可以在一个给定类型上实现 trait 。因为函数是在运行时被调用，同时 trait 需要在运行时实现，所以函数无法像宏这样。
- 实现一个宏而不是函数的消极面是宏定义要比函数定义更复杂，因为你正在为写 Rust 代码而写代码。由于这样的间接性，宏定义通常要比函数定义更难阅读、理解以及维护。
- 宏和函数的另一个区别是：宏定义无法像函数定义那样出现在模块命名空间中
- 宏和函数最重要的区别是：在一个文件中，必须在调用宏`之前`定义或导入宏，然而却可以在任意地方定义或调用函数。



> 宏分类

- 通用元编程的申明式宏 `macro_rule!`
- 自定义`drive`的过程宏



## 文档/代码库学习







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



### rust 语言附录

> Primitive Types（原始类型）

```rust
array,
bool,
char,
float: f32, f64			// 浮点型
function: fn, 
int: i8, i16, i32, i64, i128, isize
pointer: 	*T			// 指针
reference： &T		   // 引用
slice:  [T]			    // 切片
str：					// 字符串切片
tuple:					// 元组
uint: u8, u16, u32, u128, unit, usize
never: !
```



> 操作符号

```rust
&expr, &mut expr    							// 借用
&type, &mut type, &'a type, &'a mut type		// 借用指针类型
*expr											// 解引用
*const type, *mut type							// 原生指针
```



表 B-2：独立语法

| 符号                                        | 解释                                           |
| ------------------------------------------- | ---------------------------------------------- |
| `'ident`                                    | 命名生命周期或循环标签                         |
| `...u8`, `...i32`, `...f64`, `...usize`, 等 | 指定类型的数值常量                             |
| `"..."`                                     | 字符串常量                                     |
| `r"..."`, `r#"..."#`, `r##"..."##`, etc.    | 原生字符串常量, 未处理的遗漏字符               |
| `b"..."`                                    | 字节字符串; 构造一个 `[u8]` 类型而非字符串     |
| `br"..."`, `br#"..."#`, `br##"..."##`, 等   | 原生字节字符串常量，原生字节和字节结合的字符串 |
| `'...'`                                     | 字符常量                                       |
| `b'...'`                                    | ASCII码字节常量                                |
| `|...| expr`                                | 结束                                           |
| `!`                                         | 对一个离散函数来说最后总是空类型               |
| `_`                                         | “忽略”模式绑定， 也用于整数常量的可读性        |

