// sql 生成器器 golang 版本
// 2017年10月23日 星期一
// @author Joshua Conero

package sqler

import (
	"strconv"
	"strings"
)

// 数据类类型 * 默认mysql
var DbType string = "mysql"

// 通过数据表入口
func Table(name string) *Sql {
	da := getDbAdapter(DbType)
	return &Sql{
		table:     da.getColName(name),
		feildList: []string{},
		da:        da,
		dataMap:   make(map[string]interface{}),
		//whereQueue:make([]interface{}, 10),
		whereQueue: []interface{}{},
	}
}

// 通过设置数据库类型入口

func Type(vType string) *Sql {
	da := getDbAdapter(DbType)
	return &Sql{
		table:     "",
		feildList: []string{},
		da:        da,
		dataMap:   make(map[string]interface{}),
		//whereQueue:make([]interface{}, 10),
		whereQueue: []interface{}{},
	}
}

// sql 类型
type Sql struct {
	table      string   // 数据表名
	alias      string   // 数据表别名
	feildList  []string // 数据列名
	dataMap    map[string]interface{}
	whereQueue []interface{} // 条件队列
	da         *db_adapter
}

// 设置数据表
func (sql *Sql) Table(tableParam ...string) *Sql {
	pLen := len(tableParam)
	if pLen > 0 {
		sql.table = sql.da.getColName(tableParam[0])
	}
	if pLen > 1 {
		sql.alias = tableParam[1]
	}
	return sql
}

// 中间件
func (sql *Sql) Type(vtype string) *Sql {
	table := sql.table
	table = sql.da.strClearCol(table)
	sql.da = getDbAdapter(vtype)
	sql.table = sql.da.getColName(table)
	return sql
}

/**
设置查询字段
fields			string
				map[string]string
*/
func (sql *Sql) Field(fields ...interface{}) *Sql {
	colStr := sql.da.getCol()
	for _, v := range fields {
		switch v.(type) {
		case string:
			value := v.(string)
			value = colStr + value + colStr
			sql.feildList = append(sql.feildList, value)
		case map[string]string:
			value := v.(map[string]string)
			//key, valueStr := value
			for key, valueStr := range value {
				nValue := colStr + key + colStr + " AS " + colStr + valueStr + colStr
				sql.feildList = append(sql.feildList, nValue)
			}
		}
	}
	return sql
}

/**
设置数据库名称
*/
func (sql *Sql) Data(datas ...interface{}) *Sql {
	for _, data := range datas {
		switch data.(type) {
		case map[string]string:
			for k, v := range data.(map[string]string) {
				sql.dataMap[k] = v
			}
		case map[string]int:
			for k, v := range data.(map[string]int) {
				sql.dataMap[k] = v
			}
		}
	}
	return sql
}

// 设置条件
// string/map[string]string/map[string]int
func (sql *Sql) Where(wheres ...interface{}) *Sql {
	for _, v := range wheres {
		sql.whereQueue = append(sql.whereQueue, v)
	}
	return sql
}

// 条件解析
func (sql *Sql) whereParse() string {
	whereStr := ""
	whereQueue := sql.whereQueue
	whereStrArr := []string{}
	for _, v := range whereQueue {
		switch v.(type) {
		case string:
			value := v.(string)
			whereStrArr = append(whereStrArr, value)
		case map[string]string:
			valueMap := v.(map[string]string)
			for key, value := range valueMap {
				value = sql.da.getColName(key) + "=" + sql.da.getValStr(value)
				whereStrArr = append(whereStrArr, value)
			}
		case map[string]int:
			valueMap := v.(map[string]int)
			for key, value := range valueMap {
				sValue := sql.da.getColName(key) + "=" + strconv.Itoa(value)
				whereStrArr = append(whereStrArr, sValue)
			}
		}
	}
	if len(whereStrArr) > 0 {
		whereStr = strings.Join(whereStrArr, " AND ")
	}
	// 获取条件以后一直清除缓存
	sql.whereQueue = []interface{}{}
	if "" != whereStr {
		whereStr = " WHERE " + whereStr
	}
	return whereStr
}

// 私有方法
// 出口
func (sql *Sql) Select() string {
	var colTxt string
	if len(sql.feildList) > 0 {
		colTxt = strings.Join(sql.feildList, ",")
		sql.feildList = []string{}
	} else {
		colTxt = "*"
	}
	sqlStr := `SELECT ` + colTxt + ` FROM ` + sql.table + sql.whereParse()
	return sqlStr
}

// 数据拆分为 列[]string, 值[]string 元组
func (sql *Sql) dataToTuple() ([]string, []string) {
	da := sql.da
	colArr := []string{}
	valArr := []string{}
	for k, v := range sql.dataMap {
		colArr = append(colArr, da.getColName(k))
		switch v.(type) {
		case string:
			valArr = append(valArr, da.getValStr(v.(string)))
		case int:
			valArr = append(valArr, strconv.Itoa(v.(int)))
		case int8:
			valArr = append(valArr, strconv.Itoa(v.(int)))
		case float32:
			valArr = append(valArr, strconv.FormatFloat(float64(v.(float32)), 'E', -1, 32))
		}
	}
	sql.dataMap = map[string]interface{}{}
	return colArr, valArr
}

// data 转变为 更新列表
func (sql *Sql) dataToUpdStr() []string {
	updArr := []string{}
	da := sql.da
	for k, v := range sql.dataMap {
		key := da.getColName(k)
		switch v.(type) {
		case string:
			value := key + "=" + da.getValStr(v.(string))
			updArr = append(updArr, value)
		case int:
			value := key + "=" + strconv.Itoa(v.(int))
			updArr = append(updArr, value)
		case int8:
			value := key + "=" + strconv.Itoa(v.(int))
			updArr = append(updArr, value)
		case float32:
			value := key + "=" + strconv.FormatFloat(float64(v.(float32)), 'E', -1, 32)
			updArr = append(updArr, value)
		}
	}
	sql.dataMap = map[string]interface{}{}
	return updArr
}

// 新增数据
func (sql *Sql) Insert() string {
	colArr, valArr := sql.dataToTuple()
	sqlText := `INSERT INTO ` + sql.table + `(` + strings.Join(colArr, ",") + `) VALUES (` + strings.Join(valArr, ",") + `)`
	return sqlText
}

// 更新数据
func (sql *Sql) Update() string {
	sqlText := `UPDATE ` + sql.table +
		` SET ` +
		strings.Join(sql.dataToUpdStr(), ",") +
		sql.whereParse()
	return sqlText
}

// 删除数据
func (sql *Sql) Delete() string {
	sqlText := `DELETE FROM ` + sql.table + sql.whereParse()
	return sqlText
}
