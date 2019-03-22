<?php
/**
 * Created by PhpStorm.
 * User: Joshua Conero
 * Date: 2019/3/22
 * Time: 14:03
 * Email: conero@163.com
 * Link: https://leetcode-cn.com/problems/valid-parentheses/
 * Name: #20. 有效的括号
 */

/*
 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
    有效字符串需满足：
        左括号必须用相同类型的右括号闭合。
        左括号必须以正确的顺序闭合。
        注意空字符串可被认为是有效字符串。

    示例 1:
        输入: "()"
        输出: true

    示例 2:
        输入: "()[]{}"
        输出: true

    示例 3:
        输入: "(]"
        输出: false

    示例 4:
        输入: "([)]"
        输出: false

    示例 5:
        输入: "{[]}"
        输出: true
 */

// @TODO StillNeedToDo-算法实现
// 解答
class Solution {

    /**
     * @param String $s
     * @return Boolean
     */
    function isValid($s) {
        // [], (), {}
        $pregTestValid = true;
        // ()
//        echo LeetCode::simpleTestInutsToStr(["-", preg_match('/[\(\)]+/', $s), $s])."\n";
        if(preg_match('/[\(\)]+/', $s) && !preg_match('/\([^\(\)]*\)/', $s)){
            $pregTestValid = false;
        }
        // []
        if($pregTestValid && preg_match('/[\[\]]+/', $s) && !preg_match('/\[[^\[\]]*\]/', $s)){
            $pregTestValid = false;
        }
        //{}
        if($pregTestValid && preg_match('/[\{\}]+/', $s) && !preg_match('/\[[^\[\]]*\]/', $s)){
            $pregTestValid = false;
        }
        return $pregTestValid;
    }
}

// 控制台
class ValidParentheses
{
    const tMETHOD = 'isValid';
    function __construct()
    {
        LeetCode::simpleTest(['()'], self::tMETHOD, true);
        LeetCode::simpleTest(['()[]{}'], self::tMETHOD, true);
        LeetCode::simpleTest(['(]'], self::tMETHOD, false);
        LeetCode::simpleTest(['([)]'], self::tMETHOD, false);
        LeetCode::simpleTest(['{[]}'], self::tMETHOD, true);
    }
}