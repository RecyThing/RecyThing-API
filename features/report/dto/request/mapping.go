package request

import "recything/features/report/entity"

func ReportRequestToReportCore(report ReportRubbishRequest) entity.ReportCore {
	reportCore := entity.ReportCore{
		ReportType:     report.ReportType,
		Longitude:      report.Longitude,
		Latitude:       report.Latitude,
		Location:       report.Location,
		TrashType:      report.TrashType,
		Description:    report.Description,
		ScaleType:      report.ScaleType,
		InsidentTime:   report.InsidentTime,
		CompanyName:    report.CompanyName,
		DangerousWaste: report.DangerousWaste,
	}
	image := ListImageRequestToImageCore(report.Images)
	reportCore.Images = image
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
