// 2018年12月28日 星期五
// @link https://leetcode-cn.com/problems/number-of-segments-in-a-string/

/*
    统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
    请注意，你可以假定字符串里不包括任何不可打印的字符。

    示例:
        输入: "Hello, my name is John"
        输出: 5
*/


pub struct Solution{
}

// -------------------------[ script ]--------------------------------------------
impl Solution {
    pub fn count_segments(s: String) -> i32 {
    }
}
// -------------------------[ script ]--------------------------------------------


#[warn(unused_mut)]
fn main() {
    let digists = vec![5,1,5,2,5,3,5,4];
    println!("{:?} => {}", digists.clone(), Solution::repeated_n_times(digists));
//    println!("{:?} => {:?}", vec![2, 3, 5], Solution::plus_one(vec![2, 3, 5]));


    // 测试用例
    assert_eq!(5, Solution::repeated_n_times("Hello, my name is John".to_string());
}
