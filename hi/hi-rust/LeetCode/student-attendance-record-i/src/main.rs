// 2019年4月30日 星期二
// @link https://leetcode-cn.com/problems/student-attendance-record-i/
// @Name 551. 学生出勤记录 I


// 题目
/*
给定一个字符串来代表一个学生的出勤记录，这个记录仅包含以下三个字符：
    'A' : Absent，缺勤
    'L' : Late，迟到
    'P' : Present，到场
如果一个学生的出勤记录中不超过一个'A'(缺勤)并且不超过两个连续的'L'(迟到),那么这个学生会被奖赏。
你需要根据这个学生的出勤记录判断他是否会被奖赏。

示例 1:
    输入: "PPALLP"
    输出: True

示例 2:
    输入: "PPALLL"
    输出: False
*/

// 申明
struct Solution{}

// 解答
impl Solution {
    pub fn check_record(s: String) -> bool {
        let mut a_ct = 0;
        let mut l_ct = 0;
        let mut last_char: char = ' ';
        for bts in s.chars() {
            if a_ct > 1 || l_ct > 2{
                return false;
            }
            // A 字符
            if bts == 'A' {
                a_ct += 1;
                last_char = bts;
                continue
            }
            // L 字符串
            if last_char == 'L' && bts == 'L'{
                l_ct += 1;
            }else {
                l_ct = if bts == 'L'{
                    1
                }else {
                    0
                };
            }
            last_char = bts;
            //println!("[{} {}], {}", l_ct, a_ct, last_char);
        }
        //println!("[{} {}]", l_ct, a_ct);
        if a_ct > 1 || l_ct > 2{
            return false;
        }
        return true
    }
}

// 控制台
fn main() {
    assert_eq!(true, Solution::check_record("PPALLP".to_string()));
    assert_eq!(false, Solution::check_record("PPALLL".to_string()));
    assert_eq!(true, Solution::check_record("LALL".to_string()));
}
