// 20181209
// unsafe - 不全的rust
// Joshua Conero

// 在 extern "C" 块中，列出了我们希望能够调用的另一个语言中的外部函数的签名和名称。
extern "C" {
    fn abs(input: i32) -> i32;
}

// 裸指针
fn raw_pointer(){
    let mut num = 9;

    // 裸指针
    let r1 = &num as *const i32;
    let r2 = &mut num as *mut i32;

    // 非 unsafe 代码
    // error[E0133]: dereference of raw pointer is unsafe and requires unsafe function or block
    //println!("{}", *r1 + *r2);

    unsafe {
        println!("使用 unsafe 块，标注非安全代码。");
        println!("{}", *r1 + *r2);
    }

}

// 非安全
unsafe fn dangerous() {
    // body
}

fn use_unsafe(){

    // error[E0133]: call to unsafe function is unsafe and requires unsafe function or block
    // dangerous();

    unsafe {
        dangerous();
    }
}

// 调用 C 代码
fn use_extern(){
    unsafe {
        println!("Absolute value of -3 according to C: {}", abs(-3));
    }
}
fn main(){
    raw_pointer();
    use_unsafe();
    use_extern();
