package model

type WallpaperImage struct {
	BaseModel
	ID         string `json:"id"`
	WallpaperID string `json:"wallpaperId"` // 外键
	ImageUrl   string `json:"imageUrl"`
	SortOrder  int    `json:"sortOrder"`    // 顺序（可选）
}
