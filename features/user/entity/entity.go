package entity

import (
	"time"

	"gorm.io/gorm"
)

type UsersCore struct {
	Id                string
	Email             string 
	Password          string 
	ConfirmPassword   string 
	Fullname          string
	Phone             string
	Address           string
	DateOfBirth       string
	Purpose           string
	Point             int
	IsVerified        bool
	VerificationToken string
	Otp               string
	NewPassword       string
	OtpExpiration     int64
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeleteAt          gorm.DeletedAt
}
