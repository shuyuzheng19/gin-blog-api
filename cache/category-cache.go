package cache

import (
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/response"
	"common-web-framework/utils"
	"github.com/go-redis/redis"
)

type CategoryCache struct {
	redis *redis.Client
}

// SetCategoryList 缓存分类列表
func (c CategoryCache) SetCategoryList(categories []response.CategoryResponse) error {
	var str = utils.ObjectToJson(categories)
	return c.redis.Set(common.CategoryListKey, str, common.CategoryListExpire).Err()
}

// RemoveCategoryCache 删除分类缓存
func (c CategoryCache) RemoveCategoryCache() error {
	config.LOGGER.Info("清空分类缓存")
	return c.redis.Del(common.CategoryListKey).Err()
}

// GetCategoryList 缓存分类列表
func (c CategoryCache) GetCategoryList() (result []response.CategoryResponse, err error) {
	var r = c.redis.Get(common.CategoryListKey)
	err = r.Err()
	if err != nil {
		return nil, err
	}
	return utils.JsonToObject[[]response.CategoryResponse](r.Val()), nil
}

func NewCategoryCache(r *redis.Client) CategoryCache {
	return CategoryCache{redis: r}
}
