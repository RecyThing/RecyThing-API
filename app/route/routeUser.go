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

	user := e.Group("", jwt.JWTMiddleware())
	user.GET("", userHandler.GetUser)
	user.PUT("", userHandler.UpdateById)
	user.PATCH("/update-password", userHandler.UpdatePassword)
	user.PATCH("/forget-password", userHandler.ForgetPassword)
	e.POST("/otp", userHandler.EmailOTP)
	e.POST("/verify-otp", userHandler.VerifyOTP)
	e.GET("/verify-token", userHandler.VerifyAccount)
	e.POST("/register", userHandler.Register)
	e.POST("/login", userHandler.Login)
}
