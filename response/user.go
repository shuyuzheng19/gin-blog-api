package response

import "common-web-framework/common"

// UserResponse 用户信息
// @Description 返回用户的概要信息
type UserResponse struct {
	Id       int    `json:"id"`       //用户ID
	Nickname string `json:"nickName"` //用户名
	Avatar   string `json:"icon"`     //用户头像
	Role     string `json:"role"`     //用户角色
	Username string `json:"username"` //用户账号
}

// SimpleUserResponse 简洁的用户信息
// @Description 返回用户的简洁概要信息
type SimpleUserResponse struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	NickName string `json:"nickName"`
}

func (*SimpleUserResponse) TableName() string {
	return common.TableNames.UserTableName
}

// UserAdminResponse 后台管理用户列表
// @Description 后台管理用户列表
type UserAdminResponse struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	NickName  string `json:"nickname"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	CreatedAt myTime `json:"createAt"`
	RoleId    int    `json:"-"`
	Role      struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"role"`
}
