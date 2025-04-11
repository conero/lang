// 简洁的 hello world 程序
// 2025年4月9日

// 编译代码
// zig build-exe hi-simple.zig --library c
// 压缩编译
// zig build-exe hi-simple.zig --library c -O ReleaseSmall
// zig build-exe hi-simple.zig --library c -O ReleaseSmall '-femit-bin=out/hi-simple.exe'
//
// 执行代码
// zig run .\hi-simple.zig --library c
const std = @import("std");

pub fn main() void {
    std.debug.print("Hello, world! \n💗 贵阳 中国\n", .{});
}
