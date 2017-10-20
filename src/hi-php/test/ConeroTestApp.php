<?php
/**
 * Auther: Joshua Conero
 * Date: 2017/10/18 0018 16:27
 * Email: brximl@163.com
 * Name: test 应用统一入口
 */

define('TEST_ROOT', dirname(__DIR__));
require_once TEST_ROOT.'/src/autoloader.php';


// 函数
function debug(){
    // 参数解析
    $data = func_get_args();
    $n = func_num_args();
    if(1 == $n){
        $data = $data[0];
        if($data instanceof Closure){
            $data = call_user_func($data);
        }
    }
    // 格式处理
    if(is_object($data) || is_array($data)){
        $data = print_r(is_object($data)? (array)$data: $data, true);
    }
    $dir = TEST_ROOT.'/.test/';
    if(!is_dir($dir)){
        mkdir($dir);
    }
    $files = $dir .date('Y-m-d') .'.log';
    $fh = fopen($files, 'a');
    fwrite($fh, $data);
    fclose($fh);
}
