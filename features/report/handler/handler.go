package handler

import (
	"log"
	"net/http"
	"recything/features/report/dto/request"
	"recything/features/report/dto/response"
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

func (report *reportHandler) CreateReport(e echo.Context) error {
	userId, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}
	// if role != helper.USER {
	// 	return e.JSON(http.StatusForbidden, helper.ErrorResponse(err.Error()))
	// }

	newReport := request.ReportRubbishRequest{}
	err = e.Bind(&newReport)
	log.Println("images ", newReport.Images)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	reportInput := request.ReportRequestToReportCore(newReport)
	createdReport, err := report.reportService.Create(reportInput, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}
	reportResponse := response.ReportCoreToReportResponse(createdReport)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("success", reportResponse))
}

func (rco *reportHandler) SelectById(e echo.Context) error {
	idParams := e.Param("id")

	result, err := rco.reportService.SelectById(idParams)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse("error reading data"))
	}

	var reportResponse = response.ReportCoreToReportResponse(result)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("success get report data", reportResponse))
}

func (rco *reportHandler) ReadAllReport(e echo.Context) error {
	userId, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}

	data, err := rco.reportService.ReadAllReport(userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse("error get report data"))
	}

	return e.JSON(http.StatusOK, map[string]any{
		"messeage": "success get all report data",
		"data":     data,
	})
}
