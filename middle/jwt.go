package middle

import (
	"common-web-framework/common"
	"common-web-framework/helper"
	"common-web-framework/models"
	"common-web-framework/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

type RoleName string

const (
	UserRole  RoleName = "USER"        //普通用户
	AdminRole RoleName = "ADMIN"       //管理员用户
	SuperRole RoleName = "SUPER_ADMIN" //超级管理员
)

const tokenType = "Bearer " //token类型

const tokenHeader = "Authorization" //token请求头

var GetJwtUser = func(id int) models.User {
	return models.User{}
}

var GetToken = func(id int) string {
	return ""
}

func ParseToken(header string, context *gin.Context) models.User {
	if header == "" || !strings.HasPrefix(header, tokenType) {
		helper.ResultFailToToResponse(common.NoLogin, context)
		return models.User{}
	}

	var token = strings.Replace(header, tokenType, "", 1)

	var uid = utils.ParseTokenToUserId(token)

	if uid == -1 {
		helper.ResultFailToToResponse(common.ParseTokenFail, context)
		return models.User{}
	}

	var redisToken = GetToken(uid)

	if redisToken != token {
		helper.ResultFailToToResponse(common.TokenExpireFail, context)
		return models.User{}
	}

	var user = GetJwtUser(uid)

	if user.ID == 0 {
		helper.ResultFailToToResponse(common.Unauthorized, context)
		return models.User{}
	}

	return user
}

// JwtMiddle 验证身份中间件
func JwtMiddle(roleName RoleName) gin.HandlerFunc {
	return func(context *gin.Context) {
		var header = context.GetHeader(tokenHeader)

		var user = ParseToken(header, context)

		if context.IsAborted() {
			return
		}

		var role = user.Role.Name

		var isAuth = false

		if roleName == UserRole || role == string(SuperRole) {
			isAuth = true
		} else if roleName == AdminRole && role == string(AdminRole) {
			isAuth = true
		}

		if isAuth {
			context.Set("user", user)
			context.Next()
		} else {
			helper.ResultFailToToResponse(common.Forbidden, context)
			return
		}

	}
}
