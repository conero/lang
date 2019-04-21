// 2019年4月20日 星期六
// @link https://leetcode-cn.com/problems/reverse-string/
// @Name 344. 反转字符串

// 题目
/*
编写一个函数，以字符串作为输入，反转该字符串中的元音字母。

示例 1:
    输入: "hello"
    输出: "holle"

示例 2:
    输入: "leetcode"
    输出: "leotcede"
    说明:
        元音字母不包含字母"y"。
*/

struct Solution {}

// 解答
impl Solution {
    pub fn reverse_vowels(s: String) -> String {
        let s_len = s.len();
        let mut p0:usize = 0;
        let mut px:usize;
        let vowels = vec!['a', 'e', 'i', 'o', 'u'];

        let ss = s.as_bytes();
        println!("{:?}", ss);

        let ssx: &[u8] = vec![];
        while p0 < s_len {
            if p0 == px{
                break;
            }
//            ssx.as_ptr()
            p0 += 1;
        }


    }
}

// 控制台
fn main() {
    println!("Hello, world!");
}
