package route

import (
	"recything/features/recybot/service"
	"recything/features/recybot/handler"
	"recything/features/recybot/repository"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteRecybot(e *echo.Group, db *gorm.DB) {
	// User
	recybotRepository := repository.NewRecybotRepository(db)
	recybotService := service.NewRecybotService(recybotRepository)
	recybotHandler := handler.NewRecybotHandler(recybotService)

	admin := e.Group("/manage/prompts", jwt.JWTMiddleware())
	admin.POST("", recybotHandler.CreateData)

}
