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

type TopicCache struct {
	redis *redis.Client
}

// SetTopicPage 缓存某页的专题
func (t TopicCache) SetTopicPage(page int, pageInfo response.PageInfo) error {
	return t.redis.Set(common.TopicPageKey+strconv.Itoa(page), &pageInfo, common.TopicPageExpire).Err()
}

// RemoveTopicCache 删除缓存专题
func (t TopicCache) RemoveTopicCache() error {
	var keys = t.redis.Keys(common.TopicPageKey + "*").Val()
	keys = append(keys, common.TagMapKey)
	config.LOGGER.Info("已清空专题缓存")
	return t.redis.Del(keys...).Err()
}

// GetTopicPage 通过页数获取某页的专题
func (t TopicCache) GetTopicPage(page int) (response.PageInfo, error) {

	var pageInfo response.PageInfo

	var err = t.redis.Get(common.TopicPageKey + strconv.Itoa(page)).Scan(&pageInfo)

	if err != nil {
		return pageInfo, err
	}

	return pageInfo, nil
}

// SetTopicToMap 将专题简要存入缓存map
func (t TopicCache) SetTopicToMap(topic response.SimpleTopicResponse) error {
	return t.redis.HSet(common.TopicMapKey, strconv.Itoa(topic.Id), utils.ObjectToJson(topic)).Err()
}

// GetTopicToMap 从缓存map中获取专题简要
func (t TopicCache) GetTopicToMap(topicId int) (topic response.SimpleTopicResponse, err error) {
	var val = t.redis.HGet(common.TopicMapKey, strconv.Itoa(topicId)).Val()

	if val == "" {
		return topic, errors.New("从缓存获取标签信息失败")
	}

	return utils.JsonToObject[response.SimpleTopicResponse](val), nil
}

func NewTopicCache(r *redis.Client) TopicCache {
	return TopicCache{redis: r}
}
