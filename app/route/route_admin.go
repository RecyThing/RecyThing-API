package route

import (
	adminHandler "recything/features/admin/handler"
	adminRepository "recything/features/admin/repository"
	adminService "recything/features/admin/service"

	//userHandler "recything/features/user/handler"
	userRepository "recything/features/user/repository"
	userService "recything/features/user/service"

	recybotHandler "recything/features/recybot/handler"
	recybotRepository "recything/features/recybot/repository"
	recybotService "recything/features/recybot/service"
	
	trashCategoryHandler "recything/features/trash_category/handler"
	trashCategoryRepository "recything/features/trash_category/repository"
	trashCategoryService "recything/features/trash_category/service"

	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteAdmin(e *echo.Group, db *gorm.DB) {

	// import user
	userRepository := userRepository.NewUserRepository(db)
	userService := userService.NewUserService(userRepository)
	//userHandler := adminHandler.NewAdminHandler(userService)

	// manage admin
	adminRepository := adminRepository.NewAdminRepository(db)
	adminService := adminService.NewAdminService(adminRepository)
	adminHandler := adminHandler.NewAdminHandler(adminService, userService)

	//manage prompt
	recybotRepository := recybotRepository.NewRecybotRepository(db)
	recybotService := recybotService.NewRecybotService(recybotRepository)
	recybotHandler := recybotHandler.NewRecybotHandler(recybotService)

	//manage trash category
	trashCategoryRepository:=trashCategoryRepository.NewTrashCategiryRepository(db)
	trashCategoryService:=trashCategoryService.NewTrashCategoryService(trashCategoryRepository)
	trashCategoryHandler:=trashCategoryHandler.NewTrashCategoryHandler(trashCategoryService)

	e.POST("/login", adminHandler.Login)

	admin := e.Group("", jwt.JWTMiddleware())
	admin.POST("", adminHandler.Create)
	admin.GET("", adminHandler.GetAll)
	admin.GET("/:id", adminHandler.GetById)
	admin.PUT("/:id", adminHandler.UpdateById)
	admin.DELETE("/:id", adminHandler.Delete)

	// Manage Users
	user := e.Group("/manage/users", jwt.JWTMiddleware())
	user.GET("", adminHandler.GetAllUser)
	user.GET("/:id", adminHandler.GetByIdUsers)
	user.DELETE("/:id", adminHandler.DeleteUsers)

	// Manage Prompt
	recybot := e.Group("/manage/prompts", jwt.JWTMiddleware())
	recybot.POST("", recybotHandler.CreateData)
	recybot.GET("", recybotHandler.GetAllData)
	recybot.GET("/:id", recybotHandler.GetById)
	recybot.PUT("/:id", recybotHandler.UpdateData)
	recybot.DELETE("/:id", recybotHandler.DeleteById)

	// Manage Reporting
	report := e.Group("/manage/reports", jwt.JWTMiddleware())
	report.GET("", adminHandler.GetByStatusReport)
	report.GET("/:id", adminHandler.GetReportById)
	report.PATCH("/:id", adminHandler.UpdateStatusReport)

	//Manage trash category
	trashCategory := e.Group("/manage/trashes", jwt.JWTMiddleware())
	trashCategory.POST("", trashCategoryHandler.CreateCategory)
	trashCategory.GET("", trashCategoryHandler.GetAllCategory)
	trashCategory.GET("/:id", trashCategoryHandler.GetById)
	trashCategory.PUT("/:id", trashCategoryHandler.UpdateCategory)
	trashCategory.DELETE("/:id", trashCategoryHandler.DeleteById)
	
}
