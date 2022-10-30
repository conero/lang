use crate::args::Args;

// 执行方法
pub trait Action {
    // 初始化
    fn init(&self) {}

    // 方法入口
    fn run(&self, arg: &Args);
}
