package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id                string         `gorm:"primaryKey;not null" json:"id"`
	Email             string         `gorm:"unique;not null" json:"email"`
	Password          string         `gorm:"not null" json:"password"`
	Fullname          string         `gorm:"not null" json:"fullname"`
	Phone             string         `json:"phone"`
	Address           string         `json:"address"`
	DateOfBirth       string         `json:"date_of_birth"`
	Purpose           string         `json:"purpose"`
	Point             int            `gorm:"default:0" json:"point"`
	IsVerified        bool           `gorm:"default:false" json:"is_verified"`
	VerificationToken string         `json:"verification_token"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeleteAt          gorm.DeletedAt `gorm:"index"`
}
