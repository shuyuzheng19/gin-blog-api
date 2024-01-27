package service

import (
	"common-web-framework/common"
	"common-web-framework/config"
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
	"time"
)

type DataBaseServiceImpl struct {
	types  map[string]interface{}
	db     *gorm.DB
	tables common.GlobalDbTableNames
}

func (d DataBaseServiceImpl) ExecDataBaseSQL(sql string) int64 {
	return d.db.Exec(sql).RowsAffected
}

func (d DataBaseServiceImpl) getGlobalResult(tableName string, page int) string {

	var records []map[string]interface{}

	d.db.Table(tableName).Offset((page - 1) * 10).
		Limit(common.DataBaseSelectInsertCount).Scan(&records)

	var length = len(records)

	if length == 0 {
		return ""
	}

	var build = strings.Builder{}

	var firstRecord = records[0]
	var typeMap = make(map[string]string)
	for column, value := range firstRecord {
		var valType = reflect.TypeOf(value)
		typeMap[column] = fmt.Sprintf("%v", valType)
	}
	for _, record := range records {
		columns := make([]string, 0)
		values := make([]string, 0)

		for column, value := range record {
			columns = append(columns, column)
			var t = d.types[typeMap[column]]
			var getValue = t.(func(t interface{}) string)
			values = append(values, getValue(value))
		}
		var sql = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", tableName, strings.Join(columns, ", "), strings.Join(values, ", "))
		build.WriteString(sql + "\n")
	}

	return build.String()
}

func (d DataBaseServiceImpl) GetBlogInsertSQL(page int) string {
	return d.getGlobalResult(d.tables.BlogTableName, page)
}

func (d DataBaseServiceImpl) GetTagInsertSQL(page int) string {
	return d.getGlobalResult(d.tables.TagTableName, page)
}

func (d DataBaseServiceImpl) GetCategoryInsertSQL(page int) string {
	return d.getGlobalResult(d.tables.CategoryTableName, page)
}

func (d DataBaseServiceImpl) GetTopicInsertSQL(page int) string {
	return d.getGlobalResult(d.tables.TopicTableName, page)
}

func (d DataBaseServiceImpl) GetUserInsertSQL(page int) string {
	return d.getGlobalResult(d.tables.UserTableName, page)
}

func (d DataBaseServiceImpl) GetRoleInsertSQL(page int) string {
	return d.getGlobalResult(d.tables.RoleTableName, page)
}

func (d DataBaseServiceImpl) GetFileInsertSQL(page int) string {
	return d.getGlobalResult(d.tables.FileTableName, page)
}

func (d DataBaseServiceImpl) GetFileMd5InsertSQL(page int) string {
	return d.getGlobalResult(d.tables.FileMd5TableName, page)
}

func (d DataBaseServiceImpl) GetBlogTagInsertSQL(page int) string {
	return d.getGlobalResult(d.tables.BlogTagTableName, page)
}

func getInt(number interface{}) string {
	if number == nil {
		return "null"
	}
	return fmt.Sprintf("%d", number)
}

func getFloat(number interface{}) string {
	if number == nil {
		return "null"
	}
	return fmt.Sprintf("%f", number)
}

func getTime(t interface{}) string {
	if t == nil {
		return "null"
	}
	return fmt.Sprintf("'%s'", t.(time.Time).Format("2006-01-02 15:04:05"))
}

func getNull(t interface{}) string {
	if t == nil {
		return "null"
	}
	return fmt.Sprintf("%s", "null")
}

func getString(t interface{}) string {
	if t == nil {
		return "null"
	}

	var str = t.(string)

	return fmt.Sprintf("'%s'", str)
}

func getBool(t interface{}) string {
	if t == nil {
		return "null"
	}
	return fmt.Sprintf("%t", t.(bool))
}

func NewDataBaseService() DataBaseService {
	var maps = map[string]interface{}{
		"int":       getInt,
		"int8":      getInt,
		"int16":     getInt,
		"int32":     getInt,
		"int64":     getInt,
		"uint":      getInt,
		"uint8":     getInt,
		"uint16":    getInt,
		"uint32":    getInt,
		"uint64":    getInt,
		"float32":   getFloat,
		"float64":   getFloat,
		"bool":      getBool,
		"time.Time": getTime,
		"<nil>":     getNull,
		"string":    getString,
	}

	return DataBaseServiceImpl{types: maps, db: config.DB, tables: common.TableNames}
}
