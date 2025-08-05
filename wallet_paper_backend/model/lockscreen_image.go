package model

type LockScreenImage struct {
	BaseModel
	ID           string `json:"id"`
	LockScreenID string `json:"lockscreenId"` // 外键
	ImageUrl     string `json:"imageUrl"`
	SortOrder    int    `json:"sortOrder"` // 顺序（可选）
}
