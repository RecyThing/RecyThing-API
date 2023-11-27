package response

type VoucherResponse struct {
	Id          string `json:"id"`
	Image       string `json:"image"`
	RewardName  string `json:"reward_name"`
	Point       int    `json:"point"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
}
