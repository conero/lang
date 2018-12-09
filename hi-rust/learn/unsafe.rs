// 20181209
// unsafe - 不全的rust
// Joshua Conero

// 裸指针
fn raw_pointer(){
    let mut num = 9;

    let r1 = &num as *const i32;
    let r2 = &num
}


fn main(){
    raw_pointer()
}