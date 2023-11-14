package response

import (
	"recything/features/report/entity"
	user "recything/features/user/entity"
)

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
		Location:       report.Location,
		TrashType:      report.TrashType,
		Description:    report.Description,
		ScaleType:      report.ScaleType,
		InsidentTime:   report.InsidentTime,
		Status:         report.Status,
		CompanyName:    report.CompanyName,
		DangerousWaste: report.DangerousWaste,
		CreatedAt:      report.CreatedAt,
		UpdatedAt:      report.UpdatedAt,
	}
	image := ListImageCoreToImageResponse(report.Images)
	reportResponse.Images = image
	return reportResponse

}

func ReportCoreToReportResponseForDataReporting(report entity.ReportCore, user user.UsersCore) ReportDetails {
	return ReportDetails{
		Id:           report.ID,
		ReportType:   report.ReportType,
		Fullname:     user.Fullname,
		Location:     report.Location,
		InsidentTime: report.InsidentTime,
		Status:       report.Status,
	}
}

func ListReportCoresToReportResponseForDataReporting(reports []entity.ReportCore, userService user.UsersUsecaseInterface) []ReportDetails {
	responReporting := []ReportDetails{}
	for _, report := range reports {
		user, _ := userService.GetById(report.UserId)
		reports := ReportCoreToReportResponseForDataReporting(report, user)
		responReporting = append(responReporting, reports)
	}
	return responReporting
}