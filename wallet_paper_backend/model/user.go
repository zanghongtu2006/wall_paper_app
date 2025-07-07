package model

type User struct {
	BaseModel
	Username   string      `json:"username"`
	Email      string      `json:"email"`
	Avatar     string      `json:"avatar"`
	Wallpapers []Wallpaper `json:"-" gorm:"foreignKey:CreatedBy;references:ID"` // 一对多
}
