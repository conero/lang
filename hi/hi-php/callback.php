<?php
/**
 * 2017年10月9日 星期一
 * php - 回调函数
 */

 function callback($callback){
     if(!$callback instanceof Closure){
         return null;
     }
     return call_user_func($callback);
 }
 class Callback{
     public function __construct(){
        //  $this->scopeTest();
         $this->callSelfTest();
     }
     private function scopeTest(){
         // 不去取地址
         $use = 2017;
         var_dump($use);
         callback(function() use($use){
            $use = rand(100000, 999999); 
            var_dump($use);
         });
         var_dump($use);
         // 取地址
         callback(function() use(&$use){
            echo "\r\n use the address: \r";
            $use = rand(100000, 999999); 
            var_dump($use);
         });
         var_dump($use);
     }
    // 自动回调函数
    function callSelfTest(){
        $rand = call_user_func(function($r){
            echo $r. " (内部函数程序)\r\n";
            return $r/100;
        }, rand(0, 100));

        echo $rand."\r\n";
    }
 }

 new Callback;