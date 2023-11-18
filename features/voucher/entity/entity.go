package entity

import "time"

type VoucherCore struct {
	Id          string
	Image       string
	RewardName  string
	Point       int
	Description string
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
