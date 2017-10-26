// sql 生成器器 golang 版本
// 2017年10月23日 星期一
// @author Joshua Conero

package sqler

import (
	"strings"
	"strconv"
)

const (
	Author  = "Joshua Conero"
	Date    = "20171026"
	Version = "0.0.3"
)

// 数据类类型 * 默认mysql
var DbType string = "mysql"

// 数据表
func Table(name string) *Sql {
	da := getDbAdapter(DbType)
	return &Sql{
		table: da.getColName(name),
		feildList: []string{},
		da: da,
		dataMap:make(map[string]interface{}),
	}
}

// sql 类型
type Sql struct {
	table     string   // 数据表名
	feildList []string // 数据列名
	dataMap   map[string]interface{}
	da *db_adapter
}

// 入口

// 中间件
func (sql *Sql) Type(vtype string) *Sql {
	table := sql.table
	table = sql.da.strClearCol(table)
	sql.da = getDbAdapter(vtype)
	sql.table = sql.da.getColName(table)
	return sql
}
/**
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
			for key, valueStr := range value{
				nValue := colStr + key+ colStr + " AS "+ colStr + valueStr + colStr
				sql.feildList = append(sql.feildList, nValue)
			}
		}
	}
	return sql
}
/**
	设置数据库名称
 */
func (sql *Sql) Data(datas ... interface{}) *Sql {
	for _,data :=  range datas{
		switch data.(type) {
		case map[string]string:
			for k, v := range data.(map[string]string){
				sql.dataMap[k] = v
			}
		}
	}
	return sql
}
/**
	where
 */
func (sql *Sql) Where(where interface{}) *Sql  {
	return sql
}
// 私有方法
// 出口
func (sql *Sql) Select() string {
	var colTxt string
	if len(sql.feildList) > 0 {
		colTxt = strings.Join(sql.feildList, ",")
	} else {
		colTxt = "*"
	}
	sqlStr := `SELECT ` + colTxt + ` FROM ` + sql.table
	return sqlStr
}

func (sql *Sql) Insert() string {
	colSgStr := sql.da.getCol()
	valSgStr := sql.da.getVal()
	colArr := []string{}
	valArr := []string{}
	for k, v := range sql.dataMap{
		colArr = append(colArr, colSgStr + k + colSgStr)
		switch v.(type) {
		case string:
			valArr = append(valArr, valSgStr + v.(string) + valSgStr)
		case int:
			valArr = append(valArr,  strconv.Itoa(v.(int)))
		case int8:
			valArr = append(valArr,  strconv.Itoa(v.(int)))
		case float32:
			valArr = append(valArr,  strconv.FormatFloat(float64(v.(float32)), 'E', -1, 32))
		}
	}
	sqlText := `INSERT INTO ` + sql.table + `(`+strings.Join(colArr, ",")+`) VALUES (`+strings.Join(valArr, ",")+`)`
	return sqlText
}

func (sql *Sql) Update() string  {
	sqlText := `UPDATE `+sql.table+` SET `
	return sqlText
}

func (sql *Sql) Delete() string  {
	sqlText := `DELETE FROM `+sql.table
	return  sqlText
}