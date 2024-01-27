package service

import (
	"common-web-framework/cache"
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/helper"
	"common-web-framework/models"
	"common-web-framework/repository"
	"common-web-framework/request"
	"common-web-framework/response"
	"common-web-framework/utils"
	"fmt"
	"go.uber.org/zap"
)

type UserServiceImpl struct {
	repository repository.UserRepository
	cache      cache.UserCache
}

func (u UserServiceImpl) GetAdminUserList(req request.OtherAdminFilter) response.PageInfo {
	var blogs, count = u.repository.GetAdminUsers(req)

	return response.PageInfo{
		Page:  req.Page,
		Count: count,
		Size:  common.AdminUserCount,
		Data:  blogs,
	}
}

func (u UserServiceImpl) UpdateUserRole(uid int, role int) int64 {
	var count = u.repository.UpdateRole(uid, role)

	if count == 0 {
		helper.ErrorToResponse(common.UpdateFail)
	} else {
		config.LOGGER.Info("修改用户角色成功", zap.Int("user_id", uid),
			zap.Int("role_id", role))
		u.cache.DeleteUser(uid)
	}

	return count
}

func (u UserServiceImpl) Logout(uid int) {
	u.cache.RemoveToken(uid)
}

func (u UserServiceImpl) GetToken(id int) string {
	return u.cache.GetToken(id)
}

func (u UserServiceImpl) Contact(req request.ContactRequest) {
	config.ValidateError(req)

	var text = fmt.Sprintf("<h3>%s</h3><p>对方名字: %s</p><p>对方邮箱: %s</p>留言内容:<p>%s</p>",
		req.Subject, req.Name, req.Email, req.Content)

	config.CONFIG.Email.SendEmail(config.CONFIG.MyEmail, req.Subject, true, text)

	config.LOGGER.Info("联系我，用户发送信息到我的邮件，请注意查收", zap.Any("info", req))
}

func (u UserServiceImpl) GetWebSiteConfig() response.BlogConfigInfo {
	return u.cache.GetWebSiteConfig()
}

func (u UserServiceImpl) SetWebSiteConfig(c response.BlogConfigInfo) {
	config.LOGGER.Info("更新网站配置")
	u.cache.SetWebSiteConfig(c)
}

func (u UserServiceImpl) GetUser(id int) models.User {
	if user := u.cache.GetUser(id); user.ID == 0 {
		var dbUser = u.repository.FindById(id)
		u.cache.SetUser(id, dbUser)
		return dbUser
	} else {
		return user
	}
}

func (u UserServiceImpl) Login(request request.LoginRequest) response.TokenResponse {
	config.ValidateError(request)

	var encodingPassword = utils.EncryptPassword(request.Password)

	var user = u.repository.FindByUsernameAndPassword(request.Username, encodingPassword)

	if user.ID == 0 {
		helper.ErrorToResponse(common.LoginFail)
	}

	var token = utils.CreateAccessToken(user.ID, user.Username)

	config.LOGGER.Info("用户登录成功", zap.Int("id", user.ID),
		zap.String("username", user.Username),
		zap.String("nickname", user.NickName))

	u.cache.SetToken(user.ID, token.Token)

	return token
}

func (u UserServiceImpl) ValidateEmailCode(email string, code string) {

	var cacheCode = u.cache.GetEmailCode(email)

	if cacheCode == "" || cacheCode != code {
		helper.ErrorToResponse(common.EmailCodeValidate)
	}
}

func (u UserServiceImpl) SendCodeToEmail(email string) {

	var code = utils.RandomNumberCode()

	config.CONFIG.Email.SendEmail(email, "注册验证码", false, code)

	config.LOGGER.Info("发送邮箱验证码", zap.String("code", code))

	u.cache.SetEmailCode(code, email)
}

func (u UserServiceImpl) RegisteredUser(request request.UserRequest) {
	config.ValidateError(request)

	u.ValidateEmailCode(request.Email, request.Code)

	var user = request.ToUserDo()

	if err := u.repository.Save(user); err != nil {
		helper.ErrorToResponse(common.RegisteredCode)
	}

	config.LOGGER.Info("用户注册成功", zap.Int("id", user.ID),
		zap.String("username", user.Username),
		zap.String("nickname", user.NickName))
}

func NewUserService() UserService {
	var repository = repository.NewUserRepository(config.DB)
	var cache = cache.NewUserCache(config.REDIS)
	return UserServiceImpl{repository: repository, cache: cache}
}
