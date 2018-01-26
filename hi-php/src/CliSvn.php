<?php
/**
 * Auther: Joshua Conero
 * Date: 2018/1/26 0026 9:27
 * Email: brximl@163.com
 * Name: Windows 下 svn 控制器
 */

class CliSvn
{
    public $rpath;       // 项目地址 RootPath
    public $wpath;       // 项目地址 WorkPath
    protected $cacheData = [];  // 缓存数据
    public function __construct($option=array())
    {
        $this->setAttr($option);
        // 项目工作地址
        if(empty($this->wpath) && $this->rpath){
            $this->wpath = $this->rpath;
        }
        $dir = $this->wpath.'/~.~/';
        if(!is_dir($dir)){
            mkdir($dir);
        }
        $this->wpath = $dir;
        $this->clearCache();
    }

    /**
     * 缓存目录清除
     */
    public function clearCache(){
        $name = $this->wpath.'clearCache.~.~';
        $isClearMk = false;
        $cData = intval(date('Ymd'));
        //$cData += 10;   // 测试语句
        if(is_file($name)){
            $date = intval(file_get_contents($name));
            if($date != $cData){
                $isClearMk = false;
            }else{
                $isClearMk = true;
            }
        }
        if(!$isClearMk){
            $wpath = $this->wpath;
            foreach (scandir($wpath) as $v){
                if(in_array($v, ['.', '..'])){
                    continue;
                }
                $idx = strpos($v, '-');
                $file = $wpath.$v;
                if($idx !== false && $cData != intval(substr($v, 0, $idx))){
                    unlink($file);
                }
            }
            file_put_contents($name, $cData);
        }
    }
    /**
     * 数据缓存器
     * @param string $key
     * @param mixed $value
     * @return $this|bool|mixed
     */
    protected function cache($key, $value=null){
        if(!$value){
            return $this->cacheData[$key] ?? false;
        }
        $this->cacheData[$key] = $value;
        return $this;
    }
    /**
     * 设置项目项属性
     * @param array $option
     * @return $this
     */
    public function setAttr($option){
        foreach ($option as $key=>$value){
            $this->$key = $value;
        }
        return $this;
    }

    /**
     * 多行命令行执行cli 程序封装
     * @param string $cli
     * @return array
     */
    public function mutiCli($cli){
        if(is_array($cli)){
            $cli = implode(' ', $cli);
        }
        $cmdContent = "@echo off\ncd "
            .$this->rpath."\n"
            .":: made by Joshua Conero on the day ".date('Y-m-d H:i:s')."\n"
            .":: start 20180126(JC)\n\n"
            .":: from the class ".get_class($this)."\n"
            .'svn '.$cli
            ;
        $bat = $this->wpath.date('Ymd').'-svn-'.rand(0, 9999999).'.bat';
        file_put_contents($bat, $cmdContent);
        exec($bat, $array);
        return $array;
    }

    /**
     * 单行命令行数据列
     * @param string $cli
     * @param bool $fullCli
     * @return mixed
     */
    public function cli($cli, $fullCli=false){
        $bat = ($fullCli? "":"svn ").$cli;
        exec($bat, $array);
        return $array;
    }

    /**
     * 获取项目的信息
     * @return array
     */
    public function info(){
        // 数据缓存
        $data = $this->cache('info');
        if(empty($data)){
            $data = [];
            $array = $this->mutiCli('info');
            foreach ($array as $v){
                $v = trim($v);
                if(empty($v)){
                    continue;
                }
                $idx = strpos($v, ':');
                if($idx > 0){
                    $key = trim(substr($v, 0, $idx));
                    $key = strtolower($key);
                    $key = str_replace(' ','_', $key);
                    $data[$key] = trim(substr($v, $idx + 1));
                }
            }
        }
        return $data;
    }

    /**
     * @return int|mixed
     */
    public function revision(){
        $data = $this->info();
        return $data['revision'] ?? 0;
    }

    /**
     * 获取数据版本信息
     * @return bool
     */
    public function version(){
        $array = $this->cli('--version --quiet');
        return $array[0] ?? false;
    }
}