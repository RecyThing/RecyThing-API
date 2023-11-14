package entity

import (
	"recything/features/report/dto"
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
		ID:                   report.Id,
		ReportType:           report.ReportType,
		UserId:               report.UsersId,
		Longitude:            report.Longitude,
		Latitude:             report.Latitude,
		Location:             report.Location,
		AddressPoint:         report.AddressPoint,
		Status:               report.Status,
		TrashType:            report.TrashType,
		Description:          report.Description,
		ScaleType:            report.ScaleType,
		InsidentDate:         report.InsidentDate,
		InsidentTime:         report.InsidentTime,
		CompanyName:          report.CompanyName,
		DangerousWaste:       report.DangerousWaste,
		RejectionDescription: report.RejectionDescription,
		CreatedAt:            report.CreatedAt,
		UpdatedAt:            report.UpdatedAt,
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
		Id:                   report.ID,
		ReportType:           report.ReportType,
		UsersId:              report.UserId,
		Longitude:            report.Longitude,
		Latitude:             report.Latitude,
		Location:             report.Location,
		AddressPoint:         report.AddressPoint,
		Status:               report.Status,
		TrashType:            report.TrashType,
		Description:          report.Description,
		ScaleType:            report.ScaleType,
		InsidentDate:         report.InsidentDate,
		InsidentTime:         report.InsidentTime,
		CompanyName:          report.CompanyName,
		DangerousWaste:       report.DangerousWaste,
		RejectionDescription: report.RejectionDescription,
		CreatedAt:            report.CreatedAt,
		UpdatedAt:            report.UpdatedAt,
	}
	image := ListImageCoreToImageModel(report.Images)
	reportModel.Images = image
	return reportModel

}

func ReportRequestToReportCore(report dto.ReportRubbishRequest) ReportCore {
	reportCore := ReportCore{
		ReportType:           report.ReportType,
		Longitude:            report.Longitude,
		Latitude:             report.Latitude,
		Location:             report.Location,
		AddressPoint:         report.AddressPoint,
		Status:               report.Status,
		TrashType:            report.TrashType,
		ScaleType:            report.ScaleType,
		InsidentDate:         report.InsidentDate,
		InsidentTime:         report.InsidentTime,
		DangerousWaste:       report.DangerousWaste,
		RejectionDescription: report.RejectionDescription,
		CompanyName:          report.CompanyName,
		Description:          report.Description,
	}
	image := ListImageRequestToImageCore(report.Images)
	reportCore.Images = image
	return reportCore
}

func ImagerequestToImageCore(image dto.ImageRequest) ImageCore {
	return ImageCore{
		Image: image.Image,
	}
}

func ListImageRequestToImageCore(images []dto.ImageRequest) []ImageCore {
	listImage := []ImageCore{}
	for _, v := range images {
		image := ImagerequestToImageCore(v)
		listImage = append(listImage, image)
	}

	return listImage
}

func ImageCoreToImageResponse(image ImageCore) dto.ImageResponse {
	return dto.ImageResponse{
		ID:        image.ID,
		Image:     image.Image,
		CreatedAt: image.CreatedAt,
		UpdatedAt: image.UpdatedAt,
	}
}

func ListImageCoreToImageResponse(images []ImageCore) []dto.ImageResponse {
	ResponseImages := []dto.ImageResponse{}
	for _, v := range images {
		image := ImageCoreToImageResponse(v)
		ResponseImages = append(ResponseImages, image)
	}
	return ResponseImages
}

func ReportCoreToReportResponse(report ReportCore) dto.ReportCreateResponse {
	reportResponse := dto.ReportCreateResponse{
		Id:                   report.ID,
		ReportType:           report.ReportType,
		Longitude:            report.Longitude,
		Latitude:             report.Latitude,
		Location:             report.Location,
		Description:          report.Description,
		AddressPoint:         report.AddressPoint,
		Status:               report.Status,
		TrashType:            report.TrashType,
		ScaleType:            report.ScaleType,
		InsidentDate:         report.InsidentDate,
		InsidentTime:         report.InsidentTime,
		DangerousWaste:       report.DangerousWaste,
		RejectionDescription: report.RejectionDescription,
		CompanyName:          report.CompanyName,
		CreatedAt:            report.CreatedAt,
		UpdatedAt:            report.UpdatedAt,
	}
	image := ListImageCoreToImageResponse(report.Images)
	reportResponse.Images = image
	return reportResponse
}
