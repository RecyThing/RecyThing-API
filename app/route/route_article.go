package route

import (
	"recything/features/article/handler"
	"recything/features/article/repository"
	"recything/features/article/service"
	trashCategoryHandler "recything/features/trash_category/handler"
	trashCategoryRepository "recything/features/trash_category/repository"
	trashCategoryService "recything/features/trash_category/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteArticle(e *echo.Group, db *gorm.DB) {
	//manage article
	articleRepo := repository.NewArticleRepository(db)
	articleServ := service.NewArticleService(articleRepo)
	articleHand := handler.NewArticleHandler(articleServ)

	//manage trash category
	trashCategoryRepository:=trashCategoryRepository.NewTrashCategoryRepository(db)
	trashCategoryService:=trashCategoryService.NewTrashCategoryService(trashCategoryRepository)
	trashCategoryHandler:=trashCategoryHandler.NewTrashCategoryHandler(trashCategoryService)

	admin := e.Group("/admins/manage/articles", jwt.JWTMiddleware())
	admin.POST("", articleHand.CreateArticle)
	admin.GET("", articleHand.GetAllArticle)
	admin.GET("/:id", articleHand.GetSpecificArticle)
	admin.GET("/category",trashCategoryHandler.GetAllCategoryForArticle)
	admin.PUT("/:id", articleHand.UpdateArticle)
	admin.DELETE("/:id", articleHand.DeleteArticle)

	user := e.Group("/articles", jwt.JWTMiddleware())
	user.GET("", articleHand.GetAllArticle)
	user.GET("/:id", articleHand.GetSpecificArticle)
}
