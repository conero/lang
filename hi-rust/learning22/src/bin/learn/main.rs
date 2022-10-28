use yang::args::Args;
use yang::cmd::Cmd;
use yang::VERSION;

// 二进制文件
fn main() {
    let mut cmd = Cmd::from(vec![]);
    cmd.empty(|args: Args| {
        println!("Rust 2022 for learning");
        println!("Rust 语言实际代码，尽可能囊括教程所有知识点");
        println!("本项目是一个命令行程序！");
        println!("v{}", VERSION);
    });
    cmd.run();
}
