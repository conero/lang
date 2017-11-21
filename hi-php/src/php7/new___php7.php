<?php
/** 
 * 2017年7月26日 星期三
 * PHP7 新特性改变测试(分支)
 */
namespace hyang{
    class T{protected static $name = 'hayng\T';}
}

namespace{
    // 不定参数的实现
    // $test = function(int ...$ints){

    $test = function(...$ints){
        Show(json_encode($ints));
    };
    call_user_func($test,1,2,4,5,6,9,8,7);
    call_user_func($test,rand(1,9),"Joshua","Conero");

    // 匿名类
    (new class{
        public function __construct(){
            $this->about();
        }
        public function run(){
            Show('匿名函数调用测试');
        }
        private function about(){
            show('php7 - 匿名函数测试');
        }
    })
        ->run();

    // Closure::call()
    class A{
        protected $x; 
        const name = 'A';
        public function __construct(){
            $this->x = rand(0,1879999);
        }    
    }
    $getX = function(){
        return $this->x;
    };
    show($getX->call(new A));


    // ??
    class B extends A{const name = 'B';}
    $jt = [
        'base' => new A,
        'ext'   => new B
    ];
    $ext = $jt['ext2'] ?? new B;
    show($ext::name);



    class T extends hyang\T{
        public static $name = "T";
        public static function getName(){
            return self::$name ."__".parent::$name;
        }
    }
    show(T::getName());

}