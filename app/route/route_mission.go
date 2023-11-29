package route

import (
	"recything/features/mission/handler"
	"recything/features/mission/repository"
	"recything/features/mission/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteMissions(e *echo.Group, db *gorm.DB) {
	missionRepository := repository.NewMissionRepository(db)
	missionService := service.NewMissionService(missionRepository)
	missionHandler := handler.NewMissionHandler(missionService)

	mission := e.Group("/manage/missions", jwt.JWTMiddleware())
	mission.GET("", missionHandler.GetAllMission)
	mission.POST("", missionHandler.CreateMission)
	mission.POST("/stages", missionHandler.CreateMissionStage)
}
