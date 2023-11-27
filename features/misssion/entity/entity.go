package model

import (
	"time"

	"gorm.io/gorm"
)

type Mission struct {
	ID            string
	Name          string
	Creator       string
	Status        string
	AdminID       string
	MissionImage  string
	Point         int
	Description   string
	StartDate     string
	EndDate       string
	MissionStages []MissionStages
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type MissionStages struct {
	ID          string
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
