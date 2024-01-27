package repository

import (
	"common-web-framework/common"
	"common-web-framework/models"
	"common-web-framework/request"
	"common-web-framework/response"
	"common-web-framework/utils"
	"fmt"
	"gorm.io/gorm"
)

type BlogRepositoryImpl struct {
	db        *gorm.DB
	tableName string
}

func (b BlogRepositoryImpl) UpdateEyeCount(id int64, count int64) int64 {
	return b.db.Model(&models.Blog{}).Where("id = ?", id).
		UpdateColumn("eye_count", count).RowsAffected
}

func (b BlogRepositoryImpl) FindAllSearchBlog() []response.SearchBlogResponse {
	var blogs = make([]response.SearchBlogResponse, 0)

	b.db.Model(&models.Blog{}).Select("id,title,description,created_at").Scan(&blogs)

	return blogs
}

func (b BlogRepositoryImpl) Update(blog *models.Blog) int64 {
	tx := b.db.Begin()

	var build = tx.Model(&models.Blog{})

	if blog.UserID > 0 {
		build = build.Where("user_id = ?", blog.UserID)
	} else {
		blog.UserID = 0
	}

	var tags = blog.Tags

	var count = build.Where("id = ?", blog.ID).Updates(&blog).RowsAffected
	if count == 0 {
		tx.Rollback()
		return 0
	}

	if count > 0 {
		var err = tx.Model(&blog).Set("gorm:association_autocreate", false).
			Association("Tags").Replace(tags)
		if err != nil {
			tx.Rollback()
			return 0
		}
		tx.Commit()
		return count
	}
	tx.Rollback()
	return 0
}

func (b BlogRepositoryImpl) FindById(id int64) models.Blog {
	var blog = models.Blog{ID: id}

	var build = b.db.Model(&models.Blog{}).Table(b.tableName + " b")

	b.beforeFindAll("b.*", build)

	build.Preload("Tags")

	build.First(&blog, "b.id=?", id)

	return blog
}

func (b BlogRepositoryImpl) Save(blog *models.Blog) error {
	return b.db.Model(&models.Blog{}).Create(&blog).Error
}

func (b BlogRepositoryImpl) GetMoreBlogInfo(deleted bool, uid int, req request.AdminBlogFilterRequest) (_ []response.AdminBlogResponse, count int64) {

	var blogs = make([]response.AdminBlogResponse, 0)

	var build = b.db.Model(&models.Blog{}).Table(b.tableName + " as b")

	if uid > 0 {
		build.Where("b.user_id = ?", uid)
	}

	if deleted {
		build.Unscoped().Where("b.deleted_at is not null")
	}

	if req.Category > 0 {
		build.Where("b.category_id = ?", &req.Category)
	} else if req.Topic > 0 {
		build.Where("b.topic_id = ?", &req.Topic)
	} else {
		build.Where("b.category_id is not null")
	}

	if req.Keyword != "" {
		var like = "%" + req.Keyword + "%"
		build.Where("b.title like ? or b.description like ?", like, like)
	}

	if req.StartDate != "" && req.EndDate != "" {
		build.Where("b.created_at BETWEEN ? AND ?", req.StartDate, req.EndDate)
	}

	if build.Count(&count); count == 0 {
		return blogs, 0
	}

	var fields = "b.id,b.title,b.description,b.created_at,b.updated_at,b.like_count," +
		"b.eye_count,b.cover_image,b.source_url"

	b.beforeFindAll(fields, build)

	var blogCount = common.AdminBlogCount

	build.Offset((req.Page - 1) * blogCount).
		Limit(blogCount).Order(req.Sort.GetOrderString("b.")).Scan(&blogs)

	return blogs, count

}

func (b BlogRepositoryImpl) FindUserTopBlog(id int) []response.SimpleBlogResponse {
	var blogs = make([]response.SimpleBlogResponse, 0)

	b.db.Model(&models.Blog{}).Table(b.tableName).Select("id,title").
		Where("user_id = ? and category_id is not null", id).
		Offset(0).Limit(common.BlogPageCount).Order(request.EYE.GetOrderString("")).
		Scan(&blogs)

	return blogs
}

func (b BlogRepositoryImpl) beforeFind(fields string, db *gorm.DB) {
	db.Joins(fmt.Sprintf("inner join %s u on u.id=b.user_id", common.TableNames.UserTableName)).
		Joins(fmt.Sprintf("inner join %s c on c.id=b.category_id", common.TableNames.CategoryTableName)).
		Select(fields, "u.id as \"User__id\",u.nick_name as \"User__nick_name\"",
			"c.id as \"Category__id\",c.name as \"Category__name\"")
}

