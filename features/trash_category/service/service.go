package service

import (
	"recything/features/trash_category/entity"
	"recything/utils/constanta"
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
	if errEmpty != nil {
		return errCheck
	}

	data.Unit = validUnit
	err := tc.trashCategoryRepo.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (tc *trashCategoryService) GetAllCategory(page, trashType, limit string) ([]entity.TrashCategoryCore, entity.PagnationInfo, error) {
	if limit == "" && page == "" {
		if trashType != "" {
			data, pagnationInfo, err := tc.trashCategoryRepo.FindByTrashType(trashType)
			if err != nil {
				return nil, entity.PagnationInfo{}, err
			}
			return data, pagnationInfo, nil
		}

		result, pagnationInfo, err := tc.trashCategoryRepo.FindAll()
		if err != nil {
			return nil, entity.PagnationInfo{}, err
		}
		return result, pagnationInfo, nil
	}

	result, paganation, err := tc.trashCategoryRepo.FindAllWithSearchAndPagnation(page, trashType, limit)

	if err != nil {
		return result, paganation, err
	}

	return result, paganation, nil
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
