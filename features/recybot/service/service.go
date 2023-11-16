package service

import (
	"errors"
	"recything/features/recybot/entity"
	"recything/utils/validation"
)

type recybotService struct {
	recybotRepo entity.RecybotRepositoryInterface
}

func NewRecybotService(rc entity.RecybotRepositoryInterface) entity.RecybotServiceInterface {
	return &recybotService{
		recybotRepo: rc,
	}
}

// CreateData implements entity.RecybotServiceInterface.
func (rb *recybotService) CreateData(data entity.RecybotCore) (entity.RecybotCore, error) {

	errEmpty := validation.CheckDataEmpty(data.Category,data.Question)
	if errEmpty != nil {
		return entity.RecybotCore{},errEmpty
	}

	if data.Category != "sampah plastik" && data.Category != "sampah organik" {
		return entity.RecybotCore{}, errors.New("jenis sampah harus diisi dengan 'sampah plastik' atau 'sampah organik'")
	}

	result, err := rb.recybotRepo.Create(data)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (rb *recybotService) GetAllData() ([]entity.RecybotCore, error) {
	result, err := rb.recybotRepo.GetAll()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (rb *recybotService) GetById(idData string) (entity.RecybotCore, error) {
	result, err := rb.recybotRepo.GetById(idData)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Delete implements entity.RecybotServiceInterface.
func (rb *recybotService) DeleteData(idData string) error {

	err := rb.recybotRepo.Delete(idData)
	if err != nil {
		return err
	}
	return nil
}

// UpdateData implements entity.RecybotServiceInterface.
func (rb *recybotService) UpdateData(idData string, data entity.RecybotCore) (entity.RecybotCore, error) {

	errEmpty := validation.CheckDataEmpty(data.Category, data.Question)
	if errEmpty != nil {
		return entity.RecybotCore{}, errEmpty
	}

	result, err := rb.recybotRepo.Update(idData, data)
	if err != nil {
		return result, err
	}
	result.ID = idData
	return result, nil
}
