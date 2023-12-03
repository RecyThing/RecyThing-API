package route

import (
	admin "recything/features/admin/repository"
	"recything/features/mission/handler"
	"recything/features/mission/repository"
	"recything/features/mission/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteMissions(e *echo.Group, db *gorm.DB) {
	adminRepository := admin.NewAdminRepository(db)
	missionRepository := repository.NewMissionRepository(db)
	missionService := service.NewMissionService(missionRepository, adminRepository)
	missionHandler := handler.NewMissionHandler(missionService)

	admin := e.Group("admins/manage/missions", jwt.JWTMiddleware())

	admin.POST("", missionHandler.CreateMission)
	admin.PUT("/:id", missionHandler.UpdateMission)
	admin.PUT("/stages/:id", missionHandler.UpdateMissionStages)
	admin.DELETE("/:id",missionHandler.DeleteMission)

	userAndAdmin := e.Group("/missions", jwt.JWTMiddleware())
	userAndAdmin.GET("", missionHandler.GetAllMission)
	userAndAdmin.POST("",missionHandler.ClaimMission)
	userAndAdmin.GET("/:id", missionHandler.FindById)

}
