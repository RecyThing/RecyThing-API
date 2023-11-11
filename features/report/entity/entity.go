package entity

import (
	"time"
)

type ReportCore struct {
	ID            string
	ReportType    string
	UserId        string
	Longitude     float64
	Latitude      float64
	Location      string
	TrashType     string
	Description   string
	Status        string
	CompanyName   string
	DangerousWaste bool
	Images        []ImageCore
	ScaleType     string
	InsidentTime  string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type ImageCore struct {
	ID        string
	ReportID  string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
