<?php
/**
 * 2017年11月27日 星期一
 * php 99乘法表
 */


// 乘法表实现
// 列显示
function Multi($bit=9){
    // 数据显示缓存器
    $stack = [];    
    $i = 1;
    while($i <= $bit){
        if(!isset($stack[$i-1])){
            $stack[$i-1] = [];
        }
        $row = [];
        $j = 1;
        // 用户显示
        while($j<$i){
            $row[] = '';
            $j++;
        }
        while($j <= $bit){
            $row[] = $i.' × '.$j.' = '. ($i*$j);
            $j++;
        }
        $stack[] = $row;
        $i++;
    }
    //print_r($stack);
    // $d = 0;
    // while($d<count($row)){        
    //     $d++;
    // }
}

// 行显示
function Multi2($bit=9){
    // 数据显示缓存器
    $stack = [];    
    $i = 1;
    while($i <= $bit){
        if(count($stack) == 0){
            $t = 0;
            while($t<=$bit){
                $stack[$t] = '';
                $t++;
            }
        }
        $row = [];
        $j = 1;
        // 用户显示
        while($j<$i){
            $row[] = '';
            $j++;
        }
        while($j <= $bit){
            $eq = $i.'×'.$j.' = '. str_pad(($i*$j).'', 2);
            $stack[$j] .= $eq.' ';
            $row[] = $eq;
            $j++;
        }
        $i++;
    }
    foreach($stack as $v){echo $v."\r\n";}
}

// 运行控制台
// Multi();
Multi2();