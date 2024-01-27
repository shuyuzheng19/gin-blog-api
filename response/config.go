package response

// Icon 图标
type Icon struct {
	Icon       string `json:"icon"`       //icon svg代码
	Title      string `json:"title"`      //icon 标题
	ModalImage string `json:"modalImage"` //鼠标悬浮显示的图片地址
	Href       string `json:"href"`       //icon链接
	Modal      bool   `json:"modal"`      //是否为模态框
}

// BlogConfigInfo 个人相关配置信息
// @Description 网站配置
type BlogConfigInfo struct {
	Name         string   `json:"name"`         //用户名称
	Avatar       string   `json:"avatar"`       //用户头像
	Icon         []Icon   `json:"icon"`         //图标集合
	MusicID      string   `json:"musicId"`      //网易云音乐ID 已弃用
	Descriptions []string `json:"descriptions"` //网站描述
	Content      string   `json:"content"`      //公告内容
}

// GetDefaultBlogConfigInfo 默认网站配置
func GetDefaultBlogConfigInfo() BlogConfigInfo {
	return BlogConfigInfo{
		Name:         "",
		Avatar:       "",
		Icon:         []Icon{},
		MusicID:      "",
		Descriptions: []string{"后端程序员一枚", "有好的需求可以联系作者"},
		Content:      "",
	}
}
