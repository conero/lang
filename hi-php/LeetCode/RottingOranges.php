<?php
/**
 * Created by PhpStorm.
 * User: Joshua Young
 * Date: 2019/3/13
 * Time: 22:18
 * Link: https://leetcode-cn.com/problems/rotting-oranges/
 * Name: 994. 腐烂的橘子
 */

// 题目

/*
     在给定的网格中，每个单元格可以有以下三个值之一：

        值 0 代表空单元格；
        值 1 代表新鲜橘子；
        值 2 代表腐烂的橘子。
    每分钟，任何与腐烂的橘子（在 4 个正方向上）相邻的新鲜橘子都会腐烂。
    返回直到单元格中没有新鲜橘子为止所必须经过的最小分钟数。如果不可能，返回 -1。

    示例 1：
        输入：[[2,1,1],[1,1,0],[0,1,1]]
        输出：4

    示例 2：
        输入：[[2,1,1],[0,1,1],[1,0,1]]
        输出：-1
    解释：左下角的橘子（第 2 行， 第 0 列）永远不会腐烂，因为腐烂只会发生在 4 个正向上。

    示例 3：
        输入：[[0,2]]
        输出：0
    解释：因为 0 分钟时已经没有新鲜橘子了，所以答案就是 0 。

    提示：
        1 <= grid.length <= 10
        1 <= grid[0].length <= 10
        grid[i][j] 仅为 0、1 或 2
 */

// 解答
class Solution {

    /**
     * @param Integer[][] $grid
     * @return Integer
     */
    function orangesRotting($grid) {
        $m = count($grid);
        $count = -1;
        while (true){
            $rottingMk = false;
            for($i=0; $i<$m; $i++){
                $n = count($grid[$i]);
                for ($j=0; $j<$n; $j++){
                    $v = $grid[$i][$j];
                    switch ($v){
                        case 2:
                            if($i>0 && $grid[$i-1][$j] == 1){ // 上
                                $grid[$i-1][$j] = 2;
                                $rottingMk = true;
                            }
                            if($i<$m-1 && $grid[$i+1][$j] == 1){ // 下
                                $grid[$i+1][$j] = 2;
                                $rottingMk = true;
                            }
                            if($j>0 && $grid[$i][$j-1] == 1){
                                $grid[$i][$j-1] = 2;
                                $rottingMk = true;
                            }
                            if($j<$n-1 && $grid[$i][$j+1] == 1){
                                $grid[$i][$j+1] = 2;
                                $rottingMk = true;
                            }
                            break;
                    }
                }
            }
            // 次数检测
            if(!$rottingMk){
                break;
            }else{
                $count = $count == -1? 1 : ($count + 1);
            }
            // print_r($grid);
            // print_r([LeetCode::simpleTestInutsToStr($grid)]);

        }
        return $count;
    }
}


// 控制台
class RottingOranges
{
    function __construct()
    {
        $method = 'orangesRotting';
        LeetCode::simpleTest([[[2,1,1],[1,1,0],[0,1,1]]], $method, 4);
        LeetCode::simpleTest([[[2,1,1],[0,1,1],[1,0,1]]], $method, -1);
        LeetCode::simpleTest([[[0,2]]], $method, -1);
    }
}