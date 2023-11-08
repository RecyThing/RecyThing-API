package entity

import (
	"time"
)

type ReportCore struct {
	ID           string
	ReportType   string
	UserId       string
	Longitude    float64
	Latitude     float64
	Location     string
	AddressPoint string
	TrashType    string
	Description  string
	Images       []ImageCore
	ScaleType    string
	InsidentTime string
	WasteType    bool
	CompanyName  string
	CreatedAt    time.Time 
	UpdatedAt    time.Time
}

type ImageCore struct {
	ID        string
	ReportID  string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
