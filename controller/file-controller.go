package controller

import (
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/helper"
	"common-web-framework/middle"
	"common-web-framework/request"
	"common-web-framework/service"
	"common-web-framework/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"path/filepath"
	"strconv"
)

type FileController struct {
	service service.FileService
}

// UploadImageFile 上传图片
// @Summary 上传图片
// @Security JWT
// @Description 上传图片
// @Tags 文件相关接口
// @Accept multipart/form-data
// @Param files formData file true "要上传的图片文件"
// @Produce json
// @Success 200 {object} common.R{data=response.SimpleFileResponse}
// @Router /file/upload/image [post]
func (f FileController) UploadImageFile(ctx *gin.Context) {
	var result = f.service.UploadImageFile(ctx)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetSystemFile 获取本地文件
func (f FileController) GetSystemFile(ctx *gin.Context) {

	var req request.SystemFileRequest

	ctx.ShouldBindJSON(&req)

	var result = f.service.GetSystemFile(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// DeleteSystemFile 删除本地文件
func (f FileController) DeleteSystemFile(ctx *gin.Context) {

	var paths []string

	ctx.ShouldBindJSON(&paths)

	if len(paths) == 0 {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	var count = f.service.DeleteSystemFile(paths)

	helper.ResultSuccessToResponse(count, ctx)
}

func getSystemFilePath(ctx *gin.Context) string {
	var token = ctx.Query("token")

	var user = middle.ParseToken(token, ctx)

	if ctx.IsAborted() {
		return ""
	}

	if user.Role.ID != common.SuperAdminRole {
		helper.ResultFailToToResponse(common.Forbidden, ctx)
		return ""
	}

	var path = ctx.Query("path")

	if path == "" {
		helper.ResultFailToToResponse(common.PathEmptyFail, ctx)
		return ""
	}

	return path
}

// DownloadSystemFile 下载本地文件
func (f FileController) DownloadSystemFile(ctx *gin.Context) {

	var path = getSystemFilePath(ctx)

	if path != "" {
		var name = filepath.Base(path)

		ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))

		ctx.Header("Content-Type", "application/octet-stream")

		ctx.Header("Content-Transfer-Encoding", "binary")

		ctx.File(path)
	}

}

// ClearSystemFileContent 清空文件内容
func (f FileController) ClearSystemFileContent(ctx *gin.Context) {

	var path = ctx.Query("path")

	if path != "" {
		if err := f.service.ClearFileContent(path); err != nil {
			helper.ResultFailToToResponse(common.ClearFileFail, ctx)
		} else {
			helper.ResultSuccessToResponse(nil, ctx)
		}
	}
}

// GetLogFileList 获取日志文件列表
func (f FileController) GetLogFileList(ctx *gin.Context) {
	var keyword = ctx.Query("keyword")

	var result = f.service.GetSystemFile(request.SystemFileRequest{
		Path:    config.CONFIG.Logger.LoggerDir,
		Keyword: keyword,
	})

	helper.ResultSuccessToResponse(result, ctx)
}

// GetCurrentLog 获取最新日志
func (f FileController) GetCurrentLog(ctx *gin.Context) {

	var logger = config.CONFIG.Logger

	var path = filepath.Join(logger.LoggerDir, logger.DefaultName)

	var name = filepath.Base(path)

	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name))

	ctx.Header("Content-Type", "application/octet-stream")

	ctx.Header("Content-Transfer-Encoding", "binary")

	ctx.File(path)
}

// GetAdminFile 后台管理文件列表
// @Summary 后台管理文件列表
// @Security JWT
// @Description 后台管理文件列表
// @Tags 文件相关接口,后台管理接口
// @Param req query request.OtherAdminFilter true "过滤文件参数"
// @Produce json
// @Success 200 {object} common.R{data=response.PageInfo{data=response.FileAdminResponse}}
// @Router /file/admin/list [get]
func (f FileController) GetAdminFile(ctx *gin.Context) {

	var req request.OtherAdminFilter

	req.Page = 1

	ctx.ShouldBindQuery(&req)

	var uid = GetUserIdIfSuper(utils.GetUserInfo(ctx))

	var result = f.service.GetAdminFileList(uid, req)

	helper.ResultSuccessToResponse(result, ctx)
}

