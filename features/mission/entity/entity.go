package entity

import (
	"time"
)

type Mission struct {
	ID            string
	Title         string
	Creator       string
	Status        string
	AdminID       string
	MissionImage  string
	Point         int
	Description   string
	StartDate     string
	EndDate       string
	MissionStages []MissionStage
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type MissionStage struct {
	MissionID   string
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Stage struct {
	ID          string
	Title       string
	Description string
}

type ClaimedMission struct {
	ID        string
	UserID    string
	MissionID string
	Claimed   bool
	CreatedAt time.Time
}

// User Upload

type UploadMissionTaskCore struct {
	ID          string
	UserID      string
	User        string
	MissionID   string
	MissionName string
	Description string
	Reason      string
	Images      []ImageUploadMissionCore
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ImageUploadMissionCore struct {
	ID                  string
	UploadMissionTaskID string
	Image               string
	CreatedAt           time.Time
}

// histories
type MissionHistories struct {
	MissionID      string         `json:"mission_id"`
	ClaimedID      string         `json:"claimed_id,omitempty"`
	TransactionID  string         `json:"transaction_id,omitempty"`
	Title          string         `json:"title"`
	StatusApproval string         `json:"status_approval,omitempty"`
	StatusMission  string         `json:"status_mission,omitempty"`
	MissionImage   string         `json:"mission_image"`
	Reason         string         `json:"reason,omitempty"`
	Point          int            `json:"point"`
	Description    string         `json:"description"`
	StartDate      string         `json:"start_date,omitempty"`
	EndDate        string         `json:"end_date,omitempty"`
	MissionStages  []MissionStage `json:"mission_stages,omitempty"`
	CreatedAt      time.Time      `json:"created_at,omitempty"`
	UpdatedAt      time.Time      `json:"updated_at,omitempty"`
}
