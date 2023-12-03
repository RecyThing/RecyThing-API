package handler

import (
	"net/http"
	"recything/features/daily_point/entity"
	"recything/utils/helper"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
)

type dailyPointHandler struct {
	dailyPointService entity.DailyPointServiceInterface
}

func NewDailyPointHandler(daily entity.DailyPointRepositoryInterface) *dailyPointHandler {
	return &dailyPointHandler{
		dailyPointService: daily,
	}
}

func (daily *dailyPointHandler) PostWeekly(e echo.Context) error {
	err := daily.dailyPointService.PostWeekly()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}
	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil menambahkan weekly daily point"))
}

func (daily *dailyPointHandler) DailyClaim(e echo.Context) error{
	Id, _, _ := jwt.ExtractToken(e)
	if Id == "" {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse("gagal mendapatkan id"))
	}
	
	err := daily.dailyPointService.DailyClaim(Id)
	if err != nil {
		return e.JSON(http.StatusBadRequest, helper.ErrorResponse(err.Error()))
	}

	return e.JSON(http.StatusCreated, helper.SuccessResponse("berhasil melakukan daily claim"))
}