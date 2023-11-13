package route

import (
	"recything/features/user/handler"
	"recything/features/user/repository"
	"recything/features/user/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteUser(e *echo.Group, db *gorm.DB) {
	// User
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandlers(userService)

	user := e.Group("/profile", jwt.JWTMiddleware())
	user.GET("", userHandler.GetUserById)
	user.PUT("", userHandler.UpdateById)
	user.PATCH("/reset-password", userHandler.UpdatePassword)
	


}
