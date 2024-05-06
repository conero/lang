// trait

trait A{
    fn show(&self);
    fn name(&self) -> String;
    fn base(){
        println!(" for A trait.")
    }
}


struct A1{
    name: String,
}
impl A for A1{
    fn show(&self){
        println!(" the name: {}.", self.name)
    }
    fn name(&self) -> String{
        self.name
    }
}

fn main(){
    let a11 = A1{ name: String::from("Joshua"), };
    a11.show();
    A1::base();

}