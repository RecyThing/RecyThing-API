package model

import (
	"time"

	"gorm.io/gorm"
)

type Article struct {
	Id      string `gorm:"primary key"`
	Title   string
	Image   string
	Content string
	// Category  []model.TrashCategory `gorm:"foreignkey:TrashType"`
	Category  string
	Like      int            `gorm:"default:0"`
	Share     int            `gorm:"default:0"`
	CreatedAt time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
