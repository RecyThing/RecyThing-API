package dto

import "recything/features/report/entity"

type RubbishRequest struct {
	Longitude    float64        `json:"longitude" form:"longitude"`
	Latitude     float64        `json:"latitude" form:"latitude"`
	Location     string         `json:"location" form:"location"`
	TrashType    string         `json:"trash_type" form:"trash_type"`
	AddressPoint string         `json:"address_point" form:"address_point"`
	Description  string         `json:"description" form:"description"`
	Images       []ImageRequest `json:"images" form:"images"`
}

type LitteringSmallRequest struct {
	Longitude    float64        `json:"longitude" form:"longitude"`
	Latitude     float64        `json:"latitude" form:"latitude"`
	Location     string         `json:"location" form:"location"`
	AddressPoint string         `json:"address_point" form:"address_point"`
	Description  string         `json:"description" form:"description"`
	InsidentTime string         `json:"insident_time" form:"insident_time"`
	Images       []ImageRequest `json:"images" form:"images"`
}

type LitteringBigRequest struct {
	Longitude    float64        `json:"longitude" form:"longitude"`
	Latitude     float64        `json:"latitude" form:"latitude"`
	Location     string         `json:"location" form:"location"`
	WasteType    bool           `json:"waste_type"`
	CompanyName  string         `json:"company_name"`
	AddressPoint string         `json:"address_point" form:"address_point"`
	Description  string         `json:"description" form:"description"`
	InsidentTime string         `json:"insident_time" form:"insident_time"`
	Images       []ImageRequest `json:"images" form:"images"`
}

type ImageRequest struct {
	Image string `json:"image"`
}

func RubbishRequestToReportCore(report RubbishRequest) entity.ReportCore {
	reportCore := entity.ReportCore{
		Longitude:    report.Longitude,
		Latitude:     report.Latitude,
		Location:     report.Location,
		AddressPoint: report.AddressPoint,
		TrashType:    report.TrashType,
		Description:  report.Description,
	}
	image := ListImageRequestToImageCore(report.Images)
	reportCore.Images = image
	return reportCore
}

func LitteringSmallRequestToReportCore(report LitteringSmallRequest) entity.ReportCore {
	reportCore := entity.ReportCore{
		Longitude:    report.Longitude,
		Latitude:     report.Latitude,
		Location:     report.Location,
		AddressPoint: report.AddressPoint,
		Description:  report.Description,
		InsidentTime: report.InsidentTime,
	}
	images := ListImageRequestToImageCore(report.Images)
	reportCore.Images = images
	return reportCore
}

func LitteringBigRequestToReportCore(report LitteringBigRequest) entity.ReportCore {
	reportCore := entity.ReportCore{
		Longitude:    report.Longitude,
		Latitude:     report.Latitude,
		WasteType:    report.WasteType,
		CompanyName:  report.CompanyName,
		Location:     report.Location,
		AddressPoint: report.AddressPoint,
		Description:  report.Description,
		Images:       []entity.ImageCore{},
		InsidentTime: report.InsidentTime,
	}
	images := ListImageRequestToImageCore(report.Images)
	reportCore.Images = images
	return reportCore
}

func ImagerequestToImageCore(image ImageRequest) entity.ImageCore {
	return entity.ImageCore{
		Image: image.Image,
	}
}

func ListImageRequestToImageCore(images []ImageRequest) []entity.ImageCore {
	listImage := []entity.ImageCore{}
	for _, v := range images {

		image := ImagerequestToImageCore(v)
		listImage = append(listImage, image)
	}

	return listImage
}
