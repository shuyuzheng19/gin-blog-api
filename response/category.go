package response

import "common-web-framework/common"

// CategoryResponse 博客分类概要
// @Description 博客分类概要
type CategoryResponse struct {
	Id   int    `json:"id"`   //分类ID
	Name string `json:"name"` //分类名称
}

// AdminOtherResponse 后台管理分类或标签模型
// @Description 后台管理分类或标签模型
type AdminOtherResponse struct {
	CreatedAt myTime `json:"createAt"`
	UpdatedAt myTime `json:"updateAt"`
	Id        int    `json:"id"`
	Name      string `json:"name"`
}

func (*CategoryResponse) TableName() string {
	return common.TableNames.CategoryTableName
}
