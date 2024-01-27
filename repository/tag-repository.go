package repository

import (
	"common-web-framework/common"
	"common-web-framework/models"
	"common-web-framework/request"
	"common-web-framework/response"
	"fmt"
	"gorm.io/gorm"
)

type TagRepositoryImpl struct {
	db        *gorm.DB
	tableName string
}

func (t TagRepositoryImpl) Create(tag models.Tag) int64 {
	return t.db.Model(&models.Tag{}).Create(&tag).RowsAffected
}

func (t TagRepositoryImpl) Update(tag models.Tag) int64 {
	return t.db.Model(&models.Tag{}).Unscoped().
		Where("id = ?", tag.ID).
		UpdateColumns(&tag).RowsAffected
}

func (t TagRepositoryImpl) GetAdminTags(req request.OtherAdminFilter) (_ []response.AdminOtherResponse, count int64) {
	return AdminOtherFilter(req, t.db.Table(t.tableName))
}

func (t TagRepositoryImpl) FindBlogByTagId(req request.TagBlogRequest) (_ []response.BlogResponse, count int64) {
	var blogs = make([]response.BlogResponse, 0)

	var build = t.db.Model(&models.Blog{}).Table(common.TableNames.BlogTableName + " as b")

	build.Joins(fmt.Sprintf("inner join %s bt on bt.blog_id = b.id", common.TableNames.BlogTagTableName)).
		Where("bt.tag_id = ? ", req.Id)

	if build.Count(&count); count == 0 {
		return blogs, 0
	}

	build.Joins(fmt.Sprintf("inner join %s u on u.id=b.user_id", common.TableNames.UserTableName)).
		Joins(fmt.Sprintf("inner join %s c on c.id=b.category_id", common.TableNames.CategoryTableName)).
		Select("b.id,b.title,b.description,b.cover_image,b.created_at", "u.id as User__id,u.nick_name as User__nick_name",
			"c.id as \"Category__id\",c.name as \"Category__name\"")

	var blogCount = common.BlogPageCount

	build.Offset((req.Page - 1) * blogCount).Limit(blogCount).
		Order(request.CREATE.GetOrderString("b.")).Find(&blogs)

	return blogs, count
}

func (t TagRepositoryImpl) FindById(id int) (r *response.TagResponse) {
	t.db.Model(&models.Tag{}).
		Table(t.tableName).Select("id,name").First(&r, "id = ?", id)

	return r
}

func (t TagRepositoryImpl) FindAll() []response.TagResponse {

	var tags = make([]response.TagResponse, 0)

	t.db.Model(&models.Tag{}).Table(t.tableName).Select("id,name").Scan(&tags)

	return tags
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return TagRepositoryImpl{db: db, tableName: common.TableNames.TagTableName}
}
