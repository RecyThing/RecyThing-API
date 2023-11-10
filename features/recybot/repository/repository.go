package repository

import (
	"recything/features/recybot/entity"

	"gorm.io/gorm"
)

type recybotRepository struct {
	db *gorm.DB
}

func NewRecybotRepository(db *gorm.DB) entity.RecybotRepositoryInterface {
	return &recybotRepository{
		db: db,
	}
}

func (rb *recybotRepository) Create(recybot entity.RecybotCore) (entity.RecybotCore, error) {
	input := entity.CoreRecybotToModelRecybot(recybot)
	err := rb.db.Create(&input).Error
	if err != nil {
		return entity.RecybotCore{}, err
	}
	result := entity.ModelRecybotToCoreRecybot(input)
	return result, err
}
