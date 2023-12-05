package service

import (
	"recything/features/daily_point/entity"
)

type dailyPointService struct {
	DailyPointRepository entity.DailyPointRepositoryInterface
}

func NewDailyPointService(daily entity.DailyPointRepositoryInterface) entity.DailyPointServiceInterface {
	return &dailyPointService{
		DailyPointRepository: daily,
	}
}

// DailyClaim implements entity.DailyPointServiceInterface.
func (dailyS *dailyPointService) DailyClaim(userId string) error {
	tx := dailyS.DailyPointRepository.DailyClaim(userId)
	if tx != nil {
		return tx
	}

	return nil
}

// PostWeekly implements entity.DailyPointServiceInterface.
func (dailyS *dailyPointService) PostWeekly() error {
	err := dailyS.DailyPointRepository.PostWeekly()
	if err != nil {
		return err
	}

	return nil
}
