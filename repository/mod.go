package repository

import (
	"common-web-framework/cache"
	"common-web-framework/common"
	"common-web-framework/config"
	"common-web-framework/models"
	"common-web-framework/request"
	"common-web-framework/response"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func DeleteData(uid int, table string, ids []int) int64 {
	if uid > 0 {
		var sql = fmt.Sprintf("update %s set deleted_at = ? where id in ? and user_id = ?", table)
		return config.DB.Exec(sql, time.Now(), ids, uid).RowsAffected
	} else {
		var sql = fmt.Sprintf("update %s set deleted_at = ? where id in ?", table)
		return config.DB.Exec(sql, time.Now(), ids).RowsAffected
	}
}

func DeleteBlog(uid int, deleted bool, cid []int, tid []int) int64 {
	var build = config.DB.Table(common.TableNames.BlogTableName)

	if uid > 0 {
		build.Where("user_id = ?", uid)
	}

	if tid != nil {
		build.Where("topic_id in ?", tid)
	} else {
		build.Where("category_id in ?", cid)
	}

	var count int64

	if deleted {
		count = build.Delete(&models.Blog{}).RowsAffected
	} else {
		count = build.UpdateColumn("deleted_at", nil).RowsAffected
	}

	if count > 0 {
		cache.ClearBlogPageInfo()
		cache.DeleteBlogMap()
	}

	return count
}

func UnDeleteData(uid int, table string, ids []int) int64 {
	if uid > 0 {
		var sql = fmt.Sprintf("update %s set deleted_at = null where id in ? and user_id = ?", table)
		return config.DB.Exec(sql, ids, uid).RowsAffected
	} else {
		var sql = fmt.Sprintf("update %s set deleted_at = null where id in ?", table)
		return config.DB.Exec(sql, ids).RowsAffected
	}
}

// AdminOtherFilter 统一过滤标签 分类 专题
func AdminOtherFilter(req request.OtherAdminFilter, db *gorm.DB) (_ []response.AdminOtherResponse, count int64) {

	var result = make([]response.AdminOtherResponse, 0)

	if req.Deleted {
		db.Where("deleted_at is not null")
	} else {
		db.Where("deleted_at is null")
	}

	if req.Start != "" && req.End != "" {
		db.Where("created_at BETWEEN ? AND ?", req.Start, req.End)
	}

	if req.Keyword != "" {
		db.Where("name like ?", "%"+req.Keyword+"%")
	}

	if db.Count(&count); count == 0 {
		return result, count
	}

	var pageCount = common.AdminOtherCount

	db.Offset((req.Page - 1) * pageCount).
		Limit(pageCount).
		Order(req.Sort.GetOrderString("")).
		Find(&result)

	return result, count
}
