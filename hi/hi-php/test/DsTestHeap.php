<?php
/**
 * Auther: Joshua Conero
 * Date: 2017/11/24 0024 13:04
 * Email: brximl@163.com
 * Name: 堆测试
 */
use conero\ds\Heap;

require_once "./ConeroTestApp.php";

class DsTestHeap
{
    // 测试运行器
    public function __construct()
    {
        $this->baseHeap();
    }
    protected function baseHeap(){
        $heap = new Heap();
        $heap->insert(rand(0, 100))
            ->insert(rand(0, 100))
            ->insert(rand(0, 100))
            ->insert(rand(0, 100))
            ->insert(rand(0, 100))
            ->insert(rand(0, 100))
            ->insert(rand(0, 100))
        ;
        print_r($heap);
    }
}

new DsTestHeap();