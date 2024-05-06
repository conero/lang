use std::collections::HashMap;

// 系统参数
pub struct Args {
    // 命令行
    pub command: String,
    // 二级命令
    pub sub_command: String,
    // 参数选项
    pub option: Vec<String>,
    // 请求参数
    pub data: HashMap<String, Vec<String>>,
    pub raw: Vec<String>,
}

impl Args {
    // 获取命令行参数
    pub fn get_data(&self, key: &str) -> Option<&Vec<String>> {
        if !self.data.is_empty() {
            if self.data.contains_key(key) {
                let value = self.data.get(key);
                if value != None {
                    let vs = value.unwrap();
                    return Some(vs);
                }
            }
        }
        None
    }

    // 实例化函数
    pub fn new(args: &Vec<String>) -> Args {
        // 字段信息
        let mut command = String::new(); // 命令
        let mut sub_command = String::new(); // 子命令
        let mut option = Vec::new();
        let mut data: HashMap<String, Vec<String>> = HashMap::new();

        // 辅助变量
        let mut count = 0;
        let mut last_opt = String::new(); // 最新的opt缓存
        let mut last_value: Vec<String> = Vec::new(); // 最新的value缓存

        // 使用闭包函数
        // let update_data_value = |v_opt: &String| {
        //     if !v_opt.is_empty(){
        //         if !data.contains_key(v_opt.as_str()){
        //             data.insert(v_opt, last_value);
        //         }
        //         last_value = Vec::new();
        //     }
        // };

        for arg in args {
            let long_opt = arg.find("--");
            let shot_opt = arg.find('-');
            if count == 0 {
                // 命令
                if shot_opt.is_none() {
                    command = String::from(arg);
                }
            } else if sub_command.len() == 0 && shot_opt == None {
                // 子命令
                sub_command = String::from(arg);
            } else {
                if long_opt != None {
                    let (_, opt) = arg.split_at(2);
                    if !last_opt.is_empty() {
                        if !data.contains_key(last_opt.as_str()) {
                            data.insert(last_opt, last_value);
                        }
                        last_value = Vec::new();
                    }
                    last_opt = String::from(opt);
                    option.push(String::from(&last_opt));
                } else if shot_opt != None {
                    let (_, opt) = arg.split_at(1);
                    for by in opt.chars() {
                        last_opt = String::from(by);
                        option.push(String::from(&last_opt));
                    }
                } else {
                    last_value.push(String::from(arg));
                }
            }
            count += 1;
        }

        Args {
            command,
            sub_command,
            option,
            data,
            raw: args.to_vec(),
        }
    }
}
