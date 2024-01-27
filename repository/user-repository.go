package repository

import (
	"common-web-framework/common"
	"common-web-framework/models"
	"common-web-framework/request"
	"common-web-framework/response"
	"fmt"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db        *gorm.DB
	tableName string
}

func (u UserRepositoryImpl) GetAdminUsers(req request.OtherAdminFilter) (_ []response.UserAdminResponse, count int64) {
	var db = u.db.Table(u.tableName + " as u").Where("u.role_id != 3")

	var result = make([]response.UserAdminResponse, 0)

	if req.Start != "" && req.End != "" {
		db.Where("u.created_at BETWEEN ? AND ?", req.Start, req.End)
	}

	if req.Keyword != "" {
		var like = "%" + req.Keyword + "%"
		db.Where("u.username like ? or u.nick_name like ?", like, like)
	}

	if db.Count(&count); count == 0 {
		return result, count
	}

	var pageCount = common.AdminUserCount

	db.Joins(fmt.Sprintf("join %s r on r.id = u.role_id", common.TableNames.RoleTableName)).
		Select("u.id,u.username,u.nick_name,u.email,u.avatar,u.created_at",
			"r.id as \"Role__id\",r.name as \"Role__name\"").
		Offset((req.Page - 1) * pageCount).
		Limit(pageCount).
		Order(req.Sort.GetOrderString("u.")).
		Scan(&result)

	return result, count
}

func (u UserRepositoryImpl) UpdateRole(id int, role int) int64 {
	return u.db.Table(u.tableName).Where("id = ?", id).UpdateColumn("role_id", role).RowsAffected
}

func (u UserRepositoryImpl) FindByUsernameAndPassword(username, password string) models.User {
	var result models.User

	u.db.Preload("Role").
		First(&result, "username = ? and password = ?", username, password)

	return result
}

func (u UserRepositoryImpl) Save(user models.User) error {
	return u.db.Model(&models.User{}).Create(&user).Error
}

func (u UserRepositoryImpl) Update(user models.User) error {
	return u.db.Model(&models.User{}).Save(&user).Error
}

func (u UserRepositoryImpl) FindById(id int) models.User {
	var result models.User

	u.db.Model(&models.User{}).
		Preload("Role").First(&result, "id = ?", id)

	return result
}

func (u UserRepositoryImpl) FindAll() []models.User {
	var result = make([]models.User, 0)

	u.db.Model(&models.User{}).Preload("Role").Find(&result)

	return result
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepositoryImpl{db: db, tableName: common.TableNames.UserTableName}
}
