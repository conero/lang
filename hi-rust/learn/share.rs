// 20181209
// Rust 与 C 交互；生成 dll/so 等
// Joshua Coner
// 使用rustc 编译
// $ rustc --crate-type=cdylib share.rs

// #![allow(unused_variables)]
// fn main() {
//     // extern 的使用无需 unsafe。
//     #[no_mangle]
//     pub extern "C" fn call_from_c() {
//         println!("Just called a Rust function from C!");
//     }
// }

const CAL_XY: &'static str = "x*x - y"; 

#[no_mangle]
pub extern "C" fn joshua_conero() {
    println!("Joshua Conero demo it.");
}

#[no_mangle]
pub extern "C" fn cal(x:i32, y:i32) -> i32 {
    x*x - y
}

#[no_mangle]
pub extern "C" fn cal_about() -> String {
    println!("{}", CAL_XY);
    format!("{}", CAL_XY)
}
