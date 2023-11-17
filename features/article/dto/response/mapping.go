package response

import "recything/features/article/entity"

func ArticleCoreToArticleResponse(article entity.ArticleCore) ArticleCreateResponse {
	return ArticleCreateResponse{
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
