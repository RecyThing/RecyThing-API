package route

import (
	"recything/features/admin/handler"
	"recything/features/admin/repository"
	"recything/features/admin/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteAdmin(e *echo.Echo, db *gorm.DB) {

	adminRepository := repository.NewAdminRepository(db)
	adminService := service.NewAdminService(adminRepository)
	adminHandler := handler.NewAdminHandler(adminService)

	e.POST("/admins/login", adminHandler.Login)

	admin := e.Group("/admins", jwt.JWTMiddleware())
	admin.POST("/admins", adminHandler.Create)
	admin.GET("", adminHandler.GetAll)
	admin.GET("/:id", adminHandler.GetById)
	admin.PUT("/:id", adminHandler.UpdateById)
	admin.DELETE("/:id", adminHandler.Delete)

}
