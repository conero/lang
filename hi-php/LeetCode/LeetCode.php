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
    static $opt_gettimes = false;       // 选项获取运行的时间
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
        $getSconds = self::getSeconds();
        $output = call_user_func([self::$solution, $method], ...$inputs);
        $right = $out === $output;
        if(self::$opt_gettimes){
            $getSconds = round($getSconds(), 4);
            $getSconds = ' <'.$getSconds. 's> ';
        }else{
            $getSconds = '';
        }
        $info = self::$count.'. '.($right? 'true': 'false'). ' )'.$getSconds.' ['.$out.'] (VS) ('.$output.') <<= ['.self::simpleTestInutsToStr($inputs).']'.self::BR;
        if($onlyFalsePrint) {
            if(!$right){
                echo $info;
            }
        }else{
            echo $info;
        }
        self::$count += 1;
    }

    /**
     * 输入格式化
     * @param $inputs
     * @return string
     */
    static function simpleTestInutsToStr($inputs){
        $cls2Array = [];
        foreach ($inputs as $v){
            if(is_array($v)){
                // $cls2Array[] = '['.implode(' , ', $v).']';
                $cls2Array[] = '['.self::simpleTestInutsToStr($v).']';
            }else{
                $cls2Array[] = $v;
            }
        }
        return implode(' , ', $cls2Array);
    }

    /**
     * @return float
     */
    static function microtime_float()
    {
        list($usec, $sec) = explode(" ", microtime());
        return ((float)$usec + (float)$sec);
    }

    /**
     * @param null $startTm
     * @return Closure|float|null
     */
    static function getSeconds($startTm=null){
        if(empty($startTm)){
            $startTm = self::microtime_float();
            return function () use ($startTm){
                $end = self::microtime_float();
                return $end - $startTm;
            };
        }else{
            $end = self::microtime_float();
            return $end - $startTm;
        }

    }
}

// 本地测试
// 本地测试^
LeetCodeVar::Kv('c', '<cmd>');
LeetCodeVar::Kv('c', 'SortColors');
//LeetCodeVar::Kv('c', 'LemonadeChange');
//
// 控制台
LeetCodeVar::init($argv);
new LeetCode();