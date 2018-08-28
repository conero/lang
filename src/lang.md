# lang

> 2018年8月25日 星期六
>
> 原 *word* 转 *markdown* 文件





## 摘要说明



**时间开始于：**2016年4月16日13:40:23@-

    软件行业是一个发展迅速的行业，也为全面的发展，加上长期以来工作上形成的困惑。



### 摘要内容：

计算机语言介绍，内容包括该语言的发展历史，应用场景，基本特点等。



### 编程语言范式

面向对象、面向过程、函数式编程、元编程、泛型编程、命令式、逻辑式、结构化编程。

       面向对象，一些人总结了该设计模式，认为其**[6]**：相左与当前流行的“模块化”、“并行性”特性(如今流行的函数式编程)。



#### *元编程* 

1994年，Erwin Unruh

       这类计算机程序编写或者操纵其他程序（或者自身）作为它们的数据，或者在运行时完成部分本应在编译时完成的工作。很多情况下与手工编写全部代码相比工作效率更高。编写元程序的语言称之为元语言，被操作的语言称之为目标语言。一门语言同时也是自身的元语言的能力称之为反射。反射是促进元编程的一种很有价值的语言特性。把编程语言自身作为头等对象（如Lisp或Rebol）也很有用。支持泛型编程的语言也使用元编程能力。
    
       元编程通常有两种方式起作用。一种方式是通过应用程序接口（API）来暴露运行时引擎的内部信息。另一种方法是动态执行包含编程命令的字符串。因此，“程序能编写程序”。虽然两种方法都能用，但大多数方法主要靠其中一种。
    
       编译期实现而非运行期



#### 函数式编程

采用一种基于函数的递归定义的计算模型。本质上，程序被看作是一种从输入到输出的函数，基于一些更简单的函数，通过一种逐步精化的过程定义。

函数式编程是一种编程范式，它有下面的一些特征：

1. 函数是一等公民，可以像数据一样传来传去
2. 高阶函数
3. 递归
4. pipeline
5. 惰性求值
6. 柯里化
7. 偏应用函数

C++新特效(C++11/14)可满足要求



#### 数据流

语言将计算看成在一些基本的功能结点之间流动的信息流。

### 其他



> 图 计算语言结构

![计算语言结构](./image/lang/struct.png)

注: 目标语言包括机器语言，机器语言是计算机可识别的01二进制，汇编语言使用英文指令来代替二进制并通过汇编为二进制。



| 机器抽象(虚拟机) | 高级语言 | 解释型   | 编译型 |
| ---------------- | -------- | -------- | ------ |
| 硬件/机器依赖    | 低级语言 | 汇编语言 |        |
| 机器语言         |          |          |        |
| 硬件/机器        |          |          |        |



> 计算机语言层级

![](./image/lang/types.png)



> 编译器

![](./image/lang/complier.png)

>  C# / VB    SourceCode → ByteCode(MSIL/CIL中间件) → NativeCode(CPU)

 

对于C#、VB等高级语言而言，此时编译器完成的功能是把源码（SourceCode）编译成通用中间语言（MSIL/CIL）的字节码（ByteCode）。最后运行的时候通过通用语言运行库的转换，编程最终可以被CPU直接计算的机器码（NativeCode）

 

- 编译                  高级语言 -> 低级语言
- 反编译                低级语言 -> 高级语言
- 级联                  高级语言1 ->高级语言2

 

本地编译器        编译器可以生成用来在与编译器本身所在的计算机和操作系统（平台）相同的环境下运行的目标代码
交叉编译器     编译器也可以生成用来在其它平台上运行的目标代码
语法分析器
语意分析器
抽象的语法树（abstract syntax tree，或 AST）





> C/C++ 编译器

- windows  ``VC``
- linux ``gcc``
- osx ``clang``



#### windows/工具

``Cygwin`` 、``WinGW`` 分别为 *中间转换工具/平台*



> Cygwin

*Get that Linux feeling - on Windows*



> WinGW

*a contraction of "Minimalist GNU for Windows", is a minimalist development environment for native Microsoft Windows applications.*







## 编程语言


