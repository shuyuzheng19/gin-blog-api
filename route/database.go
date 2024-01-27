package route

import (
	"common-web-framework/controller"
	"common-web-framework/middle"
	"common-web-framework/service"
)

// LoadDataBaseController 加载数据库相关api
func (r *RouterSetup) LoadDataBaseController(name string) *RouterSetup {
	var group = r.apiGroup.Group(name)

	var databaseService = service.NewDataBaseService()

	var databaseController = controller.NewDataBaseController(databaseService)

	{
		group.GET("get", middle.LoggerMiddleware, middle.JwtMiddle(middle.SuperRole), databaseController.GetTableInsertSQL)
		group.POST("exec", middle.LoggerMiddleware, middle.JwtMiddle(middle.SuperRole), databaseController.ExecSQL)
	}

	return r
}
