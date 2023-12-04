package repository

import (
	"errors"
	"recything/features/achievement/entity"
	"recything/features/achievement/model"

	//users "recything/features/user/model"
	"recything/utils/constanta"

	"gorm.io/gorm"
)

type achievementRepository struct {
	db *gorm.DB
}

func NewAchievementRepository(db *gorm.DB) entity.AchievementRepositoryInterface {
	return &achievementRepository{
		db: db,
	}
}

// GetAllAchievement implements entity.AchievementRepositoryInterface.
func (ar *achievementRepository) GetAllAchievement() ([]entity.AchievementCore, error) {
	dataAchievement := []model.Achievement{}

	tx := ar.db.
		Table("achievements").
		Select("achievements.*, COUNT(users.id) as total_claimed").
		Joins("LEFT JOIN users ON achievements.name = users.badge").
		Group("achievements.id").
		Find(&dataAchievement)
		
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataResponse := entity.ListAchievementModelToAchievementCore(dataAchievement)
	return dataResponse, nil
}

// UpdateById implements entity.AchievementRepositoryInterface.
func (ar *achievementRepository) UpdateById(id int, data entity.AchievementCore) error {
	dataAchievement := entity.AchievementCoreToAchievementModel(data)

	tx := ar.db.Where("id = ?", id).Updates(&dataAchievement)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}