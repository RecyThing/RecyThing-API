package route

import (
	"recything/features/community/handler"
	"recything/features/community/repository"
	"recything/features/community/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteCommunity(e *echo.Group, db *gorm.DB) {

	communityRepo := repository.NewCommunityRepository(db)
	communityService := service.NewCommunityService(communityRepo)
	communityHandler := handler.NewCommunityHandler(communityService)

	admin := e.Group("/admins/manage/communities", jwt.JWTMiddleware())
	admin.POST("", communityHandler.CreateCommunity)
	admin.GET("", communityHandler.GetAllCommunity)
	admin.GET("/:id", communityHandler.GetCommunityById)
	admin.PUT("/:id", communityHandler.UpdateCommunityById)
	admin.DELETE("/:id", communityHandler.DeleteCommunityById)

	user := e.Group("/communities", jwt.JWTMiddleware())
	user.GET("", communityHandler.GetAllCommunity)
	user.GET("/:id", communityHandler.GetCommunityById)
}
