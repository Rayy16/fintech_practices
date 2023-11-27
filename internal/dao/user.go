package dao

import (
	"errors"

	"gorm.io/gorm"

	"fintechpractices/global"
	"fintechpractices/internal/model"
)

func IsAccountExisted(account string) (bool, error) {
	ui := model.UserInfo{}
	err := global.DB.Model(&model.UserInfo{}).
		Where("uni_account = ?", account).
		First(&ui).
		Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, err
}

func ComparePassword(account, password string) (bool, error) {
	ui := model.UserInfo{}
	err := global.DB.
		Where("uni_account = ?", account).
		First(&ui).Error

	return password == ui.PassWord, err
}

func GetUserInfo(account string) (*model.UserInfo, error) {
	ui := model.UserInfo{}
	err := global.DB.
		Where("uni_account = ?", account).
		First(&ui).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ui, err
}

func CreateUser(ui *model.UserInfo) error {
	return global.DB.Create(ui).Error
}

func UpdateUser(account string, dict map[string]interface{}) error {
	return global.DB.Model(&model.UserInfo{}).
		Where("uni_account = ?", account).
		Updates(dict).
		Error
}
