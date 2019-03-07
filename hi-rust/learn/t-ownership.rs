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
    // 变量与数据交互的方式（一）：移动
    let sv = String::from("String 字符串");
    let sv1 = sv;   // sv -> sv1
    println!("sv1 = {}", sv1);
    // println!("sv = {}", sv);    // 变量已被借用  error[E0382]: borrow of moved value: `sv`

    // 变量与数据交互的方式（二）：克隆
    let sv = String::from("String 字符串： 隐藏");
    let sv1 = sv.clone();
    println!("{}, {}", sv1, sv);
    // sv =>
    let sv2 = get_str_onwership(sv);
    //println!("{}, {}", sv2, sv);    // error[E0382]: borrow of moved value: `sv`
    println!("{}", sv2);
}

// 值获取
fn get_str_onwership(bs: String) -> String{
    println!(" get the string dsta [{}]", bs);
    let bs = bs + &" Maker".to_string();
    bs
}