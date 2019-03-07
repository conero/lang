/**
 *  时间： 2017年5月25日 星期四
 *  版本号： rust-lang V-1.17.0 
 *  命名： rust - 入门程序编写示例
 *  参考：
 *      >>.  https://doc.rust-lang.org/book/                官网程序编程设计
 *      >>.  https://doc.rust-lang.org/reference/           语言参考
 *      >>.  https://doc.rust-lang.org/stable/std/          标准库
 *      >>.  http://rustbyexample.com/                      编程示例
 *      >>.  http://rustbook.cn/                            中文字段文档
 **/
use std::env;       // 环境变量
// use std::string;

const VERSION:&'static str = "v0.0.1(20170525)";
const AUTHOR:&'static str = "Joshua Conero";

// 入门处理
// warning: function `HelloWorld` should have a snake case name such as `hello_world`   
// 编译器会检查命名规则
fn hello_world(){
    let idx = 5;
    let mtime = "2017年5月25日 星期四";

    // 输入语句
    println!("数字 : {}", idx);
    println!("时间 : {}", mtime);
    println!("{}",mtime);
    // 命令参数获取
    for argument in env::args() {
        println!("{}", argument);
    }
    println!("env::args().len()长度： {}",env::args().len());
}



// 输出命令
struct Console{}        
// 自定义结构体
struct Conero{}     

impl Console{
    // 关于
    fn about(self){
        let text = "
            @name   rust-hi 入门脚本(helloworld +)/命令行程序
            @date   2017年5月25日 星期四
            @author Joshua Coenro

            $ 命令
                $ version|v     版本号
                $ author|a     作者
                $ test|t        测试脚本
        ";
        println!("{}",text);
        println!("  @version {}",VERSION);
    }
}
// &str -> String
fn to_string(x:&str) -> String{
    return x.to_string();
}

// 继承
impl Conero{
    // 测试方法
    fn test(&self){
        println!("Conero 内的test方法")
    }
    /**
     *  把字符串分成 字符串片段 &str 和 字符串 String 两类
     *  (&str).to_string => String
    **/
    // 路由器
    fn router(&self){   
        let args = env::args();
        let csl = Console{};

        // let mut x = "";
        // let mut x = "".to_string();
        // let mut x = String::from("");
        // 获取命令值
        let mut x = String::new();
        if args.len() > 1{
            let mut count = 0;
            for argc in args{
                if count == 1{
                    //x = argc;   //  str <-> string 报错： expected &str, found struct `std::string::String`
                    x = argc;
                    // x = argc.as_slice();
                    //x = argc.as_str();   //  `argc` does not live long enough
                    //x = String::from(argc);
                    //println!("{}",argc.to_string)

                    // println!("{}",count);
                    //println!("{}",argc);
                    // println!("{}",x == "version");
                    //println!("{}",args[1]); // cannot index a value of type `std::env::Args`
                }
                count = count + 1;
            }
        }
        /*
        // 开始使用该脚本，但是报错
        match x {
            //to_string("v") | to_string("version") => println!("855"),
            // "v" | "version" => println!("855"),
            // "v".to_string() | "version".to_string() => println!("855"),
            "v".to_string() => println!("855"),
            _ => csl.about(),
        }
        */

        if x == "version" || x == "v"{  // 版本号
            println!("  @version {}",VERSION);
        }else if x == "test" || x == "t"{
            hello_world();
            self.test();
            println!("{}",to_string("测试二面"))
        }else if x == "author" || x == "a"{
            println!("  @author {}",AUTHOR);
        }else{            
            csl.about();
        }
        //hello_world();
    }
}

// 主方法
fn main(){
    let cro = Conero{};
    cro.router();

    /*
    hello_world();
    cro.test();
    */
}

