package entity

import (
	"recything/features/article/model"
)

// func CategoryModelToCategoryCore(category model.ArticleTrashCategory) CategoryCore {
// 	return CategoryCore{
// 		ArticleID:       category.ArticleID,
// 		TrashCategoryID: category.TrashCategoryID,
// 	}
// }

// func ListCategoryModelToCategoryCore(category []model.ArticleTrashCategory) []CategoryCore {
// 	coreCategory := []CategoryCore{}
// 	for _, v := range category {
// 		category := CategoryModelToCategoryCore(v)
// 		coreCategory = append(coreCategory, category)
// 	}
// 	return coreCategory
// }

func ArticleModelToArticleCore(article model.Article) ArticleCore {
	articleCore := ArticleCore{
		ID:          article.Id,
		Title:       article.Title,
		Image:       article.Image,
		Content:     article.Content,
		Like:        article.Like,
		Share:       article.Share,
		// Category_id: article.Categories,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}
	// category := ListCategoryModelToCategoryCore(article.Categories)
	// articleCore.Category_id = category
	return articleCore
}

// func CategoryCoreToCategoryModel(category CategoryCore) model.ArticleTrashCategory {
// 	return model.ArticleTrashCategory{
// 		ArticleID: category.ArticleID,
// 		TrashCategoryID: category.TrashCategoryID,
// 	}
// }

// func ListCategoryCoreToCategoryModel(category []CategoryCore) []model.ArticleTrashCategory {
// 	coreCategorys := []model.ArticleTrashCategory{}
// 	for _, v := range category {
// 		categorys := CategoryCoreToCategoryModel(v)
// 		coreCategorys = append(coreCategorys, categorys)
// 	}
// 	return coreCategorys
// }

func ArticleCoreToArticleModel(article ArticleCore) model.Article {
	articleModel := model.Article{
		Id:         article.ID,
		Title:      article.Title,
		Image:      article.Image,
		Content:    article.Content,
		Like:       article.Like,
		Share:      article.Share,
		// Categories: article.Category_id,
		CreatedAt:  article.CreatedAt,
		UpdatedAt:  article.UpdatedAt,
	}
	// category := ListCategoryCoreToCategoryModel(article.Category_id)
	// articleModel.Categories = category
	return articleModel
}
