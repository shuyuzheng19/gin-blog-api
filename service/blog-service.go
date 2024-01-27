package service

import (
	"common-web-framework/cache"
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/helper"
	"common-web-framework/repository"
	"common-web-framework/request"
	"common-web-framework/response"
	"common-web-framework/search"
	"common-web-framework/utils"
	"go.uber.org/zap"
	"strconv"
)

type BlogServiceImpl struct {
	repository  repository.BlogRepository
	cache       cache.BlogCache
	search      *search.MeiliSearchClient
	searchIndex string
}

func (b BlogServiceImpl) InitEyeCount() {
	var maps = b.cache.GetAllBlogEyeCount()
	for id, count := range maps {
		var idNumber, _ = strconv.ParseInt(id, 10, 64)
		var countNumber, _ = strconv.ParseInt(count, 10, 64)
		b.UpdateBlogEyeCount(idNumber, countNumber)
	}
	config.LOGGER.Info("初始化浏览量")
	b.cache.DeleteBlogEyeCount()
}

func (b BlogServiceImpl) UpdateBlogEyeCount(id int64, count int64) int64 {
	return b.repository.UpdateEyeCount(id, count)
}

func (b BlogServiceImpl) InitSearch() {
	b.search.DeleteAllDocument(b.searchIndex)
	var blogs = b.repository.FindAllSearchBlog()
	var jsonStr = utils.ObjectToJson(blogs)
	b.search.SaveDocument(b.searchIndex, jsonStr)
	config.LOGGER.Info("初始化搜索")
}

func (b BlogServiceImpl) GetUpdateBlog(id int64) response.BlogContentResponse {
	return b.repository.FindById(id).ToContentBlogResponse()
}

func (b BlogServiceImpl) UpdateBlog(id int64, uid int, req request.BlogRequest) {
	b.saveOrUpdateBlog(id, uid, req)
}

func (b BlogServiceImpl) GetBlog(id int64) response.BlogContentResponse {

	var blog, err = b.cache.GetBlogInfo(id)

	if err != nil {
		blog = b.repository.FindById(id)

		b.cache.SetBlogInfo(blog)
	}

	if blog.User.ID == 0 {
		helper.ErrorToResponse(common.NotFoundBlog)
	}

	config.LOGGER.Info("获取博客", zap.Int64("id", id), zap.String("title", blog.Title))

	blog.EyeCount = b.cache.GetBlogEyeCount(blog.EyeCount, blog.ID)

	return blog.ToContentBlogResponse()
}

func (b BlogServiceImpl) saveOrUpdateBlog(id int64, uid int, req request.BlogRequest) {
	config.ValidateError(req)

	var do = req.ToBlogDo(uid)

	if id > 0 {
		do.ID = id
		var count = b.repository.Update(&do)
		if count == 0 {
			helper.ErrorToResponse(common.UpdateBlogFail)
		}
		config.LOGGER.Info("修改博客成功", zap.Int("user_id", uid),
			zap.String("title", do.Title))
	} else {
		var err = b.repository.Save(&do)

		if err != nil {
			helper.ErrorToResponse(common.SaveBlogFail)
		}
		config.LOGGER.Info("添加博客成功", zap.Int("user_id", uid),
			zap.String("title", do.Title))
	}

	b.cache.ClearBlogPageInfo()

	var searchBlog = response.SearchBlogResponse{
		Id:          do.ID,
		Title:       do.Title,
		Description: do.Description,
	}

	var str = utils.ObjectToJson(searchBlog)

	b.search.SaveDocument(config.CONFIG.Search.BlogIndex, str)
}

func (b BlogServiceImpl) SaveBlog(uid int, req request.BlogRequest) {
	b.saveOrUpdateBlog(0, uid, req)
}

func (b BlogServiceImpl) SaveEditBlog(uid int, content string) {
	b.cache.SetUserEdit(uid, content)
}

func (b BlogServiceImpl) GetSaveEditBlog(uid int) string {
	return b.cache.GetUserEdit(uid)
}

func (b BlogServiceImpl) SimilarBlog(keyword string) []any {
	var req = getBlogSearchRequest(request.SearchBlogRequest{Page: 1, Keyword: keyword})

	var response = b.search.SearchDocument(config.CONFIG.Search.BlogIndex, req)

	return response.Hits
}

func (b BlogServiceImpl) DeleteBlogByIds(uid int, ids []int) int64 {
	var count = repository.DeleteData(uid, common.TableNames.BlogTableName, ids)
	config.LOGGER.Info("删除博客", zap.Int("user_id", uid), zap.Ints("ids", ids))
	if count > 0 {
		b.cache.ClearBlogPageInfo()
	}
	return count
}

func (b BlogServiceImpl) UnDeleteBlogByIds(uid int, ids []int) int64 {
	var count = repository.UnDeleteData(uid, common.TableNames.BlogTableName, ids)
	config.LOGGER.Info("恢复删除的博客", zap.Int("user_id", uid), zap.Ints("ids", ids))
	if count > 0 {
		b.cache.ClearBlogPageInfo()
	}
	return count
}

