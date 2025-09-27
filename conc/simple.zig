const std = @import("std");

pub fn main() !void {
    std.debug.print("concurrent test for zig.\n\n", .{});

    const start_nano = std.time.nanoTimestamp();

    const cct: usize = 1_987_654;
    const max_concurrent: usize = 5000;
    //const vData = simple_conc(cct, max_concurrent);
    const vData = simple_conc_pool(cct, max_concurrent);

    // 计算耗时
    const diff_nano = std.time.nanoTimestamp() - start_nano;
    const diff_Sec = @as(f64, @floatFromInt(diff_nano)) / 1_000_000_000;

    std.debug.print("并发： {}，线程池：{}, sum: {}，耗时：{}s\n", .{ cct, max_concurrent, vData, diff_Sec });
}

// 创建并发
//
// count: 数据总数, 存在问题内存泄漏
fn simple_conc(count: usize, max_concurrent: usize) isize {
    var wg: std.Thread.WaitGroup = .{}; // 线程等待
    var mut: std.Thread.Mutex = .{}; // 数据锁
    var initValue: isize = 0;
    const sum: *isize = &initValue;
    // 信号量：控制同时运行的线程数（初始值为 max_concurrent）
    var semaphore = std.Thread.Semaphore{ .permits = max_concurrent };

    // _ = max_concurrent;
    // var semaphore = std.Thread.Semaphore{};

    //sum.* = 0;
    for (0..count) |idx| {
        semaphore.wait();
        //wg.start();
        _ = std.Thread.spawn(.{}, calc_data, .{ &wg, idx, &mut, sum, &semaphore }) catch |err| {
            std.debug.print("error: index/{}, {}\n", .{ idx, err });
            semaphore.post();
        };
        // 获取信号量许可（若已达最大并发，则阻塞等待）
        // semaphore.wait();
    }

    // for (0..count) |_| {
    //     semaphore.wait();
    // }
    // 等待所有线程完成
    //wg.wait();
    return sum.*;
}

// 线程函数
fn calc_data(wg: *std.Thread.WaitGroup, idx: usize, mutex: *std.Thread.Mutex, sum: *isize, semaphore: *std.Thread.Semaphore) void {
    defer {
        _ = wg;
        //wg.finish();
        semaphore.post(); // 任务完成，释放信号量许可
    }

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

// 使用线程池测试
fn simple_conc_pool(count: usize, max_concurrent: usize) isize {
    var pool: std.Thread.Pool = undefined;
    std.Thread.Pool.init(&pool, .{
        .n_jobs = max_concurrent,
        .allocator = std.heap.page_allocator,
    }) catch |err| {
        std.debug.print("error: {}\n", .{err});
        return 0;
    };
    defer pool.deinit();

    // 并发执行
    var wg = std.Thread.WaitGroup{};
    var mut: std.Thread.Mutex = .{}; // 数据锁
    var initValue: isize = 0;
    const sum: *isize = &initValue;

    for (0..count) |idx| {
        wg.start();
        pool.spawn(calc_data_pool, .{ &wg, idx, &mut, sum }) catch |err| {
            std.debug.print("error: index/{}, {}\n", .{ idx, err });
            wg.finish();
        };
    }

    wg.wait();
    return sum.*;
}

fn calc_data_pool(wg: *std.Thread.WaitGroup, idx: usize, mutex: *std.Thread.Mutex, sum: *isize) void {
    defer {
        wg.finish();
    }
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
