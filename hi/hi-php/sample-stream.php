<?php
/**
 * 简单流测试
 * 2017年9月20日 星期三
 */

 class PHPStreamTest{
    public function __construct(){
        $this->stdStream();
    }
    public function stdStream(){
        // 相等于系统- 等待输入， 界面一直等待。程序挂起
        // echo file_get_contents('php://stdin');
        
        $stdin = fopen('php://stdin', 'r'); // 只读
        

        $stdout = fopen('php://stdout', 'a'); // 只写

        fwrite($stdout, rand(100000, 9999999));

        echo "\r\n".fgets($stdin);
    }
 }
 class Console{
     public function __construct(){
         new PHPStreamTest;
     }
 }


 new Console;