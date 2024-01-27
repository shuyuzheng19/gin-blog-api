package cache

import (
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/models"
	"common-web-framework/request"
	"common-web-framework/response"
	"common-web-framework/utils"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

type BlogCache struct {
	redis *redis.Client
}

// GetBlogEyeCount 获取自增的博客浏览量
func (b BlogCache) GetBlogEyeCount(defaultCount int64, id int64) int64 {
	var count = b.redis.HIncrBy(common.BlogEyeCountMapKey, strconv.FormatInt(id, 10), 1).Val()
	if count == 1 {
		return defaultCount + 1
	} else {
		return count
	}
}

// GetAllBlogEyeCount 获取所有的浏览量
func (b BlogCache) GetAllBlogEyeCount() map[string]string {
	var result = b.redis.HGetAll(common.BlogEyeCountMapKey).Val()
	return result
}

// DeleteBlogEyeCount 删除浏览量
func (b BlogCache) DeleteBlogEyeCount() error {
	return b.redis.Del(common.BlogEyeCountMapKey).Err()
}

// SetLatestBlog 缓存最新的10条博客
func (b BlogCache) SetLatestBlog(blogs []response.SimpleBlogResponse) error {
	return b.setSimpleBlog(common.LatestBlogKey, common.LatestBlogExpire, blogs)
}

// GetLatestBlog 从缓存获取最新博客
func (b BlogCache) GetLatestBlog() ([]response.SimpleBlogResponse, error) {
	return b.getSimpleBlog(common.LatestBlogKey)
}

// SetUserEdit 缓存用户写的博客
func (b BlogCache) SetUserEdit(uid int, content string) error {
	return b.redis.HSet(common.UserEditKey, strconv.Itoa(uid), content).Err()
}

// GetUserEdit 获取用户写的博客
func (b BlogCache) GetUserEdit(uid int) string {
	return b.redis.HGet(common.UserEditKey, strconv.Itoa(uid)).Val()
}

// SetHotBlog 缓存热门的10条博客
func (b BlogCache) SetHotBlog(blogs []response.SimpleBlogResponse) error {
	return b.setSimpleBlog(common.HotBlogKey, common.HotBlogExpire, blogs)
}

// GetHotBlog 从缓存获取热门博客
func (b BlogCache) GetHotBlog() ([]response.SimpleBlogResponse, error) {
	return b.getSimpleBlog(common.HotBlogKey)
}

// SetBlogInfo 缓存博客详情信息
func (b BlogCache) SetBlogInfo(blog models.Blog) error {
	var str = utils.ObjectToJson(blog)
	return b.redis.HSet(common.BlogMapKey, strconv.FormatInt(blog.ID, 10), str).Err()
}

// GetBlogInfo 获取博客详情信息
func (b BlogCache) GetBlogInfo(id int64) (models.Blog, error) {
	var r = b.redis.HGet(common.BlogMapKey, strconv.FormatInt(id, 10))
	if r.Err() != nil {
		return models.Blog{}, errors.New("获取博客信息失败")
	}
	return utils.JsonToObject[models.Blog](r.Val()), nil
}

// ClearBlogPageInfo 清空缓存的博客列表
func (b BlogCache) ClearBlogPageInfo() error {
	return ClearBlogPageInfo()
}

func ClearBlogPageInfo() error {
	var r = config.REDIS
	var keys = r.Keys(common.PageInfoPrefixKey + "*").Val()
	config.LOGGER.Info("已清空博客页面缓存列表")
	return r.Del(keys...).Err()
}

func DeleteBlogMap() error {
	return config.REDIS.Del(common.BlogMapKey).Err()
}

func (b BlogCache) setSimpleBlog(key string, expire time.Duration, blogs []response.SimpleBlogResponse) error {

	var blogsJson = utils.ObjectToJson(blogs)

	return b.redis.Set(key, blogsJson, expire).Err()
}

func (b BlogCache) getSimpleBlog(key string) ([]response.SimpleBlogResponse, error) {

	var str = b.redis.Get(key).Val()

	if str == "" {
		return nil, errors.New("获取失败")
	}

	return utils.JsonToObject[[]response.SimpleBlogResponse](str), nil
}

// SetPageInfo 将博客列表存入redis
func (b BlogCache) SetPageInfo(req request.BlogListRequest, pageInfo response.PageInfo) error {
	var key = fmt.Sprintf("%s:page-%d-cid-%d-sort:%s", common.PageInfoPrefixKey, req.Page, req.Cid, req.Sort)
	return b.redis.Set(key, &pageInfo, common.PageInfoExpire).Err()
}

// GetPageInfo 从redis获取博客列表
func (b BlogCache) GetPageInfo(req request.BlogListRequest) (pageInfo response.PageInfo, err error) {
	var key = fmt.Sprintf("%s:page-%d-cid-%d-sort:%s", common.PageInfoPrefixKey, req.Page, req.Cid, req.Sort)

	err = b.redis.Get(key).Scan(&pageInfo)

	if err != nil {
		return response.PageInfo{}, err
	}

	return pageInfo, nil
}

// SetRecommend 缓存推荐博客
func (b BlogCache) SetRecommend(blogs []response.RecommendBlogResponse) error {

	var blogsJson = utils.ObjectToJson(blogs)

	return b.redis.Set(common.RecommendKey, blogsJson, -1).Err()
}

// GetRecommend 从缓存获取推荐博客
func (b BlogCache) GetRecommend() []response.RecommendBlogResponse {

	var str = b.redis.Get(common.RecommendKey).Val()

	if str == "" {
		return make([]response.RecommendBlogResponse, 0)
	}

	return utils.JsonToObject[[]response.RecommendBlogResponse](str)
}

func NewBlogCache(r *redis.Client) BlogCache {
	return BlogCache{redis: r}
}
