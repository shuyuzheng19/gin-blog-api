package response

import (
	"fmt"
	"time"
)

type timeStamp time.Time

func (t *timeStamp) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("%d", tTime.UnixMilli())), nil
}

// BlogResponse 博客概要。通常是博客列表信息
// @Description 博客概要。通常是博客列表信息
type BlogResponse struct {
	ID          int64              `json:"id"`                 //博客ID
	Title       string             `json:"title"`              //博客标题
	Description string             `json:"desc"`               //博客描述
	CoverImage  string             `json:"coverImage"`         //博客封面图片
	CreatedAt   timeStamp          `json:"timeStamp"`          //博客发布时间戳
	Category    *CategoryResponse  `json:"category,omitempty"` //博客的分类概要
	User        SimpleUserResponse `json:"user"`               // 博客用户概要
	CategoryId  int                `json:"-"`                  //分类ID
	UserId      int                `json:"-"`                  //用户ID
}

// SimpleBlogResponse 简单的博客概要
// @Description 简洁的博客信息 通常用于最新博客,热门博客
type SimpleBlogResponse struct {
	Id    int    `json:"id"`    //博客ID
	Title string `json:"title"` //博客标题
}

// RecommendBlogResponse 推荐博客
// @Description 推荐博客概要
type RecommendBlogResponse struct {
	Id         string `json:"id"`         //博客ID
	Title      string `json:"title"`      //博客标题
	CoverImage string `json:"coverImage"` //博客封面
}

type myTime time.Time

func (t *myTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

// ArchiveBlogResponse 归档博客
// @Description 归档博客概要
type ArchiveBlogResponse struct {
	Id          int64  `json:"id"`     //博客ID
	Title       string `json:"title"`  //博客标题
	Description string `json:"desc"`   //博客描述
	CreateTime  myTime `json:"create"` //博客创建日期
}

// SearchBlogResponse 搜索博客的模型
// @Description 搜索博客的模型
type SearchBlogResponse struct {
	Id          int64  `json:"id"`          //博客ID
	Title       string `json:"title"`       //博客标题
	Description string `json:"description"` //博客描述
}

// AdminBlogResponse 后台管理博客列表
// @Description 后台管理博客列表
type AdminBlogResponse struct {
	Id          int64                `json:"id"`          //博客id
	Title       string               `json:"title"`       //博客标题
	Description string               `json:"description"` //博客描述
	CoverImage  string               `json:"coverImage"`  //博客封面
	EyeCount    int64                `json:"eyeCount"`    //博客浏览量
	LikeCount   int64                `json:"likeCount"`   //博客点赞量
	Category    *CategoryResponse    `json:"category"`    //博客分类
	Topic       *SimpleTopicResponse `json:"topic"`       //博客专题
	CreatedAt   myTime               `json:"createAt"`    //创建时间
	UpdatedAt   myTime               `json:"updateAt"`    //修改时间
	Original    bool                 `json:"original"`    //博客原文链接
	User        SimpleUserResponse   `json:"user"`        //博客用户信息
	UserId      int                  `json:"-"`
	CategoryId  int                  `json:"-"`
	TopicId     int                  `json:"-"`
}

// BlogContentResponse 博客的详细信息
// @Description 博客的详细信息
type BlogContentResponse struct {
	Id          int64                `json:"id"`
	Description string               `json:"description"`
	Title       string               `json:"title"`
	CoverImage  string               `json:"coverImage"`
	SourceUrl   *string              `json:"source_url"`
	Content     string               `json:"content"`
	EyeCount    int64                `json:"eyeCount"`
	LikeCount   int64                `json:"likeCount"`
	Category    *CategoryResponse    `json:"category"`
	Topic       *SimpleTopicResponse `json:"topic"`
	Tags        []TagResponse        `json:"tags"`
	User        SimpleUserResponse   `json:"user"`
	CreateTime  string               `json:"createTime"`
	UpdateTime  string               `json:"updateTime"`
}
