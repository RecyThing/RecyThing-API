package entity

type AchievementRepositoryInterface interface {
	GetAllAchievement() ([]AchievementCore, error)
	UpdateById(id int, data AchievementCore) error
	FindById(id int) (AchievementCore, error)
}

type AchievementServiceInterface interface {
	GetAllAchievement() ([]AchievementCore, error)
	UpdateById(id int, data AchievementCore) error
}