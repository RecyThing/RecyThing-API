package route

import (
	adminHandler "recything/features/admin/handler"
	adminRepository "recything/features/admin/repository"
	adminService "recything/features/admin/service"
	recybotHandler "recything/features/report/handler"
	recybotRepository "recything/features/report/repository"
	recybotService "recything/features/report/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteAdmin(e *echo.Group, db *gorm.DB) {

	adminRepository := adminRepository.NewAdminRepository(db)
	adminService := adminService.NewAdminService(adminRepository)
	adminHandler := adminHandler.NewAdminHandler(adminService)

	//manage prompt
	recybotRepository := recybotRepository.NewReportRepository(db)
	recybotService := recybotService.NewReportService(recybotRepository)
	recybotHandler := recybotHandler.NewReportHandler(recybotService)

	e.POST("/login", adminHandler.Login)

	admin := e.Group("", jwt.JWTMiddleware())
	admin.POST("", adminHandler.Create)
	admin.GET("", adminHandler.GetAll)
	admin.GET("/:id", adminHandler.GetById)
	admin.PUT("/:id", adminHandler.UpdateById)
	admin.DELETE("/:id", adminHandler.Delete)

	//Manage Users
	user := e.Group("/manage/users", jwt.JWTMiddleware())
	user.GET("/users", adminHandler.GetAllUser)
	user.GET("/users/:id", adminHandler.GetByIdUsers)
	user.DELETE("/users/:id", adminHandler.DeleteUsers)

	//Manage Prompt
	recybot := e.Group("/manage/prompts", jwt.JWTMiddleware())
	recybot.POST("", recybotHandler.CreateReport)
}
