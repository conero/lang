<?php
/**
 * 2017年9月27日 星期三
 * sql 语句生成器(oci)
 */
 
 class sqlBuilder{
     public $table;     // 数据表名
     protected $data;   // 提交数据
     protected $where;  // 数据条件
     protected $registerEvents = [];    // 注册事件
     protected $runtimeSqlStack = [];   // 运行时sql执行栈
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
      * @param string $table 数据表
      * @param string $type 数据库类型 oci, mysql
      */
     public function __construct($table=null, $type='oci'){         
         $this->table = $table;
         $this->type = $type? $type: 'oci';
     }
     /**
      * sqlBuilder constructor(单例对象).
      * @param string $table
      * @param string $type
      */
     public static function table($table, $type=null){
         return new self($table, $type);
     }
     /**
      * 数据表设置
      * @param array $data
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
        $where = $where? $where: $this->where;
        $cQuotes = $this->mutilateDbs('col_quotes', '"');
         if(empty($cols)) $cols = '*';
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
         $sql = 'SELECT '.$cols.' FROM '.$cQuotes.$this->table.$cQuotes;
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
 }