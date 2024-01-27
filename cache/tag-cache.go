package cache

import (
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/response"
	"common-web-framework/utils"
	"errors"
	"github.com/go-redis/redis"
	"strconv"
)

type TagCache struct {
	redis *redis.Client
}

// GetRandomTags 随机从缓存中取标签
func (t TagCache) GetRandomTags() (result []response.TagResponse, _ error) {

	var tags = t.redis.SRandMemberN(common.RandomTagListKey, common.RandomTagCount).Val()

	if len(tags) == 0 {
		return nil, errors.New("获取随机标签失败")
	}

	for _, str := range tags {
		result = append(result, utils.JsonToObject[response.TagResponse](str))
	}

	return result, nil
}

// RemoveTagCache 删除缓存的标签
func (t TagCache) RemoveTagCache() error {
	config.LOGGER.Info("已清空标签缓存")
	return t.redis.Del(common.RandomTagListKey, common.TagMapKey).Err()
}

// SetTags 将标签列表存入缓存
func (t TagCache) SetTags(tags []response.TagResponse) error {

	var jsons = make([]string, 0)

	for _, tag := range tags {
		jsons = append(jsons, utils.ObjectToJson(tag))
	}

	var key = common.RandomTagListKey

	var e = t.redis.SAdd(key, jsons).Err()

	if e != nil {
		return e
	}

	t.redis.Expire(key, common.RandomTagListExpire)
	return nil
}

// SetTagToMap 将标签简要存入缓存map
func (t TagCache) SetTagToMap(tag response.TagResponse) error {
	return t.redis.HSet(common.TagMapKey, strconv.Itoa(tag.Id), utils.ObjectToJson(tag)).Err()
}

// GetTagToMap 从缓存map获取标签简要
func (t TagCache) GetTagToMap(tagId int) (tag response.TagResponse, err error) {
	var val = t.redis.HGet(common.TagMapKey, strconv.Itoa(tagId)).Val()

	if val == "" {
		return tag, errors.New("从缓存获取标签信息失败")
	}

	return utils.JsonToObject[response.TagResponse](val), nil
}

func NewTagCache(r *redis.Client) TagCache {
	return TagCache{redis: r}
}
