package controller

import (
	"common-web-framework/common"
	"common-web-framework/helper"
	"common-web-framework/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type ToolsController struct {
	service service.ToolsService
}

func getBodyStr(ctx *gin.Context) []byte {
	var body = ctx.Request.Body

	defer body.Close()

	var buff, err = ioutil.ReadAll(body)

	if err != nil {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return nil
	}

	return buff
}

// JsonToYaml Json转Yaml
// @Summary Json转Yaml
// @Description Json转Yaml
// @Tags 工具相关接口
// @Param jsonData body string true "json字符串"
// @Produce json
// @Success 200 {object} common.R
// @Router /tools/json2yaml [post]
func (t ToolsController) JsonToYaml(ctx *gin.Context) {
	var buff = getBodyStr(ctx)

	if buff != nil {
		var result = t.service.JsonToYaml(buff)
		helper.ResultSuccessToResponse(result, ctx)
	}
}

// YamlToJson Yaml转JSON
// @Summary Yaml转JSON
// @Description Yaml转JSON
// @Tags 工具相关接口
// @Param jsonData body string true "yaml字符串"
// @Produce json
// @Success 200 {object} common.R
// @Router /tools/yaml2json [post]
func (t ToolsController) YamlToJson(ctx *gin.Context) {
	var buff = getBodyStr(ctx)

	if buff != nil {
		var result = t.service.YamlToJson(buff)
		helper.ResultSuccessToResponse(result, ctx)
	}
}

// FormatJson 格式化JSON
// @Summary 格式化JSON
// @Description 格式化JSON
// @Tags 工具相关接口
// @Param jsonData body string true "json字符串"
// @Produce json
// @Success 200 {object} common.R
// @Router /tools/format_json [post]
func (t ToolsController) FormatJson(ctx *gin.Context) {
	var buff = getBodyStr(ctx)

	if buff != nil {
		var result = t.service.FormatJson(buff)
		helper.ResultSuccessToResponse(result, ctx)
	}
}

func (t ToolsController) StringQuote(ctx *gin.Context) {
	var buff = getBodyStr(ctx)

	if buff != nil {
		var result = strconv.Quote(string(buff))
		helper.ResultSuccessToResponse(result, ctx)
	}
}

func (t ToolsController) UnStringQuote(ctx *gin.Context) {
	var buff = getBodyStr(ctx)

	if buff != nil {
		var result, err = strconv.Unquote(string(buff))
		if err != nil {
			helper.ResultFailToToResponse(common.ConvertFail, ctx)
		} else {
			helper.ResultSuccessToResponse(result, ctx)
		}
	}
}

// CompressJson 压缩json
// @Summary 压缩json
// @Description 压缩json
// @Tags 工具相关接口
// @Param jsonData body string true "json字符串"
// @Produce json
// @Success 200 {object} common.R
// @Router /tools/compress_json [post]
func (t ToolsController) CompressJson(ctx *gin.Context) {
	var buff = getBodyStr(ctx)

	if buff != nil {
		var result = t.service.CompressJson(buff)
		helper.ResultSuccessToResponse(result, ctx)
	}
}

// Json2Struct json转Go结构体
// @Summary json转Go结构体
// @Description json转Go结构体
// @Tags 工具相关接口
// @Param jsonData body string true "json字符串"
// @Produce json
// @Success 200 {object} common.R
// @Router /tools/json2struct [post]
func (t ToolsController) Json2Struct(ctx *gin.Context) {
	var buff = getBodyStr(ctx)

	if buff != nil {
		var result = t.service.JsonToStruct(buff)
		helper.ResultSuccessToResponse(result, ctx)
	}
}

// StructToJSON Go结构体转json
// @Summary Go结构体转json
// @Description Go结构体转json
// @Tags 工具相关接口
// @Param jsonData body string true "go结构体"
// @Produce json
// @Success 200 {object} common.R
// @Router /tools/struct2json [post]
func (t ToolsController) StructToJSON(ctx *gin.Context) {
	var buff = getBodyStr(ctx)

	if buff != nil {
		var result = t.service.StructToJSON(buff)
		helper.ResultSuccessToResponse(result, ctx)
	}
}

// Compress 压缩HTML|CSS|JS
// @Summary 压缩HTML|CSS|JS
// @Description 压缩HTML|CSS|JS
// @Tags 工具相关接口
// @Param type query string true "html|css|js"
// @Param data body string true "json字符串"
// @Produce json
// @Success 200 {object} common.R
// @Router /tools/compress [post]
func (t ToolsController) Compress(ctx *gin.Context) {
	var _type = ctx.Query("type")

	if _type == "" {
		helper.ResultFailToToResponse(common.ConvertFail, ctx)
		return
	}

	var body = getBodyStr(ctx)

	var result = t.service.Compress(_type, string(body))

	helper.ResultSuccessToResponse(result, ctx)
}

// CompressFromFile 压缩HTML|CSS|JS
// @Summary 上传压缩HTML|CSS|JS
// @Description 上传压缩HTML|CSS|JS
// @Tags 工具相关接口
// @Param file formData file true "上传的html|css|js"
// @Produce json
// @Success 200 {object} common.R
// @Router /tools/compress_file [post]
func (t ToolsController) CompressFromFile(ctx *gin.Context) {
	var file, err = ctx.FormFile("file")

	if err != nil {
		helper.ResultFailToToResponse(common.BadRequestCode, ctx)
		return
	}

	var _type = filepath.Ext(file.Filename)

	if _type == ".json" || _type == ".html" || _type == ".js" || _type == ".css" {
		var f, err2 = file.Open()

		if err2 != nil {
			helper.ResultFailToToResponse(common.OpenFileFail, ctx)
			return
		}

		var buff, _ = ioutil.ReadAll(f)

		var result = t.service.Compress(strings.Replace(_type, ".", "", 1), string(buff))

		helper.ResultSuccessToResponse(result, ctx)
	} else {
		ctx.JSON(200, common.Fail(21313, "只支持js,css,html,json"))
	}

}

// Base64EncodingOrDecoding base64编码解码
// @Summary base64编码解码
// @Description base64编码解码
// @Tags 工具相关接口
// @Param decoding query bool true "true解码 false编码"
// @Param data body string true "base64字符串"
// @Produce json
// @Success 200 {object} common.R
// @Router /tools/base64 [post]
func (t ToolsController) Base64EncodingOrDecoding(ctx *gin.Context) {
	var _type = ctx.Query("decoding")

	var flag, _ = strconv.ParseBool(_type)

	var body = getBodyStr(ctx)

	var result string

	if flag {
		result = t.service.Base64Decoding(body)
	} else {
		result = t.service.Base64Encoding(body)
	}

	helper.ResultSuccessToResponse(result, ctx)
}

func NewToolsController(service service.ToolsService) ToolsController {
	return ToolsController{service: service}
}
