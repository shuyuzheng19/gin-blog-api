package service

import (
	"common-web-framework/models"
	"common-web-framework/request"
	"common-web-framework/response"
	"github.com/gin-gonic/gin"
)

// FileService 文件服务
type FileService interface {
	// AddFile 添加一个文件
	AddFile(fileInfo models.FileInfo)
	// GlobalUploadFile 全局上传文件
	GlobalUploadFile(ctx *gin.Context, isImage bool, uid *int, isPub bool) []response.SimpleFileResponse
	// UploadImageFile 上传图片文件
	UploadImageFile(ctx *gin.Context) []response.SimpleFileResponse
	// UploadAvatarFile 上传头像文件
	UploadAvatarFile(ctx *gin.Context) []response.SimpleFileResponse
	// UploadFile 上传其他文件
	UploadFile(ctx *gin.Context) []response.SimpleFileResponse
	// GetPublicFile 获取所有公开文件
	GetPublicFile(req request.FileRequest) response.PageInfo
	// GetUserFile 获取用户的文件
	GetUserFile(uid int, req request.FileRequest) response.PageInfo
	// GetAdminFileList 获取后台管理文件列表
	GetAdminFileList(uid int, req request.OtherAdminFilter) response.PageInfo
	// UpdatePublic 更新是否为公开
	UpdatePublic(uid int, id int, isPub bool)
	// DeleteFile 删除文件
	DeleteFile(uid int, ids []int) int64
	// GetSystemFile 获取本地文件
	GetSystemFile(req request.SystemFileRequest) []response.SystemFileResponse
	// DeleteSystemFile 删除本地文件
	DeleteSystemFile(paths []string) int64
	// ClearFileContent 清空文件
	ClearFileContent(path string) error
}

// UserService 用户服务
type UserService interface {
	// RegisteredUser 注册用户
	RegisteredUser(request request.UserRequest)
	// SendCodeToEmail 发送验证码到用户邮箱
	SendCodeToEmail(email string)
	// ValidateEmailCode 验证验证码是否正确
	ValidateEmailCode(email string, code string)
	// Login 登录
	Login(request request.LoginRequest) response.TokenResponse
	// GetUser 通过ID获取用户
	GetUser(id int) models.User
	// GetToken 通过用户id获取该用户的token
	GetToken(id int) string
	// GetWebSiteConfig 获取网站配置信息
	GetWebSiteConfig() response.BlogConfigInfo
	// SetWebSiteConfig 修改网站配置信息
	SetWebSiteConfig(c response.BlogConfigInfo)
	// Contact 联系我
	Contact(req request.ContactRequest)
	// Logout 退出登录
	Logout(uid int)
	// GetAdminUserList 获取后台管理用户列表
	GetAdminUserList(req request.OtherAdminFilter) response.PageInfo
	// UpdateUserRole 修改用户的角色
	UpdateUserRole(uid int, role int) int64
}

// BlogService 博客服务
type BlogService interface {
	// FindBlogByCategory 查询分类博客
	FindBlogByCategory(req request.BlogListRequest) response.PageInfo
	// GetLatestBlogs 获取最新的前10条博客
	GetLatestBlogs() []response.SimpleBlogResponse
	// GetHostBlogs 获取前10条热门博客
	GetHostBlogs() []response.SimpleBlogResponse
	//SaveRecommend 存入推荐博客
	SaveRecommend(ids []int)
	//GetRecommend 获取推荐博客
	GetRecommend() []response.RecommendBlogResponse
	// GetArchiveBlogList 获取归档博客
	GetArchiveBlogList(req request.RangBlogRequest) response.PageInfo
	// SearchBlog 搜索博客
	SearchBlog(req request.SearchBlogRequest) response.PageInfo
	// SimilarBlog 获取相关的博客内容
	SimilarBlog(keyword string) []any
	// GetUserBlog 获取用户的博客
	GetUserBlog(req request.UserBlogRequest) response.PageInfo
	// GetUserTopBlog 获取用户前10名榜单
	GetUserTopBlog(id int) []response.SimpleBlogResponse
	// GetAdminBlogs 获取后台管理博客列表
	GetAdminBlogs(uid int, req request.AdminBlogFilterRequest) response.PageInfo
	// GetAdminDeleteBlogs 获取后台管理已经删除的博客列表
	GetAdminDeleteBlogs(uid int, req request.AdminBlogFilterRequest) response.PageInfo
	// GetAllAdminBlogs 获取所有的博客列表
	GetAllAdminBlogs(req request.AdminBlogFilterRequest) response.PageInfo
	// GetAllAdminDeleteBlogs 获取所有的已删除博客列表
	GetAllAdminDeleteBlogs(req request.AdminBlogFilterRequest) response.PageInfo
	// DeleteBlogByIds 通过id删除博客
	DeleteBlogByIds(uid int, ids []int) int64
	// UnDeleteBlogByIds 通过id恢复删除的博客
	UnDeleteBlogByIds(uid int, ids []int) int64
	// SaveEditBlog 保存用户编写的博客内容
	SaveEditBlog(uid int, content string)
	// GetSaveEditBlog 获取用户编写的博客内容
	GetSaveEditBlog(uid int) string
	// SaveBlog 保存博客
	SaveBlog(uid int, req request.BlogRequest)
	// UpdateBlog 修改博客
	UpdateBlog(id int64, uid int, req request.BlogRequest)
	// GetBlog 获取博客详情内容
	GetBlog(id int64) response.BlogContentResponse
	// GetUpdateBlog 获取要修改博客的详情内容
	GetUpdateBlog(id int64) response.BlogContentResponse
	// UpdateBlogEyeCount 更新博客浏览量
	UpdateBlogEyeCount(id int64, count int64) int64
	// InitSearch 初始化搜索
	InitSearch()
	// InitEyeCount 初始化浏览量
	InitEyeCount()
}

