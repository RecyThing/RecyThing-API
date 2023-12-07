package route

import (
	"recything/features/dashboard/handler"
	"recything/features/dashboard/repository"
	"recything/features/dashboard/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteDashboard(e *echo.Group, db *gorm.DB) {

	dashboardRepository := repository.NewDashboardRepository(db)
	dashboardService := service.NewDashboardService(dashboardRepository)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)

	dashboard := e.Group("/count-user", jwt.JWTMiddleware())
	dashboard.GET("", dashboardHandler.CountUserActive)
}
