package service

import (
	"recything/features/recybot/entity"
)

type recybotService struct {
	recybotRepo entity.RecybotRepositoryInterface
}

func NewRecybotService(recybotRepo entity.RecybotRepositoryInterface) entity.RecybotServiceInterface {
	return &recybotService{
		recybotRepo: recybotRepo,
	}
}

// CreateData implements entity.RecybotServiceInterface.
func (rb *recybotService) CreateData(recybot entity.RecybotCore) (entity.RecybotCore, error) {
	result, err := rb.recybotRepo.Create(recybot)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Delete implements entity.RecybotServiceInterface.
func (rb *recybotService) DeleteData(idData string) error {
	_, err := rb.SelectById(idData)
	if err != nil {
		return err
	}
	err = rb.recybotRepo.Delete(idData)
	if err != nil {
		return err
	}
	return nil
}

// UpdateData implements entity.RecybotServiceInterface.
func (rb *recybotService) UpdateData(idData string, recybot entity.RecybotCore) (entity.RecybotCore, error) {
	_, err := rb.recybotRepo.SelectById(idData)
	if err != nil {
		return entity.RecybotCore{}, err
	}

	result, err := rb.recybotRepo.Update(idData, recybot)
	if err != nil {
		return result, err
	}
	result.ID = idData
	return result, nil
}

func (rb *recybotService) SelectAllData() ([]entity.RecybotCore, error) {
	result, err := rb.recybotRepo.SelectAll()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (rb *recybotService) SelectById(idData string) (entity.RecybotCore, error) {
	result, err := rb.recybotRepo.SelectById(idData)
	if err != nil {
		return result, err
	}
	return result, nil
}
