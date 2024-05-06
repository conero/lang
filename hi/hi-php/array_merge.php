<?php
// 2017年8月30日 星期三
// Joshua Conero

class ArrayTest{
    public static $StdArray = [
        'k1' => null,
        'k2' => 20170830,
        'k3' => null,
        'k4' => null,
        'k4' => null
    ];
    public static function merge(){
        // array_merge  集合1 > 集合2 可用于数组顺序化(基前者)
        print_r(array_merge(
            self::$StdArray, ['k3'=> rand(1, 9), 'k1' => 'Susanna']
        ));
    }
}


ArrayTest::merge();