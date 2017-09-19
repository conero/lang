<?php
/**
 * 2017年9月19日 星期二
 * php - 死循环
 */

 class SleepWhile{
     function __construct(){
         $this->baseSleepWhileTest();
        //  $this->badWhileTest();
     }
     // 基本良性死循环
     function baseSleepWhileTest(){
         // 死循环，而不会因为死循环造成内存卡死
         $start = 1;
         while(1){
             // 延缓执行
             // usleep 微妙及
             sleep(1);
            //  print("\r\n".rand(1, 9));
             print("\r\t".rand(100000, 999999).'/'.$start);
             $start++;
         }
     }
     // 致命死循环
     function badWhileTest(){
         while(1){
            print("\r\n".rand(1, 9));
         }
     }
 }

 new SleepWhile;