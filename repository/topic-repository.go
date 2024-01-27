package repository

import (
	"common-web-framework/common"
	"common-web-framework/models"
	"common-web-framework/request"
	"common-web-framework/response"
	"fmt"
	"gorm.io/gorm"
)

type TopicRepositoryImpl struct {
	db        *gorm.DB
	tableName string
}

func (t TopicRepositoryImpl) GetAdminTopic(req request.OtherAdminFilter) (_ []response.AdminTopicResponse, count int64) {

	var db = t.db.Model(&models.Topic{}).Table(t.tableName + " as t")

	var result = make([]response.AdminTopicResponse, 0)

	if req.Deleted {
		db.Unscoped().Where("t.deleted_at is not null")
	} else {
		db.Where("t.deleted_at is null")
	}

	if req.Start != "" && req.End != "" {
		db.Where("t.created_at BETWEEN ? AND ?", req.Start, req.End)
	}

	if req.Keyword != "" {
		db.Where("t.name like ?", "%"+req.Keyword+"%")
	}

	if db.Count(&count); count == 0 {
		return result, count
	}

	t.beforeFind("t.*", db)

	var pageCount = common.AdminOtherCount

	db.Offset((req.Page - 1) * pageCount).
		Limit(pageCount).
		Order(req.Sort.GetOrderString("t.")).
		Find(&result)

	return result, count
}

func (t TopicRepositoryImpl) Create(topic models.Topic) int64 {
	return t.db.Model(&models.Topic{}).Create(&topic).RowsAffected
}

func (t TopicRepositoryImpl) Update(topic models.Topic) int64 {
	return t.db.Model(&models.Topic{}).Unscoped().
		Where("id = ?", topic.ID).
		UpdateColumns(&topic).RowsAffected
}

func (t TopicRepositoryImpl) FindAllTopicBlog(id int) []response.SimpleBlogResponse {
	var blogs = make([]response.SimpleBlogResponse, 0)

	t.db.Model(&models.Blog{}).Select("id,title").
		Where("topic_id = ?", id).
		Order(request.BACK.GetOrderString("")).Scan(&blogs)

	return blogs
}

func (t TopicRepositoryImpl) GetAllUserSimpleTopics(uid int) []response.SimpleTopicResponse {

	var topics = make([]response.SimpleTopicResponse, 0)

	var build = t.db.Model(&models.Topic{}).Table(t.tableName)

	if uid > 0 {
		build.Where("user_id = ?", uid)
	}

	build.Select("id,name").Scan(&topics)

	return topics

}

func (t TopicRepositoryImpl) beforeFind(fields string, db *gorm.DB) {
	db.Joins(fmt.Sprintf("inner join %s as u on u.id = t.user_id", common.TableNames.UserTableName)).
		Select(fields, "u.id as \"User__id\",u.nick_name as \"User__nick_name\"")
}

func (t TopicRepositoryImpl) FindUserTopics(uid int) []response.UserTopicResponse {

	var topics = make([]response.UserTopicResponse, 0)

	t.db.Model(&models.Topic{}).Table(t.tableName).Where("user_id = ?", uid).
		Select("id,name,description,cover_image").
		Order(request.CREATE.GetOrderString("")).Scan(&topics)

	return topics

}

func (t TopicRepositoryImpl) FindBlogByTopicId(req request.TopicBlogRequest) (_ []response.BlogResponse, count int64) {
	var blogs = make([]response.BlogResponse, 0)

	var build = t.db.Model(&models.Blog{}).
		Table(common.TableNames.BlogTableName+" as t").
		Where("t.topic_id = ?", req.Id)

	if build.Count(&count); count == 0 {
		return blogs, 0
	}

	var fields = "t.id,t.title,t.description,t.cover_image,t.created_at"

	t.beforeFind(fields, build)

	var blogCount = common.BlogPageCount

	build.Offset((req.Page - 1) * blogCount).Limit(blogCount).Order(request.BACK.GetOrderString("t.")).Find(&blogs)

	return blogs, count
}

func (t TopicRepositoryImpl) FindById(id int) (r *response.SimpleTopicResponse) {
	t.db.Model(&models.Topic{}).
		Table(t.tableName).Select("id,name").First(&r, "id = ?", id)
	return r
}

func (t TopicRepositoryImpl) FindTopicByPage(page int) (_ []response.TopicResponse, count int64) {

	var topics = make([]response.TopicResponse, 0)

	var topicCount = common.TopicPageCount

	var build = t.db.Model(&models.Topic{}).Table(t.tableName + " as t").Count(&count)

	if count == 0 {
		return topics, 0
	}

	var fields = "t.id,t.name,t.cover_image,t.created_at,t.description"

	t.beforeFind(fields, build)

	build.Offset((page - 1) * topicCount).Limit(topicCount).Find(&topics)

	return topics, count
}

func NewTopicRepository(db *gorm.DB) TopicRepository {
	db.Model(&models.Topic{})
	return TopicRepositoryImpl{db: db, tableName: common.TableNames.TopicTableName}
}
