<?php
/**
 * Created by PhpStorm.
 * User: Joshua Young
 * Date: 2019/3/13
 * Time: 23:06
 * Link: https://leetcode-cn.com/problems/reverse-only-letters/
 * Name: 917. 仅仅反转字母
 */

// 题目

/*
    给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。
    示例 1：
        输入："ab-cd"
        输出："dc-ba"

    示例 2：
        输入："a-bC-dEf-ghIj"
        输出："j-Ih-gfE-dCba"

    示例 3：
        输入："Test1ng-Leet=code-Q!"
        输出："Qedo1ct-eeLg=ntse-T!"
 */

// 解答
class Solution {

    /**
     * @param String $S
     * @return String
     */
    function reverseOnlyLetters($S) {
        $data = str_split($S);
        $ct = count($data);
        $last = $ct;
        for($i=0; $i<$ct; $i++){
            $is = $data[$i];
            if($i == $last){
                break;
            }
            $isBreak = false;
            if(preg_match('/[a-zA-Z]{1}/', $is)){
                while(true){
                    $last -= 1;
                    if($i == $last){
                        $isBreak = true;
                        break;
                    }
                    $lastS = $data[$last];
                    //print_r(LeetCode::simpleTestInutsToStr([$is, $lastS, $i, $last, $i == $last, implode('', $data)]));
                    //echo "\n";
                    if(preg_match('/[a-zA-Z]{1}/', $lastS)){
                       $data[$i] = $lastS;
                       $data[$last] = $is;
                       break;
                    }
                }
                if($isBreak){
                    break;
                }
            }
        }
        return implode('', $data);
    }
}

// 控制台
class ReverseOnlyLetters
{
    const method = 'reverseOnlyLetters';
    function __construct()
    {
        LeetCode::simpleTest(['ab-cd'], self::method, 'dc-ba');
        LeetCode::simpleTest(['a-bC-dEf-ghIj'], self::method, 'j-Ih-gfE-dCba');
        LeetCode::simpleTest(['Test1ng-Leet=code-Q!'], self::method, 'Qedo1ct-eeLg=ntse-T!');
    }
}