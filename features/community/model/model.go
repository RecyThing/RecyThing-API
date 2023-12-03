package model

import (
	"time"

	"gorm.io/gorm"
)

type Community struct {
	Id          string
	Name        string `gorm:"not null;unique"`
	Description string `gorm:"not null"`
	Location    string `gorm:"not null"`
	Members     int    `gorm:"default:0"`
	MaxMembers  int    `gorm:"not null"`
	Image       string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
