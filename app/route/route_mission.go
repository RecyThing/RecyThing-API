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
	admin.GET("", missionHandler.GetAllMission)
	admin.DELETE("/:id", missionHandler.DeleteMission)
	admin.GET("/:id", missionHandler.FindById)
	admin.PUT("/:id", missionHandler.UpdateMission)
	admin.PUT("/:id/stages", missionHandler.UpdateMissionStage)

	user := e.Group("/missions", jwt.JWTMiddleware())
	user.GET("", missionHandler.GetAllMission)
	user.GET("/:id", missionHandler.FindById)
	user.POST("", missionHandler.ClaimMission)
	user.POST("/proof",missionHandler.CreateUploadMission)
	user.PUT("/proof/:id",missionHandler.UpdateUploadMission)
}
