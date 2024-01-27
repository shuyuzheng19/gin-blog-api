package repository

import (
	"common-web-framework/common"
	"common-web-framework/models"
	"common-web-framework/request"
	"common-web-framework/response"
	"fmt"
	"gorm.io/gorm"
)

type FileRepositoryImpl struct {
	db        *gorm.DB
	tableName string
}

func (u FileRepositoryImpl) UpdatePublic(uid int, id int, isPub bool) int64 {
	var build = u.db.Model(&models.FileInfo{}).Table(u.tableName).Where("id = ?", id)

	if uid > 0 {
		build.Where("user_id = ?", uid)
	}

	return build.Update("is_pub", isPub).RowsAffected
}

func (u FileRepositoryImpl) DeleteFile(uid int, id []int) int64 {
	var build = u.db.Model(&models.FileInfo{}).Table(u.tableName).Unscoped().Where("id in ?", id)

	if uid > 0 {
		build.Where("user_id = ?", uid)
	}

	return build.Delete(&models.FileInfo{}).RowsAffected
}

func (u FileRepositoryImpl) GetAdminFile(uid int, req request.OtherAdminFilter) (_ []response.FileAdminResponse, count int64) {

	var db = u.db.Model(&models.FileInfo{}).Table(u.tableName + " as f")

	var result = make([]response.FileAdminResponse, 0)

	if uid > 0 {
		db.Where("f.user_id = ?", uid)
	}

	if req.Start != "" && req.End != "" {
		db.Where("f.created_at BETWEEN ? AND ?", req.Start, req.End)
	}

	if req.Keyword != "" {
		db.Where("f.old_name like ?", "%"+req.Keyword+"%")
	}

	if db.Count(&count); count == 0 {
		return result, count
	}

	var pageCount = common.AdminOtherCount

	db.Joins(fmt.Sprintf("join %s fm on f.md5 = fm.md5", common.TableNames.FileMd5TableName)).
		Joins(fmt.Sprintf("join %s u on u.id = f.user_id", common.TableNames.UserTableName)).
		Select("f.id as id,f.old_name as name,f.created_at,f.is_pub as public,f.size as size",
			"u.id as uid,u.nick_name as nickname",
			"fm.url as url,fm.md5 as md5").
		Offset((req.Page - 1) * pageCount).
		Limit(pageCount).
		Order(req.Sort.GetOrderString("f.")).
		Scan(&result)

	return result, count
}

func (u FileRepositoryImpl) FindFileInfos(uid int, req request.FileRequest) (_ []response.FileResponse, count int64) {

	var files = make([]response.FileResponse, 0)

	var build = u.db.Model(&models.FileInfo{}).Table(u.tableName + " f")

	if uid > 0 {
		build.Where("f.user_id = ?", uid)
	} else {
		build.Where("f.is_pub = ?", true)
	}

	if req.Keyword != "" {
		build.Where("old_name like ?", "%"+req.Keyword+"%")
	}

	if build.Count(&count); count == 0 {
		return files, 0
	}

	build.Joins(fmt.Sprintf("join %s fm on fm.md5 = f.md5", common.TableNames.FileMd5TableName))

	var pageCount = common.FilePageCount

	build.Select("f.id,f.old_name as name,f.created_at,f.suffix,f.size",
		"fm.md5 as md5,fm.url as url").Offset((req.Page - 1) * pageCount).Limit(pageCount)

	if req.Sort == "size" {
		build.Order("f.size desc")
	} else {
		build.Order("f.created_at desc")
	}

	build.Scan(&files)

	return files, count
}

func (u FileRepositoryImpl) BatchSave(files []models.FileInfo) error {
	return u.db.Model(&models.FileInfo{}).Save(&files).Error
}

func (u FileRepositoryImpl) FindByMd5(md5 string) string {
	var r string
	u.db.Model(&models.FileMd5Info{}).Select("url").Where("md5=?", md5).Scan(&r)
	return r
}

func (u FileRepositoryImpl) Save(file models.FileInfo) error {
	return u.db.Model(&models.FileInfo{}).Create(&file).Error
}

func (u FileRepositoryImpl) FindById(id int) *models.FileInfo {
	var result models.FileInfo

	if err := u.db.Model(&models.FileInfo{}).
		Preload("FileMd5").First(&result, "id = ?", id).Error; err != nil {
		return nil
	}

	return &result
}

func (u FileRepositoryImpl) FindAll() []models.FileInfo {
	var result = make([]models.FileInfo, 0)

	u.db.Preload("FileMd5").Find(&result)

	return result
}

func NewFileInfoRepository(db *gorm.DB) FileRepository {
	return FileRepositoryImpl{db: db, tableName: common.TableNames.FileTableName}
}
