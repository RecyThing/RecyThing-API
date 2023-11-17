package repository

import (
	"errors"
	"mime/multipart"
	"recything/features/article/entity"
	"recything/features/article/model"
	"recything/utils/storage"

	"gorm.io/gorm"
)

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) entity.ArticleRepositoryInterface {
	return &articleRepository{
		db: db,
	}
}

// DeleteArticle implements entity.ArticleRepositoryInterface.
func (article *articleRepository) DeleteArticle(id string) error {
	checkId := model.Article{}

	tx := article.db.Where("id = ?", id).Delete(&checkId)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("tidak ada data yang dihapus")
	}

	return nil
}

// GetSpecificArticle implements entity.ArticleRepositoryInterface.
func (article *articleRepository) GetSpecificArticle(idArticle string) (entity.ArticleCore, error) {
	articleData := model.Article{}

	tx := article.db.Where("id = ?", idArticle).Preload("image").First(&articleData)
	if tx.Error != nil {
		return entity.ArticleCore{}, tx.Error
	}

	dataResponse := entity.ArticleModelToArticleCore(articleData)
	return dataResponse, nil
}

// UpdateArticle implements entity.ArticleRepositoryInterface.
func (article *articleRepository) UpdateArticle(idArticle string, articleInput entity.ArticleCore, image *multipart.FileHeader) (entity.ArticleCore, error) {
	var articleData model.Article

	check := article.db.Where("id = ?", idArticle).First(&articleData)
	if check.Error != nil {
		return entity.ArticleCore{}, check.Error
	}

	if image != nil {
		imageURL, uploadErr := storage.UploadThumbnail(image)
		if uploadErr != nil {
			return entity.ArticleCore{}, uploadErr
		}
		articleData.Image = imageURL
	}

	articleData.Title = articleInput.Title
	articleData.Content = articleInput.Content
	articleData.Category = articleInput.Category

	tx := article.db.Updates(&articleData)
	if tx.Error != nil {
		return entity.ArticleCore{}, tx.Error
	}

	articleUpdate := entity.ArticleModelToArticleCore(articleData)

	return articleUpdate, nil
}

// GetAllArticle implements entity.ArticleRepositoryInterface.
func (article *articleRepository) GetAllArticle() ([]entity.ArticleCore, error) {
	var articleData []model.Article

	tx := article.db.Find(&articleData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	mapData := make([]entity.ArticleCore, len(articleData))
	for i, value := range articleData {
		mapData[i] = entity.ArticleModelToArticleCore(value)
	}

	// var datas []model.Article

	// for i, v := range articleData {
	// 	datas[i] = entity.ArticleCoreToArticleModel(v)
	// }

	return mapData, nil

}

// CreateArticle implements entity.ArticleRepositoryInterface.
func (article *articleRepository) CreateArticle(articleInput entity.ArticleCore, image *multipart.FileHeader) (entity.ArticleCore, error) {
	articleData := entity.ArticleCoreToArticleModel(articleInput)

	imageURL, uploadErr := storage.UploadThumbnail(image)
	if uploadErr != nil {
		return entity.ArticleCore{}, uploadErr
	}

	articleData.Image = imageURL
	tx := article.db.Create(&articleData)
	if tx.Error != nil {
		return entity.ArticleCore{}, tx.Error
	}

	articleCreated := entity.ArticleModelToArticleCore(articleData)

	return articleCreated, nil
}
