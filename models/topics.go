package models

import (
	"common-web-framework/common"
)

// Topic 专题模型
type Topic struct {
	MyModel
	ID          int    `gorm:"primary_key;type:int;comment:专题ID"`
	Name        string `gorm:"size:255;unique;not null;comment:专题名"`
	Description string `gorm:"size:255;not null;comment:专题描述"`
	CoverImage  string `gorm:"size:255;not null;comment:专题封面"`
	UserID      int    `gorm:"column:user_id;type:integer;comment:创建专题的用户"`
	User        User
}

// TableName 返回与模型对应的数据库表名
func (Topic) TableName() string { return common.TableNames.TopicTableName }
