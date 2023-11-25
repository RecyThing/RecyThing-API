package model

import (
	"time"

	"gorm.io/gorm"
)

type Voucher struct {
	Id          string `gorm:"primary key"`
	Image       string
	RewardName  string
	Point       int
	Description string
	StartDate   string
	EndDate     string
	CreatedAt   time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt   time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
