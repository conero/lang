// sql 生成器器 golang 版本
// 2017年10月23日 星期一
// @author Joshua Conero

package sqler

import (
	"strings"
)

const (
	Author  = "Joshua Conero"
	Date    = "2017年10月23日 星期一"
	Version = "0.0.1"
)

// 数据类类型 * 默认mysql
var DbType string = "mysql"

// 数据表
func Table(name string) *Sql {
	da := getDbAdapter(DbType)
	return &Sql{table: da.getColName(name), feildList: []string{}, da: da}
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
	sql.da = getDbAdapter(vtype)
	sql.table = sql.da.getColName(sql.table)
	return sql
}

func (sql *Sql) Field(fields ...string) *Sql {
	for _, v := range fields {
		sql.feildList = append(sql.feildList, v)
	}
	return sql
}
/**
	设置数据库名称
 */
func (sql *Sql) Data(data map[string]interface{}) *Sql {
	for k,v :=  range data{
		sql.dataMap[k] = v
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
	sqlText := `INSERT INTO ` + sql.table + `() VALUES ()`
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