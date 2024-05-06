<?php
/**
 * Auther: Joshua Conero
 * Date: 2017/11/23 0023 23:33
 * Email: brximl@163.com
 * Name: 数据结构(堆)
 */

namespace conero\ds;

/**
 * 参考： https://baike.baidu.com/item/%E5%A0%86/20606834
 * Class Heap
 * 堆分为： 大根堆, 小根堆
 * 堆的定义如下：n个元素的序列{k1,k2,ki,…,kn}当且仅当满足下关系时，称之为堆.
 *      =>  (ki <= k2i,ki <= k2i+1)或者(ki >= k2i,ki >= k2i+1), (i = 1,2,3,4...n/2)
 * @package conero\ds
 */
class Heap
{
    /**
     * @var Heap
     */
    protected static $instance;
    protected $heapStackQueue = array();
    // [{v: 值, i:序列, b:位数, t:L/R/T 左右位标记/堆顶}] => {v: 78, i:, b:位数}
    /**
     * Heap constructor.
     */
    public function __construct()
    {
    }

    /**
     * 建立一个空堆
     * 单列类
     * @return Heap
     */
    public static function build(){
        if(!self::$instance){
            self::$instance = new self();
        }
        return self::$instance;
    }

    /**
     * 向堆中插入一个新元素
     * @param mixed $value
     * @return $this
     */
    public function insert($value){
        $data = ['v'=>$value];
        $this->heapStackQueue[] = $data;
        return $this;
    }
    public function update(){}
    public function get(){}
    public function delete(){}
    public function heapify(){}
}