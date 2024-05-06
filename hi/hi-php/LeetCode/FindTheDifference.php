<?php
/**
 * Created by PhpStorm.
 * User: Joshua Young
 * Date: 2019/3/14
 * Time: 15:34
 */

//题目
/*
    给定两个字符串 s 和 t，它们只包含小写字母。
    字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
    请找出在 t 中被添加的字母。

    示例:

        输入：
        s = "abcd"
        t = "abcde"

        输出：
        e

    解释：
    'e' 是那个被添加的字母。
 */


// 解答
class Solution {

    /**
     * @param String $s
     * @param String $t
     * @return String
     */
    function findTheDifference($s, $t) {
        $sq = str_split($s);
        $st = str_split($t);
        // 字母对比
        foreach ($sq as $sqv){
            foreach ($st as $stk => $stv){
                if($sqv == $stv){
                    $st[$stk] = '';
                    break;
                }
            }
        }
        return implode('', $st);
    }
}

// 可控制台
class FindTheDifference
{
    const method = 'findTheDifference';
    function __construct()
    {
        LeetCode::simpleTest(['abcd', 'abcde'], self::method, 'e');
        LeetCode::simpleTest(['a', 'aa'], self::method, 'a');
    }
}