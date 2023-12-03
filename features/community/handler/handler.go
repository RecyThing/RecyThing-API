package handler

import (
	"net/http"
	"recything/features/community/dto/request"
	"recything/features/community/dto/response"
	"recything/features/community/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"strings"

	"github.com/labstack/echo/v4"
)

type communityHandler struct {
	communityService entity.CommunityServiceInterface
}

func NewCommunityHandler(community entity.CommunityServiceInterface) *communityHandler {
	return &communityHandler{
		communityService: community,
	}
}

func (ch *communityHandler) CreateCommunity(e echo.Context) error {

	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	input := request.CommunityRequest{}
	err := helper.BindFormData(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	image, err := e.FormFile("image")
	if err != nil {
		if err == http.ErrMissingFile {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(constanta.ERROR_EMPTY_FILE))
		}
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal upload file"))
	}

	request := request.RequestCommunityToCoreCommunity(input)
	err = ch.communityService.CreateCommunity(image, request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse(constanta.SUCCESS_CREATE_DATA))
}

func (ch *communityHandler) GetAllCommunity(e echo.Context) error {
	idUser, _, errExtract := jwt.ExtractToken(e)

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	if idUser == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	search := e.QueryParam("search")

	result, pagination, count, err := ch.communityService.GetAllCommunity(page, limit, search)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_INVALID_TYPE) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_NULL))
	}

	response := response.ListCoreCommunityToResponseCommunity(result)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount(constanta.SUCCESS_GET_DATA, response, pagination, count))
}

func (ch *communityHandler) GetCommunityById(e echo.Context) error {
	idUser, _, errExtract := jwt.ExtractToken(e)
	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	if idUser == "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	id := e.Param("id")
	result, err := ch.communityService.GetCommunityById(id)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))

		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreCommunityToResponCommunityForDetails(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_GET_DATA, response))
}

func (ch *communityHandler) DeleteCommunityById(e echo.Context) error {
	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	id := e.Param("id")
	err := ch.communityService.DeleteCommunityById(id)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_DATA_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse(constanta.SUCCESS_DELETE_DATA))
}

func (ch *communityHandler) UpdateCommunityById(e echo.Context) error {
	input := request.CommunityRequest{}

	_, role, errExtract := jwt.ExtractToken(e)
	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if errExtract != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	err := helper.BindFormData(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	id := e.Param("id")
	image, _ := e.FormFile("image")

	request := request.RequestCommunityToCoreCommunity(input)
	err = ch.communityService.UpdateCommunityById(id, image, request)
	if err != nil {
		if helper.HttpResponseCondition(err, constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		if helper.HttpResponseCondition(err, constanta.ERROR_MESSAGE...) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("berhasil update data"))
}