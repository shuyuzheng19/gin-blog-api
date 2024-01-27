package response

// TagResponse 博客标签概要
// @Description 博客标签概要
type TagResponse struct {
	Id   int    `json:"id"`   //标签ID
	Name string `json:"name"` //标签名称
}
