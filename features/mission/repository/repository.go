package repository

import (
	"errors"
	"fmt"
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
	getMission := model.Mission{}
	tx := mr.db.Where("id = ?", missionID).First(&getMission)
	if tx.Error != nil {
		return errors.New("missi tidak ditemukan")
	}

	tx = mr.db.Where("id = ?", missionID).Updates(&dataMission)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (mr *MissionRepository) UpdateMissionStage(missionStageID string, data entity.MissionStage) error {

	tx := mr.db.Where("id = ?", missionStageID).Take(&model.MissionStage{})
	if tx.Error != nil {
		return errors.New("missi tidak ditemukan")
	}

	missionStage := entity.MissionStagesCoreToMissionStagesModel(data)
	tx = mr.db.Where("id = ?", missionStageID).Updates(&missionStage)
	if tx.Error != nil {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}

func (mr *MissionRepository) GetImageURL(missionID string) (string, error) {
	mission := model.Mission{}
	err := mr.db.Take(&mission, missionID).Error
	if err != nil {
		return "", err
	}

	return mission.MissionImage, nil
}

// Claimed Mission
func (mr *MissionRepository) ClaimMission(userID string, data entity.ClaimedMission) error {
	input := entity.ClaimedCoreToClaimedMissionModel(data)

	_, err := mr.GetById(data.MissionID)
	if err != nil {
		return err
	}

	errFind := mr.FindClaimed(userID, data.MissionID)
	if errFind != nil {
		return errors.New(errFind.Error())
	}

	tx := mr.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (mr *MissionRepository) FindClaimed(userID, missionID string) error {
	dataClaimed := model.ClaimedMission{}

	tx := mr.db.Where("user_id = ? AND mission_id = ? AND claimed = 1", userID, missionID).Find(&dataClaimed)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected != 0 {
		return errors.New("error : mission sudah di klaim")
	}

	return nil
}
