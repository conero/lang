// @date 20181211
// @name marco rust 宏测试
// @author Joshua Conero


// 宏定义
// joshua conero marco test
// 类似 match 语法
macro_rules! jcmc{
    () => {
        println!(" Jcmc: welcome use the [Joshua Conero Marco Test].");
    };
    ($key:expr) => {
        println!(" Jcmc: {}", $key);
    };
}

// 运行
fn main(){
    jcmc!();
    jcmc!("eeee")
}