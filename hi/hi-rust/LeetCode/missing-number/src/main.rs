// 2019年4月18日 星期四
// @link https://leetcode-cn.com/problems/missing-number/

pub struct Solution {}

use std::collections::HashMap;
// 解答
impl Solution {
    pub fn missing_number(nums: Vec<i32>) -> i32 {
        let mut max: i32 = 0;
        let mut find_max = false;
        let mut num_dick = HashMap::new();
        let mut value = 0;
        // 获取最大值
        for v in nums {
            if v > max {
                max = v;
            }
            num_dick.insert(v, 1);
        }
        let mut i = 0;
        while i <= max {
            if !num_dick.contains_key(&i) {
                value = i;
                find_max = true;
                break;
            }
            i += 1;
        }
        if !find_max {
            value = max + 1;
        }
        value
    }
}

// 控制台
#[warn(unused_mut)]
fn main() {
    assert_eq!(2, Solution::missing_number(vec![3, 0, 1]));
    assert_eq!(8, Solution::missing_number(vec![9, 6, 4, 2, 3, 5, 7, 0, 1]));
    assert_eq!(1, Solution::missing_number(vec![0]));
    assert_eq!(2, Solution::missing_number(vec![0, 1]));
}
