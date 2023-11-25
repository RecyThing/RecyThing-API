package entity

import "time"

type VoucherCore struct {
	Id          string
	Image       string
	RewardName  string
	Point       int
	Description string
	StartDate   string
	EndDate     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
