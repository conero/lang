<?php
/**
 * 2017年9月22日 星期五
 * php 编译sql语句
 */


 // oracle ->
 class sqlBuild{
     // 新增sql安全生成编译 => ($sql, $bind)
	public function insertSql($tb, $data=null, $col= null){
		$retVel = [];
		$sql = 'INSERT INTO "__TABLE__"%COL% VALUES (%DATA%)';
		$sql = str_replace('__TABLE__', $tb, $sql);
		
		$colList = $col? $col: [];
		$bindList = [];
		$bindDt = $data? $data:null;
		$dataList = [];		
		// 有数组时
		if(is_array($data)){
			foreach($data as $k=>$v){
				if(empty($col)){
					$colList[] = $k;
					$bindList[] = ':'.$k;
					if(empty($bindDt)) $bindDt = $data;
				}else{
					$colList = $col;
					$bindList[] = ':'.$col;
					foreach($col as $vv){
						$bindDt[$vv] = isset($data[$vv])? $data[$vv]: null;
					}
				}
			}
			$sql = str_replace('%COL%', '("'.implode('","' ,$colList).'")', $sql);
			$sql = str_replace('%DATA%', implode(', ' ,$bindList), $sql);
		}
		elseif($col){
			$sql = str_replace('%COL%', '("'.implode('","' ,$colList).'")', $sql);
			$sql = str_replace('%DATA%', implode(', ' ,$colList), $sql); 
		}
		
		$retVel[] = $sql;
		$retVel[] = $bindDt;
		return $retVel;
	}
	// (string, array, string|array) => {} => (sql, bind)
	public function updateSql($tb, $data, $where){
		$bind = null;
		$sql = 'UPDATE "__TABLE__" SET %UPDATELIST% WHERE %WHERE%';
		$sql = str_replace('__TABLE__', $tb, $sql);
		$updateList = [];
		foreach($data as $k=>$v){
			$updateList[] = $k.' = \''.$v.'\'';
		}
		if(is_array($where)){
			$whList = [];
			$ctt = 1;
			$bind = [];
			foreach($where as $k=>$v){
				$key = 'args'.$ctt;
				$whList[] = $k.'=:'.$key;
				$bind[$key] = $v;
				$ctt++;
			}
			$where = implode(' AND ', $whList);
		}
		$sql = str_replace('%UPDATELIST%', implode(',', $updateList), $sql);
		$sql = str_replace('%WHERE%', $where, $sql);
		return [$sql, $bind];
	}
	// () => ($sql, $bind)
	public function deleteSql($tb, $where){
		$sql = 'DELETE FROM "__TABLE__" WHERE %WHERE%';
		$bind = null;
		$whList = [];
		$ctt = 1;
		foreach($where as $k=>$v){
			if(is_int($k)){
				$whList[] = '"'.$k.'"=:'.$k;
			}else{
				$key = 'args'.$ctt;
				$whList[] = '"'.$k.'"=:'.$key;
				if(empty($bind)) $bind = [];
				$bind[$key] = $v;
				$ctt++;
			}
		}
		return [$sql, $bind];
	}
 }