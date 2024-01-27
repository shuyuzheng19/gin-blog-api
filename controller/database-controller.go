package controller

import (
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/helper"
	"common-web-framework/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
)

type dataBaseTable string

const (
	BLOG     dataBaseTable = "BLOG"
	TAG      dataBaseTable = "TAG"
	BlogTag  dataBaseTable = "BLOG_TAG"
	FILE     dataBaseTable = "FILE"
	FileMd5  dataBaseTable = "FILE_MD5"
	Category dataBaseTable = "CATEGORY"
	Topic    dataBaseTable = "TOPICS"
	Role     dataBaseTable = "ROLE"
	USER     dataBaseTable = "USER"
)

func (d DataBaseController) ExecSQL(ctx *gin.Context) {

	var apiKey = ctx.Query("apiKey")

	if apiKey != config.CONFIG.DataBaseKey {
		ctx.AbortWithStatusJSON(200, common.Fail(403, "错误的apiKey"))
		return
	}

	var body = ctx.Request.Body

	var buff, err = ioutil.ReadAll(body)

	if err != nil {
		ctx.AbortWithStatusJSON(200, common.AutoFail(common.BadRequestCode))
		return
	}

	go func() {
		d.service.ExecDataBaseSQL(string(buff))
	}()

	helper.ResultSuccessToResponse(nil, ctx)
}

func handleInsertSQL(ctx *gin.Context, f func(page int) string, size int) {

	ctx.Header("Content-Type", "text/event-stream")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")

	var page = 0
	for {
		page++
		var sql = f(page)
		if sql == "" {
			ctx.Writer.CloseNotify()
			break
		} else {
			ctx.SSEvent("", sql)
		}
		if page*common.DataBaseSelectInsertCount > size {
			break
		}
	}
}

func (d DataBaseController) GetTableInsertSQL(ctx *gin.Context) {

	var apiKey = ctx.Query("apiKey")

	if apiKey != config.CONFIG.DataBaseKey {
		ctx.AbortWithStatusJSON(200, common.Fail(403, "错误的apiKey"))
		return
	}

	var _type = ctx.Query("TYPE")

	var size, err = strconv.Atoi(ctx.Query("size"))

	if err != nil {
		size = common.DataBaseSelectInsertCount
	}

	switch dataBaseTable(_type) {
	case BLOG:
		handleInsertSQL(ctx, d.service.GetBlogInsertSQL, size)
		break
	case TAG:
		handleInsertSQL(ctx, d.service.GetTagInsertSQL, size)
		break
	case BlogTag:
		handleInsertSQL(ctx, d.service.GetBlogTagInsertSQL, size)
		break
	case FILE:
		handleInsertSQL(ctx, d.service.GetFileInsertSQL, size)
		break
	case FileMd5:
		handleInsertSQL(ctx, d.service.GetFileMd5InsertSQL, size)
	case Category:
		handleInsertSQL(ctx, d.service.GetCategoryInsertSQL, size)
		break
	case Topic:
		handleInsertSQL(ctx, d.service.GetTopicInsertSQL, size)
		break
	case Role:
		handleInsertSQL(ctx, d.service.GetRoleInsertSQL, size)
		break
	case USER:
		handleInsertSQL(ctx, d.service.GetUserInsertSQL, size)
		break
	default:
		ctx.AbortWithStatusJSON(200, common.AutoFail(common.BadRequestCode))
	}

}

type DataBaseController struct {
	service service.DataBaseService
}

func NewDataBaseController(service service.DataBaseService) DataBaseController {
	return DataBaseController{service: service}
}
