use std::sync::{Arc, Mutex};
use std::thread;

// 并发测试
fn main() {
    println!("rust 并发测试");

    // 计算用时
    let start = std::time::Instant::now();
    let count: isize = 1_987_654;
    let sum: i32 = simple_conc(count);

    println!(
        "并发{}， 计算值 {}, 用时 {} 秒",
        count,
        sum,
        start.elapsed().as_secs()
    );
}

// 互斥锁，同一时间仅允许一个线程访问数据，支持 Sync + Send
fn simple_conc(count: isize) -> i32 {
    // 1. 创建共享数据：用 Mutex 保护 i32，用 Arc 实现多线程共享
    let sum = Arc::new(Mutex::new(0));

    let mut handles = vec![];
    for i in 0..count {
        // 克隆 Arc（增加引用计数，不复制数据）
        let counter = Arc::clone(&sum);
        handles.push(thread::spawn(move || {
            print!("thread {} ...\r", i);
            // 3. 加锁（lock() 返回 Result<MutexGuard<T>>，需 unwrap 处理）
            let mut num = counter.lock().unwrap();
            if i % 2 == 0 {
                *num += 2;
            } else {
                *num -= 1;
            }
        }));
    }

    // 等待所有线程完成
    for handle in handles {
        handle.join().unwrap();
    }

    let sum_value = sum.lock().unwrap().clone();
    sum_value
}
