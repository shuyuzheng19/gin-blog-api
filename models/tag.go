package models

import (
	"common-web-framework/common"
	"common-web-framework/response"
)

// Tag 标签模型
type Tag struct {
	MyModel
	ID   int    `gorm:"primary_key;type:int;comment:标签ID;"`
	Name string `gorm:"size:255;unique;not null;comment:标签名称"`
}

func (t Tag) ToTagResponse() response.TagResponse {
	return response.TagResponse{
		Id:   t.ID,
		Name: t.Name,
	}
}

// TableName 返回与模型对应的数据库表名
func (*Tag) TableName() string { return common.TableNames.TagTableName }
