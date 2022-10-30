use crate::action::Action;
use crate::args::Args;
use std::collections::HashMap;
use std::env;

// 类型别名
type ActionCallback = fn(&Args);

pub struct ActionApp<T>
where
    T: Action,
{
    pub command: String, // 命令
    pub alias: Vec<String>,
    pub action: T,
}

/// 二进制命令工具

pub struct Cmd<T>
where
    T: Action,
{
    raw_args: Vec<String>,
    calls: HashMap<String, ActionCallback>, // 函数集合
    actions: Vec<ActionApp<T>>,             //方法集合
    action_default: Option<ActionCallback>, // 默认执行方法
    args: Option<Args>,
}

// 为结构体添加方法
impl<T> Cmd<T>
where
    T: Action,
{
    /// 通过参数初始化命令行程序
    /// # Examples
    /// ```
    ///     use yang::cmd::Cmd;
    ///     let app = Cmd::from(vec!["log", "--stat"]);
    /// ```
    pub fn from(param: Vec<&str>) -> Cmd<T> {
        let mut args: Vec<String> = Vec::new();
        for arg in param {
            args.push(String::from(arg));
        }
        let mut app = Cmd {
            raw_args: args,
            calls: HashMap::new(),
            action_default: None,
            args: None,
            actions: Vec::new(),
        };
        app.parse_args();
        return app;
    }

    // 获取操作系统命令
    fn get_os_args(&mut self) {
        let mut args: Vec<String> = Vec::new();
        let mut idx = 0;
        for arg in env::args() {
            if idx < 1 {
                idx += 1;
                continue;
            }
            args.push(arg);
        }

        self.raw_args = args;
    }

    // 解析参数
    fn parse_args(&mut self) {
        if self.raw_args.is_empty() {
            self.get_os_args()
        }
    }

    // 方法注册
    pub fn register(&mut self, name: &str, action: ActionCallback) -> &mut Cmd<T> {
        self.calls.insert(String::from(name), action);
        self
    }

    pub fn register_action(&mut self, app: ActionApp<T>) -> &mut Cmd<T> {
        self.actions.push(app);
        self
    }

    // 默认方法
    pub fn empty(&mut self, action: ActionCallback) -> &mut Cmd<T> {
        self.action_default = Some(action);
        self
    }

    // 命令行执行
    pub fn run(&mut self) {
        let args = Args::new(&self.raw_args);
        self.args = Some(args);
        if !self.action_default.is_none() {
            (self.action_default.as_ref().unwrap())(self.args.as_ref().unwrap());
        }
    }
}
