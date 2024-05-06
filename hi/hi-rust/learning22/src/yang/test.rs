use crate::cmd::Cmd;
use std::env;

#[test]
fn cmd_from() {
    let app = Cmd::from(vec![]);
    println!("{:?}", env::args());

    for argument in env::args() {
        println!("{argument}");
    }
}
