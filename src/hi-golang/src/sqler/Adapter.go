// 数据库适应期
// 2017年10月23日 星期一
// @author Joshua Conero

package sqler

import "strings"

/**
数据库适配器
 */
type db_adapter struct {
	dbMap map[string]string
	vtype string
}

func getDbAdapter(vtype string)  *db_adapter{
	da := &db_adapter{vtype: strings.ToLower(vtype)}
	return  da.init()
}

func (da *db_adapter) init ()  *db_adapter{
	vtype := da.vtype
	if vtype == "orc"{
		vtype = "oracle"
	}
	switch vtype {
	case "mysql":
		da.dbMap = mysqlSetting
	case "orcle":
		da.dbMap = oracleSetting
	default:
		da.dbMap = map[string]string{
			"col_sign": "",
			"val_sign": "'",
		}
	}
	return da
}
func (da *db_adapter) getCol() string {
	v,ok := da.dbMap["col_sign"]
	if !ok{
		v = ""
	}
	return  v
}
func (da *db_adapter) getColName(col string) string {
	sCol := da.getCol()
	name := sCol + col + sCol
	return name
}
func (da *db_adapter) getVal() string {
	v,ok := da.dbMap["val_sign"]
	if !ok{
		v = ""
	}
	return  v
}
func (da *db_adapter) getValStr(col string) string {
	sCol := da.getVal()
	name := sCol + col + sCol
	return name
}