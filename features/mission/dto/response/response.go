package response

import (
	"time"
)

type Mission struct {
	ID            string         `json:"id"`
	Title         string         `json:"name"`
	Creator       string         `json:"creator"`
	Status        string         `json:"status"`
	MissionImage  string         `json:"mission_image"`
	Point         int            `json:"point"`
	Description   string         `json:"description"`
	StartDate     string         `json:"start_date"`
	EndDate       string         `json:"end_date"`
	MissionStages []MissionStage `json:"mission_stages"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
}

type MissionStage struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ApprovalMission struct {
	ID                      string                   `json:"id"`
	MissionName             string                   `json:"mission_name"`
	User                    string                   `json:"user"`
	Status                  string                   `json:"status"`
	Reason                  string                   `json:"reason"`
	Date                    string                   `json:"date"`
	CreatedAt               string                   `json:"created_at"`
	MissionCompletionProofs []MissionCompletionProof `json:"mission_completion_proofs"`
}

type MissionCompletionProof struct {
	ID                 string  `json:"id"`
	MissionStageID     string  `json:"mission_stage_id"`
	TitleStage         string  `json:"title_stage"`
	Description        string  `json:"description"`
	MissionImageProofs []Proof `json:"mission_Image_proofs"`
	Date               string  `json:"date"`
}

type Proof struct {
	ID   string `json:"id"`
	File string `json:"file"`
}
