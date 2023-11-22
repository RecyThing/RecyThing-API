package response

import (
	"recything/features/article/entity"
)

// func CategoryCoreToCategoryResponse(category entity.CategoryCore) TrashCategoryResponse {
// 	return TrashCategoryResponse{
// 		ArticleID: category.ArticleID,
// 		TrashCategoryID: category.TrashCategoryID,
// 	}
// }

// func ListCategoryCoreToCategoryResponse(categories []entity.CategoryCore) []TrashCategoryResponse {
// 	ResponseCategory := []TrashCategoryResponse{}
// 	for _, v := range categories {
// 		category := CategoryCoreToCategoryResponse(v)
// 		ResponseCategory = append(ResponseCategory, category)
// 	}
// 	return ResponseCategory
// }

func ArticleCoreToArticleResponse(article entity.ArticleCore) ArticleCreateResponse {
	articleResp := ArticleCreateResponse{
		Id:          article.ID,
		Title:       article.Title,
		Image:       article.Image,
		Content:     article.Content,
		Like:        article.Like,
		Share:       article.Share,
		Category_id: article.Category_id,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}
	// category := ListCategoryCoreToCategoryResponse(article.Category_id)
	// articleResp.Category_id = category
	return articleResp
}
