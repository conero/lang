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
         $this->scopeTest();
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
 }

 new Callback;