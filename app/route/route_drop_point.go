package route

import (
	"recything/features/droppoint/handler"
	"recything/features/droppoint/repository"
	"recything/features/droppoint/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteDropPoint(e *echo.Group, db *gorm.DB) {

	dropPointRepository := repository.NewDropPointRepository(db)
	dropPointService := service.NewDropPointService(dropPointRepository)
	dropPointHandler := handler.NewDropPointHandler(dropPointService)

	user := e.Group("/manage/drop-points", jwt.JWTMiddleware())
	user.POST("", dropPointHandler.CreateDropPoint)
	user.GET("", dropPointHandler.GetAllDropPoint)
	user.GET("/:id", dropPointHandler.GetDropPointById)
	user.PUT("/:id", dropPointHandler.UpdateDropPoint)
	user.DELETE("/:id", dropPointHandler.DeleteDropPoint)
}