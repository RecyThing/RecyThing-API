package handler

import (
	"net/http"
	"recything/features/dashboard/dto"
	"recything/features/dashboard/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
)

type dashboardHandler struct {
	dashboardService entity.DashboardServiceInterface
}

func NewDashboardHandler(dashboard entity.DashboardServiceInterface) *dashboardHandler {
	return &dashboardHandler{
		dashboardService: dashboard,
	}
}

func (dh *dashboardHandler) Dashboard(e echo.Context) error {

	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}

	result, voucherResult, reportResult, trashExchangeResult, scalaResult, userRanking, err := dh.dashboardService.Dashboard()
	if err != nil {
		return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
	}

	combinedResponse := dto.MapToCombinedResponse(result, voucherResult, reportResult, trashExchangeResult, scalaResult, userRanking)
	return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_GET_DATA, combinedResponse))
}

func (dh *dashboardHandler) CountWeeklyTrashAndScalaTypes(e echo.Context) error {

	_, role, err := jwt.ExtractToken(e)

	if role != constanta.SUPERADMIN && role != constanta.ADMIN {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_AKSES_ROLE))
	}

	if err != nil {
		return e.JSON(http.StatusForbidden, helper.ErrorResponse(constanta.ERROR_EXTRA_TOKEN))
	}
	
    weeklyStats, err := dh.dashboardService.CountWeeklyTrashAndScalaTypes()
    if err != nil {
        return e.JSON(http.StatusInternalServerError, helper.ErrorResponse(err.Error()))
    }

    weeklyStatsResponse := dto.ListMapToWeeklyStatsResponses(weeklyStats)

    return e.JSON(http.StatusOK, helper.SuccessWithDataResponse(constanta.SUCCESS_GET_DATA, weeklyStatsResponse))
}

