package request

type AchievementRequest struct {
	Name        string `json:"name"`
	TargetPoint int    `json:"target_point"`
}