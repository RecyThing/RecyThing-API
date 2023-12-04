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

type ExchangeVoucher struct {
	Id        string `gorm:"primary key"`
	IdUser    string
	IdVoucher string `gorm:"index"`
	Vouchers Voucher `gorm:"foreignKey:IdVoucher"`
	Phone     string
	Status    string `gorm:"type:enum('terbaru', 'diproses', 'selesai');default:terbaru"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}