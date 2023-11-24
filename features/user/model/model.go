package model

import (
	badge "recything/features/achievement/model"
	"recything/features/report/model"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id                string            `gorm:"primaryKey;not null" json:"id"`
	Email             string            `gorm:"unique;not null" json:"email"`
	Password          string            `gorm:"not null" json:"password"`
	Fullname          string            `gorm:"not null" json:"fullname"`
	Phone             string            `json:"phone"`
	Address           string            `json:"address"`
	DateOfBirth       string            `json:"date_of_birth"`
	Purpose           string            `json:"purpose"`
	Point             int               `gorm:"default:0" json:"point"`
	Badge             string            `gorm:"foreignKey:Name;default:'bronze'"`
	IsVerified        bool              `gorm:"default:false" json:"is_verified"`
	VerificationToken string            `json:"verification_token"`
	Otp               string            `json:"otp"`
	OtpExpiration     int64             `json:"otp_expiration"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
	DeleteAt          gorm.DeletedAt    `gorm:"index"`
	Reports           []model.Report    `gorm:"foreignKey:UsersId"`
	Badges            badge.Achievement `gorm:"foreignKey:Badge;references:Name"`
}
