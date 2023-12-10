package entity

type DailyPointRepositoryInterface interface {
	PostWeekly() error
	DailyClaim(userId string) error
	GetAllHistoryPoint(userID string) ([]map[string]interface{}, error)
	GetByIdHistoryPoint(userID,idTransaction string) (map[string]interface{}, error) 
}

type DailyPointServiceInterface interface {
	PostWeekly() error
	DailyClaim(userId string) error
	GetAllHistoryPoint(userID string) ([]map[string]interface{}, error)
	GetByIdHistoryPoint(userID,idTransaction string) (map[string]interface{}, error) 
}
