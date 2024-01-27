package request

import (
	"common-web-framework/models"
)

// BlogListRequest 博客过滤条件
// @Description 博客过滤条件 只限分类
type BlogListRequest struct {
	Page int  `form:"page"` //第几页博客
	Cid  int  `form:"cid"`  //博客的分类ID
	Sort Sort `form:"sort"` //博客的排序方式
}

// Sort 博客排序方式
// @Description 博客排序方式
type Sort string

const (
	CREATE Sort = "CREATE" //通过创建日期排序
	UPDATE Sort = "UPDATE" //通过修改日期排序
	EYE    Sort = "EYE"    //通过浏览量排序
	LIKE   Sort = "LIKE"   //通过点赞量排序
	BACK   Sort = "BACK"   //通过创建日期倒叙

	SIZE  Sort = "SIZE"  //文件大小正序
	BSIZE Sort = "BSIZE" //文件大小倒叙
)

// GetOrderString 博客列表排序方式
func (sort Sort) GetOrderString(prefix string) string {
	switch sort {
	case CREATE:
		return prefix + "created_at desc"
	case UPDATE:
		return prefix + "updated_at  desc"
	case EYE:
		return prefix + "eye_count desc"
	case LIKE:
		return prefix + "like_count desc"
	case BACK:
		return prefix + "created_at asc"
	case SIZE:
		return prefix + "size desc"
	case BSIZE:
		return prefix + "size asc"
	default:
		return prefix + "created_at desc"
	}
}

// RangBlogRequest 日期区间过滤条件
// @Description 日期区间过滤条件
type RangBlogRequest struct {
	StartTime int64 `form:"start"` //开始日期的时间戳
	EndTime   int64 `form:"end"`   //结束日期的时间戳
	Page      int   `form:"page"`  //第几页
}

// UserBlogRequest 用户博客过滤条件
// @Description 用户博客请求模型
type UserBlogRequest struct {
	Id   int `form:"id"`   //用户id
	Page int `form:"page"` //第几页
}

// SearchBlogRequest 搜索博客
// @Description 搜索博客
type SearchBlogRequest struct {
	Keyword string `form:"keyword"` //搜索关键字
	Page    int    `form:"page"`    //第几页
}

// AdminBlogFilterRequest 后台管理博客过滤条件
// @Description 后台管理博客过滤条件
type AdminBlogFilterRequest struct {
	Page      int    `form:"page"`     //第几页
	Keyword   string `form:"keyword"`  //要搜索的关键字
	Sort      Sort   `form:"sort"`     //排序方式
	Category  int    `form:"category"` //指定分类
	Topic     int    `form:"topic"`    //指定专题
	StartDate string `form:"date[0]"`  //开始日期
	EndDate   string `form:"date[1]"`  //结束日期
}

// BlogRequest 添加博客请求模型
// @Description 添加博客请求模型
type BlogRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Content     string  `json:"content" validate:"required"`
	CoverImage  string  `json:"coverImage" validate:"required"`
	SourceUrl   *string `json:"sourceUrl"`
	Tags        []int   `json:"tags"`
	Category    *int    `json:"category"`
	Topic       *int    `json:"topic"`
}

// ToBlogDo 将请求模型转为数据库模型
func (b BlogRequest) ToBlogDo(uid int) models.Blog {
	var tags []models.Tag
	for _, tag := range b.Tags {
		tags = append(tags, models.Tag{
			ID: tag,
		})
	}
	return models.Blog{
		Description: b.Description,
		Title:       b.Title,
		CoverImage:  b.CoverImage,
		SourceURL:   b.SourceUrl,
		Content:     b.Content,
		CategoryID:  b.Category,
		UserID:      uid,
		TopicID:     b.Topic,
		Tags:        tags,
	}
}
