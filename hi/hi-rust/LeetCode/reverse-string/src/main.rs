// 2019年4月20日 星期六
// @link https://leetcode-cn.com/problems/reverse-string/
// @Name 344. 反转字符串

/*
// 题目
编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 char[] 的形式给出。
不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
你可以假设数组中的所有字符都是 ASCII 码表中的可打印字符。


示例 1：
    输入：["h","e","l","l","o"]
    输出：["o","l","l","e","h"]

示例 2：
    输入：["H","a","n","n","a","h"]
    输出：["h","a","n","n","a","H"]
*/

struct Solution {}

// 解答
impl Solution {
    pub fn reverse_string(s: &mut Vec<char>) {
        let vlen = s.len();
        let mut i = 0;
        let mid = (vlen - vlen % 2) / 2;
        while vlen > 0 && i < mid {
            let last = (vlen - i - 1);
            if last == i {
                break;
            }
            // println!("{}/{} ({}, {})", vlen, mid, i, last);
            let c0 = s[i];
            let cx = s[last];
            s[last] = c0;
            s[i] = cx;
            i += 1;
        }
    }
}

// 控制台
#[warn(unused_variables)]
#[warn(unused_mut)]
fn main() {
    let s = &mut vec!['h', 'e', 'l', 'l', 'o'];

    Solution::reverse_string(s);
    assert_eq!(vec!['o', 'l', 'l', 'e', 'h'], *s);

    let s = &mut vec!['H', 'a', 'n', 'n', 'a', 'h'];
    Solution::reverse_string(s);
    assert_eq!(vec!['h', 'a', 'n', 'n', 'a', 'H'], *s);
    // println!("{:?}", *s);

    let s = &mut vec!['i', 'm', 'n', 'o', 't'];
    Solution::reverse_string(s);
    assert_eq!(vec!['t', 'o', 'n', 'm', 'i'], *s);

    let s = &mut vec![
        'A', ' ', 'm', 'a', 'n', ',', ' ', 'a', ' ', 'p', 'l', 'a', 'n', ',', ' ', 'a', ' ', 'c',
        'a', 'n', 'a', 'l', ':', ' ', 'P', 'a', 'n', 'a', 'm', 'a',
    ];
    Solution::reverse_string(s);
    assert_eq!(
        vec![
            'a', 'm', 'a', 'n', 'a', 'P', ' ', ':', 'l', 'a', 'n', 'a', 'c', ' ', 'a', ' ', ',',
            'n', 'a', 'l', 'p', ' ', 'a', ' ', ',', 'n', 'a', 'm', ' ', 'A'
        ],
        *s
    );
}
