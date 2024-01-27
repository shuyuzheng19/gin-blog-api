package models

import (
	"common-web-framework/common"
	"common-web-framework/response"
)

// Blog 博客模型
type Blog struct {
	MyModel
	ID          int64   `gorm:"primary_key;comment:博客ID"`
	Description string  `gorm:"size:255;not null;comment:博客描述"`
	Title       string  `gorm:"size:255;not null;comment:博客标题"`
	CoverImage  string  `gorm:"not null;comment:博客封面"`
	SourceURL   *string `gorm:"comment:博客原文链接"`
	Content     string  `gorm:"type:text;comment:博客正文"`
	EyeCount    int64   `gorm:"default:0;comment:浏览量"`
	LikeCount   int64   `gorm:"default:0;comment:点赞量"`
	CategoryID  *int    `gorm:"column:category_id;type:integer;comment:博客分类ID"`
	UserID      int     `gorm:"column:user_id;type:integer;comment:创建的用户ID"`
	TopicID     *int    `gorm:"column:topic_id;type:integer;comment:博客专题ID"`
	Tags        []Tag   `gorm:"many2many:blogs_tags"`
	Category    *Category
	User        User
	Topic       *Topic
}

// TableName 返回与模型对应的数据库表名
func (Blog) TableName() string { return common.TableNames.BlogTableName }

// ToAdminBlogResponse 将DO对象转为vo
//func (b Blog) ToAdminBlogResponse() response.AdminBlogResponse {
//
//	var category *response.CategoryResponse
//
//	if b.Category != nil {
//		category = &response.CategoryResponse{
//			Id:   b.Category.ID,
//			Name: b.Category.Name,
//		}
//	}
//
//	var topic *response.SimpleTopicResponse
//
//	if b.Topic != nil {
//		topic = &response.SimpleTopicResponse{
//			Id:   b.Topic.ID,
//			Name: b.Topic.Name,
//		}
//	}
//
//	return response.AdminBlogResponse{
//		Id:          b.ID,
//		Title:       b.Title,
//		Description: b.Description,
//		CoverImage:  b.CoverImage,
//		EyeCount:    b.EyeCount,
//		LikeCount:   b.LikeCount,
//		Category:    category,
//		Topic:       topic,
//		CreatedAt:   FormatDate(b.CreatedAt),
//		UpdatedAt:   FormatDate(b.UpdatedAt),
//		Original:    b.SourceURL == nil,
//		User:        b.User.ToSimpleUserResponse(),
//	}
//}

// ToContentBlogResponse 将DO对象转为vo
func (b Blog) ToContentBlogResponse() response.BlogContentResponse {

	var category *response.CategoryResponse

	if b.Category != nil {
		category = &response.CategoryResponse{
			Id:   b.Category.ID,
			Name: b.Category.Name,
		}
	}

	var topic *response.SimpleTopicResponse

	if b.Topic != nil {
		topic = &response.SimpleTopicResponse{
			Id:   b.Topic.ID,
			Name: b.Topic.Name,
		}
	}

	return response.BlogContentResponse{
		Id:          b.ID,
		Title:       b.Title,
		Content:     b.Content,
		Description: b.Description,
		CoverImage:  b.CoverImage,
		EyeCount:    b.EyeCount,
		LikeCount:   b.LikeCount,
		Category:    category,
		Topic:       topic,
		CreateTime:  FormatDate(b.CreatedAt),
		UpdateTime:  FormatDate(b.UpdatedAt),
		SourceUrl:   b.SourceURL,
		User:        b.User.ToSimpleUserResponse(),
		Tags:        TagListToBlogResponseList(b.Tags),
	}
}
