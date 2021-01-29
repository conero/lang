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
                $this->msg($ret);
                if(is_array($ret)){
                    header('Content-Type: application/json; charset=utf-8');
                    echo json_encode($ret);
                }else if(is_string($ret)){
                    echo $ret;
                }
            }else{
                throw new Exception("Method `$s` is not fund.", 1);
            }
        }
    }

    //日志调试
    protected function msg($data){
        if($data){
            if(is_object($data) || is_array($data)){
                $data = print_r($data, true);
            }
            $fp = fopen('./debug.log', 'a');
            fwrite($fp, $data);
            fclose($fp);
        }
    }

    //初始化
    function init(){
        $this->method = $_SERVER['REQUEST_METHOD'];
    }

    //文件上传
    function Upload(){
        $method = $this->method;
        $data = [
            'post' => $_POST,
            'files' => $_FILES,
            'method' => $method,
        ];
        return $data;
        if('POST' == $method){
            $data = [
                'post' => $_POST,
                'files' => $_FILES
            ];
            return $data;
        }
        else{}
    }

    function Help(){
        phpinfo();
    }
}

// 实例化
new Handler();