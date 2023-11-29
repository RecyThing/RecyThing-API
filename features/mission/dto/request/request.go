package request

type Mission struct {
	Title       string `form:"title"`
	Point       int    `form:"point"`
	Description string `form:"description"`
	Start_Date   string `form:"start_date"`
	End_Date     string `form:"end_date"`
}

type MissionStages struct {
	MissionID string  `json:"mission_id"`
	Stages    []Stage `json:"stages"`
}

type Stage struct {
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
}
