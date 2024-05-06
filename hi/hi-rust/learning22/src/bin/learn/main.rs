use yang::action::Action;
use yang::args::Args;
use yang::cmd::{ActionApp, Cmd};
use yang::{PROJECT, VERSION};

struct Version {
    arg: Option<Args>,
}

impl Action for Version {
    fn run(&self, _: &Args) {
        println!("{}/v{}", PROJECT, VERSION)
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

    // 命令不存在
    cmd.un_found(|args: &Args| println!("{} 命令未存在！", args.command));

    // test
    cmd.register("test", |args: &Args| {
        println!("command: {}", args.command);
        println!("sub_command: {}", args.sub_command);
        println!("option: {:?}", args.option);
        println!("data: {:?}", args.data);
        println!("raw: {:?}", args.raw);
    });

    // help
    cmd.register("help", |_: &Args| {
        println!("test      参数解析测试");
        println!("version   版本号输出");
    });

    // 默认方法
    cmd.empty(|_: &Args| {
        println!("Rust 2022 for learning");
        println!("Rust 语言实际代码，尽可能囊括教程所有知识点");
        println!("本项目是一个命令行程序！");
        println!("v{}", VERSION);
    });
    cmd.run();
}
