// 2017年5月25日 星期四
// rust-lang V-1.17.0 
// rust - 入门程序编写
use std::env;

fn main(){
    // 输入语句
    println!("Hello World");
    // 命令参数获取
    for argument in env::args() {
        println!("{}", argument);
    }

}

