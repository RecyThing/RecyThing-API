package request

type Mission struct {
	Title         string         `form:"title" json:"title"`
	MissionImage  string         `form:"mission_image" json:"mission_image"`
	Point         int            `form:"point" json:"point"`
	Description   string         `form:"description" json:"description"`
	Start_Date    string         `form:"start_date" json:"start_date"`
	End_Date      string         `form:"end_date" json:"end_date"`
	MissionStages []MissionStage `form:"mission_stages" json:"mission_stages"`
}

type MissionStage struct {
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
}

type AddMissionStage struct {
	MissionID string         `json:"mission_id"`
	Stages    []MissionStage `json:"stages"` 
}

type Claim struct {
	MissionID string `json:"mission_id"`
}
