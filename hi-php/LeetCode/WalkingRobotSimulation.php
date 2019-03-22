<?php
/**
 * Created by PhpStorm.
 * User: Joshua Conero
 * Date: 2019/3/16
 * Time: 17:24
 * Email: conero@163.com
 * Link: https://leetcode-cn.com/problems/walking-robot-simulation/
 * Name: #874. 模拟行走机器人
 */


//题目
/*
 机器人在一个无限大小的网格上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令：
    -2：向左转 90 度
    -1：向右转 90 度
    1 <= x <= 9：向前移动 x 个单位长度
 在网格上有一些格子被视为障碍物。
    第 i 个障碍物位于网格点  (obstacles[i][0], obstacles[i][1])
    如果机器人试图走到障碍物上方，那么它将停留在障碍物的前一个网格方块上，但仍然可以继续该路线的其余部分。
    返回从原点到机器人的最大欧式距离的平方。


    示例 1：
        输入: commands = [4,-1,3], obstacles = []
        输出: 25
    解释: 机器人将会到达 (3, 4)

    示例 2：
        输入: commands = [4,-1,4,-2,4], obstacles = [[2,4]]
        输出: 65
    解释: 机器人在左转走到 (1, 8) 之前将被困在 (1, 4) 处

    提示：
        0 <= commands.length <= 10000
        0 <= obstacles.length <= 10000
        -30000 <= obstacle[i][0] <= 30000
        -30000 <= obstacle[i][1] <= 30000
        答案保证小于 2 ^ 31
 */

//解答
class Solution {

    /**
     * @param Integer[] $commands
     * @param Integer[][] $obstacles
     * @return Integer
     */
    function robotSim($commands, $obstacles) {
        $point = [0, 0];    // x, y
        $xy = 1;            // 转向, 默认北
        $dt = 1;            // 增量
        foreach ($commands as $c){
            // 移动转向
            if ($c == -2 || $c == -1){
                $xy = ($xy == 1)? 0 : 1;
                if($c == -2){       // 左，逆时针
                    if ($xy == 0){  // y -> x
                        $dt = ($dt == 1)? -1: 1;
                    }
                }else{   // 右
                    if ($xy == 1){ // x -> y
                        $dt = ($dt == 1)? -1: 1;
                    }
                }
//                echo LeetCode::simpleTestInutsToStr([$xy, $dt, '<=']). "\n";
                continue;
            }

            // 障碍物检测
            $cp = $point[$xy];
            $moveTo = $point[$xy] + $dt*$c;
            echo LeetCode::simpleTestInutsToStr(['start>', "c:$c", $point, "xy: $xy", "=>$dt", [$cp, '~', $moveTo]]). "\n";
            $obBreak = false;
            $getObstacles = [];             // 用于测试，障碍点
            foreach ($obstacles as $op){
                list($opx, $opy) = $op;
                if($xy){    // y 轴上移动
//                    if($opx == $point[0] && ($cp <= $point[1] && $point[1] <= $moveTo)){
                    if($opx == $point[0] && ($cp <= $point[1] && $point[1] <= $moveTo*$dt)){
//                        $point[$xy] = $point[1] - 1 * $dt;
//                        echo LeetCode::simpleTestInutsToStr([$c, $point, $op, "$opy - 1 * $dt", [$cp, $moveTo]]). "\n";
                        $point[$xy] = $opy - 1 * $dt;
                        $getObstacles = $op;
                        $obBreak = true;
                        break;
                    }
                }else{
//                    if($opy == $point[1] && ($cp <= $point[0] && $point[0] <= $moveTo*$dt)){
                    if($opy == $point[1] && ($cp <= $point[0] && $point[0] <= $moveTo)){
//                        $point[$xy] = $point[0] - 1 * $dt;
                        $point[$xy] = $opx - 1 * $dt;
                        $obBreak = true;
                        $getObstacles = $op;
                        break;
                    }
                }
            }
            echo LeetCode::simpleTestInutsToStr(['end>', $point, $getObstacles]). "\n";
            if($obBreak){
                continue;
            }
            $point[$xy] = $moveTo;
        }
        //echo LeetCode::simpleTestInutsToStr($point). "\n";
        return pow($point[0], 2) + pow($point[1], 2);
    }
}

//控制台
class WalkingRobotSimulation
{
    const METHOD = 'robotSim';
    function __destruct()
    {
//        LeetCode::simpleTest([[4,-1,3], []], self::METHOD, 25);
//        LeetCode::simpleTest([[4,-1,4,-2,4], [[2,4]]], self::METHOD, 65);
//        LeetCode::simpleTest([[-2,-1,8,9,6], [[-1,3],[0,1],[-1,5],[-2,-4],[5,4],[-2,-3],[5,-1],[1,-1],[5,5],[5,2]]], self::METHOD, 0);
//        LeetCode::simpleTest([[7,-2,-2,7,5], [[-3,2],[-2,1],[0,1],[-2,4],[-1,0],[-2,-3],[0,-3],[4,4],[-3,3],[2,2]]], self::METHOD, 4);
        LeetCode::simpleTest([[-2,8,3,7,-1], [[-4,-1],[1,-1],[1,4],[5,0],[4,5],[-2,-1],[2,-5],[5,1],[-3,-1],[5,-3]]], self::METHOD, 324);
    }
}