// 数据库适应期
// 2017年10月23日 星期一
// @author Joshua Conero

package sqler

import (
	"regexp"
	"strings"
)

// 数据库适配器
type db_adapter struct {
	dbMap map[string]string // 数据库配置信息
	vtype string            // 数据类型
}

// 数据适配器初始化
func getDbAdapter(vtype string) *db_adapter {
	da := &db_adapter{vtype: strings.ToLower(vtype)}
	return da.init()
}

// 获取数据库配置信息
func getDbSetting(name string) map[string]string {
	if name == "orc" {
		name = "oracle"
	}
	dbMap := map[string]string{}
	switch name {
	case "mysql":
		dbMap = mysqlSetting
	case "oracle":
		dbMap = oracleSetting
	default:
		dbMap = map[string]string{
			"col_sign": "",
			"val_sign": "'",
		}
	}
	return dbMap
}

// 适配器初始化
func (da *db_adapter) init() *db_adapter {
	vtype := da.vtype
	da.dbMap = getDbSetting(vtype)
	return da
}

// 后去类型标点符号
func (da *db_adapter) getCol() string {
	v, ok := da.dbMap["col_sign"]
	if !ok {
		v = ""
	}
	return v
}

// 获取数据表名
func (da *db_adapter) getColName(col string) string {
	col = strings.TrimSpace(col)
	sCol := da.getCol()
	// 已经存在则不重复添加
	re := regexp.MustCompile("^[a-zA-Z0-9-_]+.*")
	if re.MatchString(col) {
		col = sCol + col + sCol
	}
	return col
}

// 清除数据库中被键值标注的字符串
func (da *db_adapter) strClearCol(str string) string {
	re := regexp.MustCompile("^" + da.getCol() + "+.*")
	if re.MatchString(str) {
		strLen := len(str)
		str = str[1 : strLen-1]
	}
	return str
}

// 获取值标识符号
func (da *db_adapter) getVal() string {
	v, ok := da.dbMap["val_sign"]
	if !ok {
		v = ""
	}
	return v
}

// 后去值名称
func (da *db_adapter) getValStr(col string) string {
	sCol := da.getVal()
	name := sCol + col + sCol
	return name
}
