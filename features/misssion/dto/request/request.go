package request

type Mission struct {
	Name          string          `json:"name"`
	Creator       string          `json:"creator"`
	MissionImage  string          `json:"mission_image"`
	Point         int             `json:"point"`
	Description   string          `json:"description"`
	StartDate     string          `json:"start_date"`
	EndDate       string          `json:"end_date"`
	MissionStages []MissionStages `json:"mission_stages"`
}

type MissionStages struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