func (b BlogRepositoryImpl) beforeFindAll(fields string, db *gorm.DB) {
	db.Joins(fmt.Sprintf("inner join %s u on u.id=b.user_id", common.TableNames.UserTableName)).
		Joins(fmt.Sprintf("left join %s c on c.id=b.category_id", common.TableNames.CategoryTableName)).
		Joins(fmt.Sprintf("left join %s t on t.id=b.topic_id", common.TableNames.TopicTableName)).
		Select(fields, "u.id as \"User__id\",u.nick_name as \"User__nick_name\"",
			"t.id as \"Topic__id\",t.name as \"Topic__name\"",
			"c.id as \"Category__id\",c.name as \"Category__name\"")
}

func (b BlogRepositoryImpl) FindByUserId(req request.UserBlogRequest) (_ []response.BlogResponse, count int64) {
	var blogs = make([]response.BlogResponse, 0)

	var build = b.db.Model(&models.Blog{}).
		Table(b.tableName+" as b").Where("b.user_id = ? ", req.Id)

	if build.Count(&count); count == 0 {
		return blogs, 0
	}

	var fields = "b.id, b.title, b.description, b.cover_image, b.created_at"

	b.beforeFind(fields, build)

	build.Offset((req.Page - 1) * common.BlogPageCount).
		Limit(common.BlogPageCount).
		Order(request.CREATE.GetOrderString("b.")).
		Scan(&blogs)

	return blogs, count
}

func (b BlogRepositoryImpl) RangeBlog(req request.RangBlogRequest) (_ []response.ArchiveBlogResponse, count int64) {

	var blogs = make([]response.ArchiveBlogResponse, 0)

	var archiveCount = common.ArchivePageCount

	var build = b.db.Model(&models.Blog{}).Table(b.tableName).
		Select("id,title,created_at as create_time,description")

	build.Where("created_at BETWEEN ? AND ?",
		utils.TimeStampToTime(req.StartTime), utils.TimeStampToTime(req.EndTime)).
		Count(&count)

	if count == 0 {
		return blogs, 0
	}

	build.Offset((req.Page - 1) * archiveCount).Limit(archiveCount).
		Order(request.CREATE.GetOrderString("")).Scan(&blogs)

	return blogs, count
}

func (b BlogRepositoryImpl) FindByIdIn(ids []int) []response.RecommendBlogResponse {
	var blogs = make([]response.RecommendBlogResponse, 0)
	b.db.Model(&models.Blog{}).Table(b.tableName).
		Select("id,title,cover_image").
		Where("id in ?", ids).Scan(&blogs)
	return blogs
}

func (b BlogRepositoryImpl) GetHotBlog() []response.SimpleBlogResponse {
	var blogs = make([]response.SimpleBlogResponse, 0)
	b.db.Model(&models.Blog{}).Table(b.tableName).Select("id,title").Limit(10).Order(request.EYE.GetOrderString("")).Scan(&blogs)
	return blogs
}

func (b BlogRepositoryImpl) GetLatestBlog() []response.SimpleBlogResponse {
	var blogs = make([]response.SimpleBlogResponse, 0)
	b.db.Model(&models.Blog{}).Table(b.tableName).Select("id,title").
		Limit(10).Order(request.CREATE.GetOrderString("")).Scan(&blogs)
	return blogs
}

func (b BlogRepositoryImpl) FindCategoryBlogs(req request.BlogListRequest) (_ []response.BlogResponse, count int64) {

	var blogs = make([]response.BlogResponse, 0)

	var build = b.db.Model(&models.Blog{}).Table(b.tableName + " as b")

	if req.Cid > 0 {
		build.Where("b.category_id = ?", req.Cid)
	}

	if build.Count(&count); count == 0 {
		return blogs, 0
	}

	var fields = "b.id, b.title, b.description, b.cover_image, b.created_at"

	b.beforeFind(fields, build)

	var blogCount = common.BlogPageCount

	build.Offset((req.Page - 1) * blogCount).
		Limit(blogCount).
		Order(req.Sort.GetOrderString("b.")).
		Find(&blogs)

	return blogs, count
}

func NewBlogRepository(db *gorm.DB) BlogRepository {
	return BlogRepositoryImpl{db: db, tableName: common.TableNames.BlogTableName}
}
