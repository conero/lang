// 2018年12月27日 星期四
// @link https://leetcode-cn.com/problems/plus-one/

/*
    给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

    最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

    你可以假设除了整数 0 之外，这个整数不会以零开头。

    示例 1:

    输入: [1,2,3]
    输出: [1,2,4]
    解释: 输入数组表示数字 123。
    示例 2:

    输入: [4,3,2,1]
    输出: [4,3,2,2]
    解释: 输入数组表示数字 4321。
*/

pub struct Solution{
}

// -------------------------[ script ]--------------------------------------------
impl Solution {
    pub fn plus_one(digits: Vec<i32>) -> Vec<i32> {
        let mut i  = digits.len();
        let mut new_dgt = digits.clone();
        loop {
            i -= 1;
            let n = new_dgt[i] + 1;
            if n >= 10{
                new_dgt[i] = n - 10;

                // 末尾进一
                if i == 0{
                    let mut add_dgt = vec![1];
                    for nx in &new_dgt{
                        add_dgt.push(*nx);
                    }
                    new_dgt = add_dgt;
                    break;
                }
            }else {
                new_dgt[i] = n;
                break
            }
        }
        new_dgt
    }
}
// -------------------------[ script ]--------------------------------------------


#[warn(unused_mut)]
fn main() {
//    let digists = vec![1, 2, 4];
//    println!("{:?} => {:?}", digists.clone(), Solution::plus_one(digists));
//    println!("{:?} => {:?}", vec![2, 3, 5], Solution::plus_one(vec![2, 3, 5]));


    // 测试用例
//    digists = vec![1, 2, 4];
//    assert_eq!(vec![1, 2, 5], Solution::plus_one(digists));


    assert_eq!(vec![1,2,4], Solution::plus_one(vec![1,2,3]));
    assert_eq!(vec![1,0], Solution::plus_one(vec![9]));
    assert_eq!(vec![1,0,0], Solution::plus_one(vec![9,9]));
//    Solution::plus_one(vec![9,9]);
}
