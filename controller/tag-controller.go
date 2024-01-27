package controller

import (
	"common-web-framework/common"
	"common-web-framework/helper"
	"common-web-framework/request"
	"common-web-framework/response"
	"common-web-framework/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TagController struct {
	service service.TagService
}

// GetRandomTags 获取随机标签
// @Summary 随机标签
// @Description 获取随机标签
// @Tags 标签相关接口
// @Produce json
// @Success 200 {object} common.R{data=response.TagResponse}
// @Router /tags/random [get]
func (t TagController) GetRandomTags(ctx *gin.Context) {
	var result = t.service.RandomTags()

	helper.ResultSuccessToResponse(result, ctx)
}

// GetAdminTagList 获取后台管理标签
// @Summary 后台管理标签
// @Description 获取后台管理标签信息
// @Tags 标签相关接口,后台管理接口
// @Param request query request.OtherAdminFilter true "过滤条件"
// @Security JWT
// @Produce json
// @Success 200 {object} common.R{data=response.PageInfo}
// @Router /tags/admin/list [get]
func (t TagController) GetAdminTagList(ctx *gin.Context) {

	var req request.OtherAdminFilter

	req.Page = 1

	ctx.ShouldBindQuery(&req)

	var result = t.service.GetAdminTag(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetTagBlogList 获取某个标签下的博客
// @Summary 标签博客列表
// @Description 获取某个标签下的博客
// @Tags 标签相关接口
// @Produce json
// @Param request query request.TagBlogRequest true "过滤条件"
// @Success 200 {object} common.R{data=response.PageInfo{data=[]response.TagResponse}}
// @Router /tags/blog [get]
func (t TagController) GetTagBlogList(ctx *gin.Context) {

	var req request.TagBlogRequest

	req.Page = 1

	ctx.ShouldBindQuery(&req)

	var result = t.service.GetTagBlogList(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetTagByIdInfo 获取某个标签的简要信息
// @Summary 标签简要信息
// @Description 获取某个标签的简要信息
// @Tags 标签相关接口
// @Produce json
// @Param id path int true "标签id"
// @Success 200 {object} common.R{data=response.TagResponse}
// @Router /tags/get/{id} [get]
func (t TagController) GetTagByIdInfo(ctx *gin.Context) {

	var idStr = ctx.Param("id")

	var id, err = strconv.Atoi(idStr)

	if err != nil {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	var result = t.service.GetTagInfo(id)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetAllTagInfo 获取所有标签
// @Summary 所有标签
// @Description 获取所有标签简要信息
// @Tags 标签相关接口
// @Produce json
// @Success 200 {object} common.R{data=[]response.TagResponse}
// @Router /tags/list [get]
func (t TagController) GetAllTagInfo(ctx *gin.Context) {
	helper.ResultSuccessToResponse(t.service.GetAllTag(), ctx)
}

// CreateTag 创建一个标签
// @Summary 创建标签
// @Description 创建标签
// @Tags 标签相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param name query string true "标签名字"
// @Success 200 {object} common.R
// @Router /tags/admin/save [post]
func (t TagController) CreateTag(ctx *gin.Context) {

	var name = ctx.Query("name")

	if name == "" {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	t.service.AddTag(name)

	helper.ResultSuccessToResponse(nil, ctx)
}

// UpdateTag 修改标签
// @Summary 修改标签
// @Description 修改标签
// @Tags 标签相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param req query response.TagResponse true "修改标签的模型"
// @Success 200 {object} common.R
// @Router /tags/admin/update [put]
func (t TagController) UpdateTag(ctx *gin.Context) {

	var req response.TagResponse

	ctx.ShouldBindJSON(&req)

	if req.Id == 0 || req.Name == "" {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	t.service.UpdateTag(req)

	helper.ResultSuccessToResponse(nil, ctx)
}

// DeleteTag 删除标签
// @Summary 删除标签
// @Description 删除标签
// @Tags 标签相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param ids body []int true "要删除标签的id"
// @Success 200 {object} common.R
// @Router /tags/admin/delete [put]
func (t TagController) DeleteTag(ctx *gin.Context) {

	var ids []int

	ctx.ShouldBindJSON(&ids)

	var count = t.service.DeleteTagByIds(ids)

	helper.ResultSuccessToResponse(count, ctx)
}

// UnDeleteTag 恢复删除标签
// @Summary 恢复删除标签
// @Description 恢复删除标签
// @Tags 标签相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param ids body []int true "要恢复删除标签的id"
// @Success 200 {object} common.R
// @Router /tags/admin/un_delete [put]
func (t TagController) UnDeleteTag(ctx *gin.Context) {

	var ids []int

	ctx.ShouldBindJSON(&ids)

	var count = t.service.UnDeleteTagByIds(ids)

	helper.ResultSuccessToResponse(count, ctx)
}

func NewTagController(service service.TagService) TagController {
	return TagController{service: service}
}
