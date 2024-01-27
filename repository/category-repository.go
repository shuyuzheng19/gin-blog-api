package repository

import (
	"common-web-framework/common"
	"common-web-framework/models"
	"common-web-framework/request"
	"common-web-framework/response"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	db        *gorm.DB
	tableName string
}

func (c CategoryRepositoryImpl) Create(category models.Category) int64 {
	return c.db.Model(&models.Category{}).Create(&category).RowsAffected

}

func (c CategoryRepositoryImpl) Update(category models.Category) int64 {
	return c.db.Model(&models.Category{}).Unscoped().
		Where("id = ?", category.ID).
		UpdateColumns(&category).RowsAffected
}

func (c CategoryRepositoryImpl) GetAdminCategories(req request.OtherAdminFilter) (_ []response.AdminOtherResponse, count int64) {
	return AdminOtherFilter(req, c.db.Table(c.tableName))
}

func (c CategoryRepositoryImpl) FindAll() []response.CategoryResponse {

	var categories = make([]response.CategoryResponse, 0)

	c.db.Model(&models.Category{}).Table(c.tableName).Select("id,name").Scan(&categories)

	return categories
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return CategoryRepositoryImpl{db: db, tableName: common.TableNames.CategoryTableName}
}
