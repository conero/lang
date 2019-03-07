<?php
/**
 * Auther: Joshua Conero
 * Date: 2018/1/18 0018 23:51
 * Email: brximl@163.com
 * Name: 文件对比工具
 */

//print_r($argv);

$dTar = $argv[1] ?? false;
$dSrc = $argv[2] ?? false;

if($dTar && $dSrc){
    if(is_dir($dSrc) && is_dir($dTar)){
        $dTarArray = scandir($dTar);
        $dSrcArray = scandir($dSrc);
        //print_r(array_intersect($dTarArray, $dSrcArray));
        print_r(array_diff($dTarArray, $dSrcArray));
    }
}