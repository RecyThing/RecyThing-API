package service

import (
	"mime/multipart"
	admin "recything/features/admin/entity"
	"recything/features/mission/entity"
	"recything/utils/pagination"
	"recything/utils/storage"
	"recything/utils/validation"
	"time"
)

type missionService struct {
	missionRepo entity.MissionRepositoryInterface
	AdminRepo   admin.AdminRepositoryInterface
}

func NewMissionService(missionRepo entity.MissionRepositoryInterface, adminRepo admin.AdminRepositoryInterface) entity.MissionServiceInterface {
	return &missionService{
		missionRepo: missionRepo,
		AdminRepo:   adminRepo,
	}
}

// CreateData implements entity.trashCategoryServiceInterface.
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

	err = ms.missionRepo.CreateMission(data)
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

	err := ms.missionRepo.CreateMissionStages(data)
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

	data, pagnationInfo, count, err := ms.missionRepo.FindAllMission(pageInt, limitInt, search, filter)
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

func (ms *missionService) UpdateMission(image *multipart.FileHeader, idMission string, data entity.Mission) error {
	errEmpty := validation.CheckDataEmpty(data.Title, data.Description, data.StartDate, data.EndDate)
	if errEmpty != nil {
		return errEmpty
	}

	errDate := validation.ValidateDate(data.StartDate, data.EndDate)
	if errDate != nil {
		return errEmpty
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
	err := ms.missionRepo.UpdateMission(idMission, data)
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

	err := ms.missionRepo.UpdateMissionStage(missionStageID, data)
	if err != nil {
		return err
	}
	return nil
}

// func (tc *trashCategoryService) GetById(idTrash string) (entity.TrashCategoryCore, error) {

// 	result, err := tc.trashCategoryRepo.GetById(idTrash)
// 	if err != nil {
// 		return result, err
// 	}
// 	return result, nil
// }

// // Delete implements entity.trashCategoryServiceInterface.
// func (tc *trashCategoryService) DeleteCategory(idTrash string) error {

// 	err := tc.trashCategoryRepo.Delete(idTrash)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// // UpdateData implements entity.trashCategoryServiceInterface.
// func (tc *trashCategoryService) UpdateCategory(idTrash string, data entity.TrashCategoryCore) (entity.TrashCategoryCore, error) {

// 	errEmpty := validation.CheckDataEmpty(data.TrashType, data.Unit)
// 	if errEmpty != nil {
// 		return entity.TrashCategoryCore{}, errEmpty
// 	}

// 	result, err := tc.trashCategoryRepo.Update(idTrash, data)
// 	if err != nil {
// 		return result, err
// 	}
// 	result.ID = idTrash
// 	return result, nil
// }
