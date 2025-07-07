package model

type Category struct {
	BaseModel
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Wallpapers  []Wallpaper `json:"-" gorm:"foreignKey:CategoryID"` // 一对多
}
