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

type TopicController struct {
	service service.TopicService
}

// GetTopicList 获取某页专题
// @Summary 专题列表
// @Description 获取某页专题列表
// @Tags 专题相关接口
// @Produce json
// @Param page query int true "专题的第几页"
// @Success 200 {object} common.R{data=response.PageInfo{data=[]response.TopicResponse}}
// @Router /topics/list [get]
func (t TopicController) GetTopicList(ctx *gin.Context) {

	var pageQuery = ctx.DefaultQuery("page", "1")

	var page, err = strconv.Atoi(pageQuery)

	if err != nil {
		page = 1
	}

	var result = t.service.GetTopicList(page)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetAllTopicBlogs 获取专题下的所有博客
// @Summary 专题所有博客
// @Description 专题所有博客
// @Tags 专题相关接口
// @Produce json
// @Param id path int true "专题的id"
// @Success 200 {object} common.R{data=[]response.SimpleBlogResponse}
// @Router /topics/{id}/blog [get]
func (t TopicController) GetAllTopicBlogs(ctx *gin.Context) {
	var idStr = ctx.Param("id")
	if id, err := strconv.Atoi(idStr); err == nil {
		var result = t.service.GetAllTopicBlogs(id)
		helper.ResultSuccessToResponse(result, ctx)
	} else {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
	}
}

// GetAdminUserTopicList 获取用户的专题
// @Summary 用户专题列表
// @Security JWT
// @Description 获取用户的简洁的专题概要
// @Tags 专题相关接口,后台管理接口
// @Produce json
// @Success 200 {object} common.R{data=[]response.SimpleTopicResponse}
// @Router /topics/admin/list [get]
func (t TopicController) GetAdminUserTopicList(ctx *gin.Context) {

	var uid = GetUserIdIfSuper(utils.GetUserInfo(ctx))

	var result = t.service.GetAllUserSimpleTopics(uid)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetAllUserTopics 获取用户所有专题
// @Summary 用户专题
// @Description 获取该用户的所有专题
// @Tags 专题相关接口
// @Produce json
// @Param id path int true "用户的id"
// @Success 200 {object} common.R{data=[]response.UserTopicResponse}
// @Router /topics/user/{id} [get]
func (t TopicController) GetAllUserTopics(ctx *gin.Context) {

	var idParam = ctx.Param("id")

	var uid, _ = strconv.Atoi(idParam)

	var result = t.service.GetAllUserTopics(uid)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetTopicBlogList 获取某个专题下的博客
// @Summary 专题博客列表
// @Description 获取某个专题下的博客
// @Tags 专题相关接口
// @Produce json
// @Param request query request.TopicBlogRequest true "过滤专题博客"
// @Success 200 {object} common.R{data=response.PageInfo{data=[]response.TopicResponse}}
// @Router /topics/blog [get]
func (t TopicController) GetTopicBlogList(ctx *gin.Context) {

	var req request.TopicBlogRequest

	req.Page = 1

	ctx.ShouldBindQuery(&req)

	var result = t.service.GetTopicBlogList(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetTopicByIdInfo 获取某个专题的简要信息
// @Summary 专题简要信息
// @Description 获取某个专题的简要信息
// @Tags 专题相关接口
// @Produce json
// @Param id path int true "专题id"
// @Success 200 {object} common.R{data=response.SimpleTopicResponse}
// @Router /topics/get/{id} [get]
func (t TopicController) GetTopicByIdInfo(ctx *gin.Context) {

	var idStr = ctx.Param("id")

	var id, err = strconv.Atoi(idStr)

	if err != nil {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	var result = t.service.GetTopicInfo(id)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetAdminTopicList 获取后台管理专题信息
// @Summary 后台管理专题信息
// @Description 获取后台管理专题信息
// @Tags 专题相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param request query request.OtherAdminFilter true "过滤条件"
// @Success 200 {object} common.R{data=response.PageInfo}
// @Router /topics/admin/list [get]
func (t TopicController) GetAdminTopicList(ctx *gin.Context) {
	var req request.OtherAdminFilter

	req.Page = 1

	ctx.ShouldBindQuery(&req)

	var result = t.service.GetAdminTopic(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// CreateTopic 创建一个专题
// @Summary 创建专题
// @Description 创建专题
// @Tags 专题相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param req body request.TopicRequest true "专题模型"
// @Success 200 {object} common.R
// @Router /topics/admin/save [post]
func (t TopicController) CreateTopic(ctx *gin.Context) {

	var req request.TopicRequest

	ctx.ShouldBindJSON(&req)

	var uid = utils.GetUserInfo(ctx).ID

	t.service.AddTopic(uid, req)

	helper.ResultSuccessToResponse(nil, ctx)
}

// UpdateTopic 修改专题
// @Summary 修改专题
// @Description 修改专题
// @Tags 专题相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param req body request.TopicRequest true "修改专题的模型"
// @Success 200 {object} common.R
// @Router /topics/admin/update [put]
func (t TopicController) UpdateTopic(ctx *gin.Context) {

	var req request.TopicRequest

	ctx.ShouldBindJSON(&req)

	if req.Id == 0 {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	t.service.UpdateTopic(req)

	helper.ResultSuccessToResponse(nil, ctx)
}

// DeleteTopic 删除专题
// @Summary 删除专题
// @Description 删除专题
// @Tags 专题相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param ids body []int true "要删除专题的id"
// @Success 200 {object} common.R
// @Router /topics/admin/delete [put]
func (t TopicController) DeleteTopic(ctx *gin.Context) {

	var ids []int

	ctx.ShouldBindJSON(&ids)

	var uid = GetUserIdIfSuper(utils.GetUserInfo(ctx))

	var count = t.service.DeleteTopicByIds(uid, ids)

	helper.ResultSuccessToResponse(count, ctx)
}

// UnDeleteTopic 恢复删除专题
// @Summary 恢复删除专题
// @Description 恢复删除专题
// @Tags 专题相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param ids body []int true "要恢复删除专题的id"
// @Success 200 {object} common.R
// @Router /topics/admin/un_delete [put]
func (t TopicController) UnDeleteTopic(ctx *gin.Context) {

	var ids []int

	ctx.ShouldBindJSON(&ids)

	var uid = GetUserIdIfSuper(utils.GetUserInfo(ctx))

	var count = t.service.UnDeleteTopicByIds(uid, ids)

	helper.ResultSuccessToResponse(count, ctx)
}

func NewTopicController(service service.TopicService) TopicController {
	return TopicController{service: service}
}
