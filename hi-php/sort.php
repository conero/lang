<?php
/**
 * 2017年11月14日 星期二
 * 排序算法测试实现
 */

 const BR = "\r\n";
 // 运行技术
function microtime_float()
{
    list($usec, $sec) = explode(" ", microtime());
    return ((float)$usec + (float)$sec);
}

// 返回获取计算统计
function startRun(){
    $timeStart = microtime_float();
    $getRunTime = function() use ($timeStart){
        return microtime_float() - $timeStart;
    };
    return $getRunTime;
}

// 获取随机数组
function getRandArray($bit=10){
    $array = [];
    $i = 1;
    while($i<=$bit){
        $array[] = rand(1, 1000);
        $i++;
    }
    return $array;
}

// 看不懂程序的写法
// 参考网上 - 百度百科
function shell_sort($arr)
{
    if(!is_array($arr)) return;
    $n = count($arr);
    for($gap = floor($n/2);$gap > 0;$gap = floor($gap/=2))
    {
        for($i = $gap;$i < $n;++$i)
        {
            for($j = $i - $gap;$j >= 0 && $arr[$j + $gap] < $arr[$j]; $j -= $gap)
            {
                $temp = $arr[$j];
                $arr[$j] = $arr[$j+$gap];
                $arr[$j + $gap] = $temp;
            }
        }
    }
    return $arr;
}


class Sort{
    // array() => array()
    public static function shellSort($array){
        $len = count($array);
        $i = $len;
        echo 'array length is '.$len.BR;
        while($i > 1){
            // $i = ceil($i/2);
            $i = floor($i/2);
            $j = 0;
            echo $i.BR;
            //while($j<$i){
            while(($j+$i) < $len){
                $m1 = $array[$j];
                $m2 = $array[$j+$i];
                // echo $j.", ".($j+$i)."; ".$m1.", ".$m2."; ".BR;
                // echo $j.", ".($j+$i)."; ".BR;
                if($m2 < $m1){
                    $array[$j] = $m2;
                    $array[$j+$i] = $m1;
                }
                $j++;
            }
        }
        return $array;
    }
    // 测试写法二
    // 希尔排序在数组中采用跳跃式分组的策略，通过某个增量将数组元素划分为若干组，然后分组进行插入排序，随后逐步缩小增量，继续按组进行插入排序操作，直至增量为1。
    public static function shell($arr){
        // 取增量
        $len = count($arr);
        for($inc = floor($len/2); $inc>0; $inc = floor($inc/2)){
            echo $inc.BR;
            // 按照增量分组
            for($i = 0; ($i+$inc) < $len; $i++){
                // 组内插入排序
                for($j=$i; $j<($i+$inc); $j++){        
                    $n1 = $arr[$j];
                    $n2 = $arr[$j+1];
                    if($n2 < $n1){
                        $arr[$j] = $n2;
                        $arr[$j+1] = $n1;
                    }
                }
            }
        }
        return $arr;
    }
}
// 测试函数
class Test{
    public function shellSort(){
        // $ba = getRandArray(20*10000);
        $ba = getRandArray(10);
        $start = startRun();
        // $sa = Sort::shellSort($ba);
        // $sa = Sort::shell($ba);
        $sa = shell_sort($ba);
        $sec = $start();
        echo BR;
        $msgArr = [
            'basearray' => json_encode($ba),
            'sell_sort' => json_encode($sa),
            'runtimes'  => $sec
        ];
        // echo json_encode($msgArr).BR;
        print_r($msgArr);
    }
}
// 运行统计数
function console(){
    $test = new Test();
    $test->shellSort();
}
console();