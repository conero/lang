<?php
/**
 * Auther: Joshua Conero
 * Date: 2017/10/18 0018 15:47
 * Email: brximl@163.com
 * Name: 自动类库加载
 */


spl_autoload_register(function ($class){
    $dir = __DIR__.'/';
    $file = $dir . $class.'.php';
    if(is_file($file)){
        //echo $file;
        require_once($file);
    }else{
        $nps = 'conero\\';
        $file = $dir. str_replace($nps, '', $class). '.php';
        //echo $file;
        if(is_file($file)){
            require_once $file;
        }
    }
});