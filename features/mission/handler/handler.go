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
	err = e.Bind(&requestMission)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	input := request.MissionRequestToMissionCore(requestMission)
	image, _ := e.FormFile("image")

	err = mh.missionService.UpdateMission(image, id, input)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_NOT_FOUND))
		}
		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))

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
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	input := request.MissionStagesRequestToMissionStagesCore(requestStage)
	err = mh.missionService.UpdateMissionStage(id, input)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_NOT_FOUND))
		}

		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))

		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil mengupdate tahapan misi"))

}

func (mh *missionHandler) AddNewMissionStage(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.ADMIN && role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	requestData := request.AddMissionStage{}
	err = e.Bind(&requestData)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))

	}

	data := request.AddMissionStageToMissionStageCore(requestData)
	err = mh.missionService.AddNewMissionStage(requestData.MissionID, data)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil menambahkan tahapan misi"))

}
func (mh *missionHandler) DeleteMissionStage(e echo.Context) error {
	_, role, err := jwt.ExtractToken(e)
	if role != constanta.ADMIN && role != constanta.SUPERADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}
	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	id := e.Param("id")
	err = mh.missionService.DeleteMissionStage(id)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_NOT_FOUND))
		}
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil menghapus tahapan misi"))
}

// membuat admin, hanya untuk super admin
func (mh *missionHandler) ClaimMission(e echo.Context) error {
	userID, role, err := jwt.ExtractToken(e)

	if role != "" {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	input := request.Claim{}
	err = helper.DecodeJSON(e, &input)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	request := request.ClaimRequestToClaimCore(input)

	err = mh.missionService.ClaimMission(userID, request)
	if err != nil {
		if strings.Contains(err.Error(), constanta.ERROR_RECORD_NOT_FOUND) {
			return e.JSON(http.StatusNotFound, helper.ErrorResponse(constanta.ERROR_DATA_NOT_FOUND))
		}
		if strings.Contains(err.Error(), constanta.ERROR) {
			return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
		}

		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil melakukan klaim"))

}
