// generics-struct(t)
// struct 结构体

// 面积
struct Area<T, U>{
    w: T,
    h: T,
    s: U,
}

impl<T, U> Area<T, U> {
    fn get_w(&self) -> &T {
        &self.w
    }
    fn get_s(&self) -> &U {
        &self.s
    }
//    fn to_show(self){
//        println!("width: {}, height: {}. name: {}", self.w, self.h, self.s);
//    }
}

fn apply(){
    let a = Area {
        w: 2,
        h: 10,
        s: String::from("Base 1"),
    };
    //a.to_show();
    println!("width: {}, height: {}. name: {}", a.w, a.h, a.s);
    println!("get_w {}, get_s {}", a.get_w(), a.get_s())
}

fn main(){
    apply();
}