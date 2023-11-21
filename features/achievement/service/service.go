package service

import (
	"errors"
	"recything/features/achievement/entity"
	"recything/utils/constanta"
	"strings"
)

type achievementService struct {
	achievementRepository entity.AchievementRepositoryInterface
}

func NewAchievementService(achievement entity.AchievementRepositoryInterface) entity.AchievementServiceInterface {
	return &achievementService{
		achievementRepository: achievement,
	}
}

// GetAllAchievement implements entity.AchievementServiceInterface.
func (as *achievementService) GetAllAchievement() ([]entity.AchievementCore, error) {
	achievement, err := as.achievementRepository.GetAllAchievement()
	if err != nil {
		return nil, err
	}

	return achievement, nil
}

// UpdateById implements entity.AchievementServiceInterface.
func (as *achievementService) UpdateById(id int, data entity.AchievementCore) error {
	if id == 0 {
		return errors.New(constanta.ERROR_ID_INVALID)
	}

	data.Name = strings.ToLower(data.Name)

	existingAchievement, err := as.achievementRepository.GetByName(data.Name)
	if err != nil {
		return errors.New("gagal memeriksa keberadaan data achievement")
	}

	if existingAchievement.Id != 0 && existingAchievement.Id != id {
		return errors.New("achievement dengan nama yang sama sudah ada di database")
	}

	switch data.Name {
	case "platinum":
		if data.TargetPoint != 250000 {
			return errors.New("nilai target point untuk platinum harus 250.000")
		}
	case "gold":
		if data.TargetPoint != 100000 {
			return errors.New("nilai target point untuk gold harus 100.000")
		}
	case "silver":
		if data.TargetPoint != 50000 {
			return errors.New("nilai target point untuk silver harus 50.000")
		}
	case "bronze":
		if data.TargetPoint != 0 {
			return errors.New("nilai target point untuk bronze harus 0")
		}
	default:
		return errors.New("data achievement tidak valid")
	}

	errUpdate := as.achievementRepository.UpdateById(id, data)
	if errUpdate != nil {
		return errors.New("gagal melakukan update data")
	} 

	return nil	
}
