<?php
/**
 * Auther: Joshua Conero
 * Date: 2018/1/19 0019 0:04
 * Email: brximl@163.com
 * Name: 随机数并发测试
 */



call_user_func(function (){
    //return false;   // 开关
    $count = 10000;
    $i = 1;
    $array = [];
    while ($i<$count){
        $t = uniqid();
        $t1 = uniqid();
        $array[] = $t1;
        echo $t.'  '.(in_array($t, $array)? 'Y':'N') ."\r\n";
        $array[] = $t;
        $i++;
    }
});