<?php
// Joshua Conero
// 2020.10.27
// 目录扫描测试基准
define('HOME_DIR', str_replace('\\', '/', __DIR__));


class Util{
    private static $_topdir = '';
    static function topDir(){
        if(empty(self::$_topdir)){
            switch(PHP_OS){
                case 'WINNT':
                    self::$_topdir = substr(HOME_DIR, 0, 3);
                break;
                default: 
                return HOME_DIR;
            }
        }
        return self::$_topdir;
    }
}

/**
 * 目录扫描
 */
function scandirAllDir($dir){
    $data = [
        'dir' => 0,
        'file' => 0,
        'size' => 0
    ];
    $dir = substr($dir, -1) !== '/'? $dir.'/': $dir;
    if(is_dir($dir)){
        $scandDir = scandir($dir);
        $scandDir = is_array($scandDir)? $scandDir: [];
        foreach($scandDir as $name){
            if(in_array($name, ['.', '..'])){
                continue;
            }
            $path = $dir . $name;
            if(is_file($path)){
                $data['file'] += 1;
                $data['size'] += filesize($path);
            }else if(is_dir($path)){
                $res = scandirAllDir($path);
                $data['dir'] += 1 + $res['dir'];
                $data['file'] += $res['file'];
                $data['size'] += $res['size'];
            }
        }
    }
    return $data;
}


function microtime_float()
{
    list($usec, $sec) = explode(" ", microtime());
    $startTime=((float)$usec + (float)$sec);
    return $startTime;
}

function time_get_usage()
{
    $startTime = microtime_float();
    return function() use($startTime){
        return microtime_float() - $startTime;
    };
}

class Console{
    function __construct(){
        $timer = time_get_usage();
        echo "系统进入测试基准！\n";
        
        $this->scand_dir_put();

        echo "\n.  程序内存消耗(kB)：". (memory_get_usage()/1024).
             "\n.  程序用时（秒）：". $timer().
             "\n.  php 版本：". PHP_VERSION
        ;
    }

    function input($tip){
        echo "$tip :";
        $input = fread(STDIN, 1024);
        echo "\n";
        return trim($input);
    }

    // 扫描输入的目录
    function scand_dir_put(){
        $scandDir = $this->input('请输入您需要变量的目录');
        if(empty($scandDir) || !is_dir($scandDir)){
            echo "\n您输入的目录（ $scandDir ）不是有效目录，系统将默认遍历: ".Util::topDir()."\n";
            $scandDir = Util::topDir();
        }

        $line = $this->input('您确定要遍历吗？(y/n)');
        if('n' == $line){
            echo "\n您已终止程序遍历！！";
            return;
        }else{
            echo "程序正在扫码： $scandDir ... \n";
            $result = scandirAllDir($scandDir);
            echo "目录：". $scandDir."已完成扫描\n";
            echo "结果如下：\n". 
                '目录总数：'.$result['dir']."\n".
                '文件总数：'.$result['file']."\n".
                '总数大小：'.$result['size']."\n"
                ;
        }
    }
}

new Console();
