package route

import (
	"recything/features/daily_point/handler"
	"recything/features/daily_point/repository"
	"recything/features/daily_point/service"

	missionRepository "recything/features/mission/repository"
	trashExRepository "recything/features/trash_exchange/repository"
	userRepository "recything/features/user/repository"
	voucherRepository "recything/features/voucher/repository"
	achievement"recything/features/achievement/repository"

	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteDailyPoint(e *echo.Group, db *gorm.DB) {
	missionRepo := missionRepository.NewMissionRepository(db)
	trashRepo := trashExRepository.NewTrashExchangeRepository(db)
	achievementRepository := achievement.NewAchievementRepository(db)
	userRepo := userRepository.NewUserRepository(db,achievementRepository)
	voucherRepo := voucherRepository.NewVoucherRepository(db)
	
	dailyRepo := repository.NewDailyPointRepository(db, missionRepo, trashRepo, userRepo, voucherRepo)
	dailyServ := service.NewDailyPointService(dailyRepo)
	dailyHand := handler.NewDailyPointHandler(dailyServ)

	e.POST("", dailyHand.PostWeekly)

	daily := e.Group("/point", jwt.JWTMiddleware())
	daily.POST("/daily", dailyHand.DailyClaim)
	daily.GET("/claimed", dailyHand.ClaimPointHistory)
	daily.GET("/history", dailyHand.PointHistory)
	daily.GET("/history/:idTransaction", dailyHand.PointHistoryById)
}
