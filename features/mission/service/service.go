package service

import (
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
	err, status := ms.ChangesStatusMission(data.EndDate)
	if err != nil {
		return err
	}
	data.Status = status

	return nil
}

func (ms *missionService) FindAllMission(page, limit, search, status string) ([]entity.Mission, pagination.PageInfo, int, error) {
	pageInt, limitInt, err := validation.ValidateParamsPagination(page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	data, pagnationInfo, count, err := ms.MissionRepo.FindAllMission(pageInt, limitInt, search, status)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	for i := range data {
		admin, err := ms.AdminRepo.SelectById(data[i].AdminID)
		if err != nil {
			return nil, pagination.PageInfo{}, 0, err
		}

		data[i].Creator = admin.Fullname

		err, status := ms.ChangesStatusMission(data[i].EndDate)
		if err != nil {
			return nil, pagination.PageInfo{}, 0, err
		}

		data[i].Status = status
	}

	return data, pagnationInfo, count, nil
}

func (ms *missionService) ChangesStatusMission(endDate string) (error, string) {
	endDateValid, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		return err, ""
	}
	status := constanta.OVERDUE
	currentTime := time.Now().Truncate(24 * time.Hour)
	if endDateValid.Before(currentTime) {
		return err, status
	}
	return nil, constanta.ACTIVE
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

	err := ms.MissionRepo.UpdateMission(missionID, data)
	if err != nil {
		return err
	}

	return nil
}

func (ms *missionService) UpdateMissionStage(missionStageID string, data entity.MissionStage) error {
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
