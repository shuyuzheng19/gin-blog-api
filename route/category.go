package route

import (
	"common-web-framework/controller"
	"common-web-framework/middle"
	"common-web-framework/service"
)

// LoadCategoryController 加载分类相关api
func (r *RouterSetup) LoadCategoryController(name string) *RouterSetup {
	var group = r.apiGroup.Group(name)

	var categoryService = service.NewCategoryService()

	var categoryController = controller.NewCategoryController(categoryService)

	{
		group.GET("list", categoryController.GetCategoryList)
		group.POST("admin/save", middle.JwtMiddle(middle.AdminRole), categoryController.CreateCategory)
		group.PUT("admin/update", middle.JwtMiddle(middle.SuperRole), categoryController.UpdateCategory)
		group.PUT("admin/delete", middle.JwtMiddle(middle.SuperRole), categoryController.DeleteCategory)
		group.PUT("admin/un_delete", middle.JwtMiddle(middle.SuperRole), categoryController.UnDeleteCategory)
		group.GET("admin/list", middle.JwtMiddle(middle.AdminRole), categoryController.GetAdminCategoryList)
	}

	return r
}
