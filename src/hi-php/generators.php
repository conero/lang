<?php
// 生成器学习与测试
// 2017年8月30日 星期三


class Generators{
    public static function BaseTest(){
        $baseFn = function(){
            $i = 2;
            while($i<1000){
                yield $i;
            }
        };
        foreach($baseFn() as $v)(
            var_dump($v);
        )
    }
}


Generators::BaseTest();