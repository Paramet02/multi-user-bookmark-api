package model

import (
	"time"

	"gorm.io/gorm"
)

// Collection represents a collection of bookmarks.
// entity for Collection
type CollectionModel struct {
	ID        int           `gorm:"primaryKey;autoIncrement"`
	UserID    int           `gorm:"not null;index"` // FK -> users.id
	Name      string         `gorm:"size:100;not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"` // ใช้ soft delete

	// relationship กับ UserModel
	User UserModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// TableName overrides the table name used by GORM.
func (CollectionModel) TableName() string {
	return "collections"
}