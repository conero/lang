<?php
/**
 * Auther: Joshua Conero
 * Date: 2017年9月27日 星期三 0018 15:52
 * Email: brximl@163.com
 * Name: 数据库生成器
 */

namespace conero\learn;

 class SqlerV2{
    public $table;     // 数据表名
    public $alias;     // 数据库别名，用于多表连接查询
    protected $data;   // 提交数据
    protected $where;  // 数据条件
    protected $cols;   // 系统列表
    protected $runtimeSqlStack = [];   // 运行时sql执行栈
    protected $tableJoinList = [];     // 数据链表
    protected $isBindMark = false;     // 字符绑定标识
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
     * @param string|array $data
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
      * 字符安全绑定设置
      * @return $this
      */
    public function isBind(){
        $this->isBindMark = true;
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
     * insert, 支持注册事件 => ($col, $data)
     * @param null|array $data 提交数组
     * @return string|array [$sql, $bind]
     */
    public function insert($data=null){
        $sql = 'INSERT INTO '.$this->getBuilderTable();
        // 支持子查询语句
        if(is_string($data)){
            $sql .= $data;
            return $this->isBindMark? [$sql, []]: $sql;
        }
        $data = is_array($data)? $data:array();
        if($this->data){
            if(is_array($this->data)){
                $data = array_merge($this->data, $data);
            }
            $this->data = null;
        }
        $cIst = [];
        $vIst = [];
        $ctt = 1;
        $bind = [];
        foreach ($data as $k=>$v){
            if(is_int($k)){
                // col = value
                if(strpos($v, '=') > 0){
                    list($key, $value) = explode('=', $v);
                    $cIst[] = $this->getDbNoun(trim($key));
                    $vIst[] = $this->valueSgParse(trim($value));
                }
            }else{
                $cIst[] = $this->getDbNoun($k);
                if($this->isBindMark){
                    $key = 'argv'.$ctt;
                    $vIst[] = ':'.$key;
                    $ctt += 1;
                    $bind[$key] = $v;
                }else{
                    $vIst[] = $this->valueSgParse($v);
                }
            }
        }
        $sql = 'INSERT INTO '.$this->getBuilderTable().'('.implode(',', $cIst).') VALUES ('.implode(',', $vIst).')';
        if($this->isBindMark){
            // 运行时记录
            $this->stack([
                'sql' => $sql,
                'bind'=> $bind
            ]);
            return [$sql, $bind];
        }
        return $sql;
    }

    /**
     * update, 支持注册事件 => ($updateList/更新列表, $where)
     * @param null|array $data
     * @param null|string|array $where
     * @return string|array [$sql, $bind]
     */
    public function update($data=null, $where=null){
        $whereParsed = $this->parseWhere($where);
        $data = is_array($data)? $data:array();
        if($this->data){
            if(is_array($this->data)){
                $data = array_merge($this->data, $data);
            }
            $this->data = null;
        }
        if($this->isBindMark){
            $sql = '';
            $bind = [];
            if(is_array($data)){
                $bind = [];
                $key = 'argv';$ctt = 1;
                $updateList = [];
                foreach($data as $k=>$v){
                    if(is_int($k)){
                        $updateList[] = $this->pointSgParse($v);
                        continue;
                    }
                    // 过滤指定的数组
                    $newKey = $key.$ctt;
                    $updateList[] = $this->pointSgParse($k).'=:'.$newKey;
                    $bind[$newKey] = $v;
                    $ctt++;
                }
                list($whStr, $wBind) = $whereParsed;
                $sql = 'UPDATE '.$this->getBuilderTable().' SET '.implode(',', $updateList).$whStr;
                $bind = array_merge($bind, $whereParsed);
                // 运行时记录
                $this->stack([
                    'sql' => $sql,
                    'bind'=> $bind
                ]);
            }
            return [$sql, $bind];
        }
        else{
            $sql = '';
            if(is_array($data)){
                $updateList = [];
                foreach ($data as $k=>$v){
                    if(is_int($k)){
                        $updateList[] = $this->pointSgParse($v);
                    }else{
                        $updateList[] = $this->pointSgParse($k).'='.$this->valueSgParse($v);
                    }
                }
                $sql = 'UPDATE '.$this->getBuilderTable().' SET '.implode(',', $updateList).$whereParsed;
            }
            return $sql;
        }
    }
    /**
     * @param array|string $where 删除条件
     * @return string|array [$sql, $bind]
     */
    public function delete($where=null){
        $whereParsed = $this->parseWhere($where);
        $sql = 'DELETE FROM '.$this->getBuilderTable();
        if($this->isBindMark){
            $sql .= ' '.$whereParsed[0];
            return[$sql, $whereParsed[1]];
        }else{
            $sql .= ' '.$whereParsed[0];
        }
        return $sql;
    }
    /**
     * @param null|array|string $cols
     * @param null|string|array $where，传入值时会与类的where数据合并(如果存在)
     * @return array|string [$sql, $bind]
     */
    public function select($cols=null, $where = null){
        $sql = 'SELECT '.$this->fieldParse($cols).' FROM '.$this->getBuilderTable();
        $whereParsed = $this->parseWhere($where);
        if($this->isBindMark){
            list($whStr, $bind) = $whereParsed;
            $sql .= $whStr;
            // 运行时记录
            $this->stack([
                'sql' => $sql,
                'bind'=> $bind
            ]);
            return [$sql, $bind];
        }else{
            $sql .= $whereParsed;
            return $sql;
        }
    }
     /**
      * 列求次数统计查询
      * @param null|string|array $col
      * @return array|string
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
      * @return array|string
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
      * @return array|string
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
     /**
      * @param string $str 系统参数值
      * @return string
      */
    protected function valueSgParse($str){
        $vQuotes = $this->mutilateDbs('val_quotes', '\'');
        if(substr_count($str, '\'') < 2){
            $str = $vQuotes. $str . $vQuotes;
        }
        return $str;
    }
     /**
      * @param null|string|array $where 提哦啊就爱你
      * @param bool $onlyInput 仅仅解析输入的数据
      * @param bool $isRaw 不自动添加Where等字符串
      * @return array|string
      */
    public function parseWhere($where=null, $onlyInput=false, $isRaw=false){
        // where 条件自动解析为array、string
        if($where){
            $where = is_array($where)? $where: [$where];
        }
        if(!$onlyInput && $this->where){
            $where = $where? array_merge($this->where, $where): $this->where;
            $this->where = null;
        }
        $str = '';
        $isBindMark = $this->isBindMark;
        $bind = [];
        $ctt = 1;
        foreach ($where as $k=>$v){
            if(is_int($k)){
                $tmpStr = $this->pointSgParse($v);
                // 连接词查看
                if(!preg_match("/(and)|(or)/i", $tmpStr)){
                    $str .= ($str? ' AND ': '');
                }
                $str .= $tmpStr;
            }else{
                if($isBindMark){
                    $key = 'argw'.$ctt;
                    $bind[$key] = $v;
                    $str .= ($str? ' AND ': ''). $this->pointSgParse($k).'=:'. $key;
                    $ctt += 1;
                }
                else{
                    $str .= ($str? ' AND ': ''). $this->pointSgParse($k).'='. $this->valueSgParse($v);
                }
            }
        }
        $str = ($str && !$isRaw ? ' WHERE ': '').$str;
        return $isBindMark? [$str, $bind]: $str;
    }

     /**
      * 查询列名解析
      * @param null|string|array $cols 列名
      * @return string
      */
    public function fieldParse($cols=null){
        if(is_string($cols)){
            $cols = [$cols];
        }
        if($this->cols){
            // 合并系统 col且自动销毁
            $cols = $cols? array_merge($this->cols, $cols): $this->cols;
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
                    $colList[] = $this->pointSgParse($k).' as '.$this->getDbNoun($v);
                }
            }
            $cols = implode(',', $colList);
        }
        return $cols;
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
            Sqler::count            列统计
            Sqler::sum              列合并
            Sqler::avg              列平均数

    V2 ：/20171023
        取消：
            1.事件注册，如 insert/update/delete
            2.删除部分方法中的 filter 参数(由于使用率很小)
        新特性：
            1.sql绑定与非绑定特性开发， 前者返回 array 或者返回 string

*/