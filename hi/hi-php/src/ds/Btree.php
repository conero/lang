<?php
/**
 * 2017年11月23日 星期四
 * 二叉树数据结构模型： Binary Tree
 * Joshua Conero
 */
namespace conero\ds;
/**
 * 基本二叉树数及结构实现(特殊二位数组(x,y))； 二位数组模拟型号(顺序模型)
 * 子树分分为： left subtree; right subtree:  map[top]: => {left:value, right:value}
 * 深度：  deep 
 */ 
class BTree{
    const NONE = '__NONE__';    // 空值
    public $deep;               // 深度
    protected $dataOrder = []; // 深度序列建表
    public $top;    
    /**
     * @param mixed $top 
     */
    public function __construct($top=null){
        if($top){
            $this->setTop($top);
        }
    }    
    /**
     * 设置二叉树树顶
     * @param mixed $top
     */
    public function setTop($top){
        $this->top = $top;
        if(empty($this->dataOrder)){
            $this->dataOrder[] = array();
        }
        $this->dataOrder[0][0] = $top;
        if($this->deep < 1){
            $this->deep = 1;
        }
    }
    /**
     * 数据推送值
     * @param mixed ...$value
     */
    public function push(){
        foreach(func_get_args() as $v){
            $this->_push($v);
        }
    }
    /**
     * 单值推送
     * @param mixed $value
     */
    protected function _push($value){
        if(empty($this->dataOrder)){
            $this->setTop($value);
            return null;
        }
        // Btree(x, y) = λ
        $x = count($this->dataOrder);
        $y = count($this->dataOrder[$x-1]);
        // x => maxY
        // 二叉树， $x => 满树 count = 2^x-1; 每一层容纳的数量 C(λ) = 2^(λ-1), λ<1
        $maxY = pow(2, $x-1);
        if($y >= $maxY){
            $this->dataOrder[] = [];
            $y = 0;            
        }else{
            $this->deep = $x;
            $x -= 1;
            $y -= 1;
        }
        //$this->deep
        $this->dataOrder[$x][] = $value;
    }
}