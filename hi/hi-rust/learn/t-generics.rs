//2019年1月1日 星期二
//Rust 范类测试

trait StackAction<T>{
    // 进栈
    fn push();
    fn pop() -> T;
}

//实现接口
impl StackAction for Stack<T>{
    fn push(){}
    fn pop() -> T{}
}

//栈，后进先出
struct Stack<T>{}

//堆，先进先出
struct Heap{}


fn main(){
    println!("{}", "所有权")
}