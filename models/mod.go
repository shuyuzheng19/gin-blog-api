package models

import (
	"common-web-framework/response"
	"gorm.io/gorm"
	"time"
)

type MyModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

//func BlogListToBlogResponseList(blogs []Blog) (list []response.AdminBlogResponse) {
//
//	if len(blogs) == 0 {
//		return []response.AdminBlogResponse{}
//	}
//
//	for _, blog := range blogs {
//		list = append(list, blog.ToAdminBlogResponse())
//	}
//	return list
//}

func TagListToBlogResponseList(tags []Tag) (list []response.TagResponse) {
	if len(tags) == 0 {
		return []response.TagResponse{}
	}

	for _, tag := range tags {
		list = append(list, tag.ToTagResponse())
	}
	return list
}
