use std::env;

/// 二进制命令工具

pub struct Cmd {
    args: Vec<String>,
}

// 为结构体添加方法
impl Cmd {
    /// 通过参数初始化命令行程序
    /// # Examples
    /// ```
    ///     let app = Cmd::from(vec!["log", "--stat"]);
    /// ```
    pub fn from(param: Vec<&str>) -> Cmd {
        let mut args: Vec<String> = Vec::new();
        for arg in param {
            args.push(String::from(arg));
        }
        let mut app = Cmd { args };
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

        self.args = args;
    }

    // 解析参数
    fn parse_args(&mut self) {
        if self.args.is_empty() {
            self.get_os_args()
        }
    }
}
