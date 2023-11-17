package request

import "recything/features/article/entity"

func ArticleRequestToArticleCore(article ArticleRequest) entity.ArticleCore{
	return entity.ArticleCore{
		Title: article.Title,
		Image: article.Image,
		Content: article.Content,
		Category: article.Category,
	}
}

