package route

import (
	"recything/features/user/handler"
	"recything/features/user/repository"
	"recything/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteLoginPage(e *echo.Group, db *gorm.DB) {
	// User
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandlers(userService)

	e.POST("register", userHandler.Register)
	e.GET("verify-token", userHandler.VerifyAccount)
	e.POST("login", userHandler.Login)
	
	e.POST("forgot-password", userHandler.ForgotPassword)
	e.POST("verify-otp", userHandler.VerifyOTP)
	e.PATCH("new-password", userHandler.NewPassword)
	

}
