package route

import (
	"common-web-framework/controller"
	"common-web-framework/middle"
	"common-web-framework/service"
)

// LoadTopicController 加载专题相关api
func (r *RouterSetup) LoadTopicController(name string) *RouterSetup {
	var group = r.apiGroup.Group(name)

	var topicService = service.NewTopicService()

	var topicController = controller.NewTopicController(topicService)

	{
		group.GET("list", middle.LoggerMiddleware, topicController.GetTopicList)
		group.GET(":id/blog", topicController.GetAllTopicBlogs)
		group.GET("blog", middle.LoggerMiddleware, topicController.GetTopicBlogList)
		group.GET("get/:id", topicController.GetTopicByIdInfo)
		//group.GET("blog/:id", topicController.GetAllTopicBlogs)
		group.GET("user/:id", topicController.GetAllUserTopics)
		group.GET("admin/simple_list", middle.JwtMiddle(middle.AdminRole), topicController.GetAdminUserTopicList)
		group.POST("admin/save", middle.JwtMiddle(middle.AdminRole), topicController.CreateTopic)
		group.PUT("admin/update", middle.JwtMiddle(middle.SuperRole), topicController.UpdateTopic)
		group.PUT("admin/delete", middle.JwtMiddle(middle.SuperRole), topicController.DeleteTopic)
		group.PUT("admin/un_delete", middle.JwtMiddle(middle.SuperRole), topicController.UnDeleteTopic)
		group.GET("admin/list", middle.JwtMiddle(middle.AdminRole), topicController.GetAdminTopicList)
	}
	return r
}
