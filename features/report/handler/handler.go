package handler

import (
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

func (rco *reportHandler) CreateReportRubbish(e echo.Context) error {
	userId, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}

	newReport := dto.RubbishRequest{}
	err = e.Bind(&newReport)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	reportInput := dto.RubbishRequestToReportCore(newReport)
	reportInput.ReportType = "Tumpukan Sampah"
	createdReport, err := rco.reportService.Create(reportInput, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	reportResponse := dto.ReportCoreToReportResponse(createdReport)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("success", reportResponse))
}

func (rco *reportHandler) CreateReportSMallLittering(e echo.Context) error {
	userId, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}

	newReport := dto.LitteringSmallRequest{}
	err = e.Bind(&newReport)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	reportInput := dto.LitteringSmallRequestToReportCore(newReport)
	reportInput.ReportType = "Pelanggaran Sampah"
	reportInput.ScaleType = "Skala Besar"

	createdReport, err := rco.reportService.Create(reportInput, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	reportResponse := dto.ReportCoreToReportResponse(createdReport)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("success", reportResponse))
}

func (rco *reportHandler) CreateReportBigLittering(e echo.Context) error {
	userId, _, err := jwt.ExtractToken(e)
	if err != nil {
		return e.JSON(http.StatusUnauthorized, helper.ErrorResponse(err.Error()))
	}

	newReport := dto.LitteringBigRequest{}
	err = e.Bind(&newReport)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	reportInput := dto.LitteringBigRequestToReportCore(newReport)
	reportInput.ReportType = "Pelanggaran Sampah"
	reportInput.ScaleType = "Skala Kecil"

	createdReport, err := rco.reportService.Create(reportInput, userId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	reportResponse := dto.ReportCoreToReportResponse(createdReport)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("success", reportResponse))
}

func (rco *reportHandler) SelectById(e echo.Context) error {
	idParams := e.Param("id")

	result, err := rco.reportService.SelectById(idParams)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse("error reading data"))
	}

	var reportResponse = dto.ReportCoreToReportResponse(result)

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
