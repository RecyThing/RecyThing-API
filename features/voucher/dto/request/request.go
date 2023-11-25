package request

type VoucherRequest struct {
	Image       string `form:"image"`
	Reward_Name string `form:"reward_name"`
	Point       int    `form:"point"`
	Description string `form:"description"`
	Start_Date  string `form:"start_date"`
	End_Date    string `form:"end_date"`
}
