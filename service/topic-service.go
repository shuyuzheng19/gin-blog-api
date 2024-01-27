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
	"go.uber.org/zap"
)

type TopicServiceImpl struct {
	repository repository.TopicRepository
	cache      cache.TopicCache
}

func (t TopicServiceImpl) GetAdminTopic(req request.OtherAdminFilter) response.PageInfo {
	var blogs, count = t.repository.GetAdminTopic(req)
	return response.PageInfo{
		Page:  req.Page,
		Count: count,
		Size:  common.AdminOtherCount,
		Data:  blogs,
	}
}

func (t TopicServiceImpl) DeleteTopicByIds(uid int, ids []int) int64 {
	var count = repository.DeleteData(uid, common.TableNames.TopicTableName, ids)

	if count > 0 {
		config.LOGGER.Info("删除专题",
			zap.Int64("success_count", count), zap.Ints("ids", ids))
		t.cache.RemoveTopicCache()
		repository.DeleteBlog(uid, true, nil, ids)
	}

	return count
}

func (t TopicServiceImpl) UnDeleteTopicByIds(uid int, ids []int) int64 {
	var count = repository.UnDeleteData(uid, common.TableNames.TopicTableName, ids)

	if count > 0 {
		config.LOGGER.Info("恢复删除专题",
			zap.Int64("success_count", count), zap.Ints("ids", ids))
		t.cache.RemoveTopicCache()
		repository.DeleteBlog(uid, false, nil, ids)
	}

	return count
}

func (t TopicServiceImpl) AddTopic(uid int, req request.TopicRequest) {

	config.ValidateError(req)

	var topic = models.Topic{
		Name:        req.Name,
		Description: req.Description,
		CoverImage:  req.CoverImage,
		UserID:      uid,
	}

	var count = t.repository.Create(topic)

	if count == 0 {
		helper.ErrorToResponse(common.SaveFail)
	} else {
		config.LOGGER.Info("创建专题", zap.String("name", topic.Name))
		t.cache.RemoveTopicCache()
	}
}

func (t TopicServiceImpl) UpdateTopic(req request.TopicRequest) {
	config.ValidateError(req)
	var topic = models.Topic{
		ID:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		CoverImage:  req.CoverImage,
	}

	var count = t.repository.Update(topic)

	if count == 0 {
		helper.ErrorToResponse(common.SaveFail)
	} else {
		config.LOGGER.Info("修改专题", zap.Int("id", topic.ID), zap.String("name", topic.Name))
		t.cache.RemoveTopicCache()
	}
}

func (t TopicServiceImpl) GetAllTopicBlogs(id int) []response.SimpleBlogResponse {
	return t.repository.FindAllTopicBlog(id)
}

func (t TopicServiceImpl) GetAllUserSimpleTopics(uid int) []response.SimpleTopicResponse {
	return t.repository.GetAllUserSimpleTopics(uid)
}

func (t TopicServiceImpl) GetAllUserTopics(uid int) []response.UserTopicResponse {
	return t.repository.FindUserTopics(uid)
}

func (t TopicServiceImpl) GetTopicBlogList(req request.TopicBlogRequest) response.PageInfo {
	var blogs, count = t.repository.FindBlogByTopicId(req)

	return response.PageInfo{
		Page:  req.Page,
		Count: count,
		Size:  common.BlogPageCount,
		Data:  blogs,
	}
}

func (t TopicServiceImpl) GetTopicInfo(id int) response.SimpleTopicResponse {
	var result, err = t.cache.GetTopicToMap(id)

	if err != nil {
		var tag = t.repository.FindById(id)
		t.cache.SetTopicToMap(*tag)
		config.LOGGER.Info("缓存标签信息到", zap.Any("tag_info", tag))
		result, _ = t.cache.GetTopicToMap(id)
	}

	return result
}

func (t TopicServiceImpl) GetTopicList(page int) response.PageInfo {
	var pageInfo, err = t.cache.GetTopicPage(page)

	if err != nil {
		var topics, count = t.repository.FindTopicByPage(page)
		pageInfo = response.PageInfo{
			Page:  page,
			Count: count,
			Size:  common.TopicPageCount,
			Data:  topics,
		}
		t.cache.SetTopicPage(page, pageInfo)
	}

	return pageInfo
}

func NewTopicService() TopicService {

	var TopicRepository = repository.NewTopicRepository(config.DB)

	var TopicCache = cache.NewTopicCache(config.REDIS)

	return TopicServiceImpl{repository: TopicRepository, cache: TopicCache}
}
