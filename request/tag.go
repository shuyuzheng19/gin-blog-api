package request

// TagBlogRequest 过滤标签博客
// @Description 过滤标签博客
type TagBlogRequest struct {
	Id   int `form:"id"`   //要查询哪个标签下的博客
	Page int `form:"page"` //第几页
}

// OtherAdminFilter 后台管理过滤分类、标签、专题
// @Description 后台管理过滤分类、标签、专题
type OtherAdminFilter struct {
	Page    int    `form:"page"`    //第几页
	Keyword string `form:"keyword"` //关键字
	Sort    Sort   `form:"sort"`    //排序方式
	Deleted bool   `form:"deleted"` //是否过滤删除
	Start   string `form:"date[0]"` //开始时间
	End     string `form:"date[1]"` //结束时间
}