// TagService 标签服务
type TagService interface {
	// RandomTags 随机获取标签
	RandomTags() []response.TagResponse
	// GetTagInfo 获取某个标签的信息
	GetTagInfo(id int) response.TagResponse
	// GetTagBlogList 获取标签下的博客列表
	GetTagBlogList(req request.TagBlogRequest) response.PageInfo
	// GetAllTag 获取所有标签
	GetAllTag() []response.TagResponse
	// GetAdminTag 获取后台管理标签数据
	GetAdminTag(req request.OtherAdminFilter) response.PageInfo
	// DeleteTagByIds 通过id删除标签
	DeleteTagByIds(ids []int) int64
	// UnDeleteTagByIds 通过id恢复删除的标签
	UnDeleteTagByIds(ids []int) int64
	// AddTag 添加一个标签
	AddTag(name string)
	// UpdateTag 修改标签
	UpdateTag(tag response.TagResponse)
}

// CategoryService 分类服务
type CategoryService interface {
	// GetAllCategory 获取全部分类
	GetAllCategory() []response.CategoryResponse
	// GetAdminCategory 获取后台管理分类数据
	GetAdminCategory(req request.OtherAdminFilter) response.PageInfo
	// DeleteCategoryByIds 通过id删除分类
	DeleteCategoryByIds(ids []int) int64
	// UnDeleteCategoryByIds 通过id恢复删除的分类
	UnDeleteCategoryByIds(ids []int) int64
	// AddCategory 添加一个分类
	AddCategory(name string)
	// UpdateCategory 修改分类
	UpdateCategory(category response.CategoryResponse)
}

// TopicService 专题服务
type TopicService interface {
	// GetTopicList 获取专题
	GetTopicList(page int) response.PageInfo
	// GetTopicInfo 获取某个专题的信息
	GetTopicInfo(id int) response.SimpleTopicResponse
	// GetTopicBlogList 获取专题下的博客列表
	GetTopicBlogList(req request.TopicBlogRequest) response.PageInfo
	// GetAllUserTopics 获取用户的所有专题
	GetAllUserTopics(uid int) []response.UserTopicResponse
	// GetAllUserSimpleTopics 获取当前用户的所有专题
	GetAllUserSimpleTopics(uid int) []response.SimpleTopicResponse
	// GetAllTopicBlogs 获取专题下的所有博客
	GetAllTopicBlogs(id int) []response.SimpleBlogResponse
	// GetAdminTopic 获取后台管理专题数据
	GetAdminTopic(req request.OtherAdminFilter) response.PageInfo
	// DeleteTopicByIds 通过id删除专题
	DeleteTopicByIds(uid int, ids []int) int64
	// UnDeleteTopicByIds 通过id恢复删除的专题
	UnDeleteTopicByIds(uid int, ids []int) int64
	// AddTopic 添加一个专题
	AddTopic(uid int, req request.TopicRequest)
	// UpdateTopic 修改专题
	UpdateTopic(req request.TopicRequest)
}

// DataBaseService 数据库迁移服务
type DataBaseService interface {
	getGlobalResult(tableName string, page int) string
	GetBlogInsertSQL(page int) string
	GetTagInsertSQL(page int) string
	GetCategoryInsertSQL(page int) string
	GetTopicInsertSQL(page int) string
	GetUserInsertSQL(page int) string
	GetRoleInsertSQL(page int) string
	GetFileInsertSQL(page int) string
	GetFileMd5InsertSQL(page int) string
	GetBlogTagInsertSQL(page int) string
	ExecDataBaseSQL(sql string) int64
}
