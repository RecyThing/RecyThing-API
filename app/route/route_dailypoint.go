package route

import (
	"recything/features/daily_point/handler"
	"recything/features/daily_point/repository"
	"recything/features/daily_point/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteDailyPoint(e *echo.Group, db *gorm.DB) {
	dailyRepo := repository.NewDailyPointRepository(db)
	dailyServ := service.NewDailyPointService(dailyRepo)
	dailyHand := handler.NewDailyPointHandler(dailyServ)

	e.POST("", dailyHand.PostWeekly)

	daily := e.Group("/point", jwt.JWTMiddleware())
	daily.POST("/daily", dailyHand.DailyClaim)
}
