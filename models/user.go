package models

import (
	"common-web-framework/common"
	"common-web-framework/response"
)

// User 用户表模型
type User struct {
	MyModel
	ID       int    `gorm:"primary_key;type:int;comment:用户ID"`
	Username string `gorm:"size:16;unique;not null;comment:用户账号"`
	Password string `gorm:"size:255;not null;comment用户密码:"`
	Email    string `gorm:"size:255;unique;not null;comment:用户邮箱"`
	Avatar   string `gorm:"default:'test.png';comment:用户头像"`
	NickName string `gorm:"size:50;not null;comment:用户名称"`
	RoleID   uint   `gorm:"column:role_id;type:integer;comment:角色ID"`
	Role     Role
}

func (*User) TableName() string { return common.TableNames.UserTableName }

func (u User) ToUserResponse() response.UserResponse {

	return response.UserResponse{
		Id:       u.ID,
		Nickname: u.NickName,
		Avatar:   u.Avatar,
		Role:     u.Role.Name,
		Username: u.Username,
	}
}

func (u User) ToSimpleUserResponse() response.SimpleUserResponse {
	return response.SimpleUserResponse{
		Id:       u.ID,
		NickName: u.NickName,
	}
}
