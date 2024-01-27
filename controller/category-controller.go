package controller

import (
	"common-web-framework/common"
	"common-web-framework/helper"
	"common-web-framework/request"
	"common-web-framework/response"
	"common-web-framework/service"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service service.CategoryService
}

// GetCategoryList 获取分类列表
// @Summary 分类列表
// @Description 获取分类列表
// @Tags 分类相关接口
// @Produce json
// @Success 200 {object} common.R{data=response.CategoryResponse}
// @Router /category/list [get]
func (c CategoryController) GetCategoryList(ctx *gin.Context) {
	var result = c.service.GetAllCategory()

	helper.ResultSuccessToResponse(result, ctx)
}

// GetAdminCategoryList 获取后台管理分类信息
// @Summary 后台管理分类信息
// @Description 获取后台管理分类信息
// @Tags 分类相关接口,后台管理接口
// @Param request query request.OtherAdminFilter true "过滤条件"
// @Security JWT
// @Produce json
// @Success 200 {object} common.R{data=response.PageInfo}
// @Router /category/admin/list [get]
func (c CategoryController) GetAdminCategoryList(ctx *gin.Context) {
	var req request.OtherAdminFilter

	req.Page = 1

	ctx.ShouldBindQuery(&req)

	var result = c.service.GetAdminCategory(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// CreateCategory 创建一个分类
// @Summary 创建分类
// @Description 创建分类
// @Tags 分类相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param name query string true "分类名字"
// @Success 200 {object} common.R
// @Router /category/admin/save [post]
func (c CategoryController) CreateCategory(ctx *gin.Context) {

	var name = ctx.Query("name")

	if name == "" {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	c.service.AddCategory(name)

	helper.ResultSuccessToResponse(nil, ctx)
}

// UpdateCategory 修改分类
// @Summary 修改分类
// @Description 修改分类
// @Tags 分类相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param req query response.CategoryResponse true "修改分类的模型"
// @Success 200 {object} common.R
// @Router /category/admin/update [put]
func (c CategoryController) UpdateCategory(ctx *gin.Context) {

	var req response.CategoryResponse

	ctx.ShouldBindJSON(&req)

	if req.Id == 0 || req.Name == "" {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	c.service.UpdateCategory(req)

	helper.ResultSuccessToResponse(nil, ctx)
}

// DeleteCategory 删除分类
// @Summary 删除分类
// @Description 删除分类
// @Tags 分类相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param ids body []int true "要删除分类的id"
// @Success 200 {object} common.R
// @Router /category/admin/delete [put]
func (c CategoryController) DeleteCategory(ctx *gin.Context) {

	var ids []int

	ctx.ShouldBindJSON(&ids)

	var count = c.service.DeleteCategoryByIds(ids)

	helper.ResultSuccessToResponse(count, ctx)
}

// UnDeleteCategory 恢复删除分类
// @Summary 恢复删除分类
// @Description 恢复删除分类
// @Tags 分类相关接口,后台管理接口
// @Security JWT
// @Produce json
// @Param ids body []int true "要恢复删除分类的id"
// @Success 200 {object} common.R
// @Router /category/admin/un_delete [put]
func (c CategoryController) UnDeleteCategory(ctx *gin.Context) {

	var ids []int

	ctx.ShouldBindJSON(&ids)

	var count = c.service.UnDeleteCategoryByIds(ids)

	helper.ResultSuccessToResponse(count, ctx)
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return CategoryController{service: service}
}
