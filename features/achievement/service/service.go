package service

import (
	"errors"
	"recything/features/achievement/entity"
	"recything/utils/constanta"
	"recything/utils/validation"
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

	errEmpty := validation.CheckDataEmpty(data.TargetPoint)
	if errEmpty != nil {
		return errEmpty
	}

	errUpdate := as.achievementRepository.UpdateById(id, data)
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
