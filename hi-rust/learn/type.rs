// 20181116
// data-type learn
// Joshua Conero  > type.rs
use std::env;

fn tArray(){}

fn main(){
    println!("Rust Data-Type Learning");
    let mut ctt = 0;
    for arg in env::args(){
        ctt += 1;
        if ctt == 1{
            continue;
        }
        match arg {
            "array" => {
                tArray();
            }
        }
    }
}

