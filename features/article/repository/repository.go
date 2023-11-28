package repository

import (
	"errors"
	"mime/multipart"
	"recything/features/article/entity"
	"recything/features/article/model"
	"recything/utils/pagination"
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

	categoryId := model.ArticleTrashCategory{}
	categoryDel := article.db.Where("article_id = ?", id).Delete(&categoryId)
	if categoryDel.Error != nil {
		return categoryDel.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("tidak ada data yang dihapus")
	}

	return nil
}

// GetSpecificArticle implements entity.ArticleRepositoryInterface.
func (article *articleRepository) GetSpecificArticle(idArticle string) (entity.ArticleCore, error) {
	articleData := model.Article{}

	tx := article.db.Preload("Categories").Where("id = ?", idArticle).First(&articleData)
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

	tx := article.db.Begin()

	// Hapus kategori yang terkait dengan artikel
	categoryId := model.ArticleTrashCategory{}
	categoryDel := tx.Where("article_id = ?", idArticle).Delete(&categoryId)
	if categoryDel.Error != nil {
		return entity.ArticleCore{}, categoryDel.Error
	}

	if err := tx.Save(&articleData).Error; err != nil {
		tx.Rollback()
		return entity.ArticleCore{}, err
	}

	// Tambahkan kategori yang baru
	for _, categoryId := range articleInput.Category_id {
		categories := new(model.ArticleTrashCategory)
		categories.ArticleID = idArticle
		categories.TrashCategoryID = categoryId

		txLink := tx.Create(&categories)
		if txLink.Error != nil {
			tx.Rollback()
			return entity.ArticleCore{}, errors.New("kategori tidak ditemukan")
		}
	}

	tx.Commit()

	articleUpdate := entity.ArticleModelToArticleCore(articleData)

	return articleUpdate, nil
}

// GetAllArticle implements entity.ArticleRepositoryInterface.
func (article *articleRepository) GetAllArticle(page, limit int, tittle string) ([]entity.ArticleCore, pagination.PageInfo, error) {
	var articleData []model.Article

	offset := (page - 1) * limit
	query := article.db.Model(&model.Article{}).Preload("Categories")

	if tittle != "" {
		query = query.Where("title LIKE ?", "%"+tittle+"%")
	}

	var totalCount int64
	tx := query.Count(&totalCount).Find(&articleData)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, tx.Error
	}

	query = query.Offset(offset).Limit(limit)

	// txData := article.db.Preload("Categories").Find(&articleData)
	// if txData.Error != nil {
	// 	return nil, pagination.PageInfo{}, txData.Error
	// }

	tx = query.Find(&articleData)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, tx.Error
	}

	mapData := make([]entity.ArticleCore, len(articleData))
	for i, value := range articleData {
		mapData[i] = entity.ArticleModelToArticleCore(value)
	}

	pageInfo := pagination.CalculateData(int(totalCount), limit, page)

	return mapData, pageInfo, nil

}

// CreateArticle implements entity.ArticleRepositoryInterface.
func (article *articleRepository) CreateArticle(articleInput entity.ArticleCore, image *multipart.FileHeader) (entity.ArticleCore, error) {
	articleData := entity.ArticleCoreToArticleModel(articleInput)

	imageURL, uploadErr := storage.UploadThumbnail(image)
	if uploadErr != nil {
		return entity.ArticleCore{}, uploadErr
	}

	articleData.Image = imageURL

	txOuter := article.db.Begin()

	if err := txOuter.Save(&articleData).Error; err != nil {
		txOuter.Rollback()
		return entity.ArticleCore{}, err
	}

	articleCreated := entity.ArticleModelToArticleCore(articleData)

	for _, categoryId := range articleInput.Category_id {

		 // Check if the category exists
		 var categoryCount int64
		 if err := txOuter.Model(&model.ArticleTrashCategory{}).Where("article_id = ?", categoryId).Count(&categoryCount).Error; err != nil {
			 txOuter.Rollback()
			 return entity.ArticleCore{}, err
		 }
	 
		//  // If the category doesn't exist, return an error
		//  if categoryCount == 0 {
		// 	 txOuter.Rollback()
		// 	 return entity.ArticleCore{}, errors.New("kategori tidak ditemukan")
		//  }

		categories := new(model.ArticleTrashCategory)
		categories.ArticleID = articleCreated.ID
		categories.TrashCategoryID = categoryId

		txInner := txOuter.Create(&categories)
		if txInner.Error != nil {
			txOuter.Rollback()
			return entity.ArticleCore{}, errors.New("kategori tidak ditemukan")
		}

	}

	txOuter.Commit()

	return articleCreated, nil
}
