<?php
/** 
 * 2017年7月26日 星期三
 * PHP7 新特性改变测试
 */

$v = phpversion();
list($v1,$v2,$v3) = explode('.',$v);
$v1 = intval($v1);

function Show($msg,$isError=false){
    echo "\r\t".($isError? ':-( ':':)- ').$msg."\r\n";
    // echo "\r\t".($isError? ':-( ':':)- ').$msg."\r\v";
}
function Showv($msg,$isError=false){
    echo "\r\t".($isError? ':-( ':':)- ').$msg."\r\v";
}
show($v);

class Test{
    public static function php5(){}
    public static function php7(){
        // 新特性
        require_once __DIR__.'/new___php7.php';
        
    }
}

switch($v1){
    case 5:
        Test::php5();
        break;
    case 7:
        Test::php7();
        break;
}
