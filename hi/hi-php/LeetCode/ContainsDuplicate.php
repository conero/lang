<?php
/**
 * Created by PhpStorm.
 * User: Joshua Young
 * Date: 2019/3/14
 * Time: 14:20
 * Link: https://leetcode-cn.com/problems/contains-duplicate/
 * Name: 217. 存在重复元素
 */

// 题目
/*
    给定一个整数数组，判断是否存在重复元素。
    如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。

    示例 1:
        输入: [1,2,3,1]
        输出: true

    示例 2:
        输入: [1,2,3,4]
        输出: false

    示例 3:
        输入: [1,1,1,3,3,4,3,2,4,2]
        输出: true
 */

// 解答
class Solution {

    /**
     * @param Integer[] $nums
     * @return Boolean
     */
    function containsDuplicate($nums) {
        $dick = []; // 字典
        foreach ($nums as $v){
            if(!($dick[$v] ?? false)){
                $dick[$v] = true;
                continue;
            }
            return true;
        }
        return false;
    }
}

// 控制台
class ContainsDuplicate
{
    const method = 'containsDuplicate';
    function __construct()
    {
        LeetCode::simpleTest([[1,2,3,1]], self::method, true);
        LeetCode::simpleTest([[1,2,3,4]], self::method, false);
        LeetCode::simpleTest([[1,1,1,3,3,4,3,2,4,2]], self::method, true);
    }

}