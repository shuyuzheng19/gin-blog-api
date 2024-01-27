package controller

import (
	"common-web-framework/common"
	"common-web-framework/helper"
	"common-web-framework/request"
	"common-web-framework/service"
	"common-web-framework/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type BlogController struct {
	service service.BlogService
}

// GetCategoryBlogList 分类博客列表
// @Summary 分类博客列表
// @Description 分类博客列表
// @Tags 博客相关接口
// @Param request query request.BlogListRequest true "博客分类过滤条件"
// @Produce json
// @Success 200 {object} common.R{data=response.PageInfo{data=[]response.BlogResponse}}
// @Router /blog/list [get]
func (b BlogController) GetCategoryBlogList(ctx *gin.Context) {
	var req request.BlogListRequest

	req.Page = 1

	ctx.ShouldBindQuery(&req)

	var result = b.service.FindBlogByCategory(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetLatestBlog 获取最新博客
// @Summary 最新博客
// @Description 前10条最新博客
// @Tags 博客相关接口
// @Produce json
// @Success 200 {object} common.R{data=response.SimpleBlogResponse}
// @Router /blog/latest [get]
func (b BlogController) GetLatestBlog(ctx *gin.Context) {
	var result = b.service.GetLatestBlogs()

	helper.ResultSuccessToResponse(result, ctx)
}

// GetHotBlog 获取热门博客
// @Summary 热门博客
// @Description 前10条热门博客
// @Tags 博客相关接口
// @Produce json
// @Success 200 {object} common.R{data=response.SimpleBlogResponse}
// @Router /blog/hots [get]
func (b BlogController) GetHotBlog(ctx *gin.Context) {
	var result = b.service.GetHostBlogs()

	helper.ResultSuccessToResponse(result, ctx)
}

// GetRangBlog 获取某个日期区间的博客
// @Summary 归档博客
// @Description 获取归档博客
// @Tags 博客相关接口
// @Produce json
// @Param request query request.RangBlogRequest true "博客归档过滤条件"
// @Success 200 {object} common.R{data=response.SimpleBlogResponse}
// @Router /blog/range [get]
func (b BlogController) GetRangBlog(ctx *gin.Context) {

	var req request.RangBlogRequest

	req.Page = 1

	ctx.ShouldBindQuery(&req)

	var result = b.service.GetArchiveBlogList(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetRecommendBlog 获取推荐博客
// @Summary 推荐博客
// @Description 获取推荐博客
// @Tags 博客相关接口
// @Produce json
// @Success 200 {object} common.R{data=response.RecommendBlogResponse}
// @Router /blog/recommend [get]
func (b BlogController) GetRecommendBlog(ctx *gin.Context) {
	var result = b.service.GetRecommend()
	helper.ResultSuccessToResponse(result, ctx)
}

// SetRecommendBlog 更新推荐博客
// @Summary 更新推荐博客
// @Description 更新推荐博客
// @Tags 博客相关接口,后台管理接口
// @Param ids body []int true "要添加的推荐博客id,只能添加4个"
// @Security JWT
// @Produce json
// @Success 200 {object} common.R
// @Router /blog/admin/recommend [post]
func (b BlogController) SetRecommendBlog(ctx *gin.Context) {

	var ids []int

	var err = ctx.ShouldBindJSON(&ids)

	if err != nil {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	b.service.SaveRecommend(ids)

	helper.ResultSuccessToResponse(nil, ctx)
}

// GetBlogById 获取博客详情
// @Summary 博客详情
// @Description 博客详情
// @Tags 博客相关接口
// @Param id path int true "博客的id"
// @Produce json
// @Success 200 {object} common.R
// @Router /blog/get/{id} [get]
func (b BlogController) GetBlogById(ctx *gin.Context) {

	var idQuery = ctx.Param("id")

	if id, err := strconv.ParseInt(idQuery, 10, 64); err == nil {
		helper.ResultSuccessToResponse(b.service.GetBlog(id), ctx)
	} else {
		helper.ResultFailToToResponse(common.NotFoundBlog, ctx)
	}
}

// SearchBlog 搜索博客
// @Summary 搜索博客
// @Description 搜索博客
// @Tags 博客相关接口
// @Produce json
// @Param request query request.SearchBlogRequest true "博客归档过滤条件"
// @Success 200 {object} common.R{data=response.PageInfo{data=[]response.ArchiveBlogResponse}}
// @Router /blog/search [get]
func (b BlogController) SearchBlog(ctx *gin.Context) {

	var req request.SearchBlogRequest

	req.Page = 1

	ctx.ShouldBindQuery(&req)

	var result = b.service.SearchBlog(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetAdminBlog 后台管理获取用户博客列表
// @Summary 用户博客列表
// @Description 获取用户博客列表
// @Security JWT
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Param request query request.AdminBlogFilterRequest true "博客过滤条件"
// @Success 200 {object} common.R{data=response.PageInfo{data=[]response.AdminBlogResponse}}
// @Router /blog/admin/list [get]
func (b BlogController) GetAdminBlog(ctx *gin.Context) {

	var req = request.AdminBlogFilterRequest{Page: 1}

	ctx.ShouldBindQuery(&req)

	var uid = utils.GetUserInfo(ctx).ID

	var result = b.service.GetAdminBlogs(uid, req)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetAdminDeleteBlog 后台管理获取已删除的博客列表
// @Summary 用户删除博客列表
// @Description 获取用户已删除的博客列表
// @Security JWT
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Param request query request.AdminBlogFilterRequest true "博客过滤条件"
// @Success 200 {object} common.R{data=response.PageInfo{data=[]response.AdminBlogResponse}}
// @Router /blog/admin/delete_list [get]
func (b BlogController) GetAdminDeleteBlog(ctx *gin.Context) {

	var req = request.AdminBlogFilterRequest{Page: 1}

	ctx.ShouldBindQuery(&req)

	var uid = utils.GetUserInfo(ctx).ID

	var result = b.service.GetAdminDeleteBlogs(uid, req)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetAllAdminBlog 后台管理获取所有的博客列表
// @Summary 所有的博客列表
// @Description 获取所有的博客列表
// @Security JWT
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Param request query request.AdminBlogFilterRequest true "博客过滤条件"
// @Success 200 {object} common.R{data=response.PageInfo{data=[]response.AdminBlogResponse}}
// @Router /blog/admin/all/list [get]
func (b BlogController) GetAllAdminBlog(ctx *gin.Context) {

	var req = request.AdminBlogFilterRequest{Page: 1}

	ctx.ShouldBindQuery(&req)

	var result = b.service.GetAllAdminBlogs(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// DeleteBlogs 通过id删除博客
// @Summary 删除博客
// @Description 通过博客id删除博客
// @Security JWT
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Param ids body []int true "要删除的博客id"
// @Success 200 {object} common.R "删除成功条数"
// @Router /blog/admin/deletes [put]
func (b BlogController) DeleteBlogs(ctx *gin.Context) {

	var ids []int

	ctx.ShouldBindJSON(&ids)

	var uid = GetUserIdIfSuper(utils.GetUserInfo(ctx))

	var result = b.service.DeleteBlogByIds(uid, ids)

	helper.ResultSuccessToResponse(result, ctx)
}

// UnDeleteBlogs 通过id恢复删除博客
// @Summary 恢复删除的博客
// @Description 通过博客id恢复删除的博客
// @Security JWT
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Param ids body []int true "要删除的博客id列表"
// @Success 200 {object} common.R "恢复成功条数"
// @Router /blog/admin/un_deletes [put]
func (b BlogController) UnDeleteBlogs(ctx *gin.Context) {

	var ids []int

	ctx.ShouldBindJSON(&ids)

	var uid = GetUserIdIfSuper(utils.GetUserInfo(ctx))

	var result = b.service.UnDeleteBlogByIds(uid, ids)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetAllDeleteAdminBlog 后台管理获取所有已删除的博客列表
// @Summary 所有已删除的博客列表
// @Description 获取所有已删除博客列表
// @Security JWT
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Param request query request.AdminBlogFilterRequest true "博客过滤条件"
// @Success 200 {object} common.R{data=response.PageInfo{data=[]response.AdminBlogResponse}}
// @Router /blog/admin/all/delete_list [get]
func (b BlogController) GetAllDeleteAdminBlog(ctx *gin.Context) {

	var req = request.AdminBlogFilterRequest{Page: 1}

	ctx.ShouldBindQuery(&req)

	var result = b.service.GetAllAdminDeleteBlogs(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetUserBlogList 获取用户博客
// @Summary 用户博客
// @Description 获取用户博客
// @Tags 博客相关接口
// @Produce json
// @Param request query request.UserBlogRequest true "用户博客过滤条件"
// @Success 200 {object} common.R{data=response.PageInfo{data=[]response.BlogResponse}}
// @Router /blog/user [get]
func (b BlogController) GetUserBlogList(ctx *gin.Context) {

	var req request.UserBlogRequest

	req.Page = 1

	ctx.ShouldBindQuery(&req)

	var result = b.service.GetUserBlog(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetUserBlogTopList 获取用户博客榜单
// @Summary 用户博客榜单
// @Description 获取用户博客榜单
// @Tags 博客相关接口
// @Produce json
// @Param id query int true "用户ID"
// @Success 200 {object} common.R{data=[]response.SimpleBlogResponse}
// @Router /blog/user_top [get]
func (b BlogController) GetUserBlogTopList(ctx *gin.Context) {

	var idStr = ctx.Query("id")

	if id, err := strconv.Atoi(idStr); err == nil {
		helper.ResultSuccessToResponse(b.service.GetUserTopBlog(id), ctx)
	} else {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
	}

}

// SaveUserEditor 保存用户编写的博客内容
// @Summary 用户保存草稿
// @Description 用户保存的草稿
// @Security JWT
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Param id query int true "用户ID"
// @Success 200 {object} common.R
// @Router /blog/admin/user_edit [post]
func (b BlogController) SaveUserEditor(ctx *gin.Context) {

	var maps map[string]string

	ctx.ShouldBindJSON(&maps)

	var content = maps["content"]

	if content == "" {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	var uid = utils.GetUserInfo(ctx).ID

	b.service.SaveEditBlog(uid, content)

	helper.ResultSuccessToResponse(nil, ctx)
}

// GetSaveUserEditor 获取用户编写的博客内容
// @Summary 获取用户保存草稿
// @Description 获取用户保存的草稿
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Security JWT
// @Param id query int true "用户ID"
// @Success 200 {object} common.R "用户草稿"
// @Router /blog/admin/user_edit [get]
func (b BlogController) GetSaveUserEditor(ctx *gin.Context) {

	var uid = utils.GetUserInfo(ctx).ID

	helper.ResultSuccessToResponse(b.service.GetSaveEditBlog(uid), ctx)
}

// SaveBlog 添加博客
// @Summary 添加博客
// @Description 添加博客
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Security JWT
// @Param request body request.BlogRequest true "添加博客的请求模型"
// @Success 200 {object} common.R
// @Router /blog/admin/save [post]
func (b BlogController) SaveBlog(ctx *gin.Context) {
	var req request.BlogRequest

	ctx.ShouldBindJSON(&req)

	var uid = utils.GetUserInfo(ctx).ID

	b.service.SaveBlog(uid, req)

	helper.ResultSuccessToResponse(nil, ctx)
}

// UpdateBlog 修改博客
// @Summary 修改博客
// @Description 修改博客
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Security JWT
// @Param id path int true "要修改的博客id"
// @Param request body request.BlogRequest true "修改博客的请求模型"
// @Success 200 {object} common.R
// @Router /blog/admin/update/{id} [put]
func (b BlogController) UpdateBlog(ctx *gin.Context) {
	var req request.BlogRequest

	ctx.ShouldBindJSON(&req)

	var idParam = ctx.Param("id")

	if id, err := strconv.ParseInt(idParam, 10, 64); err == nil {
		var uid = GetUserIdIfSuper(utils.GetUserInfo(ctx))
		b.service.UpdateBlog(id, uid, req)
		helper.ResultSuccessToResponse(nil, ctx)
	} else {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
	}

}

// InitSearch 初始化搜索
// @Summary 初始化搜索
// @Description 初始化搜索
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Security JWT
// @Success 200 {object} common.R
// @Router /blog/admin/init_search [get]
func (b BlogController) InitSearch(ctx *gin.Context) {
	b.service.InitSearch()
	helper.ResultSuccessToResponse(nil, ctx)
}

// InitEyeCount 初始化浏览量
// @Summary 初始化浏览量
// @Description 初始化浏览量
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Security JWT
// @Success 200 {object} common.R
// @Router /blog/admin/init_view [get]
func (b BlogController) InitEyeCount(ctx *gin.Context) {
	b.service.InitEyeCount()
	helper.ResultSuccessToResponse(nil, ctx)
}

// GetUpdateBlog 获取要修改的博客
// @Summary 获取修改博客信息
// @Description 获取修改博客信息
// @Tags 博客相关接口,后台管理接口
// @Produce json
// @Security JWT
// @Param id path int true "要修改的博客id"
// @Success 200 {object} common.R
// @Router /blog/admin/update/{id} [get]
func (b BlogController) GetUpdateBlog(ctx *gin.Context) {
	var idParam = ctx.Param("id")

	var uid = GetUserIdIfSuper(utils.GetUserInfo(ctx))

	if id, err := strconv.ParseInt(idParam, 10, 64); err == nil {
		var blog = b.service.GetUpdateBlog(id)

		if blog.Id == 0 {
			helper.ResultFailToToResponse(common.NotFoundBlog, ctx)
			return
		}

		if uid > 0 && blog.User.Id != uid {
			helper.ResultFailToToResponse(common.GetUpdateFail, ctx)
			return
		}
		helper.ResultSuccessToResponse(blog, ctx)
	} else {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
	}

}

// SimilarBlog 相关博客
// @Summary 搜索相关博客
// @Description 获取搜索的相关博客
// @Tags 博客相关接口
// @Produce json
// @Param keyword query string true "博客关键字"
// @Success 200 {object} common.R
// @Router /blog/similar [get]
func (b BlogController) SimilarBlog(ctx *gin.Context) {

	var keyword = ctx.Query("keyword")

	if keyword == "" {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	helper.ResultSuccessToResponse(b.service.SimilarBlog(keyword), ctx)
}

func NewBlogController(service service.BlogService) BlogController {
	return BlogController{service: service}
}
