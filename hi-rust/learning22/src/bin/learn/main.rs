use yang::action::Action;
use yang::args::Args;
use yang::cmd::{ActionApp, Cmd};
use yang::VERSION;

struct Version {
    arg: Option<Args>,
}

impl Action for Version {
    fn run(&self, arg: &Args) {
        println!("{}", VERSION)
    }
}

// 二进制文件
fn main() {
    let mut cmd = Cmd::from(vec![]);
    let version = ActionApp {
        command: String::from("version"),
        alias: vec![],
        action: Box::new(Version { arg: None }),
    };

    cmd.register_action(Box::new(version));
    // 默认方法
    cmd.empty(|args: &Args| {
        println!("Rust 2022 for learning");
        println!("Rust 语言实际代码，尽可能囊括教程所有知识点");
        println!("本项目是一个命令行程序！");
        println!("v{}", VERSION);
    });
    cmd.run();
}
