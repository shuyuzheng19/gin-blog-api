package middle

import (
	"common-web-framework/config"
	"common-web-framework/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// LoggerMiddleware 记录对方IP等信息日志中间件
func LoggerMiddleware(context *gin.Context) {
	defer func() {
		start := time.Now()

		context.Next()

		latency := time.Since(start)

		var ip = utils.GetIPAddress(context.Request)

		var city = utils.GetIpCity(ip)

		var userAgent = utils.GetClientPlatformInfo(context.GetHeader("User-Agent"))

		config.LOGGER.Info("客户端信息",
			zap.String("ip", ip),
			zap.String("city", city),
			zap.String("user_agent", userAgent),
			zap.String("latency", fmt.Sprintf("%v", latency)),
			zap.String("path", context.Request.RequestURI),
			zap.Int("status", context.Writer.Status()),
			zap.String("method", context.Request.Method),
		)
	}()
}
