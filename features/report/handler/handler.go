package handler

import (
	"log"
	"net/http"
	"recything/features/report/dto"
	"recything/features/report/entity"
	"recything/utils/helper"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
)

type reportHandler struct {
	reportService entity.ReportServiceInterface
}

func NewReportHandler(report entity.ReportServiceInterface) *reportHandler {
	return &reportHandler{reportService: report}
}

func (report *reportHandler) CreateReportRubbish(e echo.Context) error {
	userId, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}
	// if role != helper.USER {
	// 	return e.JSON(http.StatusForbidden, helper.ErrorResponse(err.Error()))
	// }

	newReport := dto.ReportRubbishRequest{}
	err = e.Bind(&newReport)
	log.Println("images ", newReport.Images)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	reportInput := entity.ReportRequestToReportCore(newReport)
	createdReport, err := report.reportService.Create(reportInput, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}
	reportResponse := entity.ReportCoreToReportResponse(createdReport)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("success", reportResponse))
}
