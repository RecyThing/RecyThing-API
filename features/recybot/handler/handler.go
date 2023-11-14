package handler

import (
	"net/http"
	req "recything/features/recybot/dto/request"
	"recything/features/recybot/dto/response"
	"recything/features/recybot/entity"
	"recything/utils/helper"

	"github.com/labstack/echo/v4"
)

type recybotHandler struct {
	Recybot entity.RecybotServiceInterface
}

func NewRecybotHandler(Recybot entity.RecybotServiceInterface) *recybotHandler {
	return &recybotHandler{Recybot: Recybot}
}

func (rb *recybotHandler) CreateData(e echo.Context) error {
	request := req.RecybotRequest{}
	err := helper.DecodeJSON(e, &request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	input := req.RequestRecybotToCoreRecybot(request)
	result, err := rb.Recybot.CreateData(input)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreRecybotToResponRecybot(result)
	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("Berhasil menambahkan data", response))
}

func (rb *recybotHandler) GetAllData(e echo.Context) error {
	result, err := rb.Recybot.SelectAllData()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.ListCoreRecybotToCoreRecybot(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mendapatkan seluruh data", response))
}

func (rb *recybotHandler) GetById(e echo.Context) error {
	id := e.Param("id")
	result, err := rb.Recybot.SelectById(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreRecybotToResponRecybot(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mendapatkan data", response))
}

func (rb *recybotHandler) DeleteById(e echo.Context) error {
	id := e.Param("id")
	err := rb.Recybot.DeleteData(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusOK, helper.SuccessResponse("Berhasil menghapus data"))
}

func (rb *recybotHandler) UpdateData(e echo.Context) error {
	id := e.Param("id")
	request := req.RecybotRequest{}
	err := helper.DecodeJSON(e, &request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	input := req.RequestRecybotToCoreRecybot(request)
	result, err := rb.Recybot.UpdateData(id, input)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreRecybotToResponRecybot(result)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse("Berhasil mengupdate data", response))
}
