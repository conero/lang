<?php
/**
 * Created by PhpStorm.
 * User: Joshua Conero
 * Date: 2019/4/5
 * Time: 21:52
 * Email: conero@163.com
 * Link: https://leetcode-cn.com/problems/sort-colors/
 * Name: 75. 颜色分类
 */



// 题目
/*
 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
    此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

注意:
不能使用代码库中的排序函数来解决这道题。

示例:
    输入: [2,0,2,1,1,0]
    输出: [0,0,1,1,2,2]

进阶：
    一个直观的解决方案是使用计数排序的两趟扫描算法。
    首先，迭代计算出0、1 和 2 元素的个数，然后按照0、1、2的排序，重写当前数组。
    你能想出一个仅使用常数空间的一趟扫描算法吗？
 */


// 解答
class Solution {

    /**
     * @param Integer[] $nums
     * @return NULL
     */
    function sortColors(&$nums) {
        // 0, 1, 2
        $dick = [
            0 => [],
            1 => [],
            2 => []
        ];
        foreach($nums as $v){
            if(isset($dick[$v])){
                $dick[$v][] = $v;
            }
        }
        $tmpArry = [];
        foreach ($dick as $dk){
            $tmpArry = array_merge($tmpArry, $dk);
        }
        $nums = $tmpArry;
        return null;
    }
}


// 控制台
class SortColors
{
    const METHOD = 'sortColors';
    function __construct()
    {
        // test 1
        $num = [2,0,2,1,1,0];
        $instance = new Solution();
        $input = $num;
        $instance->sortColors($num);
        print_r(LeetCode::simpleTestInutsToStr([$input, $num]));


        // test 2
        echo "\n";
        $num = [2, 0, 1, 0, 2, 0, 1, 0, 2, 0 ,1, 2, 0, 2, 1];
        $instance = new Solution();
        $input = $num;
        $instance->sortColors($num);
        print_r(LeetCode::simpleTestInutsToStr([$input, $num]));
    }
}