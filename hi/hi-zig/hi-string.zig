// string 类型测试
// 2025年4月29日

// 编译代码
// zig build-exe hi-string.zig --library c
// 压缩编译
// zig build-exe hi-string.zig --library c -O ReleaseSmall
// zig build-exe hi-string.zig --library c -O ReleaseSmall '-femit-bin=out/hi-string.exe';.\out\hi-string.exe
//
// 执行代码
// zig run .\hi-string.zig --library c
const std = @import("std");

pub fn main() void {
    std.debug.print("Zig string 测试，since@2025年4月29日 \n\n", .{});
    simpleStringTest();
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

//  简单字符串测试
fn simpleStringTest() void {
    std.debug.print("-------------------- [ simpleStringTest/begin ] --------------------\n", .{});

    // 字符串变量
    const testString = "天行健，君子以自强不息。\n A gentleman strives for self-improvement without ceasing.";
    std.debug.print("testString/字面字符串定义（{?}）: {s}, len={d}, lastWord={s}\n", .{ @TypeOf(testString), testString, testString.len, testString[testString.len - 1 .. testString.len] });
    // const utf9Len: usize = std.unicode.calcUtf16LeLen(testString) catch |err| {
    //     std.debug.print("testString/calcUtf16LeLen err: {?}\n", .{err});
    //     //0;
    //     @as(usize, 0);
    // };
    const utf9Len: usize = std.unicode.calcUtf16LeLen(testString) catch 0;
    std.debug.print("testString/utf9Len: {d}\n", .{utf9Len});

    // index计算
    const uIndex: usize = std.mem.lastIndexOf(u8, testString, "强") orelse 0;
    std.debug.print("字符串index（强）: {d}\n", .{uIndex});
    strZeroSplit(testString);
    utf8ZeroSplit(testString);

    std.debug.print("-------------------- [ simpleStringTest/end ] --------------------\n\n", .{});
}

// utf8 字符串分割
fn utf8ZeroSplit(vStr: []const u8) void {
    const uView = std.unicode.Utf8View.init(vStr) catch |err| {
        std.debug.print("utf8Split err: {?}\n", .{err});
        return;
    };
    var iter = uView.iterator();
    std.debug.print("UTF 分割：", .{});
    while (iter.nextCodepointSlice()) |codepoint| {
        std.debug.print("{s}-", .{codepoint});
    }
    std.debug.print("\n", .{});
}

// 字符串分割
// @todo utf8 空字符串分割无效
fn strZeroSplit(vStr: []const u8) void {
    var iter = std.mem.splitAny(u8, vStr, "");
    std.debug.print("字符串分割：", .{});
    while (iter.next()) |codepoint| {
        std.debug.print("{s}-", .{codepoint});
    }
    std.debug.print("\n", .{});
}
