<?php
// 分片上传

class Handler{
    protected $method = '';
    function __construct(){
        $s = $_GET['s'] ?? null;
        if($s){
            $s = ucfirst($s);
            if(method_exists($this, $s)){
                $this->init();
                $ret = $this->$s();
                if(is_array($ret)){
                    header('Content-Type: application/json; charset=utf-8');
                    echo json_encode($ret);
                }
            }else{
                throw new Exception("Method `$s` is not fund", 1);
            }
        }
    }

    //初始化
    function init(){
        $this->method = $_SERVER['REQUEST_METHOD'];
    }

    //文件上传
    function Upload(){
        $method = $this->method;
        if('POST' == $method){}
        else{}
    }

    function Help(){
        phpinfo();
    }
}

// 实例化
new Handler();