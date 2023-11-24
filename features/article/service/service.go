package service

import (
	"errors"
	"fmt"
	"mime/multipart"
	"recything/features/article/entity"
)

type articleService struct {
	ArticleRepository entity.ArticleRepositoryInterface
}

func NewArticleService(article entity.ArticleRepositoryInterface) entity.ArticleServiceInterface {
	return &articleService{
		ArticleRepository: article,
	}
}

// DeleteArticle implements entity.ArticleServiceInterface.
func (ac *articleService) DeleteArticle(id string) error {
	if id == "" {
		return errors.New("id artikel tidak ditemukan")
	}

	errArticle := ac.ArticleRepository.DeleteArticle(id)
	if errArticle != nil {
		return errors.New("gagal menghapus artikel " + errArticle.Error())
	}

	return nil
}

// GetSpecificArticle implements entity.ArticleServiceInterface.
func (ac *articleService) GetSpecificArticle(idArticle string) (entity.ArticleCore, error) {
	if idArticle == "" {
		return entity.ArticleCore{}, errors.New("id tidak cocok")
	}

	articleData, err := ac.ArticleRepository.GetSpecificArticle(idArticle)
	if err != nil {
		fmt.Println("service", err)
		return entity.ArticleCore{}, errors.New("gagal membaca data")
	}

	return articleData, nil
}

// UpdateArticle implements entity.ArticleServiceInterface.
func (article *articleService) UpdateArticle(idArticle string, articleInput entity.ArticleCore, image *multipart.FileHeader) (entity.ArticleCore, error) {

	if idArticle == "" {
		return entity.ArticleCore{}, errors.New("id tidak ditemukan")
	}

	if articleInput.Title == "" || articleInput.Content == "" {
		return entity.ArticleCore{}, errors.New("artikel tidak boleh kosong")
	}

	if len(articleInput.Category_id) == 0 {
		return entity.ArticleCore{}, errors.New("kategori tidak boleh kosong")
	}

	if image != nil && image.Size > 5*1024*1024 {
		return entity.ArticleCore{}, errors.New("ukuran file tidak boleh lebih dari 5 MB")
	}

	articleUpdate, errinsert := article.ArticleRepository.UpdateArticle(idArticle, articleInput, image)
	if errinsert != nil {
		return entity.ArticleCore{}, errinsert
	}

	return articleUpdate, nil
}

// GetAllArticle implements entity.ArticleServiceInterface.
func (ac *articleService) GetAllArticle() ([]entity.ArticleCore, error) {
	article, err := ac.ArticleRepository.GetAllArticle()
	if err != nil {
		return []entity.ArticleCore{}, errors.New("gagal mendapatkan artikel")
	}

	return article, nil
}

// CreateArticle implements entity.ArticleServiceInterface.
func (article *articleService) CreateArticle(articleInput entity.ArticleCore, image *multipart.FileHeader) (entity.ArticleCore, error) {

	if articleInput.Title == "" || articleInput.Content == "" {
		return entity.ArticleCore{}, errors.New("judul dan konten artikel tidak boleh kosong")
	}

	if len(articleInput.Category_id) == 0 {
		return entity.ArticleCore{}, errors.New("kategori tidak boleh kosong")
	}

	if image != nil && image.Size > 5*1024*1024 {
		return entity.ArticleCore{}, errors.New("ukuran file tidak boleh lebih dari 5 MB")
	}

	articleCreate, errinsert := article.ArticleRepository.CreateArticle(articleInput, image)
	if errinsert != nil {
		return entity.ArticleCore{}, errinsert
	}

	return articleCreate, nil
}
