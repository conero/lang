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
}

/*
 * LeetCode 入口
 */
class LeetCode
{
    function __construct()
    {
    }
}


// 控制台
LeetCodeVar::init($argv);
new LeetCode();