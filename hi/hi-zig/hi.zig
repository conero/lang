// zig 实例
// 2023年4月20日 星期四
// window 下无 libc 时，使用如下进行编译。
// --library c   表明不依赖任何系统文件！
// zig build-exe hi.zig --library c
//
// 极小代码编译
// zig build-exe hi.zig --library c -O ReleaseSmall -fstrip -fsingle-threaded -target x86_64-windows

const std = @import("std");
const process = std.process;
const print = std.debug.print;

pub fn main() !void {
    const args = try std.process.argsAlloc(std.heap.c_allocator);
    defer std.process.argsFree(std.heap.c_allocator, args);

    print("Hello, world! {s}\n", .{args[0]});

    // 遍历并打印所有命令行参数
    if (args.len > 1) {
        for (args[1..]) |arg| {
            print("Argument {s}\n", .{arg});
        }
    } else {
        print("No arguments provided.\n", .{});
    }
}
