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
        $sLen = strlen($s);
        for($i=0; $i < $sLen; $i++){
            $iw = substr($s, $i, ($i == 0? 1: $i));
            echo "$iw , $i, ".($i+1)." $s~\n";
            $tLen = strlen($t);
            for ($j=0; $j < $tLen; $j++){
                $jw = substr($t, $j, $j+1);
                print_r(LeetCode::simpleTestInutsToStr([$iw, $jw, $t]));
                echo "\n";
                if($iw == $jw){
                    //$t = substr($t, 0, $j).' '.substr($t, $j+1);
                    $t = substr($t, 0, $j).substr($t, $j+1);
                    break;
                }
            }
        }
        //$t = str_replace(' ', '', $t);
        return $t;
    }
}

// 可控制台
class FindTheDifference
{
    const method = 'findTheDifference';
    function __construct()
    {
        LeetCode::simpleTest(['abcd', 'abcde'], self::method, 'e');
    }
}