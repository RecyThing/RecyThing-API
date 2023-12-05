package service

import (
	"errors"
	"mime/multipart"
	admin "recything/features/admin/entity"
	"recything/features/mission/entity"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/storage"
	"recything/utils/validation"
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

	if len(data.MissionStages) < constanta.MIN_STAGE {
		return errors.New("tahapan misi tidak boleh kosong")

	}
	if len(data.MissionStages) > constanta.MAX_STAGE {
		return errors.New(constanta.ERROR_MISSION_LIMIT)
	}

	errEmpty := validation.CheckDataEmpty(data.Title, data.Description, data.StartDate, data.EndDate, data.Point)
	if errEmpty != nil {
		return errEmpty
	}

	for _, stage := range data.MissionStages {
		err := validation.CheckDataEmpty(stage.Description, data.Title)
		if err != nil {
			return err
		}
	}

	err := validation.ValidateDate(data.StartDate, data.EndDate)
	if err != nil {
		return err
	}

	imageURL, errUpload := storage.UploadThumbnail(image)
	if errUpload != nil {
		return err
	}

	data.MissionImage = imageURL
	err = ms.MissionRepo.CreateMission(data)
	if err != nil {
		return err
	}

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
	}

	return data, pagnationInfo, count, nil
}

func (ms *missionService) UpdateMission(image *multipart.FileHeader, missionID string, data entity.Mission) error {

	err := validation.ValidateDateForUpdate(data.StartDate, data.EndDate)
	if err != nil {
		return err
	}
	err = validation.CheckDataEmpty(data.Title, data.Description, data.Point, data.EndDate, data.StartDate)
	if err != nil {
		return err
	}

	imageURL, err := ms.MissionRepo.GetImageURL(missionID)
	if err != nil {
		return err
	}

	if image != nil {
		newImageURL, errUpload := storage.UploadThumbnail(image)
		if errUpload != nil {
			return err
		}
		data.MissionImage = newImageURL
	} else {
		data.MissionImage = imageURL
	}

	err = ms.MissionRepo.UpdateMission(missionID, data)
	if err != nil {
		return err
	}

	return nil
}

func (ms *missionService) UpdateMissionStage(missionID string, data []entity.MissionStage) error {
	for _, stage := range data {
		err := validation.CheckDataEmpty(stage.Description, stage.Title)
		if err != nil {
			return err
		}
	}

	err := ms.MissionRepo.UpdateMissionStage(missionID, data)
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

func (ms *missionService) FindById(missionID string) (entity.Mission, error) {

	dataMission, err := ms.MissionRepo.FindById(missionID)
	if err != nil {
		return entity.Mission{}, err
	}

	return dataMission, nil
}

func (ms *missionService) DeleteMission(missionID string) error {

	err := ms.MissionRepo.DeleteMission(missionID)
	if err != nil {
		return err
	}

	return nil
}
