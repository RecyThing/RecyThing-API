package entity

type DailyPointRepositoryInterface interface {
	PostWeekly() error
	DailyClaim(userId string) error
}

type DailyPointServiceInterface interface {
	PostWeekly() error
	DailyClaim(userId string) error
}
