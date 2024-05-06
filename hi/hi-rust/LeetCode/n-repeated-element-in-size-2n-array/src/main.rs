// 2018年12月27日 星期四
// @link https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array/

/*
   在大小为 2N 的数组 A 中有 N+1 个不同的元素，其中有一个元素重复了 N 次。
    返回重复了 N 次的那个元素。

    示例 1：
        输入：[1,2,3,3]
        输出：3

    示例 2：
        输入：[2,1,2,5,3,2]
        输出：2

    示例 3：
        输入：[5,1,5,2,5,3,5,4]
        输出：5

    提示：

    4 <= A.length <= 10000
    0 <= A[i] < 10000
    A.length 为偶数
*/

pub struct Solution{
}

// -------------------------[ script ]--------------------------------------------
use std::collections::HashMap;
impl Solution {
    pub fn repeated_n_times(a: Vec<i32>) -> i32 {
        let len = a.len();
        let mut dicks: HashMap<i32, i32> = HashMap::new();
        for n in &a{
            if dicks.contains_key(n){
                // let n1 = *n;
                // dicks[&n1] += 1;
                let add_n = *dicks.get_mut(n).unwrap()+1;
                dicks.insert(*n, add_n);
            }else {
                dicks.insert(*n, 1);
            }
        }

        // 截取值
        // 字典变量
        for (n, time) in dicks{
            let half = (len/2) as i32;
            if time == half{
                return n;
            }
        }
        0
    }


}
// -------------------------[ script ]--------------------------------------------


#[warn(unused_mut)]
fn main() {
    let digists = vec![5,1,5,2,5,3,5,4];
    println!("{:?} => {}", digists.clone(), Solution::repeated_n_times(digists));
//    println!("{:?} => {:?}", vec![2, 3, 5], Solution::plus_one(vec![2, 3, 5]));


    // 测试用例
//    digists = vec![1, 2, 4];
//    assert_eq!(vec![1, 2, 5], Solution::plus_one(digists));


//    assert_eq!(vec![1,2,4], Solution::plus_one(vec![1,2,3]));
//    assert_eq!(vec![1,0], Solution::plus_one(vec![9]));
//    assert_eq!(vec![1,0,0], Solution::plus_one(vec![9,9]));
//    Solution::plus_one(vec![9,9]);

    assert_eq!(3, Solution::repeated_n_times(vec![1,2,3,3]));
    assert_eq!(2, Solution::repeated_n_times(vec![2,1,2,5,3,2]));
    assert_eq!(5, Solution::repeated_n_times(vec![5,1,5,2,5,3,5,4]));
}
