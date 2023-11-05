package route

import (
	userHandler "recything/features/user/handler"
	userRepository "recything/features/user/repository"
	userService "recything/features/user/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	// User
	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	userHandler := userHandler.NewUserHandlers(userService)

	user := e.Group("/users")
	user.POST("/register", userHandler.Register)
	user.POST("/login", userHandler.Login)
	user.GET("", userHandler.GetUser, jwt.JWTMiddleware())
	user.PUT("", userHandler.UpdateById, jwt.JWTMiddleware())
	e.GET("/verify", userHandler.VerifyAccount)
}
