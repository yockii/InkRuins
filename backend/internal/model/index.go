package model

import (
	"github.com/yockii/yunjin/pkg/util"
	"gorm.io/gorm"
)

var Models []any

type BaseModel struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement:false" json:"id,string"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli" json:"updated_at"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	b.ID = util.NextID()
	return nil
}

type BaseDeleteModel struct {
	BaseModel
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
