package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModal struct {
	Id        string         `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	IsDeleted bool           `json:"is_deleted"`
}
