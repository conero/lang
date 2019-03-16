<?php
/**
 * Created by PhpStorm.
 * User: Joshua Conero
 * Date: 2019/3/16
 * Time: 10:32
 * Link: https://leetcode-cn.com/problems/palindrome-number/
 * Name:  #9. 回文数
 */

//题目
/*
 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

    示例 1:
        输入: 121
        输出: true

    示例 2:
        输入: -121
        输出: false
    解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。

    示例 3:
        输入: 10
        输出: false
    解释: 从右向左读, 为 01 。因此它不是一个回文数。

    进阶:
        你能不将整数转为字符串来解决这个问题吗？
 */

//解答
class Solution {

    /**
     * @param Integer $x
     * @return Boolean
     */
    function isPalindrome($x) {
        if($x < 0){
            return false;
        }
        $sArr = str_split($x.'');
        $sArrLen = count($sArr);
        $isPalin = true;
        for($i = 0; $i < $sArrLen; $i++){
            if($i == ($sArrLen - $i - 1)){
                break;
            }
            if($sArr[$i] != $sArr[$sArrLen-$i-1]){
                $isPalin = false;
                break;
            }
        }
        return $isPalin;
    }
}


// 控制台
class PalindromeNumber
{
    const method = 'isPalindrome';
    function __construct()
    {
        LeetCode::simpleTest([121], self::method, true);
        LeetCode::simpleTest([-121], self::method, false);
        LeetCode::simpleTest([-10], self::method, false);
    }
}