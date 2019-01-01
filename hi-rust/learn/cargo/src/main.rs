//fn main() {
//    println!("Hello, world!");
//}
// 单文件写入
//-------

//2019年1月1日 星期二
//Rust 所有权概念测试



fn main(){
    println!("{}", "所有权");

    // 栈上数据处理
    let s = "字面值字符串-str";
    let cs = s;
    println!("cs = {}", cs);
    println!("c = {}", s);

    // 堆上数据操作
    let sv = String::from("String 字符串");
    let sv1 = sv;
    println!("sv1 = {}", sv1);
    println!("sv = {}", sv);
}