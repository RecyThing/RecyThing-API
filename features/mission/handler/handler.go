package handler

import (
	"net/http"
	"recything/features/mission/dto/request"
	"recything/features/mission/dto/response"
	"recything/features/mission/entity"

	"strings"

	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
)

type missionHandler struct {
	missionService entity.MissionServiceInterface
}

func NewMissionHandler(missionService entity.MissionServiceInterface) *missionHandler {
	return &missionHandler{missionService: missionService}
}

func (mh *missionHandler) CreateMission(e echo.Context) error {
	id, role, err := jwt.ExtractToken(e)
	if role != constanta.ADMIN && role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	requestMission := request.Mission{}
	err = e.Bind(&requestMission)
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

	input := request.MissionRequestToMissionCore(requestMission)
	input.AdminID = id
	err = mh.missionService.CreateMission(image, input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("Berhasil menambahkan missi"))
}

func (mh *missionHandler) GetAllMission(e echo.Context) error {

	page := e.QueryParam("page")
	limit := e.QueryParam("limit")
	search := e.QueryParam("search")
	status := e.QueryParam("status")

	result, pagnation, count, err := mh.missionService.FindAllMission(page, limit, search, status)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_INVALID_TYPE) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	if len(result) == 0 {
		return e.JSON(http.StatusOK, helper.SuccessResponse("Belum ada missi"))
	}

	response := response.ListMissionCoreToMissionResponse(result)
	return e.JSON(http.StatusOK, helper.SuccessWithPagnationAndCount("Berhasil mendapatkan seluruh missi", response, pagnation, count))
}

func (mh *missionHandler) UpdateMission(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.ADMIN && role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	id := e.Param("id")
	requestMission := request.Mission{}
	err = helper.BindFormData(e, &requestMission)

	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	input := request.MissionRequestToMissionCore(requestMission)
	image, err := e.FormFile("image")
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	err = mh.missionService.UpdateMission(image, id, input)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_DATA_ID) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil mengupdate missi"))
}

func (mh *missionHandler) UpdateMissionStages(e echo.Context) error {

	_, role, err := jwt.ExtractToken(e)
	if role != constanta.ADMIN && role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	id := e.Param("id")
	requestStage := request.MissionStage{}
	err = helper.BindFormData(e, &requestStage)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	input := request.MissionStagesRequestToMissionStagesCore(requestStage)
	err = mh.missionService.UpdateMissionStage(id, input)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_DATA_ID) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil mengupdate mission stages"))

}