// UpdateFilePublic 设置文件是否公开
// @Summary 设置文件是否公开
// @Security JWT
// @Description 设置文件是否公开
// @Tags 文件相关接口,后台管理接口
// @Param is_pub query bool true "true公开 false不公开"
// @Param id query int true "修改的文件id"
// @Produce json
// @Success 200 {object} common.R
// @Router /file/admin/public [put]
func (f FileController) UpdateFilePublic(ctx *gin.Context) {

	var idStr = ctx.Query("id")

	var id, err = strconv.Atoi(idStr)

	if err != nil {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	var pubStr = ctx.DefaultQuery("is_pub", "false")

	var pub, _ = strconv.ParseBool(pubStr)

	var uid = GetUserIdIfSuper(utils.GetUserInfo(ctx))

	f.service.UpdatePublic(uid, id, pub)

	helper.ResultSuccessToResponse(nil, ctx)
}

// DeleteFile 删除文件
// @Summary 删除文件
// @Security JWT
// @Description 删除文件
// @Tags 文件相关接口,后台管理接口
// @Param ids query []int true "要删除的文件id"
// @Produce json
// @Success 200 {object} common.R
// @Router /file/admin/delete [put]
func (f FileController) DeleteFile(ctx *gin.Context) {

	var ids []int

	ctx.ShouldBindJSON(&ids)

	var uid = GetUserIdIfSuper(utils.GetUserInfo(ctx))

	var count = f.service.DeleteFile(uid, ids)

	helper.ResultSuccessToResponse(count, ctx)
}

// UploadAvatarFile 上传头像
// @Summary 上传头像
// @Description 上传头像
// @Tags 文件相关接口
// @Accept multipart/form-data
// @Param files formData file true "要上传的头像图片文件"
// @Produce json
// @Success 200 {object} common.R{data=response.SimpleFileResponse}
// @Router /file/upload/avatar [post]
func (f FileController) UploadAvatarFile(ctx *gin.Context) {
	var result = f.service.UploadAvatarFile(ctx)

	helper.ResultSuccessToResponse(result, ctx)
}

// UploadOtherFile 上传文件
// @Summary 上传文件
// @Description 上传文件
// @Tags 文件相关接口
// @Security JWT
// @Accept multipart/form-data
// @Param files formData file true "要上传的文件"
// @Produce json
// @Success 200 {object} common.R{data=response.SimpleFileResponse}
// @Router /file/upload/other [post]
func (f FileController) UploadOtherFile(ctx *gin.Context) {
	var result = f.service.UploadFile(ctx)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetPublicFiles 获取公开上传的文件
// @Summary 公开文件列表
// @Description 获取公开上传的文件
// @Tags 文件相关接口
// @Param request query request.FileRequest true "过滤条件"
// @Produce json
// @Success 200 {object} common.R{data=response.PageInfo{data=response.FileResponse}}
// @Router /file/public [get]
func (f FileController) GetPublicFiles(ctx *gin.Context) {

	var req = request.FileRequest{Page: 1, Sort: "date"}

	ctx.ShouldBindQuery(&req)

	var result = f.service.GetPublicFile(req)

	helper.ResultSuccessToResponse(result, ctx)
}

// GetCurrentUserFiles 获取当前用户的的文件
// @Summary 当前用户文件列表
// @Description 获取当前用户的的文件
// @Tags 文件相关接口
// @Security JWT
// @Param request query request.FileRequest true "过滤条件"
// @Produce json
// @Success 200 {object} common.R{data=response.PageInfo{data=response.FileResponse}}
// @Router /file/current [get]
func (f FileController) GetCurrentUserFiles(ctx *gin.Context) {
	var req = request.FileRequest{Page: 1, Sort: "date"}

	ctx.ShouldBindQuery(&req)

	var uid = utils.GetUserInfo(ctx).ID

	var result = f.service.GetUserFile(uid, req)

	helper.ResultSuccessToResponse(result, ctx)
}

func NewFileController(service service.FileService) FileController {
	return FileController{service: service}
}
