package repository

import (
	"errors"
	"mime/multipart"
	"recything/features/article/entity"
	"recything/features/article/model"
	trashcategory "recything/features/trash_category/entity"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/storage"

	"gorm.io/gorm"
)

type articleRepository struct {
	db            *gorm.DB
	trashcategory trashcategory.TrashCategoryRepositoryInterface
}

func NewArticleRepository(db *gorm.DB, trashcategory trashcategory.TrashCategoryRepositoryInterface) entity.ArticleRepositoryInterface {
	return &articleRepository{
		db:            db,
		trashcategory: trashcategory,
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
	input := entity.ArticleCoreToArticleModel(articleInput)
	var articleData model.Article

	check := article.db.Where("id = ?", idArticle).First(&articleData)
	if check.Error != nil {
		return entity.ArticleCore{}, check.Error
	}

	if image != nil {
		imageURL, errUpload := storage.UploadThumbnail(image)
		if errUpload != nil {
			return entity.ArticleCore{}, errUpload
		}
		articleData.Image = imageURL

	} else {
		input.Image = articleData.Image
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
func (article *articleRepository) GetAllArticle(page, limit int, search string) ([]entity.ArticleCore, pagination.PageInfo, int, error) {
	var articleData []model.Article

	offset := (page - 1) * limit
	query := article.db.Model(&model.Article{}).Preload("Categories")

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	var totalCount int64
	tx := query.Count(&totalCount).Find(&articleData)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	query = query.Offset(offset).Limit(limit)

	// txData := article.db.Preload("Categories").Find(&articleData)
	// if txData.Error != nil {
	// 	return nil, pagination.PageInfo{}, txData.Error
	// }

	tx = query.Find(&articleData)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	dataResponse := entity.ListArticleModelToArticleCore(articleData)
	pageInfo := pagination.CalculateData(int(totalCount), limit, page)

	return dataResponse, pageInfo, int(totalCount), nil
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

	for i, categoryId := range articleInput.Category_id {
		_, tx := article.trashcategory.GetById(categoryId)
		if tx != nil {
			txOuter.Rollback()
			return entity.ArticleCore{}, errors.New(constanta.ERROR_RECORD_NOT_FOUND)
		}

		// Check if the category exists
		// var categoryCount int64
		// if err := txOuter.Model(&model.ArticleTrashCategory{}).Where("article_id = ?", categoryId).Count(&categoryCount).Error; err != nil {
		// 	txOuter.Rollback()
		// 	return entity.ArticleCore{}, err
		// }

		//  // If the category doesn't exist, return an error
		//  if categoryCount == 0 {
		// 	 txOuter.Rollback()
		// 	 return entity.ArticleCore{}, errors.New("kategori tidak ditemukan")
		//  }

		categories := new(model.ArticleTrashCategory)
		categories.ArticleID = articleCreated.ID
		categories.TrashCategoryID = categoryId

		for j := i + 1; j < len(articleInput.Category_id); j++ {
			if categoryId == articleInput.Category_id[j] {
				return entity.ArticleCore{}, errors.New("error : kategori tidak boleh sama")
			}
		}
		txInner := txOuter.Create(&categories)
		if txInner.Error != nil {
			txOuter.Rollback()
			return entity.ArticleCore{}, txInner.Error
		}

	}

	txOuter.Commit()

	return articleCreated, nil
}

// PostLike implements entity.ArticleRepositoryInterface.
func (article *articleRepository) PostLike(idArticle string) error {
	var articleData model.Article

	check := article.db.Where("id = ?", idArticle).First(&articleData)
	if check.Error != nil {
		return check.Error
	}

	articleData.Like += 1

	tx := article.db.Updates(articleData)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// PostShare implements entity.ArticleRepositoryInterface.
func (article *articleRepository) PostShare(idArticle string) error {
	var articleData model.Article

	check := article.db.Where("id = ?", idArticle).First(&articleData)
	if check.Error != nil {
		return check.Error
	}

	articleData.Share += 1

	tx := article.db.Updates(articleData)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// GetPopularArticle implements entity.ArticleRepositoryInterface.
func (article *articleRepository) GetPopularArticle(search string) ([]entity.ArticleCore, error) {
	var articleData []model.Article

	tx := article.db.Order("`like` DESC").Limit(10).Find(&articleData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataResponse := entity.ListArticleModelToArticleCore(articleData)

	return dataResponse, nil
}

// GetArticleByCategory implements entity.ArticleRepositoryInterface.
func (article *articleRepository) GetArticleByCategory(idCategory string) ([]entity.ArticleCore, error) {
	var articleData []model.Article

	data := article.db.
		Table("articles").
		Select("articles.*").
		Joins("INNER JOIN article_trash_categories ON articles.id = article_trash_categories.article_id").
		Where("article_trash_categories.trash_category_id = ?", idCategory).
		Preload("Categories").
		Find(&articleData)

	if data.Error != nil {
		return nil, data.Error
	}

	articles := entity.ListArticleModelToArticleCore(articleData)

	return articles, nil
}