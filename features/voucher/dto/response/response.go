package response

import "time"

type VoucherResponse struct {
	Id         string    `json:"id"`
	RewardName string    `json:"reward_name"`
	Point      int       `json:"point"`
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
}
