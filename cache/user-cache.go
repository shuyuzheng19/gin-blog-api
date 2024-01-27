package cache

import (
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/models"
	"common-web-framework/response"
	"common-web-framework/utils"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strconv"
)

type UserCache struct {
	redis *redis.Client
}

// SetToken 将token缓存
func (u UserCache) SetToken(id int, token string) error {
	return u.redis.Set(fmt.Sprintf(common.UserTokenKey+"%d", id), token, common.TokenExpire).Err()
}

// GetToken 从缓存中获取token
func (u UserCache) GetToken(id int) string {
	return u.redis.Get(fmt.Sprintf(common.UserTokenKey+"%d", id)).Val()
}

// RemoveToken 删除用户的token
func (u UserCache) RemoveToken(uid int) error {
	return u.redis.Del(common.UserTokenKey + strconv.Itoa(uid)).Err()
}

// SetUser 将用户信息缓存
func (u UserCache) SetUser(id int, user models.User) error {
	var str = utils.ObjectToJson(user)
	return u.redis.Set(common.UserInfoKey+strconv.Itoa(id), str, common.UserInfoKeyExpire).Err()
}

func (u UserCache) DeleteUser(id int) error {
	config.LOGGER.Info("清除用户信息", zap.Int("uid", id))
	return u.redis.Del(common.UserInfoKey + strconv.Itoa(id)).Err()
}

// GetUser 从缓存中获取用户信息
func (u UserCache) GetUser(id int) models.User {
	var result = u.redis.Get(common.UserInfoKey + strconv.Itoa(id)).Val()
	return utils.JsonToObject[models.User](result)
}

// SetEmailCode 将邮箱验证码缓存
func (u UserCache) SetEmailCode(code, email string) error {
	return u.redis.Set(fmt.Sprintf(common.EmailCodeKey+"%s", email), code, common.EmailCodeKeyExpire).Err()
}

// GetEmailCode 从缓存中获取邮箱验证码
func (u UserCache) GetEmailCode(email string) string {
	return u.redis.Get(fmt.Sprintf(common.EmailCodeKey+"%s", email)).Val()
}

// GetWebSiteConfig 获取网站信息
func (u UserCache) GetWebSiteConfig() response.BlogConfigInfo {

	var str = u.redis.Get(common.WebSiteConfigKey).Val()

	if str == "" {
		return response.GetDefaultBlogConfigInfo()
	}

	return utils.JsonToObject[response.BlogConfigInfo](str)
}

// SetWebSiteConfig 缓存网站信息
func (u UserCache) SetWebSiteConfig(siteConfig response.BlogConfigInfo) error {

	var str = utils.ObjectToJson(siteConfig)

	return u.redis.Set(common.WebSiteConfigKey, str, -1).Err()
}

func NewUserCache(r *redis.Client) UserCache {
	return UserCache{redis: r}
}
