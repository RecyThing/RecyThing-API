package repository

import (
	"errors"
	"recything/features/mission/entity"
	"recything/utils/constanta"
	"recything/utils/validation"

	"gorm.io/gorm"
)

type MissionRepository struct {
	db *gorm.DB
}

func NewMissionRepository(db *gorm.DB) entity.MissionRepositoryInterface {
	return &MissionRepository{
		db: db,
	}
}

// Create implements entity.MissionRepositoryInterface.
func (mr *MissionRepository) Create(input entity.Mission) error {
	data := entity.MissionCoreToMissionModel(input)

	tx := mr.db.Create(&data)
	if tx.Error != nil {
		if validation.IsDuplicateError(tx.Error) {
			return errors.New(constanta.ERROR_DATA_EXIST)
		}
		return tx.Error
	}
	return nil
}

// func (tc *trashCategoryRepository) Create(data entity.TrashCategoryCore) error {
// 	input := entity.CoreTrashCategoryToModelTrashCategory(data)

// 	tx := tc.db.Create(&input)
// 	if tx.Error != nil {
// 		if validation.IsDuplicateError(tx.Error) {
// 			return errors.New(constanta.ERROR_DATA_EXIST)
// 		}
// 		return tx.Error
// 	}
// 	return nil
// }

// func (tc *trashCategoryRepository) FindAll(page, limit int, trashType string) ([]entity.TrashCategoryCore, pagination.PageInfo, int, error) {
// 	dataTrashCategories := []model.TrashCategory{}
// 	offsetInt := (page - 1) * limit

// 	totalCount, err := tc.GetCount(trashType)
// 	if err != nil {
// 		return nil, pagination.PageInfo{}, 0, err
// 	}

// 	paginationQuery:= tc.db.Limit(limit).Offset(offsetInt)
// 	if trashType == "" {
// 		tx := paginationQuery.Find(&dataTrashCategories)
// 		if tx.Error != nil {
// 			return nil, pagination.PageInfo{}, 0, tx.Error
// 		}
// 	}

// 	if trashType != "" {
// 		tx := paginationQuery.Where("trash_type LIKE ?", "%"+trashType+"%").Find(&dataTrashCategories)
// 		if tx.Error != nil {
// 			return nil, pagination.PageInfo{}, 0, tx.Error
// 		}
// 	}

// 	result := entity.ListModelTrashCategoryToCoreTrashCategory(dataTrashCategories)
// 	paginationInfo := pagination.CalculateData(totalCount, limit, page)
// 	return result, paginationInfo, totalCount, nil
// }

// func (tc *trashCategoryRepository) GetCount(trashType string) (int, error) {
// 	var totalCount int64
// 	model:=tc.db.Model(&model.TrashCategory{})
// 	if trashType == "" {
// 		tx :=model.Count(&totalCount)
// 		if tx.Error != nil {
// 			return 0, tx.Error
// 		}
// 	}

// 	if trashType != "" {
// 		tx := model.Where("trash_type LIKE ?", "%"+trashType+"%").Count(&totalCount)
// 		if tx.Error != nil {
// 			return 0, tx.Error
// 		}

// 	}
// 	return int(totalCount), nil
// }

// func (tc *trashCategoryRepository) GetById(idTrash string) (entity.TrashCategoryCore, error) {

// 	dataTrashCategories := model.TrashCategory{}
// 	tx := tc.db.Where("id = ?", idTrash).First(&dataTrashCategories)
// 	if tx.Error != nil {

// 		if tx.RowsAffected == 0 {
// 			return entity.TrashCategoryCore{}, errors.New(constanta.ERROR_DATA_ID)
// 		}

// 		return entity.TrashCategoryCore{}, tx.Error
// 	}

// 	result := entity.ModelTrashCategoryToCoreTrashCategory(dataTrashCategories)
// 	return result, nil
// }

// func (tc *trashCategoryRepository) Update(idTrash string, data entity.TrashCategoryCore) (entity.TrashCategoryCore, error) {
// 	dataTrashCategories := entity.CoreTrashCategoryToModelTrashCategory(data)

// 	tx := tc.db.Where("id = ?", idTrash).Updates(&dataTrashCategories)
// 	if tx.Error != nil {
// 		return entity.TrashCategoryCore{}, tx.Error
// 	}

// 	if tx.RowsAffected == 0 {
// 		return entity.TrashCategoryCore{}, errors.New(constanta.ERROR_DATA_ID)
// 	}

// 	result := entity.ModelTrashCategoryToCoreTrashCategory(dataTrashCategories)
// 	return result, nil
// }

// func (tc *trashCategoryRepository) Delete(idTrash string) error {
// 	data := model.TrashCategory{}

// 	tx := tc.db.Where("id = ?", idTrash).Delete(&data)
// 	if tx.Error != nil {
// 		return tx.Error
// 	}
// 	if tx.RowsAffected == 0 {
// 		return errors.New(constanta.ERROR_DATA_ID)
// 	}

// 	return nil
// }
