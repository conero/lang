## Zig

> Joshua Conero
>
> 2025年1月21日





ziglang 始于2016由  Andrew Kelley 开发，基于 LLVM 后端



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























### 附录

- [Zig 语言圣经](https://course.ziglang.cc/)
- [Zig 语言中文社区](https://ziglang.cc/)
- [Zig语言中文手册](https://sxwangzhiwen.github.io/zigcndoc/zigcndoc.html)