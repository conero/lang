// 20181209
// Rust 与 C 交互；生成 dll/so 等
// Joshua Coner


#![allow(unused_variables)]
fn main() {
    // extern 的使用无需 unsafe。
    #[no_mangle]
    pub extern "C" fn call_from_c() {
        println!("Just called a Rust function from C!");
    }
}