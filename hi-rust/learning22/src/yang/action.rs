use crate::args::Args;

// 执行方法
pub trait Action {
    // 初始化
    fn init(self, arg: Args);

    // 方法入口
    fn run(&self) {
        println!("命令行的 Action::run 待实现！")
    }
}
