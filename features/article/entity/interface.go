package entity

import (
	"mime/multipart"
	"recything/utils/pagination"
)

type ArticleRepositoryInterface interface {
	CreateArticle(articleInput ArticleCore, image *multipart.FileHeader) (ArticleCore, error)
	GetAllArticle(page, limit int, search string) ([]ArticleCore, pagination.PageInfo, error)
	GetSpecificArticle(idArticle string) (ArticleCore, error)
	UpdateArticle(idArticle string, articleInput ArticleCore, image *multipart.FileHeader) (ArticleCore, error)
	DeleteArticle(id string) error
}

type ArticleServiceInterface interface {
	CreateArticle(articleInput ArticleCore, image *multipart.FileHeader) (ArticleCore, error)
	GetAllArticle(page, limit int, search string) ([]ArticleCore, pagination.PageInfo, error)
	GetSpecificArticle(idArticle string) (ArticleCore, error)
	UpdateArticle(idArticle string, articleInput ArticleCore, image *multipart.FileHeader) (ArticleCore, error)
	DeleteArticle(id string) error
}
