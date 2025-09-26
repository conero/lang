const std = @import("std");

pub fn main() !void {
    std.debug.print("concurrent test for zig.\n\n", .{});

    const cct: usize = 1_978_654;
    const vData = simple_conc(cct);

    std.debug.print("sum: {}， 并发： {}\n", .{ vData, cct });
}

// 创建并发
fn simple_conc(count: usize) isize {
    var wg: std.Thread.WaitGroup = .{}; // 线程等待
    var mut: std.Thread.Mutex = .{}; // 数据锁
    var initValue: isize = 0;
    const sum: *isize = &initValue;
    //sum.* = 0;
    for (0..count) |idx| {
        wg.start();
        _ = std.Thread.spawn(.{}, calc_data, .{ &wg, idx, &mut, sum }) catch |err| {
            std.debug.print("error: {}\n", .{err});
        };
    }

    // 等待所有线程完成
    wg.wait();
    return sum.*;
}

// 线程函数
fn calc_data(wg: *std.Thread.WaitGroup, idx: usize, mutex: *std.Thread.Mutex, sum: *isize) void {
    defer wg.finish();
    mutex.lock();
    defer mutex.unlock();

    std.debug.print("thread {} ...\r", .{idx});

    var value = sum.*;
    if (idx % 2 == 0) {
        value += 2;
    } else {
        value -= 1;
    }
    sum.* = value;
}
