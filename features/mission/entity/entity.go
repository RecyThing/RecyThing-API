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
	ID         string         
	UserID     string       
	MissionID  string       
	Claimed    bool           
	CreatedAt  time.Time
}
