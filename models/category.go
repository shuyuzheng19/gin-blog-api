package models

import (
	"common-web-framework/common"
	"common-web-framework/response"
)

// Category 分类模型
type Category struct {
	MyModel
	ID   int    `gorm:"primary_key;type:int;comment:分类ID"`
	Name string `gorm:"size:255;unique;not null;comment:分类名称"`
}

// TableName 返回与模型对应的数据库表名
func (*Category) TableName() string { return common.TableNames.CategoryTableName }

func (c Category) ToCategoryResponse() response.CategoryResponse {
	return response.CategoryResponse{
		Id:   c.ID,
		Name: c.Name,
	}
}
