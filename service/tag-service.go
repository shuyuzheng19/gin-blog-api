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

type TagServiceImpl struct {
	repository repository.TagRepository
	cache      cache.TagCache
}

func (t TagServiceImpl) AddTag(name string) {

	var tag = models.Tag{
		Name: name,
	}

	var count = t.repository.Create(tag)

	if count == 0 {
		helper.ErrorToResponse(common.SaveFail)
	} else {
		config.LOGGER.Info("添加标签", zap.String("name", tag.Name))
		t.cache.RemoveTagCache()
	}

}

func (t TagServiceImpl) UpdateTag(tag response.TagResponse) {

	var tagModel = models.Tag{
		ID:   tag.Id,
		Name: tag.Name,
	}

	var count = t.repository.Update(tagModel)

	if count == 0 {
		helper.ErrorToResponse(common.UpdateFail)
	} else {
		config.LOGGER.Info("修改标签", zap.Int("id", tagModel.ID),
			zap.String("name", tag.Name))
		t.cache.RemoveTagCache()
	}
}

func (t TagServiceImpl) DeleteTagByIds(ids []int) int64 {
	var count = repository.DeleteData(-1, common.TableNames.TagTableName, ids)

	if count > 0 {
		config.LOGGER.Info("删除标签",
			zap.Int64("success_count", count), zap.Ints("ids", ids))
		t.cache.RemoveTagCache()
	}

	return count
}

func (t TagServiceImpl) UnDeleteTagByIds(ids []int) int64 {
	var count = repository.UnDeleteData(-1, common.TableNames.TagTableName, ids)

	if count > 0 {
		config.LOGGER.Info("恢复删除标签",
			zap.Int64("success_count", count), zap.Ints("ids", ids))
		t.cache.RemoveTagCache()
	}

	return count
}

func (t TagServiceImpl) GetAdminTag(req request.OtherAdminFilter) response.PageInfo {
	var blogs, count = t.repository.GetAdminTags(req)
	return response.PageInfo{
		Page:  req.Page,
		Count: count,
		Size:  common.AdminOtherCount,
		Data:  blogs,
	}
}

func (t TagServiceImpl) GetAllTag() []response.TagResponse {
	return t.repository.FindAll()
}

func (t TagServiceImpl) GetTagBlogList(req request.TagBlogRequest) response.PageInfo {
	var blogs, count = t.repository.FindBlogByTagId(req)

	return response.PageInfo{
		Page:  req.Page,
		Count: count,
		Size:  common.BlogPageCount,
		Data:  blogs,
	}
}

func (t TagServiceImpl) GetTagInfo(id int) response.TagResponse {
	var result, err = t.cache.GetTagToMap(id)

	if err != nil {
		var tag = t.repository.FindById(id)
		t.cache.SetTagToMap(*tag)
		config.LOGGER.Info("缓存标签信息", zap.Any("tag_info", tag))
		result, _ = t.cache.GetTagToMap(id)
	}

	return result
}

func (t TagServiceImpl) RandomTags() []response.TagResponse {
	var tags, err = t.cache.GetRandomTags()

	if err != nil {
		var tagList = t.repository.FindAll()
		t.cache.SetTags(tagList)
		config.LOGGER.Info("添加随机标签缓存")
		tags, err = t.cache.GetRandomTags()
	}

	return tags
}

func NewTagService() TagService {

	var tagRepository = repository.NewTagRepository(config.DB)

	var tagCache = cache.NewTagCache(config.REDIS)

	return TagServiceImpl{repository: tagRepository, cache: tagCache}
}
