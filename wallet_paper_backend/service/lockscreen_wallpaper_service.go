package service

import (
	"github.com/zanghongtu2006/wallet_paper_backend/config"
	"github.com/zanghongtu2006/wallet_paper_backend/model"
)

func GetLockscreenList() []model.LockScreen {
	var list []model.LockScreen
	config.DB.Find(&list)
	return list
}

func GetLockscreenByID(id string) (*model.LockScreen, error) {
	var item model.LockScreen
	result := config.DB.First(&item, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &item, nil
}

func CreateLockscreen(w *model.LockScreen) error {
	return config.DB.Create(w).Error
}

func UpdateLockscreen(id string, updated *model.LockScreen) error {
	return config.DB.Model(&model.LockScreen{}).Where("id = ?", id).Updates(updated).Error
}

func DeleteLockscreen(id string) error {
	return config.DB.Delete(&model.LockScreen{}, "id = ?", id).Error
}
