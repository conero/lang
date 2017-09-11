<?php
namespace Conero\Cn;
/**
 * 2017年9月11日 星期一
 * PHP-实例； 设计模式
 */

// 单例模式
class Singler{
    private $stack = [];
    private static $instance;
    private function __construct(){}
    public static function getGlobalInstance(){
        if(!self::$instance){
            self::$instance = new self();
        }
        return self::$instance;
    }
    public static function getSinInstance(){
        return new self();
    }
    public function setStack($key, $value){
        $this->stack[$key] = $value;
        return $this;
    }
    public function getStack($key=null, $def=null){
        if(null === $key) return $this->stack;
        return isset($this->stack[$key])? $this->stack[$key]: $def;
    }
}

// 工厂模式
interface ComInterface{
    public function getName();  // 工厂名称
    public function Start($param=array());    // 工厂入口
    public function End();      // 工厂结束
}
class WorkshopMySql implements ComInterface{
    public function getName(){return 'Mysql 生成部门';}
    public function Start($param=array()){}
    public function End($param=array()){}
}
class WorkshopOci implements ComInterface{
    public function getName(){return 'Oracle 生成部门';}
    public function Start($param=array()){}
    public function End($param=array()){}
}
class WorkshopPg implements ComInterface{
    public function getName(){return 'PostgreSQ 生成部门';}
    public function Start($param=array()){}
    public function End($param=array()){}
}
//
class Com{
    private static $hasWorkshop = [
        'MySql', 'Oci', 'Pg'
    ];
    public static function Workshop($name){
        if(in_array($name, self::$hasWorkshop)){
            // $className = 'Workshop'.$name;
            $className = 'Conero\Cn\Workshop'.$name;
            return new $className();
        }
        return new \StdClass();
    }
}

// 观察者模式， 依赖注入
class Su{
    public function setDI($key, $instance){
        $this->$key = $instance;
    }
}

// cmd - 控制台
class Console{
    public function __construct(){
        // $this->SinglerTest();
        // $this->ComTest();
        $this->InstanceDITest();
    }
    private function SinglerTest(){
        $sIn = Singler::getSinInstance();
        $In = Singler::getGlobalInstance();
        $sIn->setStack('t', rand(100000, 999999))
            ->setStack('a', 'Joshua Conero')
            ->setStack('l', 'Susanna')
            ;
        $In->setStack('T', rand(100000, 999999));
        print_r($sIn->getStack());
        print_r($In->getStack());
    }
    private function ComTest(){
        $mysql = Com::Workshop('MySql');
        $oci = Com::Workshop('Oci');
        $pg = Com::Workshop('Pg');
        print_r([$mysql->getName(), $oci->getName(), $pg->getName()]);
    }
    private function InstanceDITest(){
        $su = new Su;
        $su->setDI('com', Com::Workshop('MySql'));
        print_r([$su->com->getName()]);
    }
}

new Console;