package helper

import (
	"common-web-framework/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResultSuccessToResponse(result interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, common.Success(result))
}

func ResultFailToToResponse(code common.ErrorCode, ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusOK, common.AutoFail(code))
}
