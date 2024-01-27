package models

import (
	"common-web-framework/common"
)

// Role 角色模型
type Role struct {
	ID          uint   `gorm:"primary_key;type:int;comment:角色ID"`
	Name        string `gorm:"size:255;unique;not null;comment:角色名"`
	Description string `gorm:"size:255;not null;comment:角色描述"`
}

func (*Role) TableName() string { return common.TableNames.RoleTableName }