func (b BlogServiceImpl) GetAllAdminBlogs(req request.AdminBlogFilterRequest) response.PageInfo {
	return b.getAdminBlogs(false, -1, req)
}

func (b BlogServiceImpl) GetAllAdminDeleteBlogs(req request.AdminBlogFilterRequest) response.PageInfo {
	return b.getAdminBlogs(false, -1, req)
}

func (b BlogServiceImpl) GetAdminBlogs(uid int, req request.AdminBlogFilterRequest) response.PageInfo {
	return b.getAdminBlogs(false, uid, req)
}

func (b BlogServiceImpl) getAdminBlogs(deleted bool, uid int, req request.AdminBlogFilterRequest) response.PageInfo {
	var blogs, count = b.repository.GetMoreBlogInfo(deleted, uid, req)

	return response.PageInfo{
		Page:  req.Page,
		Count: count,
		Size:  common.AdminBlogCount,
		Data:  blogs,
	}
}

func (b BlogServiceImpl) GetAdminDeleteBlogs(uid int, req request.AdminBlogFilterRequest) response.PageInfo {
	return b.getAdminBlogs(true, uid, req)
}

func (b BlogServiceImpl) GetUserBlog(req request.UserBlogRequest) response.PageInfo {
	var blogs, count = b.repository.FindByUserId(req)
	return response.PageInfo{
		Page:  req.Page,
		Count: count,
		Size:  common.BlogPageCount,
		Data:  blogs,
	}
}

func (b BlogServiceImpl) GetUserTopBlog(id int) []response.SimpleBlogResponse {
	return b.repository.FindUserTopBlog(id)
}

func getBlogSearchRequest(req request.SearchBlogRequest) search.MeiliSearchRequest {
	var searchCount = common.SearchBlogPageCount
	return search.NewSearchRequest().SetQ(req.Keyword).SetOffset((req.Page - 1) * searchCount).
		SetAttributesToHighlight([]string{"*"}).
		SetLimit(searchCount).
		SetShowMatchesPosition(false).
		SetHighlightPreTag("<b>").
		SetHighlightPostTag("</b>").
		Build()
}

func (b BlogServiceImpl) SearchBlog(req request.SearchBlogRequest) response.PageInfo {
	var result = b.search.SearchDocument(b.searchIndex, getBlogSearchRequest(req))

	config.LOGGER.Info("搜索博客", zap.Any("info", req))

	return response.PageInfo{
		Page:  req.Page,
		Count: result.EstimatedTotalHits,
		Size:  common.SearchBlogPageCount,
		Data:  result.Hits,
	}
}

func (b BlogServiceImpl) GetArchiveBlogList(req request.RangBlogRequest) response.PageInfo {

	var blogs, count = b.repository.RangeBlog(req)

	return response.PageInfo{
		Page:  req.Page,
		Count: count,
		Size:  common.ArchivePageCount,
		Data:  blogs,
	}
}

func (b BlogServiceImpl) SaveRecommend(ids []int) {
	if len(ids) != 4 {
		helper.ErrorToResponse(common.FailCode)
	}

	var blogs = b.repository.FindByIdIn(ids)

	b.cache.SetRecommend(blogs)

	config.LOGGER.Info("更新推荐博客")
}

func (b BlogServiceImpl) GetRecommend() []response.RecommendBlogResponse {
	return b.cache.GetRecommend()
}

func (b BlogServiceImpl) GetHostBlogs() []response.SimpleBlogResponse {
	if blogs, err := b.cache.GetHotBlog(); err != nil {
		var latestBlog = b.repository.GetHotBlog()
		b.cache.SetHotBlog(latestBlog)
		config.LOGGER.Info("添加热门博客缓存")
		return latestBlog
	} else {
		return blogs
	}
}

func (b BlogServiceImpl) GetLatestBlogs() []response.SimpleBlogResponse {
	if blogs, err := b.cache.GetLatestBlog(); err != nil {
		var latestBlog = b.repository.GetLatestBlog()
		b.cache.SetLatestBlog(latestBlog)
		config.LOGGER.Info("添加最新博客缓存")
		return latestBlog
	} else {
		return blogs
	}
}

func (b BlogServiceImpl) FindBlogByCategory(req request.BlogListRequest) response.PageInfo {

	var pageInfo response.PageInfo

	if info, err := b.cache.GetPageInfo(req); err != nil {
		var list, count = b.repository.FindCategoryBlogs(req)

		pageInfo = response.PageInfo{
			Page:  req.Page,
			Count: count,
			Size:  common.BlogPageCount,
			Data:  list,
		}

		b.cache.SetPageInfo(req, pageInfo)
		return pageInfo
	} else {
		return info
	}
}

func NewBlogService() BlogService {

	var repository = repository.NewBlogRepository(config.DB)

	var cache = cache.NewBlogCache(config.REDIS)

	return BlogServiceImpl{repository: repository, cache: cache, search: config.SEARCH, searchIndex: config.CONFIG.Search.BlogIndex}
}
