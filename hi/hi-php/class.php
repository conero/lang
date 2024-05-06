<?php
/**
 * 2017年12月11日 星期一
 * php 类特性测试
 */
namespace Rong;
class Base{
    const ID = 'SURONG';
    public $text = '基类';
}

namespace Su;
class Base{
    const ID = 'SU_BASE';
    public $text = 'Su 命名空间中的基类';
}


// 系统运行命名空间
namespace Console;
//$RB = \Rong\Base::class;
//$SB = \Su\Base::class;

//print_r($RB);
//print_r($SB);

class Test{
    public function __construct()
    {
        var_dump($this->clsList());
        //print_r(new ($this->clsList()['RB']));
    }
    public function clsList(){
        return [
            'RB' => \Rong\Base::class,
            'SB' => \Su\Base::class
        ];
    }
}
new Test();