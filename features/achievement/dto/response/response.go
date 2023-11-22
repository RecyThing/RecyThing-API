package response

type AchievementResponse struct {
	Name        string `json:"name"`
	TargetPoint int    `json:"target_point"`
	TotalUser   int    `json:"total_user"`
}
