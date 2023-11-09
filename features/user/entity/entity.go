package entity

import (
	"time"

	"gorm.io/gorm"
)

type UsersCore struct {
	Id                string
	Email             string `validate:"required,email"`
	Password          string `validate:"required,min=8"`
	ConfirmPassword   string `validate:"required,eqfield=Password"`
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
	OtpExpiration     time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeleteAt          gorm.DeletedAt
}
