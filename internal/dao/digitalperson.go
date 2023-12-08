package dao

import (
	"errors"
	"fintechpractices/global"
	"fintechpractices/internal/model"

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

func StatusBy(status dpStatus) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("dp_status = ?", status.Int())
	}
}

func OwnerBy(owner string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("owner_id = ?", owner)
	}
}

func DpLinkBy(dpLink string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("dp_link = ?", dpLink)
	}
}

func CoverImageLinkBy(coverImage string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("cover_image_link = ?", coverImage)
	}
}

func PublishedBy(published bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("published = ?", published)
	}
}

func AuditedBy(audited bool) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("audited = ?", audited)
	}
}

func GetDigitalPersonsBy(Options ...func(*gorm.DB) *gorm.DB) ([]model.DigitalPersonInfo, int64, error) {
	db := global.DB
	for _, option := range Options {
		db = option(db)
	}
	var res []model.DigitalPersonInfo
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

func UpdateDigitalPersonByLink(link string, dict map[string]interface{}) error {
	return global.DB.Model(&model.DigitalPersonInfo{}).
		Where("dp_link = ? and deleted = false", link).
		Updates(dict).
		Error
}

func DeleteDigitalPersonByLink(link string) error {
	var row model.DigitalPersonInfo
	err := global.DB.
		Where("dp_link = ? and deleted = false", link).
		First(&row).
		Error
	if err != nil {
		return err
	}

	row.Deleted = true
	return global.DB.Save(&row).Error
}

func CreateDigitalPerson(dp *model.DigitalPersonInfo) error {
	return global.DB.Create(dp).Error
}
