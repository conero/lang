use std::{
    env
};

// array
fn t_array(){
    // 英文月
    let month = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
    println!("12 months: {:?}", month);
    let mut month_idx:[i32; 12] = [1;12];

    // 数据循环
    // for x in &month{}
    // for x in month.iter(){}
    // for i in month.len(){}

    // 数组循环2
    for i in 0..month.len(){
        // println!("{}", i);
        // 类型装换
        month_idx[i] = i as i32;
    }
    println!("12 months index: {:?}", month_idx)
}

// bool
fn t_bool(){
    let bool_val = true & false | false;
    println!(" bool type test");
    println!(" true & false | false => {}", bool_val);

}

// char
fn t_char(){
    let jc = vec!['J', 'o', 's', 'h', 'u', 'a', ',', 'C', 'o', 'n', 'e', 'r', 'o'];
    println!("jc = {:?}", jc);

    // str => char
    println!("\"Joshua,Conero\".chars() = {:?}", "Joshua,Conero".chars());
    for x in "Joshua,Conero".chars() {
        print!("{}|", x)
    }
    print!("\n")

}

fn t_f32(){
    // 系统默认位数 f64
    let f1: f32 = 7.32;
    println!("{}", f1 + (1.24 as f32))
}

//变量测试
fn t_var(){}

// 程序入口
fn index(){
    println!(" 支持测试 $ [type]");
    println!("  array  数组测试");
    println!("  bool");
    println!("  char");
    println!("  f32");
}

// `arg` is borrowed here
fn get_cmd() -> &'static str{
    let mut cmd = "";
    let mut ctt = 0;
    for arg in env::args(){
        ctt += 1;
        if ctt == 1{
            continue;
        }
        cmd = arg.as_str()
    }
    cmd
}

fn main(){
    println!("Rust Data-Type Learning: ");
    //let cmd = get_cmd();
    let cmd = format!("{}", get_cmd());
    match cmd.as_str() {
        "array" => t_array(),
        "bool"  => t_bool(),
        "char"  => t_char(),
        "f32"   => t_f32(),
        "var"   => t_var(),
        _ => index(),
    }
}
