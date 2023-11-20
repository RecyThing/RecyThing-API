package entity

import (
	"recything/features/article/model"
)

func ArticleModelToArticleCore(article model.Article) ArticleCore {
	articleCore := ArticleCore{
		ID:        article.Id,
		Title:     article.Title,
		Image:     article.Image,
		Category:  article.Category,
		Content:   article.Content,
		Like:      article.Like,
		Share:     article.Share,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
	// category := entity.ListModelTrashCategoryToCoreTrashCategory(article.Category)
	// articleCore.Category = category
	return articleCore
}

func ArticleCoreToArticleModel(article ArticleCore) model.Article {
	articleModel := model.Article{
		Id:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Category:  article.Category,
		Content:   article.Content,
		Like:      article.Like,
		Share:     article.Share,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
	// category := entity.ListCoreTrashCategoryToModelTrashCategory(article.Category)
	// articleModel.Category = category
	return articleModel
}
