package dao

import (
	"fintechpractices/global"

	"gorm.io/gorm"
)

func PageBy(pageNo int, pageSize int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((pageNo - 1) * pageSize).Limit(pageSize)
	}
}

func OrderBy(cond string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(cond)
	}
}

func Count(tableName string, Options ...func(*gorm.DB) *gorm.DB) (int64, error) {
	db := global.DB.Table(tableName)
	for _, option := range Options {
		db = option(db)
	}
	var cnt int64
	err := db.Count(&cnt).Error
	return cnt, err
}
