package request

// TopicBlogRequest 过滤专题博客
// @Description 过滤专题博客
type TopicBlogRequest struct {
	Id   int `form:"id"`   //要查询哪个专题下的博客
	Page int `form:"page"` //第几页
}

// TopicRequest 专题请求模型
// @Description 专题请求模型
type TopicRequest struct {
	Id          int    `json:"id"`                               //专题id
	Name        string `json:"name" validate:"required,max=50"`  //专题名
	Description string `json:"desc" validate:"required,max=200"` //专题描述
	CoverImage  string `json:"cover" validate:"required""`
}
