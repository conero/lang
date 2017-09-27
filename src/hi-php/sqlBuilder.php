<?php
/**
 * 2017年9月27日 星期三
 * sql 语句生成器(oci)
 */
 
 class sqlBuilder{
     protected $table;  // 数据表名
     protected $data;   // 提交数据
     protected $registerEvents = [];    // 注册事件
     /**
      * sqlBuilder constructor.
      * @param string $table
      */
     public function __construct($table){
         $this->table = $table;
     }
     /**
      * @param array $data
      * @return $this
      */
     public function setData($data){
         $this->data = $data;
         return $this;
     }
     /**
      * on
      * @param string $key 
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
            foreach($data as $k=>$v){
                // 过滤指定的数组
                if($filter && !in_array($k, $filter)){continue;}
                $cols[] = '"'.$k.'"';
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
                INSERT INTO "'.$this->table.'"('.implode(',', $cols).') VALUES ('.implode(',', $cBind).')
            ';
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
        $sql = null;
        $bind = [];
        if(is_array($data)){
            $key = 'argv';$ctt = 1;
            $updateList = [];
            foreach($data as $k=>$v){
                // 过滤指定的数组
                $newKey = $key.$ctt;
                $updateList[] = '"'.$k.'"=:'.$newKey;
                $bind[$newKey] = $v;
                $ctt++;
            }
            $whArray = is_array($where)? $where: [];
            $whList = [];
            foreach ($whArray as $k1=>$v1){
                $newKey = $key.$ctt;
                $bind[$newKey] = $v1;
                $whList[] = '"'.$k1.'"=:'.$newKey;
                $ctt++;
            }
            $updateEvent = $this->getBindEvent('update');
            if($updateEvent){
                $updateEvent($updateList, $where);
            }
            $where = count($whList)>0? implode(' AND ', $whList): $where;
            // 更新sql语句
            $sql = 'UPDATE "'.$this->table.'" SET '.implode(',', $updateList).' '.($where? ' WHERE '. $where:'');
        }
        return [$sql, $bind];
     }

     /**
      * @param string|array $where
      * @param array|null $filter 过滤列表
      * @return array [$sql, $bind]
      */
     public function delete($where, $filter=null){
         $sql = null;
         $bind = [];
         $where = is_array($where)? $where:[];
         $whList = [];
         $key = 'argv';$ctt = 1;
         $filter = is_array($filter)? $filter: null;
         foreach ($where as $k=>$v){
             if($filter && !in_array($k, $filter)){continue;}
             $newKey = $key . $ctt;
             $whList[] = '"'.$k.'"=:'.$newKey;
             $bind[$newKey] = $v;
             $ctt++;
         }
         if(count($whList) > 0){
             $deleteEvent = $this->getBindEvent('delete');
             if($deleteEvent){
                 $deleteEvent($whList);
             }
             $sql = '
                DELETE FROM "'.$this->table.'" WHERE '.implode(' AND ', $whList).'
             ';
         }
         return [$sql, $bind];
     }
     public function del($where, $filter=null){
         $sql = null;
         $bind = [];
         $where = is_array($where)? $where:[];
         $whList = [];
         $key = 'argv';$ctt = 1;
         $filter = is_array($filter)? $filter: null;
         foreach ($where as $k=>$v){
             if($filter && !in_array($k, $filter)){continue;}
             $newKey = $key . $ctt;
             $whList[] = '"'.$k.'"=:'.$newKey;
             $bind[$newKey] = $v;
             $ctt++;
         }
         if(count($whList) > 0){
             $deleteEvent = $this->getBindEvent('delete');
             if($deleteEvent){
                 $deleteEvent($whList);
             }
             $sql = '
                DELETE FROM "'.$this->table.'" WHERE '.implode(' AND ', $whList).'
             ';
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
         if(empty($cols)) $cols = '*';
         elseif (is_array($cols)){
             $colList = [];
             foreach ($cols as $k=>$v){
                 if(is_int($k)){
                     $colList[] = '"'.$v.'"';
                 }else{
                     $colList[] = '"'.$k.'" as "'.$v.'"';
                 }
             }
             $cols = implode(',', $colList);
         }
         $sql = 'SELECT '.$cols.' FROM "'.$this->table.'"';
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
                     $whList[] = '"'.$k.'"=:'.$newKey;
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
        return [$sql, $bind];
     }
 }