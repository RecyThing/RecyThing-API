package request

import (
	"recything/features/article/entity"
)

func ArticleRequestToArticleCore(article ArticleRequest) entity.ArticleCore {
	articleReq := entity.ArticleCore{
		Title:    article.Title,
		Image:    article.Image,
		Content:  article.Content,
		Category: article.Category,
	}
	// category := request.ListTrashCategoryRequestToTrashCategoryCore(article.Category)
	// articleReq.Category = category
	return articleReq
}
