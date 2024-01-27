package route

import (
	"common-web-framework/controller"
	"common-web-framework/middle"
	"common-web-framework/service"
)

// LoadTagController 加载标签相关api
func (r *RouterSetup) LoadTagController(name string) *RouterSetup {
	var group = r.apiGroup.Group(name)

	var tagService = service.NewTagService()

	var tagController = controller.NewTagController(tagService)

	{
		group.GET("random", tagController.GetRandomTags)
		group.GET("blog", tagController.GetTagBlogList)
		group.GET("list", tagController.GetAllTagInfo)
		group.GET("get/:id", tagController.GetTagByIdInfo)
		group.POST("admin/save", middle.JwtMiddle(middle.AdminRole), tagController.CreateTag)
		group.PUT("admin/update", middle.JwtMiddle(middle.SuperRole), tagController.UpdateTag)
		group.PUT("admin/delete", middle.JwtMiddle(middle.SuperRole), tagController.DeleteTag)
		group.PUT("admin/un_delete", middle.JwtMiddle(middle.SuperRole), tagController.UnDeleteTag)
		group.GET("admin/list", middle.JwtMiddle(middle.AdminRole), tagController.GetAdminTagList)
	}
	return r
}
