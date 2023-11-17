package route

import (
	"recything/features/article/handler"
	"recything/features/article/repository"
	"recything/features/article/service"
	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteArticle(e *echo.Group, db *gorm.DB) {
	//manage article
	articleRepo := repository.NewArticleRepository(db)
	articleServ := service.NewArticleService(articleRepo)
	articleHand := handler.NewArticleHandler(articleServ)

	article := e.Group("/manage/articles", jwt.JWTMiddleware())
	article.POST("", articleHand.CreateArticle)
	article.GET("", articleHand.GetAllArticle)
	article.GET("/:id", articleHand.GetSpecificArticle)
	article.PUT("/:id", articleHand.UpdateArticle)
	article.DELETE("/:id", articleHand.DeleteArticle)
}
