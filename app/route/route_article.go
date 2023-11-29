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

	article := e.Group("/manage/articles", jwt.JWTMiddleware())
	article.POST("", articleHand.CreateArticle)
	article.GET("", articleHand.GetAllArticle)
	article.GET("/:id", articleHand.GetSpecificArticle)
	article.GET("/category",trashCategoryHandler.GetAllCategoryForArticle)
	article.PUT("/:id", articleHand.UpdateArticle)
	article.DELETE("/:id", articleHand.DeleteArticle)
}
