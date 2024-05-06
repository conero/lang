// 2018年12月28日 星期五
// @link https://leetcode-cn.com/problems/verifying-an-alien-dictionary/

/*
    某种外星语也使用英文小写字母，但可能顺序 order 不同。字母表的顺序（order）是一些小写字母的排列。

    给定一组用外星语书写的单词 words，以及其字母表的顺序 order，只有当给定的单词在这种外星语中按字典序排列时，返回 true；否则，返回 false。



    示例 1：

        输入：words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
        输出：true
        解释：在该语言的字母表中，'h' 位于 'l' 之前，所以单词序列是按字典序排列的。

    示例 2：

        输入：words = ["word","world","row"], order = "worldabcefghijkmnpqstuvxyz"
        输出：false
        解释：在该语言的字母表中，'d' 位于 'l' 之后，那么 words[0] > words[1]，因此单词序列不是按字典序排列的。

    示例 3：

        输入：words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
        输出：false
        解释：当前三个字符 "app" 匹配时，第二个字符串相对短一些，然后根据词典编纂规则 "apple" > "app"，因为 'l' > '∅'，其中 '∅' 是空白字符，定义为比任何其他字符都小（更多信息）。


    提示：

        1 <= words.length <= 100
        1 <= words[i].length <= 20
        order.length == 26
        在 words[i] 和 order 中的所有字符都是英文小写字母。
*/

pub struct Solution{
}

// -------------------------[ script ]--------------------------------------------
use std::collections::HashMap;
impl Solution {
    pub fn is_alien_sorted(words: Vec<String>, order: String) -> bool {
        // 生成顺序字典
        let mut order_dick: HashMap<char, i32> = HashMap::new();
        let mut ctt = 0;
        for c in order.chars(){
            order_dick.insert(c, ctt);
            ctt += 1;
        }

        // 顺序比较
        let words_len = words.len()-1;
        let mut is_include = true;         // 是否被包含
        for i in 0..words_len {
            let word = &words[i];
            let next_word = &words[i+1];
            let word_len = word.len();
            let mut max_len = word_len;
            let next_wlen = next_word.len();
            if max_len < next_wlen{
                max_len = next_wlen;
            }

            // 字母比较
            for j in 0..max_len {
                if j < word_len && j < next_wlen{
                    let c0 = word.clone().remove(j);
                    let c1 = next_word.clone().remove(j);
                    // let c0 = word.remove(j);
                    // let c1 = next_word.remove(j);
                    if c0 != c1{
                        if order_dick.get(&c0).unwrap() > order_dick.get(&c1).unwrap(){
                            return false;
                        }else {
                            is_include = false;
                        }
                    }else {
                        if !is_include{
                            return true;
                        }
                    }
                }
                // 被包含，且次单次比上一个长
                if is_include && next_wlen < word_len{
                    return false;
                }
            }
        }

        true
    }
}
// -------------------------[ script ]--------------------------------------------

fn from_str_vec(dd: Vec<&str>) -> Vec<String>{
    let mut new_dd: Vec<String> = Vec::new();
    for s in dd{
        new_dd.push(String::from(s));
    }
    new_dd
}

#[warn(unused_mut)]
fn main() {
    let words = from_str_vec(vec!["mtkwpj","wlaees"]);
    let order = "evhadxsqukcogztlnfjpiymbwr";
    println!("{}", Solution::is_alien_sorted(words, String::from(order)));

    // 测试用例
//    assert_eq!(true, Solution::is_alien_sorted(from_str_vec(vec!["hello","leetcode"]), "hlabcdefgijkmnopqrstuvwxyz".to_string()));
//    assert_eq!(false, Solution::is_alien_sorted(from_str_vec(vec!["word","world","row"]), "worldabcefghijkmnpqstuvxyz".to_string()));
//    assert_eq!(false, Solution::is_alien_sorted(from_str_vec(vec!["apple","app"]), "abcdefghijklmnopqrstuvwxyz".to_string()));
    // //本测试用例【无法理解】
    //assert_eq!(true, Solution::is_alien_sorted(from_str_vec(vec!["mtkwpj","wlaees"]), "evhadxsqukcogztlnfjpiymbwr".to_string()));
}



