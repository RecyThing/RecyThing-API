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
func (rb *recybotService) Delete(idData string) (entity.RecybotCore, error) {
	result, err := rb.recybotRepo.Delete(idData)
	if err != nil {
		return result, err
	}
	return result, nil
}

// UpdateData implements entity.RecybotServiceInterface.
func (rb *recybotService) UpdateData(idData string) (entity.RecybotCore, error) {
	result, err := rb.recybotRepo.Update (idData)
	if err != nil {
		return result, err
	}
	return result, nil
}
