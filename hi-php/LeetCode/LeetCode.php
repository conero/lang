<?php
/**
 * Created by PhpStorm.
 * User: Joshua Young
 * Date: 2019/3/8
 * Time: 21:48
 */


/*
 * LeetCode Var 变量集
 */
class LeetCodeVar{
    static $Cmd = "";            // 命令
    static $cmdKvData = [];      // 命令行解析的KV值
    static $cmdSetting = [];     // 命令行选项
    static function init($params){
        $num = count($params);
        $lastKv = null;
        for($i=1; $i <$num; $i++){
            $kv = $params[$i];
            if (preg_match('/-[a-zA-Z]{1}/', $kv)){
                $kv = substr($kv, 1);
                $lastKv = $kv;
                self::$cmdSetting[$kv] = true;
            }elseif ($lastKv){
                self::$cmdKvData[$lastKv] = $kv;
                $lastKv = null;
            }else{
                $lastKv = $kv;
            }
        }
        // print_r(self::$cmdKvData);
        // print_r(self::$cmdSetting);
    }

    /**
     * @param $key
     * @param null $value
     * @return mixed|null
     */
    static function Kv($key, $value=null){
        if(null === $value){
            return self::$cmdKvData[$key] ?? null;
        }
        self::$cmdKvData[$key] = $value;
    }
}

/*
 * LeetCode 入口
 */
class LeetCode
{
    function __construct()
    {
        $kvC = LeetCodeVar::Kv('c');
        if($kvC){
            // license-key-formatting
            // license key formatting
            $kvC = trim($kvC);
            $kvC = str_replace(' ', '-', $kvC);
            $instance = null;
            if(substr_count($kvC, '-') > 0){
                $instance = '';
                foreach (explode($kvC, '-') as $v){
                    $instance .= ucfirst($v);
                }
            }else{
                $instance = ucfirst($kvC);
            }
            $instanceFile = __DIR__.'/'. $instance.'.php';
            if(is_file($instanceFile)){
                require_once $instanceFile;
                if(class_exists($instance)){
                    new $instance();
                }else{
                    echo "!!   $instance 类不存在，脚本运行失败。\n";
                    echo "ii   $instanceFile\n";
                }
            }else{
                echo "!!   $instanceFile 文件不存在，脚本运行失败。\n";
            }
        }else{
            echo '!!   脚本选项错误，请使用： $ file.php -c <cmd>\n';
        }
    }

    const BR = "\n";
    static $solution;
    static $count = 1;
    /**
     * @param array $inputs
     * @param string $method
     * @param mixed $out
     * @param bool $onlyFalsePrint
     */
    static function simpleTest($inputs, $method, $out, $onlyFalsePrint=false){
        if(empty(self::$solution)){
            self::$solution = new Solution();
        }
        $output = call_user_func([self::$solution, $method], ...$inputs);
        $right = $out === $output;
        $info = self::$count.'. '.($right? 'true': 'false'). ' ) ['.implode(' , ', $inputs).'] => ['.$output.'] (VS) ('.$out.')'.self::BR;
        if($onlyFalsePrint) {
            if(!$right){
                echo $info;
            }
        }else{
            echo $info;
        }
        self::$count += 1;
    }
}

// 本地测试
// 本地测试^
LeetCodeVar::Kv('c', '<cmd>');
LeetCodeVar::Kv('c', 'licenseKeyFormatting');
//
// 控制台
LeetCodeVar::init($argv);
new LeetCode();