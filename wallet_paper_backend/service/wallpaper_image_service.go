package service

import (
	"github.com/zanghongtu2006/wallet_paper_backend/config"
	"github.com/zanghongtu2006/wallet_paper_backend/model"
)

func GetImagesByWallpaperID(wallpaperID string) []model.WallpaperImage {
	var images []model.WallpaperImage
	config.DB.Where("wallpaper_id = ?", wallpaperID).Order("sort_order ASC").Find(&images)
	return images
}

func AddImageToWallpaper(image *model.WallpaperImage) error {
	return config.DB.Create(image).Error
}

func DeleteImagesByWallpaperID(wallpaperID string) error {
	return config.DB.Where("wallpaper_id = ?", wallpaperID).Delete(&model.WallpaperImage{}).Error
}
