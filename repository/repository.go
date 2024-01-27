package repository

import (
	"common-web-framework/models"
	"common-web-framework/request"
	"common-web-framework/response"
)

type BlogRepository interface {
	// FindCategoryBlogs 查询分类下的博客
	//FindCategoryBlogs(req request.BlogListRequest) (result []models.Blog, count int64)
	FindCategoryBlogs(req request.BlogListRequest) (_ []response.BlogResponse, count int64)
	// GetLatestBlog 获取最新博客
	GetLatestBlog() []response.SimpleBlogResponse
	// GetHotBlog 获取前10名热门博客
	GetHotBlog() []response.SimpleBlogResponse
	// FindByIdIn 通过ID获取推荐博客内容
	FindByIdIn(ids []int) []response.RecommendBlogResponse
	// RangeBlog 某个日期区间的博客
	RangeBlog(req request.RangBlogRequest) (_ []response.ArchiveBlogResponse, count int64)
	// FindByUserId 通过用户id获取该用户的博客
	FindByUserId(req request.UserBlogRequest) (_ []response.BlogResponse, count int64)
	// FindUserTopBlog 通过用户id获取该用户的前10热门博客
	FindUserTopBlog(id int) []response.SimpleBlogResponse
	// GetMoreBlogInfo 获取详细的博客信息或已被删除的博客信息
	GetMoreBlogInfo(deleted bool, uid int, req request.AdminBlogFilterRequest) (_ []response.AdminBlogResponse, count int64)
	// Save 保存
	Save(blog *models.Blog) error
	// Update 修改
	Update(blog *models.Blog) int64
	// FindById 通过id查询博客
	FindById(id int64) models.Blog
	// UpdateEyeCount 更新浏览量
	UpdateEyeCount(id int64, count int64) int64
	// FindAllSearchBlog 获取所有博客写入搜索
	FindAllSearchBlog() []response.SearchBlogResponse
}

type FileRepository interface {
	// Save 添加一个文件
	Save(file models.FileInfo) error
	// BatchSave 批量添加文件
	BatchSave(files []models.FileInfo) error
	// FindById 用过ID查询一个用户
	FindById(id int) *models.FileInfo
	// FindByMd5 通过md5查询url
	FindByMd5(md5 string) string
	// FindAll 查询所有用户
	FindAll() []models.FileInfo
	// FindFileInfos 获取用户的文件列表 如果id为<=0则只查询公开文件
	FindFileInfos(uid int, req request.FileRequest) (_ []response.FileResponse, count int64)
	// GetAdminFile 获取后台管理的文件信息
	GetAdminFile(uid int, req request.OtherAdminFilter) (_ []response.FileAdminResponse, count int64)
	// UpdatePublic 更新is_pub字段
	UpdatePublic(uid int, id int, isPub bool) int64
	// DeleteFile 强制删除文件
	DeleteFile(uid int, id []int) int64
}

type UserRepository interface {
	// Save 添加一个用户
	Save(user models.User) error
	// Update 修改用户信息
	Update(user models.User) error
	// FindById 用过ID查询一个用户
	FindById(id int) models.User
	// FindAll 查询所有用户
	FindAll() []models.User
	// FindByUsernameAndPassword 通过账号和密码查询
	FindByUsernameAndPassword(username, password string) models.User
	// GetAdminUsers 获取所有用户
	GetAdminUsers(req request.OtherAdminFilter) (_ []response.UserAdminResponse, count int64)
	// UpdateRole 修改用户角色
	UpdateRole(id int, role int) int64
}

type CategoryRepository interface {
	// FindAll 获取所有分类
	FindAll() []response.CategoryResponse
	// GetAdminCategories 获取后台管理的标签信息
	GetAdminCategories(req request.OtherAdminFilter) (_ []response.AdminOtherResponse, count int64)
	// Create 添加分类
	Create(category models.Category) int64
	// Update 修改分类
	Update(category models.Category) int64
}

type TagRepository interface {
	// FindAll 获取所有标签
	FindAll() []response.TagResponse
	// FindById 通过分类id获取简要
	FindById(id int) (r *response.TagResponse)
	// FindBlogByTagId 获取某个标签下的博客
	FindBlogByTagId(req request.TagBlogRequest) (_ []response.BlogResponse, count int64)
	// GetAdminTags 获取后台管理的标签信息
	GetAdminTags(req request.OtherAdminFilter) (_ []response.AdminOtherResponse, count int64)
	// Create 添加标签
	Create(tag models.Tag) int64
	// Update 修改标签
	Update(tag models.Tag) int64
}

type TopicRepository interface {
	// FindTopicByPage 获取专题列表
	FindTopicByPage(page int) (_ []response.TopicResponse, count int64)
	// FindById 通过分类id获取简要
	FindById(id int) (r *response.SimpleTopicResponse)
	// FindBlogByTopicId 获取某个专题下的博客
	FindBlogByTopicId(req request.TopicBlogRequest) (_ []response.BlogResponse, count int64)
	// FindUserTopics 通过用户id查询该用户所创建的专题
	FindUserTopics(id int) []response.UserTopicResponse
	// GetAllUserSimpleTopics 通过用户ID查询该用户的专题
	GetAllUserSimpleTopics(uid int) []response.SimpleTopicResponse
	// FindAllTopicBlog 获取该专题下的所有博客
	FindAllTopicBlog(id int) []response.SimpleBlogResponse
	// GetAdminTopic 获取后台管理的标签信息
	GetAdminTopic(req request.OtherAdminFilter) (_ []response.AdminTopicResponse, count int64)
	// Create 添加专题
	Create(topic models.Topic) int64
	// Update 修改专题
	Update(topic models.Topic) int64
}
