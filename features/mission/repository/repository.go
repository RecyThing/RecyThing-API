package repository

import (
	"errors"
	"fmt"
	"log"
	"recything/features/mission/entity"
	"recything/features/mission/model"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/validation"
	"time"

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
	log.Println(" input repo : ", input)
	data := entity.MissionCoreToMissionModel(input)
	log.Println("data repo ")
	tx := mr.db.Create(&data)
	if tx.Error != nil {
		if validation.IsDuplicateError(tx.Error) {
			return errors.New(constanta.ERROR_DATA_EXIST)
		}
		return tx.Error
	}
	return nil
}

func (mr *MissionRepository) FindAllMission(page, limit int, search, status string) ([]entity.Mission, pagination.PageInfo, int, error) {
	data := []model.Mission{}
	offsetInt := (page - 1) * limit
	paginationQuery := mr.db.Limit(limit).Offset(offsetInt)

	totalCount, err := mr.GetCount(status, search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	fmt.Println("status : ", status)

	if status != "" {
		tx := paginationQuery.Where("status LIKE ?", "%"+status+"%").Preload("MissionStages").Find(&data)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	if search != "" {
		tx := paginationQuery.Where("title LIKE?", "%"+search+"%").Preload("MissionStages").Find(&data)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	if search == "" || status == "" {
		tx := paginationQuery.Preload("MissionStages").Find(&data)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	for _, v := range data {
		endDateValid, err := time.Parse("2006-01-02", v.EndDate)
		if err != nil {
			return nil, pagination.PageInfo{}, 0, err
		}
		currentTime := time.Now().Truncate(24 * time.Hour)
		if endDateValid.Before(currentTime) {
			v.Status = constanta.OVERDUE
		} else {
			v.Status = constanta.ACTIVE
		}

		err = mr.db.Model(&v).Where("id = ?", v.ID).Update("status", &v.Status).Error
		if err != nil {
			return nil, pagination.PageInfo{}, 0, err
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

func (mr *MissionRepository) GetCount(filter, search string) (int, error) {
	var totalCount int64
	model := mr.db.Model(&model.Mission{})
	if filter == "" || search == "" {
		tx := model.Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}
	}

	if search != "" {
		tx := model.Where("title LIKE ?", "%"+search+"%").Count(&totalCount)
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

func (mr *MissionRepository) SaveChangesStatusMission(data entity.Mission) error {
	mission := entity.MissionCoreToMissionModel(data)
	err := mr.db.Model(&mission).Where("id = ?", data.ID).Update("status", &data.Status).Error
	if err != nil {
		return err
	}

	return nil
}

func (mr *MissionRepository) UpdateMission(missionID string, data entity.Mission) error {
	dataMission := entity.MissionCoreToMissionModel(data)
	tx := mr.db.Where("id = ?", missionID).First(&dataMission)
	if tx.Error != nil {
		return tx.Error
	}

	dataMission.Title = data.Title

	tx = mr.db.Where("id = ?", missionID).Save(&dataMission)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_ID)
	}
	return nil
}

func (mr *MissionRepository) UpdateMissionStage(MissionStageID string, data entity.MissionStage) error {
	missionStage := entity.MissionStagesCoreToMissionStagesModel(data)
	tx := mr.db.Where("id = ?", MissionStageID).Updates(&missionStage)
	if tx.Error != nil {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}

func (mr *MissionRepository) GetById(missionID string) (entity.Mission, error) {
	mission := model.Mission{}
	tx := mr.db.Take(&mission, "id = ?", missionID)
	if tx.Error != nil {
		return entity.Mission{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.Mission{}, errors.New(constanta.ERROR_DATA_ID)
	}

	result := entity.MissionModelToMissionCore(mission)
	return result, nil

}
