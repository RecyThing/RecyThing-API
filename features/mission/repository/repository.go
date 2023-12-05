package repository

import (
	"errors"
	"log"
	"mime/multipart"
	"recything/features/mission/entity"
	"recything/features/mission/model"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/pagination"
	"recything/utils/storage"
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

	result := []model.Mission{}

	for _, v := range data {
		newStatus, err := helper.ChangeStatusMission(v.EndDate)
		if err != nil {
			return nil, pagination.PageInfo{}, 0, err
		}

		v.Status = newStatus
		err = mr.db.Model(&v).Updates(map[string]interface{}{"status": v.Status}).Error
		if err != nil {
			return nil, pagination.PageInfo{}, 0, err
		}

		if err := mr.db.Where("id = ?", v.ID).Take(&v).Error; err != nil {
			return nil, pagination.PageInfo{}, 0, err
		}
		result = append(result, v)
	}

	dataMission := entity.ListMissionModelToMissionCore(result)
	paginationInfo := pagination.CalculateData(totalCount, limit, page)
	return dataMission, paginationInfo, totalCount, nil
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

func (mr *MissionRepository) FindById(missionID string) (entity.Mission, error) {
	dataMission := model.Mission{}

	tx := mr.db.Where("id = ? ", missionID).First(&dataMission)
	if tx.Error != nil {
		return entity.Mission{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.Mission{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dataResponse := entity.MissionModelToMissionCore(dataMission)
	return dataResponse, nil
}

func (mr *MissionRepository) UpdateMission(missionID string, data entity.Mission) error {

	dataMission := entity.MissionCoreToMissionModel(data)
	getMission := model.Mission{}
	tx := mr.db.Where("id = ?", missionID).First(&getMission)
	if tx.Error != nil {
		return tx.Error
	}

	ok := helper.FieldsEqual(getMission, data, "Title", "Description", "Point", "StartDate", "EndDate")
	if ok {
		return errors.New(constanta.ERROR_INVALID_UPDATE)
	}

	endDateValid, err := time.Parse("2006-01-02", data.EndDate)
	if err != nil {
		return err
	}
	currentTime := time.Now().Truncate(24 * time.Hour)
	if endDateValid.Before(currentTime) {
		data.Status = constanta.OVERDUE
	} else {
		data.Status = constanta.ACTIVE
	}

	tx = mr.db.Where("id = ?", missionID).Updates(&dataMission)
	if tx.Error != nil {
		if tx.Error != nil {
			if validation.IsDuplicateError(tx.Error) {
				return errors.New(constanta.ERROR_DATA_EXIST)
			}
			return tx.Error
		}
		return tx.Error
	}
	return nil
}

func (mr *MissionRepository) UpdateMissionStage(missionStageID string, data entity.MissionStage) error {
	oldMissionStage := model.MissionStage{}
	tx := mr.db.Where("id = ?", missionStageID).Take(&oldMissionStage)
	if tx.Error != nil {
		return tx.Error
	}
	ok := helper.FieldsEqual(oldMissionStage, data, "Title", "Description")
	if ok {
		return errors.New(constanta.ERROR_INVALID_UPDATE)
	}

	missionStage := entity.MissionStagesCoreToMissionStagesModel(data)
	tx = mr.db.Where("id = ?", missionStageID).Updates(&missionStage)
	if tx.Error != nil {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}
func (mr *MissionRepository) AddNewMissionStage(missionID string, data []entity.MissionStage) error {
	mission := model.Mission{}

	tx := mr.db.Where("id = ?", missionID).Take(&mission)
	if tx.Error != nil {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	if len(data) > 3 {
		return errors.New(constanta.ERROR_MISSION_LIMIT)
	}

	var countStage int64
	tx = mr.db.Model(&model.MissionStage{}).Where("mission_id = ?", missionID).Count(&countStage)
	if tx.Error != nil {
		return tx.Error
	}

	if countStage > 3 {
		return errors.New(constanta.ERROR_MISSION_LIMIT)
	}

	if int(countStage)+len(data) > 3 {
		return errors.New(constanta.ERROR_MISSION_LIMIT)
	}

	missionStages := entity.ListMissionStagesCoreToMissionStagesModel(data)
	tx = mr.db.Save(&missionStages)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (mr *MissionRepository) DeleteMission(missionID string) error {
	dataMission := model.Mission{}

	tx := mr.db.Where("id = ? ", missionID).Delete(&dataMission)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

func (mr *MissionRepository) DeleteMissionStage(stageID string) error {
	mission := model.MissionStage{}
	tx := mr.db.Where("id = ?", stageID).Take(&mission)
	if tx.Error != nil {
		return tx.Error
	}

	tx = mr.db.Where("id = ?", stageID).Delete(&mission)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (mr *MissionRepository) GetImageURL(missionID string) (string, error) {
	mission := model.Mission{}
	err := mr.db.Where("id = ?", missionID).Take(&mission).Error
	if err != nil {
		return "", err
	}

	return mission.MissionImage, nil
}

// Claimed Mission
func (mr *MissionRepository) ClaimMission(userID string, data entity.ClaimedMission) error {
	input := entity.ClaimedCoreToClaimedMissionModel(data)

	errFind := mr.FindClaimed(userID, data.MissionID)
	if errFind == nil {
		return errors.New("error : mission sudah di klaim")
	}

	input.UserID = userID
	tx := mr.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (mr *MissionRepository) FindClaimed(userID, missionID string) error {
	dataClaimed := model.ClaimedMission{}
	tx := mr.db.Where("user_id = ? AND mission_id = ? AND claimed = 1", userID, missionID).First(&dataClaimed)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}
	return nil
}

// Upload Mission User
func (mr *MissionRepository) CreateUploadMission(userID string, data entity.UploadMissionTaskCore, images []*multipart.FileHeader) error {
	request := entity.UploadMissionTaskCoreToUploadMissionTaskModel(data)
	request.UserID = userID
	tx := mr.db.Create(&request)
	if tx.Error != nil {
		return tx.Error
	}

	for _, image := range images {
		imageURL, uploadErr := storage.UploadProof(image)
		if uploadErr != nil {
			return uploadErr
		}

		ImageList := entity.ImageUploadMissionCore{}
		ImageList.UploadMissionTaskID = request.ID
		ImageList.Image = imageURL

		ImageSave := entity.ImageUploadMissionCoreToImageUploadMissionModel(ImageList)

		if err := mr.db.Create(&ImageSave).Error; err != nil {
			return err
		}

		data.Images = append(data.Images, ImageList)
	}

	return nil
}

func (mr *MissionRepository) FindUploadMission(userID, missionID,status string) error {
	dataUpload := model.UploadMissionTask{}
	log.Println("user",userID,"dada",missionID)
	tx := mr.db.Where("user_id = ? AND mission_id = ? AND status = ?", userID, missionID,status).First(&dataUpload)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return tx.Error
	}

	log.Println(tx.RowsAffected)

	return nil
}
