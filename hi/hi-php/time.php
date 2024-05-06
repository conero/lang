<?php
/**
 * 时间处理
 * 2017年10月11日 星期三
 */


class TestTime{
    public function __construct(){        
        $this->microtimeTest();
    }
    private function microtimeTest(){
        // 格式输出测试
        /*
        print_r([
            microtime(),    // list($usec, $sec) = explode(" ", microtime());
            microtime(true) // float
        ]);
        */

        $st = microtime(true);

        sleep(10);
        echo microtime(true) - $st;
    }
}

new TestTime;