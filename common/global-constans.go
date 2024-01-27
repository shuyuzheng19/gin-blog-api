package common

import (
	"time"
)

const TokenExpire = time.Hour * 24 * 7 //token过期实际

var TableNames GlobalDbTableNames

const UserRole = 1

const AdminRole = 2

const SuperAdminRole = 3

type GlobalDbTableNames struct {
	UserTableName     string //用户表名
	FileTableName     string //文件表名
	FileMd5TableName  string //文件md5表名
	CategoryTableName string //分类表名
	RoleTableName     string //角色表名
	TagTableName      string //标签表名
	TopicTableName    string //专题表名
	BlogTableName     string //博客表名
	BlogTagTableName  string //博客标签表名
}

// 初始化数据库生成的表名
func init() {
	TableNames = GlobalDbTableNames{
		UserTableName:     "users",
		FileTableName:     "file_infos",
		CategoryTableName: "categories",
		RoleTableName:     "roles",
		TagTableName:      "tags",
		TopicTableName:    "topics",
		BlogTableName:     "blogs",
		BlogTagTableName:  "blogs_tags",
		FileMd5TableName:  "file_md5_infos",
	}
}
