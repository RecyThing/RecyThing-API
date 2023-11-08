package dto

import (
	"recything/features/report/entity"
	"time"
)

type ReportCreateResponse struct {
	Id             string          `json:"Id,omitempty"`
	ReportType     string          `json:"report_type,omitempty"`
	Longitude      float64         `json:"longitude,omitempty"`
	Latitude       float64         `json:"latitude,omitempty"`
	Location       string          `json:"location,omitempty"`
	AddressPoint   string          `json:"address_point,omitempty"`
	Description    string          `json:"description,omitempty"`
	Images         []ImageResponse `json:"images,omitempty"`
	TrashType      string          `json:"trash_type,omitempty"`
	ScaleType      string          `json:"scale_type,omitempty"`
	CompanyName    string          `json:"company_name"`
	DangerousWaste bool            `json:"dangerous_waste"`
	InsidentTime   string          `json:"insident_time,omitempty"`
	CreatedAt      time.Time       `json:"created_at,omitempty"`
	UpdatedAt      time.Time       `json:"updated_at,omitempty"`
}

type ImageResponse struct {
	ID        string    `json:"id,omitempty"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func ImageCoreToImageResponse(image entity.ImageCore) ImageResponse {
	return ImageResponse{
		ID:        image.ID,
		Image:     image.Image,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}
}

func ListImageCoreToImageResponse(images []entity.ImageCore) []ImageResponse {
	ResponseImages := []ImageResponse{}
	for _, v := range images {
		image := ImageCoreToImageResponse(v)
		ResponseImages = append(ResponseImages, image)
	}
	return ResponseImages
}

func ReportCoreToReportResponse(report entity.ReportCore) ReportCreateResponse {
	reportResponse := ReportCreateResponse{
		Id:             report.ID,
		ReportType:     report.ReportType,
		Longitude:      report.Longitude,
		Latitude:       report.Latitude,
		AddressPoint:   report.AddressPoint,
		Location:       report.Location,
		TrashType:      report.TrashType,
		Description:    report.Description,
		ScaleType:      report.ScaleType,
		DangerousWaste: report.WasteType,
		CompanyName:    report.CompanyName,
		InsidentTime:   report.InsidentTime,
		CreatedAt:      report.CreatedAt,
		UpdatedAt:      report.UpdatedAt,
	}
	image := ListImageCoreToImageResponse(report.Images)
	reportResponse.Images = image
	return reportResponse

}
