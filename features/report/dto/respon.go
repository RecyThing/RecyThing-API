package dto

import (
	"time"
)

type ReportCreateResponse struct {
	Id           string          `json:"Id,omitempty"`
	ReportType   string          `json:"report_type,omitempty"`
	Longitude    float64         `json:"longitude,omitempty"`
	Latitude     float64         `json:"latitude,omitempty"`
	Location     string          `json:"location,omitempty"`
	Description  string          `json:"description,omitempty"`
	Images       []ImageResponse `json:"images,omitempty"`
	TrashType    string          `json:"trash_type,omitempty"`
	ScaleType    string          `json:"scale_type,omitempty"`
	InsidentTime string          `json:"insident_time,omitempty"`
	CreatedAt    time.Time       `json:"created_at,omitempty"`
	UpdatedAt    time.Time       `json:"updated_at,omitempty"`
}

type ImageResponse struct {
	ID        string    `json:"id,omitempty"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
