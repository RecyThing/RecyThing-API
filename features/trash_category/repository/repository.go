package repository

import (
	"errors"
	"math"

	"recything/features/trash_category/entity"
	"recything/features/trash_category/model"
	"recything/utils/constanta"
	"strconv"

	"gorm.io/gorm"
)

type trashCategoryRepository struct {
	db *gorm.DB
}

func NewTrashCategiryRepository(db *gorm.DB) entity.TrashCategoryRepositoryInterface {
	return &trashCategoryRepository{
		db: db,
	}
}

func (tc *trashCategoryRepository) Create(data entity.TrashCategoryCore) (entity.TrashCategoryCore, error) {
	input := entity.CoreTrashCategoryToModelTrashCategory(data)

	tx := tc.db.Create(&input)
	if tx.Error != nil {
		return entity.TrashCategoryCore{}, tx.Error
	}

	result := entity.ModelTrashCategoryToCoreTrashCategory(input)
	return result, nil
}

// func (tc *trashCategoryRepository) GetAll() ([]entity.TrashCategoryCore, error) {
// 	dataTrashCategories := []model.TrashCategory{}

// 	tx := tc.db.Find(&dataTrashCategories)
// 	if tx.Error != nil {
// 		return []entity.TrashCategoryCore{}, tx.Error
// 	}

// 	result := entity.ListModelTrashCategoryToCoreTrashCategory(dataTrashCategories)
// 	return result, nil
// }

func (tc *trashCategoryRepository) GetAll(page string, limit string) ([]entity.TrashCategoryCore, entity.PagnationInfo, error) {
	

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return nil, entity.PagnationInfo{}, err
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, entity.PagnationInfo{}, err
	}

	if limit == "" {
		limitInt = 10
	}

	if page == "" {
		pageInt = 10
	}

	
	offsetInt := (pageInt - 1) * limitInt

	
	dataTrashCategories := []model.TrashCategory{}

	tx := tc.db.Limit(limitInt).Offset(offsetInt).Find(&dataTrashCategories)
	if tx.Error != nil {
		return nil, entity.PagnationInfo{}, tx.Error
	}

	
	result := entity.ListModelTrashCategoryToCoreTrashCategory(dataTrashCategories)

	
	var totalCount int64
	err = tc.db.Model(&model.TrashCategory{}).Count(&totalCount).Error
	if err != nil {
		return nil, entity.PagnationInfo{}, err
	}

	lastPage := int(math.Ceil(float64(totalCount) / float64(limitInt)))

	paginationInfo := entity.PagnationInfo{
		Limit:       limitInt,
		CurrentPage: pageInt,
		LastPage:    lastPage,
	}

	// Combine result and pagination info
	return result, paginationInfo, nil
}

func (tc *trashCategoryRepository) GetById(idTrash string) (entity.TrashCategoryCore, error) {
	dataTrashCategories := model.TrashCategory{}

	tx := tc.db.Where("id = ?", idTrash).First(&dataTrashCategories)
	if tx.Error != nil {
		return entity.TrashCategoryCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.TrashCategoryCore{}, errors.New(constanta.ERROR_DATA_ID)
	}

	result := entity.ModelTrashCategoryToCoreTrashCategory(dataTrashCategories)
	return result, nil
}

func (tc *trashCategoryRepository) Update(idTrash string, data entity.TrashCategoryCore) (entity.TrashCategoryCore, error) {
	dataTrashCategories := entity.CoreTrashCategoryToModelTrashCategory(data)

	tx := tc.db.Where("id = ?", idTrash).Updates(&dataTrashCategories)
	if tx.Error != nil {
		return entity.TrashCategoryCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.TrashCategoryCore{}, errors.New(constanta.ERROR_DATA_ID)
	}

	result := entity.ModelTrashCategoryToCoreTrashCategory(dataTrashCategories)
	return result, nil
}

func (tc *trashCategoryRepository) Delete(idTrash string) error {
	data := model.TrashCategory{}

	tx := tc.db.Where("id = ?", idTrash).Delete(&data)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}
