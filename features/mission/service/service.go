package service

import (
	"mime/multipart"
	"recything/features/mission/entity"
	"recything/utils/pagination"
	"recything/utils/storage"
	"recything/utils/validation"
)

type missionService struct {
	missionRepo entity.MissionRepositoryInterface
}

func NewMissionService(missionRepo entity.MissionRepositoryInterface) entity.MissionServiceInterface {
	return &missionService{
		missionRepo: missionRepo,
	}
}

// CreateData implements entity.trashCategoryServiceInterface.
func (ms *missionService) CreateMission(image *multipart.FileHeader, data entity.Mission) error {

	errEmpty := validation.CheckDataEmpty(data.Title, data.Description, data.StartDate, data.EndDate, data.Point)
	if errEmpty != nil {
		return errEmpty
	}

	imageURL, errUpload := storage.UploadThumbnail(image)
	if errUpload != nil {
		return errUpload
	}
	data.MissionImage = imageURL

	err := ms.missionRepo.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (ms *missionService) FindAll(page, limit, filter string) ([]entity.Mission, pagination.PageInfo, int, error) {
	pageInt, limitInt, err := validation.ValidateParamsPagination(page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	data, pagnationInfo, count, err := ms.missionRepo.FindAll(pageInt, limitInt, filter)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}
	return data, pagnationInfo, count, nil
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
