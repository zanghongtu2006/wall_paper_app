package service

import (
	"github.com/zanghongtu2006/wallet_paper_backend/config"
	"github.com/zanghongtu2006/wallet_paper_backend/model"
)

// GetWallpaperList 获取所有壁纸列表
// 返回一个包含所有壁纸信息的切片
// 如果没有壁纸数据，返回一个空切片
func GetWallpaperList() []model.Wallpaper {
    var dbWallpapers []model.Wallpaper
    config.DB.Find(&dbWallpapers)
    return dbWallpapers
}

// GetWallpaperByID 根据 ID 获取壁纸信息
// 如果找不到对应的壁纸，返回 nil 和错误信息
func GetWallpaperByID(id string) (*model.Wallpaper, error) {
	var wallpaper model.Wallpaper
	result := config.DB.First(&wallpaper, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &wallpaper, nil
}

// CreateWallpaper 创建一条新的壁纸记录
func CreateWallpaper(w *model.Wallpaper) error {
	return config.DB.Create(w).Error
}

// UpdateWallpaper 更新已有壁纸信息
func UpdateWallpaper(id string, updated *model.Wallpaper) error {
	return config.DB.Model(&model.Wallpaper{}).Where("id = ?", id).Updates(updated).Error
}

// DeleteWallpaper 根据 ID 删除壁纸
func DeleteWallpaper(id string) error {
	return config.DB.Delete(&model.Wallpaper{}, "id = ?", id).Error
}
