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
    // 数据递归的分为两部分， P1 P2, P1 中的任何一个元素都不 P2 要小
    public static function quick($array){
        // print_r($array);
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
        if(is_array($left) && is_array($right)){
            $array = array_merge($left, $right);
        }
        return $array;
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
    // php 如何实现二叉树数据结构 ??
    // 堆排序
    public static function heap($array){}
    // 归并排序; 采用分治法（Divide and Conquer）
    public static function merge($array){
        //$aLen = count($array);   
        // 分解
        self::divideMerge($array, $dA);
        // if($dA) echo json_encode($dA). BR;
        // 合并
        return self::conquerMerge($dA); 
    }
    // 分解
    private static function divideMerge($a, &$dArray=array()){
        $len = count($a);
        if($len > 2){
            $dN = floor($len/2);
            $aP1 = array_slice($a, 0, $dN);
            if($aP1) self::divideMerge($aP1, $dArray);
            $aP2 = array_slice($a, $dN);
            if($aP2) self::divideMerge($aP2, $dArray);
        }else if($len == 2){
            if($a[1] < $a[0]){
                $tmpV = $a[0];
                $a[0] = $a[1];
                $a[1] = $tmpV;
            }
        }
        // 分解到长度为 1/2 
        if($len <3){
            if(!is_array($dArray)){
                $dArray = [];
            }
            $dArray[] = $a;
            return false;
        }
        return $a;
    }
    // 合并
    private static function conquerMerge($a){
        // if($a) echo json_encode($a). BR;
        $len = count($a);
        $newArray = [];
        if($len > 1 && self::cqrMergeMk($a)){
            $i = 0;
            while($i<$len-1){
                $a1 = $a[$i];
                $a2 = $a[$i+1];
                // a2 -> a1
                // 数据第一部分
                for($j=0; $j<count($a2); $j++){
                    $n1 = $a2[$j];
                    $baseA = [$n1];
                    $a1Len = count($a1);
                    $orderMk = false;       // 是否排序标记
                    // 数组第二部分
                    for($k=0; $k<$a1Len; $k++){
                        if($a1[$k] > $n1){
                            if($k == 0){
                                $a1 = array_merge($baseA, $a1);
                            }else{
                                // $mA2 = ($k==$a1Len-1)? []:array_slice($a1, $k+1);
                                $mA2 = array_slice($a1, $k);
                                $a1 = array_merge(array_slice($a1, 0, $k), $baseA, $mA2);
                            }
                            $orderMk = true;
                            break;
                        }
                    }
                    if(!$orderMk){
                        $a1[] = $n1;
                    }
                }
                $newArray[] = $a1;
                $i += 2;
            }
        }
        // 递归合并
        if(self::cqrMergeMk($newArray)){
            if(count($newArray) == 1){
                $newArray = $newArray[0];
            }else{
                $newArray = self::conquerMerge($newArray);
            }
        }
        return $newArray;
    }
    // 检查是否为可合并的数组
    private static function cqrMergeMk($a){
        $right = false;
        foreach($a as $v){
            if(is_array($v)){
                $right = true;
                break;
            }
        }
        return $right;
    }
    // 基数排序
    // 数据的个十百千万位划分
    private static function getBaseRadix(){
        $base = [];
        $i=0;
        while($i<10){
            $base[] = false;
            $i += 1;
        }
        return $base;
    }
    // 获取数组序列中最大的位数
    private static function getNbit($array){
        $maxBit = 1;
        foreach($array as $v){
            $b = ceil(log10($v));
            if($b > $maxBit){
                $maxBit = $b;
            }
        }
        return $maxBit;
    }
    /**
     * 获取数据中指定位置的数字. 有点困难， 数学计算法以及字符串处理法
     * @param $b int
     * @param $n int
     * @return int
     */
    public static function getBinN($b, $n){
        $rInt = 0;

        // 数学计算法
        $bit = ceil(log10($n));
        if($bit >= $b){
            $b0 = $n / pow(10, $b);
            $rInt = intval(($b0 - intval($b0))*10);
        }
        

        // // 字符解决法则
        // $str = $n.'';
        // $sLen = strlen($str);
        // if($sLen >= $b){
        //     $rInt = intval(substr($str, (-1 * $b), 1));
        // }

        return $rInt;
    }
    // 基数排序
    public static function radix($array){
        $maxBit = self::getNbit($array);    // 数据最大位
        // echo $maxBit. BR;
        // echo ceil(log10($v=2)). ' ('.$v.', 10)'. BR;        
        // echo json_encode(self::getBaseRadix()). BR;        
        $i = 1;        
        while($i <= $maxBit){
            $baseA = self::getBaseRadix();
            foreach($array as $v){
                $n = self::getBinN($i, $v);
                if(!$baseA[$n]) $baseA[$n] = [];
                $baseA[$n][] = $v;                                        
            }
            $array = [];
            // echo json_encode($baseA). BR;
            foreach($baseA as $a){
                if(!$a) continue;
                $array = array_merge($array, $a);
            }
            // echo json_encode($array). BR;
            $i++;
        }
        return $array;
    }
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
        $sSQuick = Sort::quick($ba);
        $sMerge = Sort::merge($ba);
        $sRadix = Sort::radix($ba);
        $sec = $start();
        echo BR;
        $msgArr = [
            'basearray' => json_encode($ba),
            'sell_sort' => json_encode($sa),
            'bubble_sort' => json_encode($sBubble),
            'insert_sort' => json_encode($sInsert),
            'select_sort' => json_encode($sSelect),
            'quick_sort' => json_encode($sSQuick),
            'merge_sort' => json_encode($sMerge),
            'radix_sort' => json_encode($sRadix),
            'runtimes'  => $sec
        ];
        // echo json_encode($msgArr).BR;
        print_r($msgArr);
    }
    public function RadixTest(){
        $ba = getRandArray(13);
        echo json_encode($ba). BR;  
        foreach($ba as $v){
            echo $v.','.Sort::getBinN(3, $v).'' . BR;
        }
        // echo json_encode(Sort::radix($ba)). BR;  
    }
}
// 运行统计数
function console(){
    $test = new Test();
    $test->shellSort();
    // $test->RadixTest();
}
console();