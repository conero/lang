use std::sync::{mpsc, Arc, Mutex};
use std::thread;

// 并发测试
fn main() {
    println!("rust 并发测试");

    // 计算用时
    let start = std::time::Instant::now();
    let count: isize = 1_987_654;
    //let count: isize = 19_987_654;
    //let sum: i32 = simple_conc(count);
    let sum: i32 = simple_conc_pool(count);

    println!(
        "并发{}， 计算值 {}, 用时 {} 秒",
        count,
        sum,
        start.elapsed().as_secs_f64()
    );
}

// 互斥锁，同一时间仅允许一个线程访问数据，支持 Sync + Send
//
// !存在内存泄漏
#[allow(dead_code)]
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

// 安全的并发计算：使用线程池复用线程
//
// 纯标准库实现：线程池 + 任务队列
fn simple_conc_pool(count: isize) -> i32 {
    let sum = Arc::new(Mutex::new(0));
    let (tx, rx) = mpsc::channel::<isize>();

    // 关键修复：用 Arc<Mutex<>> 包装接收端，实现共享（而非克隆 rx）
    let rx_shared = Arc::new(Mutex::new(rx)); // 共享的接收端
    const NUM_WORKERS: usize = 8;

    let mut workers = vec![];
    for _ in 0..NUM_WORKERS {
        let sum = Arc::clone(&sum);
        let rx_shared = Arc::clone(&rx_shared); // 克隆 Arc（共享同一个 Mutex<Receiver>）

        workers.push(thread::spawn(move || {
            // 循环获取锁并接收任务
            while let Ok(i) = rx_shared.lock().unwrap().recv() {
                print!("thread {} \r", i);
                // 每次锁定后接收
                let mut num = sum.lock().unwrap();
                if i % 2 == 0 {
                    *num += 2;
                } else {
                    *num -= 1;
                }
            }
        }));
    }

    // 执行任务
    for i in 0..count {
        tx.send(i).unwrap();
    }
    drop(tx);

    // 等待所有任务完成
    for worker in workers {
        worker.join().unwrap();
    }

    // 返回值
    let sum_value = *sum.lock().unwrap();
    sum_value
}
