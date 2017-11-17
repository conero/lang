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

// [909,446,254,446,906,258,3,878,869,723,161,891,939]
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
    // 冒泡排序
    // Joshua Conero
    public static function bubble($array){
        $len = count($array);
        for($i=0; $i<$len; $i++){
            for($j=$i+1; $j<$len; $j++){
                if($array[$j] < $array[$i]){
                    $m1 = $array[$i];     
                    $array[$i] = $array[$j];
                    $array[$j] = $m1;
                }
            }
        }
        return $array;
    }
    // 插入排序 - Joshua Conero
    public static function insert($array){
        $len = count($array);
        for($i=0; $i<$len-1; $i++){
            // 找到剩余列表中的最小值
            // 并将最小值推到末尾        
            for($j=$i; $j<$len-1; $j++){
                if($array[$j] < $array[$j+1]){
                    $m0 = $array[$j+1];
                    $array[$j+1] = $array[$j];
                    $array[$j] = $m0;
                }
            }
            // 有序数组
            // 将最小值插入到指定的位置
            $min = $array[$len-1];
            $k = 0;
            while($k<$i){
                if($array[$k] > $min){
                    $m1 = $array[$k];   // 替换值
                    $array[$k] = $min;
                    $min = $m1;
                }
                $k++;
            }
            $array[$len-1] = $array[$i];
            $array[$i] = $min;            
        }
        return $array;
    }    
    // 快速排序
    public static function quick($array){
        $aLen = count($array);
        $alice = floor($aLen/2);
        for($i=$alice; $i<$aLen; $i++){
            for($j=0; $j<$alice; $j++){
                if($array[$i] < $array[$j]){
                    $m = $array[$i];
                    $array[$i] = $array[$j];
                    $array[$j] = $m;
                }
            }
        }
        // left 任何一个值都比 right要小
        $left = array_slice($array, 0, $alice);
        $right = array_slice($array, $alice);
        if(count($left)>1){
            $left = Sort::quick($left);
        }    
        if(count($right)>1){
            $right = Sort::quick($right);
        }    
        $array = array_merge($left, $right);
        // /*
        echo BR;
        $msgArr = [
            'basearray`' => json_encode($array),
            'left' => json_encode($left),
            'right' => json_encode($right)
        ];
        // echo json_encode($msgArr).BR;
        print_r($msgArr);
        // die;
        // */

        //return $array;
    }
    // 直接选择排序
    public static function select($array){
        $aLen = count($array);
        // 确定位置
        for($i=0; $i<$aLen-1; $i++){
            // 排剩下的数组序列
            for($j=$i; $j<$aLen-1; $j++){
                if($array[$j] < $array[$j+1]){
                    $n1 = $array[$j+1];
                    $array[$j+1] = $array[$j];
                    $array[$j] = $n1;
                }
            }
            $n1 = $array[$i];
            $array[$i] = $array[$aLen-1];
            $array[$aLen-1] = $n1;
        }
        return $array;
    }
    // 堆排序
    // 归并排序
    // 基数排序
}
// 测试函数
class Test{
    public function shellSort(){
        // $ba = getRandArray(20*10000);
        $ba = getRandArray(13);
        $start = startRun();
        // $sa = Sort::shellSort($ba);
        // $sa = Sort::shell($ba);
        $sa = shell_sort($ba);
        $sBubble = Sort::bubble($ba);
        $sInsert = Sort::insert($ba);
        $sSelect = Sort::select($ba);
        $sSQuick = Sort::quick($ba); die;
        $sec = $start();
        echo BR;
        $msgArr = [
            'basearray' => json_encode($ba),
            'sell_sort' => json_encode($sa),
            'bubble_sort' => json_encode($sBubble),
            'insert_sort' => json_encode($sInsert),
            'select_sort' => json_encode($sSelect),
            'quick_sort' => json_encode($sSQuick),
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