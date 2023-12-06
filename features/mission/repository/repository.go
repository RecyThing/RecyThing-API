package repository

import (
	"errors"
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

// UpdateStatusMissionApproval implements entity.MissionRepositoryInterface.

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

func (mr *MissionRepository) FindAllMission(page, limit int, search, filter string) ([]entity.Mission, pagination.PageInfo, int, error) {
	data := []model.Mission{}
	offsetInt := (page - 1) * limit
	paginationQuery := mr.db.Limit(limit).Offset(offsetInt)

	totalCount, err := mr.GetCountMission(filter, search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	if filter != "" {
		tx := paginationQuery.Where("status LIKE ?", "%"+filter+"%").Preload("MissionStages").Find(&data)
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

	if search == "" || filter == "" {
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

		if err := mr.db.Where("id = ?", v.ID).Preload("MissionStages").Take(&v).Error; err != nil {
			return nil, pagination.PageInfo{}, 0, err
		}
		result = append(result, v)
	}

	dataMission := entity.ListMissionModelToMissionCore(result)
	paginationInfo := pagination.CalculateData(totalCount, limit, page)
	return dataMission, paginationInfo, totalCount, nil
}

func (mr *MissionRepository) GetCountMission(filter, search string) (int, error) {
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

func (mr *MissionRepository) GetCountMissionApproval(filter, search string) (int, error) {

	// counts := helper.CountMissionApproval{}
	var totalCount int64
	model := mr.db.Model(&model.UploadMissionTask{})
	if filter == "" || search == "" {
		tx := model.Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}
	}

	if search != "" {
		tx := model.
			Joins("JOIN users ON upload_mission_tasks.user_id = users.id").
			Where("users.fullname LIKE ?", "%"+search+"%").Count(&totalCount)

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

func (mr *MissionRepository) GetCountDataMissionApproval() (helper.CountMissionApproval, error) {

	counts := helper.CountMissionApproval{}
	model := mr.db.Model(&model.UploadMissionTask{})

	tx := model.Count(&counts.All)
	if tx.Error != nil {
		return counts, tx.Error
	}

	tx = model.Where("status LIKE ?", "%"+constanta.DISETUJUI+"%").Count(&counts.Approved)
	if tx.Error != nil {
		return counts, tx.Error
	}
	tx = model.Where("status LIKE ?", "%"+constanta.DITOLAK+"%").Count(&counts.NotApproved)
	if tx.Error != nil {
		return counts, tx.Error
	}
	tx = model.Where("status LIKE ?", "%"+constanta.PERLU_TINJAUAN+"%").Count(&counts.NeedReview)
	if tx.Error != nil {
		return counts, tx.Error
	}

	return counts, nil
}

func (mr *MissionRepository) FindById(missionID string) (entity.Mission, error) {
	dataMission := model.Mission{}

	tx := mr.db.Where("id = ? ", missionID).Preload("MissionStages").First(&dataMission)
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

func (mr *MissionRepository) UpdateMissionStage(missionID string, data []entity.MissionStage) error {
	tx := mr.db.Where("id = ?", missionID).Take(&model.Mission{})
	if tx.Error != nil {
		return tx.Error
	}

	if len(data) > constanta.MAX_STAGE {
		return errors.New(constanta.ERROR_MISSION_LIMIT)
	}

	var countStage int64
	tx = mr.db.Model(&model.MissionStage{}).Where("mission_id = ?", missionID).Count(&countStage)
	if tx.Error != nil {
		return tx.Error
	}

	allStages := []model.MissionStage{}
	tx = mr.db.Where("mission_id = ?", missionID).Find(&allStages)
	if tx.Error != nil {
		return tx.Error
	}

	dataIDs := make(map[string]bool)
	for _, stage := range data {
		dataIDs[stage.ID] = true
	}

	for _, stage := range allStages {
		if _, exists := dataIDs[stage.ID]; !exists {
			tx = mr.db.Unscoped().Delete(&stage)
			if tx.Error != nil {
				return tx.Error
			}
		}
	}
	for _, stage := range allStages {
		if _, exists := dataIDs[stage.ID]; !exists {
			tx = mr.db.Unscoped().Delete(&stage)
			if tx.Error != nil {
				return tx.Error
			}
		}
	}

	for _, stage := range data {
		if stage.ID == "" {
			countStage++
		}
	}

	if countStage > constanta.MAX_STAGE {
		return errors.New(constanta.ERROR_MISSION_LIMIT)
	}

	missionStage := entity.ListMissionStagesCoreToMissionStagesModel(data)

	for i, stage := range missionStage {
		if stage.ID == "" {
			tx = mr.db.Create(&stage)
			if tx.Error != nil {
				return tx.Error
			}
		}

		if stage.ID != "" {
			for j := i + 1; j < len(data); j++ {
				if stage.ID == data[j].ID {
					return errors.New(constanta.ERROR_INVALID_ID)
				}
			}

			existStage := model.MissionStage{}
			tx = mr.db.Where("id = ?", stage.ID).First(&existStage)
			if tx.Error != nil {
				return tx.Error
			}

			existStage.Title = stage.Title
			existStage.Description = stage.Description

			tx = mr.db.Save(&existStage)
			if tx.Error != nil {
				return tx.Error
			}

		}
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
func (mr *MissionRepository) CreateUploadMissionTask(userID string, data entity.UploadMissionTaskCore, images []*multipart.FileHeader) error {
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

func (mr *MissionRepository) FindUploadMissionStatus(id, missionID, userID, status string) error {
	dataUpload := model.UploadMissionTask{}

	if id == "" {
		tx := mr.db.Where("user_id = ? AND mission_id = ?", userID, missionID).First(&dataUpload)
		if tx.Error != nil {
			return tx.Error
		}

		if tx.RowsAffected == 0 {
			return tx.Error
		}
	}

	if missionID == "" {
		tx := mr.db.Where("id = ? AND user_id = ? AND status = ?", id, userID, status).First(&dataUpload)
		if tx.Error != nil {
			return tx.Error
		}

		if tx.RowsAffected == 0 {
			return tx.Error
		}

	}

	return nil
}

func (mr *MissionRepository) UpdateUploadMissionTask(id string, images []*multipart.FileHeader, data entity.UploadMissionTaskCore) error {
	dataUploadMission := model.UploadMissionTask{}
	request := entity.UploadMissionTaskCoreToUploadMissionTaskModel(data)

	tx := mr.db.Where("id = ?", id).First(&dataUploadMission)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}
	request.Status = constanta.PERLU_TINJAUAN
	errUpdate := mr.db.Model(&dataUploadMission).Updates(request)
	if errUpdate.Error != nil {
		return errUpdate.Error
	}

	ImageList := []model.ImageUploadMission{}

	tx = mr.db.Where("upload_mission_task_id = ? ", id).Find(&ImageList)
	if tx.Error != nil {
		return tx.Error
	}

	tx = mr.db.Unscoped().Delete(&ImageList)
	if tx.Error != nil {
		return tx.Error
	}

	for _, image := range images {
		Imagedata := entity.ImageUploadMissionCore{}
		imageURL, uploadErr := storage.UploadProof(image)
		if uploadErr != nil {
			return uploadErr
		}

		Imagedata.UploadMissionTaskID = id
		Imagedata.Image = imageURL
		ImageSave := entity.ImageUploadMissionCoreToImageUploadMissionModel(Imagedata)

		if err := mr.db.Create(&ImageSave).Error; err != nil {
			return err
		}

		data.Images = append(data.Images, Imagedata)
	}

	return nil
}

func (mr *MissionRepository) FindUploadById(id string) error {
	dataUploadMission := model.UploadMissionTask{}

	tx := mr.db.Where("id = ? ", id).First(&dataUploadMission)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}
func (mr *MissionRepository) FindMissionApprovalById(UploadMissionTaskID string) (entity.UploadMissionTaskCore, error) {
	data := model.UploadMissionTask{}

	tx := mr.db.Where("id = ? ", UploadMissionTaskID).Preload("Images").First(&data)
	if tx.Error != nil {
		return entity.UploadMissionTaskCore{}, tx.Error
	}

	// if tx.RowsAffected == 0 {
	// 	return entity.UploadMissionTaskCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	// }

	result := entity.UploadMissionTaskModelToUploadMissionTaskCore(data)

	return result, nil
}

func (mr *MissionRepository) FindAllMissionApproval(page, limit int, search, filter string) ([]entity.UploadMissionTaskCore, pagination.PageInfo, int, error) {
	approvalMission := []model.UploadMissionTask{}
	offsetInt := (page - 1) * limit
	paginationQuery := mr.db.Limit(limit).Offset(offsetInt)
	totalCount, err := mr.GetCountMissionApproval(filter, search)
	
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	if filter != "" {
		tx := paginationQuery.Where("status LIKE ?", "%"+filter+"%").Preload("Images").Find(&approvalMission)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	if search != "" {
		
		tx := paginationQuery.Model(&model.UploadMissionTask{}).
			Joins("JOIN users ON upload_mission_tasks.user_id = users.id").
			Where("users.fullname LIKE ?", "%"+search+"%").Preload("Images").
			Find(&approvalMission)

		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, errors.New("error disini")
		}
	}

	if search == "" && filter == "" {
		tx := paginationQuery.Preload("Images").Find(&approvalMission)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	paginationInfo := pagination.CalculateData(totalCount, limit, page)
	result := entity.ListUploadMissionTaskModelToUploadMissionTaskCore(approvalMission)
	return result, paginationInfo, totalCount, nil

}

func (mr *MissionRepository) UpdateStatusMissionApproval(uploadMissionTaskID, status, reason string) error {
	// err := mr.db.Model(&model.UploadMissionTask{}).Where("id = ?", UploadMissionTaskID).Update("status", status).Error
	// if err != nil {
	// 	return err
	// }

	err := mr.db.Model(&model.UploadMissionTask{}).Where("id = ?", uploadMissionTaskID).Updates(map[string]interface{}{
		"status": status,
		"reason": reason,
	}).Error

	if err != nil {
		return err
	}
	return nil

}
