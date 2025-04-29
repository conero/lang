// string 类型测试
// 2025年4月29日

// 编译代码
// zig build-exe hi-string.zig --library c
// 压缩编译
// zig build-exe hi-string.zig --library c -O ReleaseSmall
// zig build-exe hi-string.zig --library c -O ReleaseSmall '-femit-bin=out/hi-string.exe'
//
// 执行代码
// zig run .\hi-string.zig --library c
const std = @import("std");

pub fn main() void {
    std.debug.print("Zig string 测试，since@2025年4月29日 \n\n", .{});

    baseStringTest();
}

//  基础字符串测试
fn baseStringTest() void {
    std.debug.print("-------------------- [ baseStringTest/begin ] --------------------\n", .{});

    // 字符串变量
    const testString = "Hello Zig";
    std.debug.print("testString/字面字符串定义（{?}）: {s}\n", .{ @TypeOf(testString), testString });

    // 变量定义
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    // 创建一个空的 ArrayList 用于存储字符串
    var mutableString = std.ArrayList(u8).init(allocator);
    defer mutableString.deinit(); // 确保在函数退出时释放内存
    std.debug.print("存储可变数组的容器类型: {?}\n", .{@TypeOf(mutableString)});

    // 赋值字符串到 ArrayList
    mutableString.appendSlice("Joshua Conero, 古丞秋！") catch |err| {
        std.debug.print("mutableString.appendSlice err: {?}\n", .{err});
    };
    // 内容输出
    std.debug.print("mutableString/ArrayList: {s}\n", .{mutableString.items});

    // 清除现有内容
    mutableString.clearRetainingCapacity();
    mutableString.appendSlice("新值：元圣修，笃信好学") catch |err| {
        std.debug.print("mutableString.appendSlice err: {?}\n", .{err});
    };
    // 内容输出
    std.debug.print("mutableString/ArrayList: {s}\n", .{mutableString.items});
    printString("函数（printString）打印字符串的值", mutableString.items);

    std.debug.print("-------------------- [ baseStringTest/end ] --------------------\n\n", .{});
}

// 打印字符串
fn printString(desc: []const u8, str: []const u8) void {
    std.debug.print("{s}: {s}\n", .{ desc, str });
}
