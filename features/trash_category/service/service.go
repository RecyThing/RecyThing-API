package service

import (
	"recything/features/trash_category/entity"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/validation"
)

type trashCategoryService struct {
	trashCategoryRepo entity.TrashCategoryRepositoryInterface
}

func NewTrashCategoryService(trashCategoryRepo entity.TrashCategoryRepositoryInterface) entity.TrashCategoryServiceInterface {
	return &trashCategoryService{
		trashCategoryRepo: trashCategoryRepo,
	}
}

// CreateData implements entity.trashCategoryServiceInterface.
func (tc *trashCategoryService) CreateCategory(data entity.TrashCategoryCore) error {

	errEmpty := validation.CheckDataEmpty(data.Unit, data.TrashType, data.Point)
	if errEmpty != nil {
		return errEmpty
	}

	validUnit, errCheck := validation.CheckEqualData(data.Unit, constanta.Unit)
	if errCheck != nil {
		return errCheck
	}

	data.Unit = validUnit
	err := tc.trashCategoryRepo.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (tc *trashCategoryService) GetAllCategory(page, trashType, limit string) ([]entity.TrashCategoryCore, pagination.PageInfo, error) {
	pageInt, limitInt, err := validation.ValidateTypePaginationParameters(limit, page)
	if err != nil {
		return nil, pagination.PageInfo{}, err
	}
	
	validPage, validLimit := validation.ValidatePaginationParameters(pageInt, limitInt)
	data, pagnationInfo, err := tc.trashCategoryRepo.FindAll(validPage, validLimit, trashType)
	if err != nil {
		return nil, pagination.PageInfo{}, err
	}
	return data, pagnationInfo, nil
}

func (tc *trashCategoryService) GetById(idTrash string) (entity.TrashCategoryCore, error) {

	result, err := tc.trashCategoryRepo.GetById(idTrash)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Delete implements entity.trashCategoryServiceInterface.
func (tc *trashCategoryService) DeleteCategory(idTrash string) error {

	err := tc.trashCategoryRepo.Delete(idTrash)
	if err != nil {
		return err
	}
	return nil
}

// UpdateData implements entity.trashCategoryServiceInterface.
func (tc *trashCategoryService) UpdateCategory(idTrash string, data entity.TrashCategoryCore) (entity.TrashCategoryCore, error) {

	errEmpty := validation.CheckDataEmpty(data.TrashType, data.Unit)
	if errEmpty != nil {
		return entity.TrashCategoryCore{}, errEmpty
	}

	result, err := tc.trashCategoryRepo.Update(idTrash, data)
	if err != nil {
		return result, err
	}
	result.ID = idTrash
	return result, nil
}
