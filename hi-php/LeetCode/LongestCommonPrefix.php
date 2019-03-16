<?php
/**
 * Created by PhpStorm.
 * User: Joshua Conero
 * Date: 2019/3/16
 * Time: 11:07
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