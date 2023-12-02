package service

import (
	"errors"
	"log"
	"mime/multipart"
	admin "recything/features/admin/entity"
	"recything/features/mission/entity"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/storage"
	"recything/utils/validation"
	"time"
)

type missionService struct {
	MissionRepo entity.MissionRepositoryInterface
	AdminRepo   admin.AdminRepositoryInterface
}

func NewMissionService(missionRepo entity.MissionRepositoryInterface, adminRepo admin.AdminRepositoryInterface) entity.MissionServiceInterface {
	return &missionService{
		MissionRepo: missionRepo,
		AdminRepo:   adminRepo,
	}
}

func (ms *missionService) CreateMission(image *multipart.FileHeader, data entity.Mission) error {

	errEmpty := validation.CheckDataEmpty(data.Title, data.Description, data.StartDate, data.EndDate, data.Point, image)
	if errEmpty != nil {
		return errEmpty
	}

	err := validation.ValidateDate(data.StartDate, data.EndDate)
	if err != nil {
		return err
	}
	uploadError := make(chan error)
	var imageURL string
	go func() {
		imageURL, errUpload := storage.UploadThumbnail(image)
		if errUpload != nil {
			uploadError <- errUpload
			return
		}
		data.MissionImage = imageURL
		uploadError <- nil
	}()

	data.MissionImage = imageURL
	err = ms.ChangesStatusMission(data)
	if err != nil {
		return err
	}

	err = ms.MissionRepo.CreateMission(data)
	if err != nil {
		return err
	}
	return nil
}

func (ms *missionService) CreateMissionStages(adminID, missionID string, data []entity.MissionStage) error {
	for _, stage := range data {
		errEmpty := validation.CheckDataEmpty(stage.Title, stage.Description, stage.MissionID)
		if errEmpty != nil {
			return errEmpty
		}
	}

	err := ms.MissionRepo.CreateMissionStages(data)
	if err != nil {
		return err
	}
	return nil
}

func (ms *missionService) FindAllMission(page, limit, search, filter string) ([]entity.Mission, pagination.PageInfo, int, error) {
	pageInt, limitInt, err := validation.ValidateParamsPagination(page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	data, pagnationInfo, count, err := ms.MissionRepo.FindAllMission(pageInt, limitInt, search, filter)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	for i := range data {
		err := ms.ChangesStatusMission(data[i])
		if err != nil {
			return nil, pagination.PageInfo{}, 0, err
		}

		admin, err := ms.AdminRepo.SelectById(data[i].AdminID)
		if err != nil {
			return nil, pagination.PageInfo{}, 0, err
		}

		data[i].Creator = admin.Fullname
	}

	return data, pagnationInfo, count, nil
}

func (ms *missionService) ChangesStatusMission(data entity.Mission) error {
	endDate, err := time.Parse("2006-01-02", data.EndDate)
	if err != nil {
		return err
	}

	currentTime := time.Now().Truncate(24 * time.Hour)
	if endDate.Before(currentTime) {
		data.Status = "Melewati Tenggat"
	}
	return nil
}

func (ms *missionService) UpdateMission(image *multipart.FileHeader, missionID string, data entity.Mission) error {

	uploadError := make(chan error)
	var imageURL string

	go func() {
		imageURL, errUpload := storage.UploadThumbnail(image)
		if errUpload != nil {
			uploadError <- errUpload
			return
		}
		data.MissionImage = imageURL
		uploadError <- nil
	}()

	data.MissionImage = imageURL
	log.Println("data service after validation", data)
	err := ms.MissionRepo.UpdateMission(missionID, data)
	if err != nil {
		return err
	}

	return nil
}

func (ms *missionService) UpdateMissionStage(missionStageID string, data entity.Stage) error {
	errEmpty := validation.CheckDataEmpty(data.Title, data.Description)
	if errEmpty != nil {
		return errEmpty
	}

	err := ms.MissionRepo.UpdateMissionStage(missionStageID, data)
	if err != nil {
		return err
	}
	return nil
}

// Claimed Mission
func (ms *missionService) ClaimMission(userID string, data entity.ClaimedMission) error {
	if data.MissionID == "" {
		return errors.New(constanta.ERROR_EMPTY)
	}
	err := ms.MissionRepo.ClaimMission(userID, data)
	if err != nil {
		return err
	}

	return nil
}
