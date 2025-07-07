package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BaseModel 定义了所有模型的基础字段
// 包括 ID、创建时间、更新时间和删除时间
type BaseModel struct {
    ID        string         `gorm:"primaryKey;type:char(36)" json:"id"`
    CreatedAt time.Time      `json:"createdAt"`
    UpdatedAt time.Time      `json:"updatedAt"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy string `json:"createdBy"` // 可换成 uint
    UpdatedBy string `json:"updatedBy"`
    IsDeleted bool   `json:"isDeleted"` // 你也可以换成 int
	DeletedBy string `json:"deletedBy"`
    Status    int    `json:"status"`    // 0=草稿，1=有效，9=禁用
}

// BeforeCreate 在创建记录之前调用
// 这里我们使用 UUID 生成一个唯一的 ID
func (b *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
    if b.ID == "" {
        b.ID = uuid.New().String()
    }
    return
}
