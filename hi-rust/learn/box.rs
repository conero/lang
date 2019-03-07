// 20181130
// Joshua Conero

// cons list 是一个来源于 Lisp 编程语言及其方言的数据结构。
// 递归空间
use std::env;
use std::ops::Deref;
enum List{
    Cons(i32, Box<List>),
    Nil,
}

// 使用 List 枚举的指针
use List::{Cons, Nil};

// 实现类似于 Box 的功能
struct NBox<T>(T);
impl <T> NBox<T>{
    fn new(x: T) -> NBox<T>{
        NBox(x)
    }
}
// 使用 Deref
impl<T> Deref for NBox<T>{
    type Target = T;
    fn deref(&self) -> &T{
        &self.0
    }
}

// 使用 Drop trait
// 类似 C++ 中的解析函数
struct CustomSmartPointer {
    data: String,
}

impl Drop for CustomSmartPointer {
    fn drop(&mut self) {
        println!("Dropping CustomSmartPointer with data `{}`!", self.data);
    }
}

// 测试用例
struct TC{
}

impl TC{
    // 基本测试
    fn base(){
        let b = Box::new(5);
        println!("b={}", b)
    }
    // recursive type
    // 递归类型
    fn dglx(){
        let list = Cons(1,
                        Box::new(Cons(2,
                                      Box::new(Cons(3,
                                                    Box::new(Cons(4,
                                                                  Box::new(Nil))))))));

        //let c1 = Cons(1000, Box::new(Nil));
        //let (cx, cy) = c1;
        //println!("{}", cx)
        // println!("{}", list.0)

    }
    // 解引用
    fn jyy(){
        let x = 5;
        let y = &x;

        // 解引用
        // y == x (报错)
        println!("x = {}; y = {}; Is x = *y -> {}", x, y, *y == x);

        // 使用 Box
        let x = 5;
        let y = Box::new(x);
        // 解引用
        // y == x (报错)
        println!("x = {}; y = {}; Is x = *y -> {}", x, y, *y == x);
    }
    // 自定义
    fn deref(){
        let x = 5;
        let y = NBox::new(x);
        println!("y={}", *y);
        println!("x = {}; y = {}; Is x = *y -> {}", x, *y, x == *y);
    }
    fn drop(){
        let c = CustomSmartPointer { data: String::from("my stuff") };
        let d = CustomSmartPointer { data: String::from("other stuff") };
        let h = CustomSmartPointer { data: String::from("try my best to learning.")};
        //通过 std::mem::drop 提早丢弃值
        drop(h);
        println!("CustomSmartPointers created.");
    }
    // 默认类
    fn empty(){
        println!(" Box 测试智能指针");
        println!("  . base");
    }
    fn run(&self){
        let mut ctt = 0;
        let mut cmd = String::from("");
        for v in env::args(){
            ctt += 1;
            if ctt == 1{
                continue;
            }
            cmd = v;
            break;
        }
        match cmd.as_str() {
            "base" => TC::base(),
            "dglx" => TC::dglx(),
            "jyy"  => TC::jyy(),
            "deref" => TC::deref(),
            "drop" => TC::drop(),
            _ => TC::empty(),
        }
    }
}

fn main(){
    let tc = TC{};
    tc.run();
    //TC::dglx();
    //TC::jyy();
    //TC::deref();
    TC::drop();
}