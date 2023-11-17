package entity

import "mime/multipart"

type ArticleRepositoryInterface interface {
	CreateArticle(articleInput ArticleCore, image *multipart.FileHeader) (ArticleCore, error)
	GetAllArticle()([]ArticleCore, error)
	GetSpecificArticle(idArticle string) (ArticleCore, error)
	UpdateArticle(idArticle string, articleInput ArticleCore, image *multipart.FileHeader) (ArticleCore, error)
	DeleteArticle(id string) (error)
}

type ArticleServiceInterface interface {
	CreateArticle(articleInput ArticleCore, image *multipart.FileHeader) (ArticleCore, error)
	GetAllArticle()([]ArticleCore, error)
	GetSpecificArticle(idArticle string) (ArticleCore, error)
	UpdateArticle(idArticle string, articleInput ArticleCore, image *multipart.FileHeader) (ArticleCore, error)
	DeleteArticle(id string) (error)
}
