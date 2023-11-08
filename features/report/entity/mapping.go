package entity

import (
	"recything/features/report/model"
)

func ImageModelToImageCore(image model.Image) ImageCore {
	return ImageCore{
		ID:        image.ID,
		ReportID:  image.ReportId,
		Image:     image.Image,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}
}

func ListImageModelToImageCore(images []model.Image) []ImageCore {
	coreImages := []ImageCore{}
	for _, v := range images {
		image := ImageModelToImageCore(v)
		coreImages = append(coreImages, image)
	}
	return coreImages
}

func ReportModelToReportCore(report model.Report) ReportCore {
	reportCore := ReportCore{
		ID:           report.Id,
		ReportType:   report.ReportType,
		UserId:       report.UsersId,
		Longitude:    report.Longitude,
		Latitude:     report.Latitude,
		AddressPoint: report.AddressPoint,
		Location:     report.Location,
		TrashType:    report.TrashType,
		Description:  report.Description,
		CompanyName:  report.CompanyName,
		WasteType:    report.WasteType,
		ScaleType:    report.ScaleType,
		InsidentTime: report.InsidentTime,
		CreatedAt:    report.CreatedAt,
		UpdatedAt:    report.UpdatedAt,
	}
	image := ListImageModelToImageCore(report.Images)
	reportCore.Images = image
	return reportCore

}

func ImageCoreToImageModel(image ImageCore) model.Image {
	return model.Image{
		ID:        image.ID,
		ReportId:  image.ReportID,
		Image:     image.Image,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}
}

func ListImageCoreToImageModel(images []ImageCore) []model.Image {
	coreImages := []model.Image{}
	for _, v := range images {
		image := ImageCoreToImageModel(v)
		coreImages = append(coreImages, image)
	}
	return coreImages
}

func ReportCoreToReportModel(report ReportCore) model.Report {
	reportModel := model.Report{
		Id:           report.ID,
		ReportType:   report.ReportType,
		UsersId:      report.UserId,
		Longitude:    report.Longitude,
		AddressPoint: report.AddressPoint,
		Latitude:     report.Latitude,
		Location:     report.Location,
		TrashType:    report.TrashType,
		Description:  report.Description,
		ScaleType:    report.ScaleType,
		WasteType:    report.WasteType,
		CompanyName:  report.ReportType,
		InsidentTime: report.InsidentTime,
		CreatedAt:    report.CreatedAt,
		UpdatedAt:    report.UpdatedAt,
	}
	image := ListImageCoreToImageModel(report.Images)
	reportModel.Images = image
	return reportModel

}
