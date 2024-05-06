<?php
/**
 * Created by PhpStorm.
 * User: Joshua Young
 * Date: 2019/3/14
 * Time: 16:14
 * Link: https://leetcode-cn.com/problems/binary-watch/
 * Name: 401. 二进制手表
 */

//题目

/*
 二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。
每个 LED 代表一个 0 或 1，最低位在右侧。

      [8, 4, 2, 1]
      [32, 16, 8, 4, 2, 1]

    [8, 4, 2', 1']
    [32, 16', 8', 4, 2, 1']
    例如，上面的二进制手表读取 “3:25”。
    给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。

    案例:
        输入: n = 1
        返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]

    注意事项:
        输出的顺序没有要求。
        小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
        分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。
 */


//解答
class Solution {

    /**
     * @param Integer $num
     * @return String[]
     */
    function readBinaryWatch($num) {
        $hour = [8, 4, 2, 1];
        $minute = [32, 16, 8, 4, 2, 1];
        // @TODO 需要实现
    }
}

//控制台
class BinaryWatch
{
    const method = 'readBinaryWatch';
    function __construct()
    {
        $intance = new Solution();
        $n = 1;
        print_r(LeetCode::simpleTestInutsToStr([$n, $intance->readBinaryWatch($n)]));
    }
}