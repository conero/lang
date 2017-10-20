<?php
/**
 * Auther: Joshua Conero
 * Date: 2017/10/18 0018 15:52
 * Email: brximl@163.com
 * Name: 数据库生成器
 */

namespace conero\learn;

 class Sqler{
     public $table;     // 数据表名
     public $alias;     // 数据库别名，用于多表连接查询
     protected $data;   // 提交数据
     protected $where;  // 数据条件
     protected $cols;   // 系统列表
     protected $registerEvents = [];    // 注册事件
     protected $runtimeSqlStack = [];   // 运行时sql执行栈
     protected $tableJoinList = [];     // 数据链表
     /*
        $tableJoinList = [
            alias => {
                type: INNER/LEFT/RIGHT      连接类型，空连接
                condition: 条件，连接条件,
                table: alias/table          数据表名称
            }
        ];
      */
     // 多数据库支持
     public $mutilateOptions = [
         // oracle
         'oci' => [
             'col_quotes' => '"',       // 字段引号
             'val_quotes' => '\''       // 值引号
         ],
         // mysql 数据库
         'mysql' => [
             'col_quotes' => '`',       // 字段引号
             'val_quotes' => '\''       // 值引号
         ]
     ];
     public $type;

     /**
      * sqlBuilder constructor.
      * @param string|array|null $table 数据表, [table alias] 或者 table
      * @param string $type 数据库类型 oci, mysql
      */
     public function __construct($table=null, $type='oci'){
         if(is_array($table)){
             list($tb, $alias) = $table;
             $this->table = $tb;
             $this->alias = $alias;
         }else{
             $this->table = $table;
         }
         $this->type = $type? $type: 'oci';
     }
     /**
      * sqlBuilder constructor(单例对象).
      * @param string|array $table
      * @param string $type
      */
     public static function table($table, $type=null){
         return new self($table, $type);
     }
     /**
      * 数据库连接表
      * @param string|array $table
      * @param string $condition
      * @param string $type INNER/LEFT/RIGHT
      * @return $this
      */
     public function join($table, $condition, $type=null){
         if(is_array($table)){
             list($tb, $alias) = $table;
             $table = $tb;
         }else{
             if(false !== strpos($table, ' ')){
                 list($tb, $alias) = explode(' ', $table);
                 $table = $tb;
             }else{
                 $alias = $table;
             }
         }
         $condition = preg_replace_callback("/\.[\d_]+/", function ($matches){
                 print_r($matches);
                 return '888888';
            }
            ,$condition);
         $this->tableJoinList[$alias] = [
             'type'         => $type,
             'condition'    => $condition,
             'table'        => $table
         ];
         return $this;
     }
     /**
      * 数据表设置
      * @param string|array $table
      * @return $this
      */
     public function setTable($table){
         $this->table = $table;
         return $this;
     }
     /**
      * 多数据库配置获取
      * @param string $key 键值
      * @param null $def 默认
      * @return mixed|null
      */
     public function mutilateDbs($key, $def=null){
         $data = $this->mutilateOptions[$this->type]? $this->mutilateOptions[$this->type]: [];
         return isset($data[$key])? $data[$key]: $def;
     }
     /**
      * 设置请求数据
      * @param array $data
      * @return $this
      */
     public function setData($data){
         $this->data = $data;
         return $this;
     }
     /**
      * 设置请求数据
      * @param array $data
      * @return $this
      */
     public function data($data){
         $this->data = $data;
         return $this;
     }
     /**
      * 设置请求数据
      * @param array $data
      * @return $this
      */
     public function where($data){
         $this->where = $data;
         return $this;
     }
     /**
      * 设置查询列表
      * @param array $data
      * @return $this
      */
     public function field($data){
         $this->cols = $data;
         return $this;
     }
     /**
      * 数据库别名
      * @param string $alias
      * @return $this
      */
     public function alias($alias){
         $this->alias = $alias;
         return $this;
     }
     /**
      * 事件绑定
      * @param string $key  insert/update/delete
      * @param callback $callback
      * @return
      */
     public function on($key, $callback){
         if($key && is_callable($callback)){
             $this->registerEvents[$key] = $callback;
         }
         return $this;
     }

     /**
      * sql 运行堆操作
      * @param null $data
      * @return $this|array
      */
     public function stack($data=null){
         if(null !== $data){
             $this->runtimeSqlStack[] = $data;
             return $this;
         }else{
             return $this->runtimeSqlStack;
         }
     }
     /**
      * 获取到运行器最后执行的sql语句
      * @param bool $isRaw
      * @return string|array
      */
     public function getLastSql($isRaw=true){
         $data = count($this->runtimeSqlStack) > 1?
             $this->runtimeSqlStack[count($this->runtimeSqlStack)-1]:
             null;
         if($isRaw){
             return $data? $this->getRawSql($data): null;
         }
         return $data? $data:[null, []];
     }
     /**
      * 获取当前的注册时间
      * @param string $key
      * @return callable|null
      */
     protected function getBindEvent($key){
         $callback = isset($this->registerEvents[$key]) && is_callable($this->registerEvents[$key])?
             $this->registerEvents[$key]: null;
         // 获取成功以后自动销毁注册函数，一次性的
         if($callback){
             unset($this->registerEvents[$key]);
         }
         return $callback;
     }
     /**
      * insert, 支持注册事件 => ($col, $data)
      * @param null|array $data 提交数组
      * @param null|array $filter 过滤数组中的字段
      * @return array [$sql, $bind]
      */
     public function insert($data=null, $filter = null){
         $data = $data? $data: $this->data;
         $sql = null;
         $bind = [];
         if(is_array($data)){
             $filter = is_array($filter)? $filter: null;
             $cols = [];$key = 'argv';$ctt = 1;
             $cBind = [];
             $cQuotes = $this->mutilateDbs('col_quotes', '"');
             foreach($data as $k=>$v){
                 // 过滤指定的数组
                 if($filter && !in_array($k, $filter)){continue;}
                 $cols[] = $cQuotes.$k.$cQuotes;
                 $newKey = $key . $ctt;
                 $cBind[] = ':'.$newKey;
                 $bind[$newKey] = $v;
                 $ctt++;
             }
             $insertEvent = $this->getBindEvent('insert');
             if($insertEvent){
                 $insertEvent($cols, $cBind);
             }
             $sql = '
                INSERT INTO '.$cQuotes.$this->table.$cQuotes.'('.implode(',', $cols).') VALUES ('.implode(',', $cBind).')
            ';
             // 运行时记录
             $this->stack([
                 'sql' => $sql,
                 'bind'=> $bind
             ]);
         }
         return [$sql, $bind];
     }

     /**
      * update, 支持注册事件 => ($updateList/更新列表, $where)
      * @param null|array $data
      * @param null|string|array $where
      * @return array [$sql, $bind]
      */
     public function update($data=null, $where=null){
         $data = $data? $data: $this->data;
         $where = $where? $where: $this->where;
         $sql = null;
         $bind = [];
         if(is_array($data)){
             $cQuotes = $this->mutilateDbs('col_quotes', '"');
             $key = 'argv';$ctt = 1;
             $updateList = [];
             foreach($data as $k=>$v){
                 if(is_int($k)){
                     $updateList[] = $v;
                     continue;
                 }
                 // 过滤指定的数组
                 $newKey = $key.$ctt;
                 $updateList[] = $cQuotes.$k.$cQuotes.'=:'.$newKey;
                 $bind[$newKey] = $v;
                 $ctt++;
             }
             $whArray = is_array($where)? $where: [];
             $whList = [];
             foreach ($whArray as $k1=>$v1){
                 $newKey = $key.$ctt;
                 $bind[$newKey] = $v1;
                 $whList[] = $cQuotes.$k1.$cQuotes.'=:'.$newKey;
                 $ctt++;
             }
             $updateEvent = $this->getBindEvent('update');
             if($updateEvent){
                 $updateEvent($updateList, $where);
             }
             $where = count($whList)>0? implode(' AND ', $whList): $where;
             // 更新sql语句
             $sql = 'UPDATE '.$cQuotes.$this->table.$cQuotes.' SET '.implode(',', $updateList).' '.($where? ' WHERE '. $where:'');
             // 运行时记录
             $this->stack([
                 'sql' => $sql,
                 'bind'=> $bind
             ]);
         }
         return [$sql, $bind];
     }
     /**
      * @param string|array $where
      * @param array|null $filter 过滤列表
      * @return array [$sql, $bind]
      */
     public function delete($where=null, $filter=null){
         $sql = null;
         $bind = [];
         $where = $where? $where: $this->where;
         $where = is_array($where)? $where:[];
         $whList = [];
         $key = 'argv';$ctt = 1;
         $filter = is_array($filter)? $filter: null;
         $cQuotes = $this->mutilateDbs('col_quotes', '"');
         foreach ($where as $k=>$v){
             if($filter && !in_array($k, $filter)){continue;}
             $newKey = $key . $ctt;
             $whList[] = $cQuotes.$k.$cQuotes.'=:'.$newKey;
             $bind[$newKey] = $v;
             $ctt++;
         }
         if(count($whList) > 0){
             $deleteEvent = $this->getBindEvent('delete');
             if($deleteEvent){
                 $deleteEvent($whList);
             }
             $sql = '
                DELETE FROM '.$cQuotes.$this->table.$cQuotes.' WHERE '.implode(' AND ', $whList).'
             ';
             // 运行时记录
             $this->stack([
                 'sql' => $sql,
                 'bind'=> $bind
             ]);
         }
         return [$sql, $bind];
     }
     /**
      * @param null|array $cols
      * @param null|string|array $where
      * @param array|null $filter 过滤列表
      * @return array [$sql, $bind]
      */
     public function select($cols=null, $where = null, $filter=null){
         if($this->where){
             $where = $where? array_merge($where, $this->where): $this->where;
             $this->where = null;
         }
         $cQuotes = $this->mutilateDbs('col_quotes', '"');
         if($this->cols){
             // 合并系统 col且自动销毁
             $cols = $cols? array_merge($cols, $this->cols): $this->cols;
             $this->cols = null;
         }
         if(empty($cols)){
             $cols = '*';
         }
         elseif (is_array($cols)){
             $colList = [];
             foreach ($cols as $k=>$v){
                 if(is_int($k)){
                     $colList[] = $cQuotes.$v.$cQuotes;
                 }else{
                     $colList[] = $cQuotes.$k.$cQuotes.' as '.$cQuotes.$v.$cQuotes;
                 }
             }
             $cols = implode(',', $colList);
         }
         $sql = 'SELECT '.$cols.' FROM '.$this->getBuilderTable();
         $bind = [];
         if($where){
             if(is_array($where)){
                 $key = 'args';$ctt = 1;
                 $whList = [];
                 $filter = is_array($filter)? $filter: null;
                 foreach ($where as $k=>$v){
                     if($filter && !in_array($k, $filter)){continue;}
                     $newKey = $key. $ctt;
                     $bind[$newKey] = $v;
                     $whList[] = $cQuotes.$k.$cQuotes.'=:'.$newKey;
                     $ctt++;
                 }
                 if(count($whList) > 0){
                     $where = ' WHERE '.implode(' AND ', $whList);
                 }else{
                     $where = '';
                 }
             }else{
                 $where = ' WHERE '.$where;
             }
         }else{
             $where = '';
         }
         $sql .= $where;
         // 运行时记录
         $this->stack([
             'sql' => $sql,
             'bind'=> $bind
         ]);
         return [$sql, $bind];
     }
     /**
      * @param string|array $sql， [$sql, $bind]
      * @param array $bind 过滤列表
      * @return string
      */
     public function getRawSql($sql, $bind=[]){
         if(is_array($sql)){ // 数组参参数传入
             list($sql, $bind) = $sql;
         }
         $vQuotes = $this->mutilateDbs('val_quotes', '\'');
         foreach($bind as $k=>$v){
             $value = $v? $vQuotes.$v.$vQuotes: null;
             if(empty($value)){
                 $value = 'NULL';
             }
             $sql = str_replace(':'.$k, $value, $sql);
         }
         return $sql;
     }
     /**
      * 获取编译后的数据表名称，可自动销毁join列表，防止数据残留， 主要用于 select
      * @return string
      */
     protected function getBuilderTable(){
         $cQuotes = $this->mutilateDbs('col_quotes', '"');
         $table = $this->table;
         if(count($this->tableJoinList) > 0){
             $tbList = [];
             foreach($this->tableJoinList as $k=>$v){
                 $type = $v['type'];
                 $type = ($type? ' '. strtoupper($type): ' ').' JOIN ';
                 $tbList[] = $type. $this->getDbNoun($v['table']).' ' .$k. ' ON '. $v['condition']
                 ;
             }
             $builderStrTb = $this->getDbNoun($table).' '. $this->alias. ' '.implode(' ', $tbList);
             $this->tableJoinList = [];
         }else{
             $builderStrTb = $this->getDbNoun($table);
         }
         return $builderStrTb;
     }
     /**
      * 数据库名称安全解析
      * @param string $noun 数据库字段名称
      * @return string
      */
     protected function getDbNoun($noun){
         $cQuotes = $this->mutilateDbs('col_quotes', '"');
         return (false !== strpos($noun, $cQuotes))? $noun: $cQuotes.$noun.$cQuotes;
     }
 }


/*
    ChangeLog   更新日志：
        1). 支持多表数据连接    /2017年10月18日 星期三
            select 时 class::$col 与 fun::$col 合并，并且自动销毁（防止下次影响）
            字符串安全监测
*/