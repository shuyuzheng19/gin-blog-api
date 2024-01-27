package route

import (
	"common-web-framework/config"
	"common-web-framework/controller"
	"common-web-framework/middle"
	"common-web-framework/service"
)

func (r *RouterSetup) LoadBlogController(name string) *RouterSetup {
	var group = r.apiGroup.Group(name)

	var blogService = service.NewBlogService()

	var blogController = controller.NewBlogController(blogService)

	if cronJob != nil {
		cronJob.AddFunc("0 1 * * *", func() {
			config.LOGGER.Info("初始化浏览量.....")
			blogService.InitEyeCount()
			config.LOGGER.Info("初始化浏览量完毕....")
		})
	}

	{
		group.GET("search", middle.LoggerMiddleware, blogController.SearchBlog)
		group.GET("list", middle.LoggerMiddleware, blogController.GetCategoryBlogList)
		group.GET("get/:id", middle.LoggerMiddleware, blogController.GetBlogById)
		group.GET("range", blogController.GetRangBlog)
		group.GET("latest", blogController.GetLatestBlog)
		group.GET("hots", blogController.GetHotBlog)
		group.GET("user", blogController.GetUserBlogList)
		group.GET("similar", blogController.SimilarBlog)
		group.GET("user_top", blogController.GetUserBlogTopList)
		group.GET("recommend", blogController.GetRecommendBlog)
		group.POST("admin/save", middle.JwtMiddle(middle.AdminRole), blogController.SaveBlog)
		group.PUT("admin/update/:id", middle.JwtMiddle(middle.AdminRole), blogController.UpdateBlog)
		group.GET("admin/update/:id", middle.JwtMiddle(middle.AdminRole), blogController.GetUpdateBlog)
		group.POST("admin/user_edit", middle.JwtMiddle(middle.AdminRole), blogController.SaveUserEditor)
		group.GET("admin/user_edit", middle.JwtMiddle(middle.AdminRole), blogController.GetSaveUserEditor)
		group.GET("admin/list", middle.JwtMiddle(middle.AdminRole), blogController.GetAdminBlog)
		group.GET("admin/delete_list", middle.JwtMiddle(middle.AdminRole), blogController.GetAdminDeleteBlog)
		group.GET("admin/all/list", middle.JwtMiddle(middle.AdminRole), blogController.GetAllAdminBlog)
		group.GET("admin/all/delete_list", middle.JwtMiddle(middle.AdminRole), blogController.GetAllDeleteAdminBlog)
		group.PUT("admin/deletes", middle.JwtMiddle(middle.AdminRole), blogController.DeleteBlogs)
		group.PUT("admin/un_deletes", middle.JwtMiddle(middle.AdminRole), blogController.UnDeleteBlogs)
		group.GET("admin/init_search", middle.JwtMiddle(middle.SuperRole), blogController.InitSearch)
		group.GET("admin/init_view", middle.JwtMiddle(middle.SuperRole), blogController.InitEyeCount)
		group.POST("admin/recommend", middle.JwtMiddle(middle.SuperRole), blogController.SetRecommendBlog)
	}

	return r
}
