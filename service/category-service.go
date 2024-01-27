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

type CategoryServiceImpl struct {
	repository repository.CategoryRepository
	cache      cache.CategoryCache
}

func (c CategoryServiceImpl) AddCategory(name string) {
	var category = models.Category{
		Name: name,
	}

	var count = c.repository.Create(category)

	if count == 0 {
		helper.ErrorToResponse(common.SaveFail)
	} else {
		config.LOGGER.Info("创建分类", zap.String("name", name))
		c.cache.RemoveCategoryCache()
	}
}

func (c CategoryServiceImpl) UpdateCategory(category response.CategoryResponse) {
	var categoryModel = models.Category{
		ID:   category.Id,
		Name: category.Name,
	}

	var count = c.repository.Update(categoryModel)

	if count == 0 {
		helper.ErrorToResponse(common.UpdateFail)
	} else {
		config.LOGGER.Info("创建分类",
			zap.Int("id", categoryModel.ID),
			zap.String("name", categoryModel.Name))
		c.cache.RemoveCategoryCache()
	}
}

func (c CategoryServiceImpl) DeleteCategoryByIds(ids []int) int64 {
	var count = repository.DeleteData(-1, common.TableNames.CategoryTableName, ids)

	if count > 0 {
		config.LOGGER.Info("删除分类",
			zap.Int64("success_count", count), zap.Ints("ids", ids))
		c.cache.RemoveCategoryCache()
		repository.DeleteBlog(-1, true, ids, nil)
	}

	return count
}

func (c CategoryServiceImpl) UnDeleteCategoryByIds(ids []int) int64 {
	var count = repository.UnDeleteData(-1, common.TableNames.CategoryTableName, ids)

	if count > 0 {
		config.LOGGER.Info("恢复删除分类",
			zap.Int64("success_count", count), zap.Ints("ids", ids))
		c.cache.RemoveCategoryCache()
		repository.DeleteBlog(-1, false, ids, nil)
	}

	return count
}

func (c CategoryServiceImpl) GetAdminCategory(req request.OtherAdminFilter) response.PageInfo {
	var blogs, count = c.repository.GetAdminCategories(req)
	return response.PageInfo{
		Page:  req.Page,
		Count: count,
		Size:  common.AdminOtherCount,
		Data:  blogs,
	}
}

func (c CategoryServiceImpl) GetAllCategory() []response.CategoryResponse {
	var categories, err = c.cache.GetCategoryList()

	if err != nil {
		var tagList = c.repository.FindAll()
		c.cache.SetCategoryList(tagList)
		config.LOGGER.Info("添加分类列表缓存")
		categories, _ = c.cache.GetCategoryList()
	}

	return categories
}

func NewCategoryService() CategoryService {

	var CategoryRepository = repository.NewCategoryRepository(config.DB)

	var CategoryCache = cache.NewCategoryCache(config.REDIS)

	return CategoryServiceImpl{repository: CategoryRepository, cache: CategoryCache}
}
