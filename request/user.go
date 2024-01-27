package request

import (
	"common-web-framework/models"
	"common-web-framework/utils"
)

// UserRequest 用户注册请求体
// @Description 用户注册请求体
type UserRequest struct {
	Username string `json:"username" validate:"required,min=8,max=16"` //用户账号
	Password string `json:"password" validate:"required,min=8,max=16"` //用户密码
	Email    string `json:"email" validate:"required,email"`           //用户邮箱
	NickName string `json:"nickName" validate:"required,max=50,min=1"` //用户名称
	Code     string `json:"code" validate:"required,min=6,max=6"`      //邮箱验证码
}

// LoginRequest 账号登录请求体
// @Description 账号登录请求体
type LoginRequest struct { //账号登录请求体
	Username string `json:"username" validate:"required"` //账号
	Password string `json:"password" validate:"required"` //密码
}

func (r UserRequest) ToUserDo() models.User {
	return models.User{
		Username: r.Username,
		Password: utils.EncryptPassword(r.Password),
		Email:    r.Email,
		NickName: r.NickName,
		RoleID:   1,
	}
}

// ContactRequest 联系我请求
// @Description 联系我请求模型
type ContactRequest struct {
	Name    string `json:"name"  validate:"required"`       //你的名字
	Email   string `json:"email" validate:"required,email"` //你的邮箱
	Subject string `json:"subject" validate:"required"`     //邮件主题
	Content string `json:"content" validate:"required"`     //邮件内容
}
