package request

// FileRequest 查询过滤文件条件
// @Description 查询过滤文件条件
type FileRequest struct {
	Page    int    `form:"page"`    //第几页文件
	Keyword string `form:"keyword"` //文件的关键字
	Sort    string `form:"sort"`    //文件排序方式 size:大小排序和date:日期排序
}

// SystemFileRequest 过滤本地文件
type SystemFileRequest struct {
	Path    string `json:"path"`    //文件路径
	Keyword string `json:"keyword"` //文件关键字
}
