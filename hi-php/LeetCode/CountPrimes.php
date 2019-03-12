<?php
/**
 * Created by PhpStorm.
 * User: Joshua Young
 * Date: 2019/3/12
 * Time: 15:00
 * Link: https://leetcode-cn.com/problems/count-primes/
 * Name: 204. 计数质数
 */


// 问题
/*
 统计所有小于非负整数 n 的质数的数量。

    示例:
        输入: 10
        输出: 4
    解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 */

// 解答
/*
    质数（prime number）又称素数，有无限个。
    质数定义为在大于1的自然数中，除了1和它本身以外不再有其他因数。
    质数又称素数。一个大于1的自然数，除了1和它自身外，不能被其他自然数整除的数叫做质数；否则称为合数。
 */
class Solution {
    /**
     * @param Integer $n
     * @return Integer
     */
    function countPrimes($n) {
        $primes = [];
        for($i=2; $i<$n; $i++){
            $isPrimes = true;
            for($j=2; $j<$i; $j++){
                if($i%$j == 0){
                    $isPrimes = false;
                }
            }
            if($isPrimes){
                $primes[] = $i;
            }
        }
        return count($primes);
    }
}


// 控制台
class CountPrimes
{
    function __construct()
    {
        $method = 'countPrimes';
        LeetCode::simpleTest([10], $method, 4);
        LeetCode::simpleTest([2], $method, 0);
    }
}