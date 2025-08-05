package service

import (
	"github.com/zanghongtu2006/wallet_paper_backend/config"
	"github.com/zanghongtu2006/wallet_paper_backend/model"
)

func GetImagesByLockScreenID(wallpaperID string) []model.LockScreenImage {
	var images []model.LockScreenImage
	config.DB.Where("wallpaper_id = ?", wallpaperID).Order("sort_order ASC").Find(&images)
	return images
}

func AddImageToLockScreen(image *model.LockScreenImage) error {
	return config.DB.Create(image).Error
}

func DeleteImagesByLockScreenID(wallpaperID string) error {
	return config.DB.Where("wallpaper_id = ?", wallpaperID).Delete(&model.LockScreenImage{}).Error
}
