package request

import (
	"recything/features/article/entity"
)

func ArticleRequestToArticleCore(article ArticleRequest) entity.ArticleCore {
	articleReq := entity.ArticleCore{
		Title:       article.Title,
		Image:       article.Image,
		Content:     article.Content,
		Category_id: article.Category_id,
	}
	// category := ListCategoryRequestToCategoryCore(article.Category_id)
	// articleReq.Category_id = category
	return articleReq
}

// func CategotyrequestToCategotyCore(category ArticleTrashCategoryRequest) entity.CategoryCore {
// 	return entity.CategoryCore{
// 		ArticleID: category.ArticleID,
// 		TrashCategoryID: category.TrashCategoryID,
// 	}
// }

// func ListCategoryRequestToCategoryCore(categories []ArticleTrashCategoryRequest) []entity.CategoryCore {
// 	listCategory := []entity.CategoryCore{}
// 	for _, v := range categories {
// 		category := CategotyrequestToCategotyCore(v)
// 		listCategory = append(listCategory, category)
// 	}

// 	return listCategory
// }