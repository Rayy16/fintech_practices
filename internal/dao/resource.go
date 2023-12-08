package dao

import (
	"fintechpractices/global"
	"fintechpractices/internal/model"

	"errors"

	"gorm.io/gorm"
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

func BelongTo(dpID string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("dp_id = ?", dpID)
	}
}

func TypeBy(resourceType resourceType) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("resource_type = ?", resourceType.String())
	}
}

func ResourceLinkBy(rsLink string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("resource_link = ?", rsLink)
	}
}

func GetResourceBy(Options ...func(*gorm.DB) *gorm.DB) ([]model.MetadataMarket, int64, error) {
	db := global.DB
	for _, option := range Options {
		db = option(db)
	}

	var res []model.MetadataMarket
	var cnt int64
	err := db.Where("deleted = false").Find(&res).Count(&cnt).Error
	if err == nil {
		return res, cnt, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return res, 0, nil
	} else {
		return res, -1, err
	}
}

func UpdateResourceByLink(link string, dict map[string]interface{}) error {
	return global.DB.Model(&model.MetadataMarket{}).
		Where("resource_link = ? and deleted = false", link).
		Updates(dict).
		Error
}

func DeleteResourceByLink(link string) error {
	var row model.MetadataMarket
	err := global.DB.
		Where("resource_link = ? and deleted = false", link).
		First(&row).
		Error
	if err != nil {
		return err
	}
	row.Deleted = true
	return global.DB.Save(&row).Error
}

func CreateMetadataMarket(resource *model.MetadataMarket) error {
	return global.DB.Create(resource).Error
}
