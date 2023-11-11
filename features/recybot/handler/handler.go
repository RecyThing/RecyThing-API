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
	err := e.Bind(&request)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}
	input := req.RequestRecybotToCoreRecybot(request)
	result, err := rb.Recybot.CreateData(input)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	response := response.CoreRecybotToResponRecybot(result)

	return e.JSON(http.StatusCreated, helper.SuccessWithDataResponse("berhasil menambahkan data", response))

}
