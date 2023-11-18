package model

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	Id          string `gorm:"primary key"`
	Image       string `gorm:"not null"`
	RewardName  string `gorm:"not null"`
	Point       int
	Description string    `gorm:"not null"`
	StartDate   time.Time `gorm:"not null"`
	EndDate     time.Time `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
