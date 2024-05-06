///  依赖库版本信息库
pub const VERSION: &'static str = "0.0.1";
/// 项目代码
pub const PROJECT: &'static str = "learn";

// 命令行工具
pub mod cmd;
// 命令行参数
pub mod args;
// 方法集合
pub mod action;

// 测试用例
#[cfg(test)]
mod test;
