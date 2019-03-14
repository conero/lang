<?php
/**
 * Created by PhpStorm.
 * User: Joshua Young
 * Date: 2019/3/14
 * Time: 14:20
 * Link: https://leetcode-cn.com/problems/contains-duplicate-ii/
 * Name: 219. 存在重复元素 II
 */

// 题目
/*
    给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。

    示例 1:
        输入: nums = [1,2,3,1], k = 3
        输出: true

    示例 2:
        输入: nums = [1,0,1,1], k = 1
        输出: true

    示例 3:
        输入: nums = [1,2,3,1,2,3], k = 2
        输出: false
 */

// 解答
class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Boolean
     */
    function containsNearbyDuplicate($nums, $k) {
        $ct = count($nums);
        for($i=0; $i < $ct; $i++){
            for($j = $i+1; $j < $ct; $j++){
                if($nums[$i] == $nums[$j] && abs($i-$j) == $k){
                    return true;
                }
            }
        }
        return false;
    }
}

// 控制台
class ContainsDuplicateII
{
    const method = 'containsNearbyDuplicate';
    function __construct()
    {
        LeetCode::simpleTest([[1,2,3,1], 3], self::method, true);
        LeetCode::simpleTest([[1,0,1,1], 1], self::method, true);
        LeetCode::simpleTest([[1,2,3,1,2,3], 2], self::method, false);
        LeetCode::simpleTest([[99, 99], 2], self::method, true);
    }

}