package response

import "time"

type VoucherResponse struct {
	Id         string `json:"id"`
	Image      string `json:"image"`
	RewardName string `json:"reward_name"`
	Point      int    `json:"point"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
