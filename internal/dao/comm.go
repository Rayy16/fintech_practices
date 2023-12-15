package dao

import (
	"errors"
	"fintechpractices/global"

	"gorm.io/gorm"
)

type dpStatus struct {
	code int
}

func (d dpStatus) Int() int {
	return d.code
}

var (
	StatusCreatable = dpStatus{0}
	StatusCreating  = dpStatus{1}
	StatusSuccess   = dpStatus{2}
	StatusFailed    = dpStatus{3}
)

type resourceType struct {
	t string
}

func (r resourceType) String() string {
	return r.t
}

func NewResourceType(t string) (resourceType, error) {
	if t == "image" || t == "tone" {
		return resourceType{t: t}, nil
	}
	return resourceType{t: "unknown"}, errors.New("unknown resource type")
}

var (
	TypeUnknown = resourceType{"unknown"}
	TypeImage   = resourceType{"image"}
	TypeTone    = resourceType{"tone"}
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
