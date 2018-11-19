<?php
/**
 * Auther: Joshua Conero
 * Date: 2018/11/19 0019 15:30
 * Email: brximl@163.com
 * Name: 控制检测
 */



class NC{
    static function zore(){
        self::check(0);
        self::check('');
    }
    static function check($z){
        echo "\n $z 值判断 if 为：". ($z? 'true': 'false');
        echo "\n is_null: ". (is_null($z)? 'true': 'false');
    }
}


class Console{
    function __construct()
    {
        NC::zore();
    }
}

new Console();