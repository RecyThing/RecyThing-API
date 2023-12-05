package response

import (
	"time"
)

type Mission struct {
	ID            string          `json:"id"`
	Title          string          `json:"name"`
	Creator       string          `json:"creator"`
	Status        string          `json:"status"`
	MissionImage  string          `json:"mission_image"`
	Point         int             `json:"point"`
	Description   string          `json:"description"`
	StartDate     string          `json:"start_date"`
	EndDate       string          `json:"end_date"`
	MissionStages []MissionStage `json:"mission_stages"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type MissionStage struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ApprovalMission struct{
	ID string
	MissionName string
	MissionActor string
	Status string
	Date string
	MissionCompletionProof []MissionCompletionProof
}

type MissionCompletionProof struct{
	ID string
	missionStageID string
	TitleStage string
	Description string
	MissionImageProf []Proof
	Date string
}

type Proof struct{
	ID string
	File string 
}


