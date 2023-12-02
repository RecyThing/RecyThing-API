package repository

import (
	"errors"
	"recything/features/mission/entity"
	"recything/features/mission/model"
	"recything/utils/constanta"
	"recything/utils/pagination"
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
func (mr *MissionRepository) CreateMission(input entity.Mission) error {
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
func (mr *MissionRepository) CreateMissionStages(input []entity.MissionStage) error {
	data := entity.ListMissionStagesCoreToMissionStagesModel(input)

	tx := mr.db.Create(&data)
	if tx.Error != nil {
		if validation.IsDuplicateError(tx.Error) {
			return errors.New(constanta.ERROR_DATA_EXIST)
		}
		return tx.Error
	}
	return nil
}

func (mr *MissionRepository) FindAll(page, limit int, filter string) ([]entity.Mission, pagination.PageInfo, int, error) {
	data := []model.Mission{}
	offsetInt := (page - 1) * limit

	totalCount, err := mr.GetCount(filter)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	if filter == "" {
		tx := mr.db.Limit(limit).Offset(offsetInt).Preload("MissionStages").Find(&data)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	if filter != "" {
		tx := mr.db.Limit(limit).Offset(offsetInt).Where("status LIKE ?", "%"+filter+"%").Preload("MissionStages").Find(&data)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	result := entity.ListMissionModelToMissionCore(data)
	paginationInfo := pagination.CalculateData(totalCount, limit, page)
	return result, paginationInfo, totalCount, nil
}

func (mr *MissionRepository) GetAdminIDbyMissionID(missionID string) (string, error) {
	mission := model.Mission{}
	err := mr.db.Take(&mission, "admin_id = ?", missionID).Error
	if err != nil {
		return mission.AdminID, err
	}
	return mission.AdminID, nil
}

func (mr *MissionRepository) GetCount(filter string) (int, error) {
	var totalCount int64
	model := mr.db.Model(&model.Mission{})
	if filter == "" {
		tx := model.Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}
	}

	if filter != "" {
		tx := model.Where("status LIKE ?", "%"+filter+"%").Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}

	}
	return int(totalCount), nil
}

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
