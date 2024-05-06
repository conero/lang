<?php
/**
 * 2017年10月12日 星期四
 * Sqler 测试用例
 */
use conero\ds\Btree;
require_once "./ConeroTestApp.php";

class TestCase{
    public function __construct(){
        $this->baseBt();
        // $this->btreeYMax(10);
    }
    private function baseBt(){
        $bt = new Btree('二叉树顶');
        print_r($bt);
        $bt->push(1, 2, 5, 8, 9, 5, 8, 41, rand(0, 100), rand(0, 100), rand(0, 100));
        print_r($bt);
    }
    // 二叉树最大值
    private function btreeYMax($bit){
        $i = 1;
        while($i <= $bit){
            echo (pow(2, $i-1)).", ".$i."\r\n";
            $i += 1;
        }
    }
}
new TestCase();