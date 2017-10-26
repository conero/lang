<?php
/**
 * Auther: Joshua Conero
 * Date: 2017年9月27日 星期三 0018 15:52
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
        $condition = $this->pointSgParse($condition);
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
     * @param bool $cover 是否覆盖原来的数据
     * @return $this
     */
    public function setData($data, $cover=false){
        if($cover){
            $this->data = $data;
        }else{
            $this->data = is_array($this->data)?
                array_merge($this->data, $data): $data;
        }
        return $this;
    }
    /**
     * 设置请求数据
     * @param array $data
     * @param bool $cover  是否覆盖原来的数据
     * @return $this
     */
    public function data($data, $cover=false){
        if($cover){
            $this->data = $data;
        }else{
            $this->data = is_array($this->data)?
                array_merge($this->data, $data): $data;
        }
        return $this;
    }
    /**
     * 设置条件数据
     * @param string|array $data 最终生成数组
     * @param bool $cover 是否覆盖原来的数据
     * @return $this
     */
    public function where($data, $cover=false){
        if(is_string($data)){
            $data = [$data];
        }
        if($cover){
            $this->where = $data;
        }else{
            $this->where = is_array($this->where)?
                array_merge($this->where, $data): $data;
        }
        return $this;
    }
    /**
     * 设置查询列表
     * @param array $data
     * @param bool $cover 是否覆盖原来的数据
     * @return $this
     */
    public function field($data, $cover=false){
        if(is_string($data)){
            $data = [$data];
        }
        if($cover){
            $this->cols = $data;
        }else{
            $this->cols = is_array($this->cols)?
                array_merge($this->cols, $data): $data;
        }
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
     * @return $this
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
        // 字段列名处理
        if(empty($cols)){
            $cols = '*';
        }
        elseif (is_array($cols)){
            $colList = [];
            foreach ($cols as $k=>$v){
                // 数字键数组
                if(is_int($k)){
                    $colList[] = $this->pointSgParse($v);
                // k-v 数组
                }else{
                    $colList[] = $this->pointSgParse($k).' as '.$cQuotes.$v.$cQuotes;
                }
            }
            $cols = implode(',', $colList);
        }
        $sql = 'SELECT '.$cols.' FROM '.$this->getBuilderTable();
        $bind = [];
        // where 条件处理
        if($where){
            if(is_array($where)){
                $key = 'args';$ctt = 1;
                $whList = [];
                $filter = is_array($filter)? $filter: null;
                foreach ($where as $k=>$v){
                    // 数字键数组
                    if(is_int($k)){
                        $whList[] = $this->pointSgParse($v);
                        continue;
                    }
                    // kv-数组处理
                    if($filter && !in_array($k, $filter)){continue;}
                    $newKey = $key. $ctt;
                    $bind[$newKey] = $v;
                    $whList[] = $this->pointSgParse($k).'=:'.$newKey;
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
      * 列求次数统计查询
      * @param null|string|array $col
      * @return array
      */
     public function count($col=null){
         if(empty($col)){
             $col = 'COUNT(*)';
         }else{
             if(is_array($col)){
                 $col = 'COUNT('.$this->getDbNoun($col[0]).') AS '.$this->getDbNoun($col[1]);
             }else{
                 $col = 'COUNT('.$this->getDbNoun($col).')';
             }
         }
         return $this->select($col);
     }
     /**
      * 列求和查询
      * @param string|array $col
      * @return array
      */
     public function sum($col){
         if(is_array($col)){
             $col = 'SUM('.$this->getDbNoun($col[0]).') AS '.$this->getDbNoun($col[1]);
         }else{
             $col = 'SUM('.$this->getDbNoun($col).')';
         }
         return $this->select($col);
     }
     /**
      * 列平均数查询
      * @param string|array $col
      * @return array
      */
     public function avg($col){
         if(is_array($col)){
             $col = 'AVG('.$this->getDbNoun($col[0]).') AS '.$this->getDbNoun($col[1]);
         }else{
             $col = 'AVG('.$this->getDbNoun($col).')';
         }
         return $this->select($col);
     }
    /**
     * @param string|array $sql/[$sql, $bind]
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
            $builderStrTb = $this->getDbNoun($table).($this->alias? ' '.$this->alias.' ':'');
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

    /**
     * 字符换点操作符号解析
     * @param string $str
     * @return string
     */
    protected function pointSgParse($str){
        if(false !== strpos($str, '.')){
            // alias.column => alias."column" oracle
            preg_match_all("/\.[a-zA-Z0-9_]+/", $str, $matches);
            $matches = isset($matches[0])? $matches[0]: [];
            $cQuotes = $this->mutilateDbs('col_quotes', '"');
            foreach ($matches as $v){
                $newStr = '.'.$cQuotes.substr($v, 1).$cQuotes;
                $str = str_replace($v, $newStr, $str);
            }
        }
        else{
            $str = $this->getDbNoun($str);
        }
        return $str;
    }
}

/*
    v1.0.1/20171021
    ChangeLog   更新日志：
        1). 支持多表数据连接    /2017年10月18日 星期三
            select 时 class::$col 与 fun::$col 合并，并且自动销毁（防止下次影响）
            字符串安全监测,
            数据表别名不知用标点符号


    类结构
        (入口)实例化 $instance
                new Sqler()
                Sqler::table()

        (中间件) -> $this
            Sqler::join             数据表连接
            Sqler::setTable         数据表设置
            Sqler::setData          设置请求数据
            Sqler::data             设置请求数据
            Sqler::where            设置条件数据
            Sqler::field            设置查询列表
            Sqler::alias            设置数据表别名
            Sqler::on               事件支持

        (出口)
            Sqler::getLastSql       获取最后一条数据
            Sqler::insert           数据新增
            Sqler::update           数据更新
            Sqler::delete           数据删除
            Sqler::select           数据查询
*/