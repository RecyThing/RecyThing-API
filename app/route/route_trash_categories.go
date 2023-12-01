package route

import (
	"recything/features/trash_category/handler"
	"recything/features/trash_category/repository"
	"recything/features/trash_category/service"

	"recything/utils/jwt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RouteTrash(e *echo.Group, db *gorm.DB) {
	//manage trash category
	trashCategoryRepository := repository.NewTrashCategoryRepository(db)
	trashCategoryService := service.NewTrashCategoryService(trashCategoryRepository)
	trashCategoryHandler := handler.NewTrashCategoryHandler(trashCategoryService)

	//Manage trash category
	trashCategory := e.Group("/manage/trashes", jwt.JWTMiddleware())
	trashCategory.POST("", trashCategoryHandler.CreateCategory)
	trashCategory.GET("", trashCategoryHandler.GetAllCategory)
	trashCategory.GET("/categories", trashCategoryHandler.GetAllCategoriesFetch)
	trashCategory.GET("/:id", trashCategoryHandler.GetById)
	trashCategory.PUT("/:id", trashCategoryHandler.UpdateCategory)
	trashCategory.DELETE("/:id", trashCategoryHandler.DeleteById)

}
