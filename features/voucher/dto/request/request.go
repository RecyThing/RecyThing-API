package request

import "time"

type VoucherRequest struct {
	Image       string    `json:"image"`
	RewardName  string    `json:"reward_name"`
	Point       int       `json:"point"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
