package handler

import (
	"net/http"
	"recything/features/drop-point/dto/request"
	"recything/features/drop-point/dto/response"
	"recything/features/drop-point/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"recything/utils/pagination"
	"strconv"

	"github.com/labstack/echo/v4"
)

type dropPointHandler struct {
	dropPointService entity.DropPointServiceInterface
}

func NewDropPointHandler(dropPoint entity.DropPointServiceInterface) *dropPointHandler {
	return &dropPointHandler{
		dropPointService: dropPoint,
	}
}

func (dph *dropPointHandler) CreateDropPoint(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	input := request.DropPointRequest{}

	errBind := helper.DecodeJSON(e, &input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	request := request.DropPointRequestToReportCore(input)

	result, errCreate := dph.dropPointService.CreateDropPoint(request)
	if errCreate != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errCreate.Error()))
	}

	response := response.DropPointCoreToDropPointResponse(result)

	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse(constanta.SUCCESS_CREATE_DATA, response))
}

func (dph *dropPointHandler) UpdateDropPoint(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	input := request.DropPointRequest{}

	errBind := helper.DecodeJSON(e, &input)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errBind.Error()))
	}

	request := request.DropPointRequestToReportCore(input)

	dropPointId := e.Param("id")
	_, errUpdate := dph.dropPointService.UpdateDropPointById(dropPointId, request)
	if errUpdate != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(errUpdate.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil melakukan update data"))

}

func (dph *dropPointHandler) DeleteDropPoint(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	dropPointId := e.Param("id")
	err = dph.dropPointService.DeleteDropPointById(dropPointId)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_DELETE_DATA))
}

func (dph *dropPointHandler) GetAllDropPoint(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	name := e.QueryParam("name")
	address := e.QueryParam("address")
	page, _ := strconv.Atoi(e.QueryParam("page"))
	limit, _ := strconv.Atoi(e.QueryParam("limit"))

	dropPoints, paginationInfo, err := dph.dropPointService.GetAllDropPoint(page, limit, name, address)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(dropPoints) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}

	totalData := len(dropPoints)

	// Ubah model ke format response yang diinginkan
	responseDropPoint := make([]response.DropPointResponse, len(dropPoints))
	for i, dp := range dropPoints {
		responseDropPoint[i] = response.DropPointCoreToDropPointResponse(dp)
	}

	return e.JSON(http.StatusOK, helper.SuccessWithPaginationAndDataResponse("berhasil mendapatkan data", pagination.PaginationMessage(paginationInfo, totalData), responseDropPoint, paginationInfo, totalData))

}

func (dph *dropPointHandler) GetDropPointById(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	idParams := e.Param("id")
	result, err := dph.dropPointService.GetDropPointById(idParams)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse("gagal membaca data"))
	}

	var reportResponse = response.DropPointCoreToDropPointResponse(result)

	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_GET_DATA, reportResponse))
}
