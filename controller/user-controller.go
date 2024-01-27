package controller

import (
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/helper"
	"common-web-framework/request"
	"common-web-framework/response"
	"common-web-framework/service"
	"common-web-framework/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct {
	service service.UserService
}

// SendCodeToEmail 发送邮件
// @Summary 发送邮件
// @Description 发送注册验证码到用户邮箱
// @Tags 用户相关接口
// @Param email query string true "用户的邮箱"
// @Produce json
// @Success 200 {object} common.R
// @Router /user/send_email [get]
func (u UserController) SendCodeToEmail(ctx *gin.Context) {

	var email = ctx.Query("email")

	if email == "" {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	u.service.SendCodeToEmail(email)

	helper.ResultSuccessToResponse(nil, ctx)
}

// Logout 退出登录
// @Summary 退出登录
// @Description 退出当前用户登录信息
// @Tags 用户相关接口
// @Produce json
// @Security JWT
// @Success 200 {object} common.R
// @Router /user/logout [get]
func (u UserController) Logout(ctx *gin.Context) {

	var uid = utils.GetUserInfo(ctx).ID

	u.service.Logout(uid)

	helper.ResultSuccessToResponse(nil, ctx)

}

// ContactMe 联系我
// @Summary 联系我
// @Description 用户联系我，将信息发送到我的邮箱
// @Tags 用户相关接口
// @Param request body request.ContactRequest true "联系我请求模型"
// @Produce json
// @Success 200 {object} common.R
// @Router /user/contact_me [post]
func (u UserController) ContactMe(ctx *gin.Context) {

	var req request.ContactRequest

	ctx.ShouldBindJSON(&req)

	u.service.Contact(req)

	helper.ResultSuccessToResponse(nil, ctx)
}

// RegisteredUser 注册用户
// @Summary 注册用户
// @Description 注册用户
// @Tags 用户相关接口
// @Param request body request.UserRequest true "注册用户的结构体"
// @Produce json
// @Accept json
// @Success 200 {object} common.R
// @Router /user/registered [post]
func (u UserController) RegisteredUser(ctx *gin.Context) {
	var userRequest request.UserRequest

	ctx.ShouldBindJSON(&userRequest)

	u.service.RegisteredUser(userRequest)

	helper.ResultSuccessToResponse(nil, ctx)
}

// GetUserInfo 获取当前登录用户信息
// @Security JWT
// @Summary 用户信息
// @Description 获取当前登录用户信息
// @Tags 用户相关接口
// @Produce json
// @Success 200 {object} common.R{data=response.UserResponse}
// @Router /user/auth/get [get]
func (u UserController) GetUserInfo(ctx *gin.Context) {
	helper.ResultSuccessToResponse(utils.GetUserInfo(ctx).ToUserResponse(), ctx)
}

// GetWebSiteConfig 获取网站配置
// @Summary 个人信息
// @Description 获取个人信息介绍
// @Tags 用户相关接口
// @Produce json
// @Success 200 {object} common.R
// @Router /user/config [get]
func (u UserController) GetWebSiteConfig(ctx *gin.Context) {
	helper.ResultSuccessToResponse(u.service.GetWebSiteConfig(), ctx)
}

// SetWebSiteConfig 修改网站配置
// @Security JWT
// @Summary 修改个人信息
// @Description 修改个人信息
// @Tags 用户相关接口
// @Produce json
// @Param request body response.BlogConfigInfo true "个人信息"
// @Success 200 {object} common.R
// @Router /user/config [put]
func (u UserController) SetWebSiteConfig(ctx *gin.Context) {
	var config response.BlogConfigInfo

	ctx.ShouldBindJSON(&config)

	u.service.SetWebSiteConfig(config)

	helper.ResultSuccessToResponse(nil, ctx)
}

// ClearCache 清空所有缓存
// @Summary 清空所有缓存
// @Description 清空所有缓存
// @Tags 博客相关接口,后台管理接口
// @param patten query string true "要清空的键"
// @Produce json
// @Security JWT
// @Success 200 {object} common.R
// @Router /user/admin/clear_cache [get]
func (u UserController) ClearCache(ctx *gin.Context) {

	var patten = ctx.Query("patten")

	if patten == "" {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	var keys = config.REDIS.Keys(patten).Val()

	config.REDIS.Del(keys...)
	helper.ResultSuccessToResponse(nil, ctx)
}

func (u UserController) IsChinaIp(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/javascript")

	var flag = utils.IsCnIp(utils.GetIPAddress(ctx.Request))

	fmt.Fprintf(ctx.Writer, "ipCallback(%t)", flag)
}

// Login 登录
// @Summary 登录
// @Description 用户登录
// @Tags 用户相关接口
// @Param request body request.LoginRequest true "用户登录请求体"
// @Produce json
// @Accept json
// @Success 200 {object} common.R
// @Router /user/login [post]
func (u UserController) Login(ctx *gin.Context) {
	var loginRequest request.LoginRequest

	ctx.ShouldBindJSON(&loginRequest)

	var tokenResponse = u.service.Login(loginRequest)

	helper.ResultSuccessToResponse(tokenResponse, ctx)
}

// GetAdminUserList 获取后台管理用户列表
// @Summary 后台管理用户列表
// @Description 获取后台管理用户列表
// @Tags 用户相关接口,后台管理接口
// @Param request query request.OtherAdminFilter true "过滤条件"
// @Produce json
// @Accept json
// @Success 200 {object} common.R
// @Router /user/admin/list [get]
func (u UserController) GetAdminUserList(ctx *gin.Context) {

	var req request.OtherAdminFilter

	req.Page = 1

	ctx.ShouldBindQuery(&req)

	var result = u.service.GetAdminUserList(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// UpdateUserRole 修改用户角色
// @Summary 修改用户角色
// @Description 修改用户角色
// @Tags 用户相关接口,后台管理接口
// @Param id query int true "用户id"
// @Param rid query int true "角色id"
// @Produce json
// @Accept json
// @Success 200 {object} common.R
// @Router /user/admin/update_role [put]
func (u UserController) UpdateUserRole(ctx *gin.Context) {

	var idStr = ctx.Query("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	var roleStr = ctx.Query("rid")

	rid, err := strconv.Atoi(roleStr)

	if err != nil {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	var result = u.service.UpdateUserRole(id, rid)

	helper.ResultSuccessToResponse(result, ctx)
}

func NewUserController(service service.UserService) UserController {
	return UserController{service: service}
}
