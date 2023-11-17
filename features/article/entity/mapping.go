package entity

import "recything/features/article/model"

func ArticleModelToArticleCore(article model.Article) ArticleCore {
	return ArticleCore{
		ID:        article.Id,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		Category:  article.Category,
		Like:      article.Like,
		Share:     article.Share,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
}

func ArticleCoreToArticleModel(article ArticleCore) model.Article {
	return model.Article{
		Id:        article.ID,
		Title:     article.Title,
		Image:     article.Image,
		Content:   article.Content,
		Category:  article.Category,
		Like:      article.Like,
		Share:     article.Share,
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
}
