//2019年1月1日 星期二
//Rust 所有权概念测试


fn main(){
    println!("{}", "所有权");
    let s = "字面值字符串-str";
    let cs = s;
    println!("{}", cs);
}