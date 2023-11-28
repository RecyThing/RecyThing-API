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

	article := e.Group("", jwt.JWTMiddleware())
	article.POST("/manage/articles", articleHand.CreateArticle)
	article.GET("/manage/articles", articleHand.GetAllArticle)
	article.GET("/manage/articles/:id", articleHand.GetSpecificArticle)
	article.PUT("/manage/articles/:id", articleHand.UpdateArticle)
	article.DELETE("/manage/articles/:id", articleHand.DeleteArticle)

	user := e.Group("",jwt.JWTMiddleware())
	user.GET("/user/articles", articleHand.GetAllArticle)
	user.GET("/user/articles/:id", articleHand.GetSpecificArticle)
}
