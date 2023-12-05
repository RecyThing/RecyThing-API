package model

import (
	badge "recything/features/achievement/model"
	dp "recything/features/daily_point/model"
	mission "recything/features/mission/model"
	"recything/features/report/model"
	trashExchange "recything/features/trash_exchange/model"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id                string                        `gorm:"primaryKey;not null" json:"id"`
	Email             string                        `gorm:"unique;not null" json:"email"` // unique saya hapus jangan lupa masukin lagi
	Password          string                        `gorm:"not null" json:"password"`
	Fullname          string                        `gorm:"not null" json:"fullname"`
	Phone             string                        `json:"phone"`
	Address           string                        `json:"address"`
	DateOfBirth       string                        `json:"date_of_birth"`
	Purpose           string                        `json:"purpose"`
	Point             int                           `gorm:"default:0" json:"point"`
	Badge             string                        `gorm:"foreignKey:Name;default:'bronze'"`
	IsVerified        bool                          `gorm:"default:false" json:"is_verified"`
	VerificationToken string                        `json:"verification_token"`
	Otp               string                        `json:"otp"`
	OtpExpiration     int64                         `json:"otp_expiration"`
	CreatedAt         time.Time                     `json:"created_at"`
	UpdatedAt         time.Time                     `json:"updated_at"`
	DeleteAt          gorm.DeletedAt                `gorm:"index"`
	Reports           []model.Report                `gorm:"foreignKey:UsersId"`
	Badges            badge.Achievement             `gorm:"foreignKey:Badge;references:Name"`
	TrashExchange     []trashExchange.TrashExchange `gorm:"foreignKey:EmailUser;references:Email"`
	ClaimedMissions   []mission.ClaimedMission      `gorm:"foreignKey:UserID"`
	DailyClaim        []dp.DailyPoint               `gorm:"many2many:UserDailyPoints"`
}

type UserDailyPoints struct {
	UsersID      string
	DailyPointID int
	CreatedAt    time.Time
}
