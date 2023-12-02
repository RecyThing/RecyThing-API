package service

import (
	"mime/multipart"
	admin "recything/features/admin/entity"
	"recything/features/mission/entity"
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
	}
	
	if image == nil {
		data.MissionImage = imageURL
	}

	err = ms.MissionRepo.UpdateMission(missionID, data)
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
