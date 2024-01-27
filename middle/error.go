package middle

import (
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/helper"
	"common-web-framework/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ErrorMiddle 全局错误中间件
func ErrorMiddle(ctx *gin.Context) {
	defer func() {
		var err interface{} = recover()

		if err != nil {
			switch t := err.(type) {
			case common.R:
				var ip = utils.GetIPAddress(ctx.Request)
				var city = utils.GetIpCity(ip)
				config.LOGGER.Warn("自定义错误",
					zap.String("path", ctx.FullPath()),
					zap.String("method", ctx.Request.Method),
					zap.String("ip", ip),
					zap.String("city", city),
					zap.Any("error", t))
				ctx.AbortWithStatusJSON(200, t)
				break
			default:
				config.LOGGER.Warn("服务器抛出错误",
					zap.String("path", ctx.FullPath()),
					zap.String("method", ctx.Request.Method), zap.String("error", err.(error).Error()))
				helper.ResultFailToToResponse(common.ServerError, ctx)
				break
			}
		}
	}()

	ctx.Next()
}
