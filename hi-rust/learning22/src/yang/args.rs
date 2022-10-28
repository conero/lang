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
    pub data: HashMap<String, Option<String>>,
    raw: Vec<String>,
}

impl Args {
    // 获取命令行参数
    pub fn get_data(&self, key: &str) -> Option<&String> {
        if !self.data.is_empty() {
            if self.data.contains_key(key) {
                let value = self.data.get(key);
                if value != None {
                    let vs = value.unwrap().as_ref().unwrap();
                    return Some(vs);
                }
            }
        }
        None
    }

    // 实例化函数
    pub fn new(args: &Vec<String>) -> Args {
        Args {
            command: "".to_string(),
            sub_command: "".to_string(),
            option: vec![],
            data: Default::default(),
            raw: args.to_vec(),
        }
    }
}
