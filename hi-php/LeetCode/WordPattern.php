<?php
/**
 * Created by PhpStorm.
 * User: Joshua Young
 * Date: 2019/3/14
 * Time: 15:12
 * Link: https://leetcode-cn.com/problems/word-pattern/
 * Name: 290. 单词模式
 */

// 题目

/*
    给定一种 pattern(模式) 和一个字符串 str ，判断 str 是否遵循相同的模式。

    这里的遵循指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应模式。

    示例1:
        输入: pattern = "abba", str = "dog cat cat dog"
        输出: true

    示例 2:
        输入:pattern = "abba", str = "dog cat cat fish"
        输出: false

    示例 3:
        输入: pattern = "aaaa", str = "dog cat cat dog"
        输出: false

    示例 4:
        输入: pattern = "abba", str = "dog dog dog dog"
        输出: false

    说明:
        你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母。
 */

//解答
class Solution {

    /**
     * @param String $pattern
     * @param String $str
     * @return Boolean
     */
    function wordPattern($pattern, $str) {
        $ps = str_split($pattern);
        $ss = explode(' ', $str);
        $psLen = count($ps);
        $ssLen = count($ss);

        // 样式不一致直接判处为错误
        if($psLen != $ssLen || strlen($pattern) != $ssLen){
            return false;
        }
        // 全局比较
        for($i = 0; $i < $psLen; $i++){
            for($j = $i+1; $j < $psLen; $j++){
                if(($ps[$i] == $ps[$j]) != ($ss[$i] == $ss[$j])){
                    return false;
                }
            }
        }
        return true;

    }
}

// 控制台
class WordPattern
{
    const method = 'wordPattern';
    function __construct()
    {
        LeetCode::simpleTest(['abba', 'dog cat cat dog'], self::method, true);
        LeetCode::simpleTest(['abba', 'dog cat cat fish'], self::method, false);
        LeetCode::simpleTest(['aaaa', 'dog cat cat dog'], self::method, false);
        LeetCode::simpleTest(['abba', 'dog dog dog dog'], self::method, false);
        LeetCode::simpleTest(['', 'beef'], self::method, false);
    }

}