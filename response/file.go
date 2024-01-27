package response

// SimpleFileResponse 上传文件返回
// @Description 返回上传的文件信息
type SimpleFileResponse struct {
	Status  string `json:"status"`  //是否上传成功
	Message string `json:"message"` //成功或失败的原因
	Name    string `json:"name"`    //文件的名称
	Create  string `json:"create"`  //文件上传的日期
	Url     string `json:"url"`     //上传成功后返回的url
}

// FileResponse 返回的文件信息
// @Description 返回的文件信息
type FileResponse struct {
	Id        int    `json:"id"`      //文件ID
	Name      string `json:"name"`    //文件名
	CreatedAt myTime `json:"dateStr"` //上传日期
	Suffix    string `json:"suffix"`  //文件后缀
	Size      int64  `json:"size"`    //文件大小
	Md5       string `json:"md5"`     //文件md5
	Url       string `json:"url"`     //文件url
}

// FileAdminResponse 后台管理文件模型
// @Description 后台管理文件模型
type FileAdminResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Size      int    `json:"size"`
	Uid       int    `json:"uid"`
	Nickname  string `json:"nickname"`
	Url       string `json:"url"`
	Md5       string `json:"md5"`
	CreatedAt myTime `json:"createAt"`
	Public    bool   `json:"public"`
}

// SystemFileResponse 本地文件响应
type SystemFileResponse struct {
	Path       string `json:"path"`        //文件路径
	Ext        string `json:"ext"`         //文件后缀
	Name       string `json:"name"`        //文件名称
	Size       int64  `json:"size"`        //文件大小
	CreateTime string `json:"create_time"` //创建日期
	UpdateTime string `json:"update_time"` //修改日期
}
