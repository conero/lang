<?php
/**
 * Created by PhpStorm.
 * User: Joshua Conero
 * Date: 2019/3/16
 * Time: 11:07
 * Link: https://leetcode-cn.com/problems/longest-common-prefix/
 * Name: #14. 最长公共前缀
 */

//题目
/*
    编写一个函数来查找字符串数组中的最长公共前缀。
    如果不存在公共前缀，返回空字符串 ""。

    示例 1:
        输入: ["flower","flow","flight"]
        输出: "fl"

    示例 2:

        输入: ["dog","racecar","car"]
        输出: ""
    解释: 输入不存在公共前缀。

    说明:
        所有输入只包含小写字母 a-z 。
 */

//解答
class Solution {

    /**
     * @param String[] $strs
     * @return String
     */
    function longestCommonPrefix($strs) {
        $prefix = '';
        if(empty($strs)){
            return $prefix;
        }

        $i = 1;
        while (true){
            $w = '';
            $isBreak = false;
            foreach ($strs as $s){
                if($i > strlen($s)){
                    $isBreak = true;
                    break;
                }
                if(empty($w)){
                    $w = substr($s, $i-1, 1);
                }else{
                    $wTmp = substr($s, $i-1, 1);
                    if($w != $wTmp){
                        $isBreak = true;
                        break;
                    }
                }
                //print_r([$w, $isBreak, $i]);
            }
            if($isBreak){
                break;
            }
            $prefix .= $w;
            $i += 1;
        }
        return $prefix;
    }
}

//控制台
class LongestCommonPrefix
{
    const METHOD = 'longestCommonPrefix';
    function __construct()
    {
        LeetCode::simpleTest([["flower","flow","flight"]], self::METHOD, 'fl');
        LeetCode::simpleTest([["dog","racecar","car"]], self::METHOD, '');
        LeetCode::simpleTest([[]], self::METHOD, '');
    }
}