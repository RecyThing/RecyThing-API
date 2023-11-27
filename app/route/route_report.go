package route

import (
	"recything/features/report/handler"
	"recything/features/report/repository"
	"recything/features/report/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteReport(e *echo.Group, db *gorm.DB) {
	// User
	repotRepository := repository.NewReportRepository(db)
	reportService := service.NewReportService(repotRepository)
	reportHandler := handler.NewReportHandler(reportService)

	user := e.Group("", jwt.JWTMiddleware())
	user.POST("", reportHandler.CreateReport)
	user.GET("/history", reportHandler.ReadAllReport)
	user.GET("/history/:id", reportHandler.SelectById)
}
