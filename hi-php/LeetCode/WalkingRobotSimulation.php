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

// @TODO StillNeedToDo-算法实现
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
//            echo LeetCode::simpleTestInutsToStr(['start>', "c:$c", $point, "xy: $xy", "=>$dt", [$cp, '~', $moveTo]]). "\n";
            $obBreak = false;
            $getObstacles = [];             // 用于测试，障碍点
            foreach ($obstacles as $op){
                list($opx, $opy) = $op;
                if($xy){    // y 轴上移动
                    $originalV = $point[1];
                    if($opx == $point[0]){
                        if($dt > 0 && $moveTo >= $opy && $cp < $opy){ // -> +
                            $point[$xy] = $opy - 1;
                            $getObstacles = $op;
                            $obBreak = true;
                            break;
                        }elseif ($dt < 0 && $moveTo <= $opy && $cp > $opy){
                            $point[$xy] = $opy + 1;
                            $getObstacles = $op;
                            $obBreak = true;
                            break;
                        }
                    }
                }else{          // x 轴上的移动
                    $originalV = $point[0];
                    if($opy == $point[1]){
                        if($dt > 0 && $moveTo >= $opx && $cp < $opx){ // -> +
                            $point[$xy] = $opx - 1;
                            $getObstacles = $op;
                            $obBreak = true;
                            break;
                        }elseif ($dt < 0 && $moveTo <= $opx && $cp > $opx){
                            $point[$xy] = $opx + 1;
                            $getObstacles = $op;
                            $obBreak = true;
                            break;
                        }
                    }
                }
            }
            if($obBreak){
//                echo LeetCode::simpleTestInutsToStr(['end>', $point, $getObstacles, $obBreak]). "\n";
                continue;
            }
            $point[$xy] = $moveTo;
//            echo LeetCode::simpleTestInutsToStr(['end>', $point, $getObstacles, $obBreak]). "\n";
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
        LeetCode::simpleTest([[4,-1,3], []], self::METHOD, 25);
        LeetCode::simpleTest([[4,-1,4,-2,4], [[2,4]]], self::METHOD, 65);
        LeetCode::simpleTest([[-2,-1,8,9,6], [[-1,3],[0,1],[-1,5],[-2,-4],[5,4],[-2,-3],[5,-1],[1,-1],[5,5],[5,2]]], self::METHOD, 0);
        LeetCode::simpleTest([[7,-2,-2,7,5], [[-3,2],[-2,1],[0,1],[-2,4],[-1,0],[-2,-3],[0,-3],[4,4],[-3,3],[2,2]]], self::METHOD, 4);
        LeetCode::simpleTest([[-2,8,3,7,-1], [[-4,-1],[1,-1],[1,4],[5,0],[4,5],[-2,-1],[2,-5],[5,1],[-3,-1],[5,-3]]], self::METHOD, 324);
        LeetCode::simpleTest([[1,1,3,4,3], [[-1,5],[-4,-4],[-3,3],[3,0],[2,5],[-4,4],[-3,1],[-2,-4],[-1,-4],[0,-3]]], self::METHOD, 144);
        LeetCode::simpleTest([[-2,-1,-2,3,7], [[1,-3],[2,-3],[4,0],[-2,5],[-5,2],[0,0],[4,-4],[-2,-5],[-1,-2],[0,2]]], self::METHOD, 100);

        // 压力测试
        $cmd = [1,2,-2,5,-1,-2,-1,8,3,-1,9,4,-2,3,2,4,3,9,2,-1,-1,-2,1,3,-2,4,1,4,-1,1,9,-1,-2,5,-1,5,5,-2,6,6,7,7,2,8,9,-1,7,4,6,9,
            9,9,-1,5,1,3,3,-1,5,9,7,4,8,-1,-2,1,3,2,9,3,-1,-2,8,8,7,5,-2,6,8,4,6,2,7,2,-1,7,-2,3,3,2,-2,6,9,8,1,-2,-1,1,4,7];
        $obs = [[-57,-58],[-72,91],[-55,35],[-20,29],[51,70],[-61,88],[-62,99],[52,17],[-75,-32],[91,-22],[54,33],[-45,-59],[47,-48],[53,-98],
            [-91,83],[81,12],[-34,-90],[-79,-82],[-15,-86],[-24,66],[-35,35],[3,31],[87,93],[2,-19],[87,-93],[24,-10],[84,-53],[86,87],[-88,-18],
            [-51,89],[96,66],[-77,-94],[-39,-1],[89,51],[-23,-72],[27,24],[53,-80],[52,-33],[32,4],[78,-55],[-25,18],[-23,47],[79,-5],[-23,-22],
            [14,-25],[-11,69],[63,36],[35,-99],[-24,82],[-29,-98],[-50,-70],[72,95],[80,80],[-68,-40],[65,70],[-92,78],[-45,-63],[1,34],[81,50],
            [14,91],[-77,-54],[13,-88],[24,37],[-12,59],[-48,-62],[57,-22],[-8,85],[48,71],[12,1],[-20,36],[-32,-14],[39,46],[-41,75],[13,-23],
            [98,10],[-88,64],[50,37],[-95,-32],[46,-91],[10,79],[-11,43],[-94,98],[79,42],[51,71],[4,-30],[2,74],[4,10],[61,98],[57,98],[46,43],
            [-16,72],[53,-69],[54,-96],[22,0],[-7,92],[-69,80],[68,-73],[-24,-92],[-21,82],[32,-1],[-6,16],[15,-29],[70,-66],[-85,80],[50,-3],
            [6,13],[-30,-98],[-30,59],[-67,40],[17,72],[79,82],[89,-100],[2,79],[-95,-46],[17,68],[-46,81],[-5,-57],[7,58],[-42,68],[19,-95],
            [-17,-76],[81,-86],[79,78],[-82,-67],[6,0],[35,-16],[98,83],[-81,100],[-11,46],[-21,-38],[-30,-41],[86,18],[-68,6],[80,75],[-96,-44],
            [-19,66],[21,84],[-56,-64],[39,-15],[0,45],[-81,-54],[-66,-93],[-4,2],[-42,-67],[-15,-33],[1,-32],[-74,-24],[7,18],[-62,84],[19,61],
            [39,79],[60,-98],[-76,45],[58,-98],[33,26],[-74,-95],[22,30],[-68,-62],[-59,4],[-62,35],[-78,80],[-82,54],[-42,81],[56,-15],[32,-19],
            [34,93],[57,-100],[-1,-87],[68,-26],[18,86],[-55,-19],[-68,-99],[-9,47],[24,94],[92,97],[5,67],[97,-71],[63,-57],[-52,-14],[-86,-78],
            [-17,92],[-61,-83],[-84,-10],[20,13],[-68,-47],[7,28],[66,89],[-41,-17],[-14,-46],[-72,-91],[4,52],[-17,-59],[-85,-46],[-94,-23],[-48,-3],
            [-64,-37],[2,26],[76,88],[-8,-46],[-19,-68]];
        LeetCode::simpleTest([$cmd, $obs], self::METHOD, 5140);
    }
}