package entity

type AchievementRepositoryInterface interface {
	GetAllAchievement() ([]AchievementCore, error)
	UpdateById(id int, data AchievementCore) error
	GetByName(name string) (AchievementCore, error)
}

type AchievementServiceInterface interface {
	GetAllAchievement() ([]AchievementCore, error)
	UpdateById(id int, data AchievementCore) error
}