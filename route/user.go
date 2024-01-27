package route

import (
	"common-web-framework/controller"
	"common-web-framework/middle"
	"common-web-framework/service"
)

// LoadUserController 加载用户相关api
func (r *RouterSetup) LoadUserController(name string) *RouterSetup {

	var group = r.apiGroup.Group(name)

	var userService = service.NewUserService()

	middle.GetJwtUser = userService.GetUser

	middle.GetToken = userService.GetToken

	var userController = controller.NewUserController(userService)

	{
		group.POST("login", middle.LoggerMiddleware, userController.Login)
		group.POST("registered", middle.LoggerMiddleware, userController.RegisteredUser)
		group.GET("send_email", middle.LoggerMiddleware, userController.SendCodeToEmail)
		group.POST("contact_me", middle.LoggerMiddleware, userController.ContactMe)
		group.GET("config", userController.GetWebSiteConfig)
		group.GET("is_cn", userController.IsChinaIp)
		group.GET("auth/get", middle.JwtMiddle(middle.UserRole), userController.GetUserInfo)
		group.GET("logout", middle.LoggerMiddleware, middle.JwtMiddle(middle.UserRole), userController.Logout)
		group.PUT("config", middle.JwtMiddle(middle.SuperRole), userController.SetWebSiteConfig)
		group.GET("admin/list", middle.JwtMiddle(middle.SuperRole), userController.GetAdminUserList)
		group.PUT("admin/update_role", middle.JwtMiddle(middle.SuperRole), userController.UpdateUserRole)
		group.GET("admin/clear_cache", middle.JwtMiddle(middle.SuperRole), userController.ClearCache)
	}

	return r
}
