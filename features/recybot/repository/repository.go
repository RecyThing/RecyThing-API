package repository

import (
	"recything/features/recybot/entity"
	"recything/features/recybot/model"

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

func (rb *recybotRepository) Update(idData string) (entity.RecybotCore, error) {
	data := model.Recybot{}
	err := rb.db.Where("id = ?", idData).Updates(&data).Error
	if err != nil {
		return entity.RecybotCore{}, err
	}
	result := entity.ModelRecybotToCoreRecybot(data)
	return result, err
}

func (rb *recybotRepository) Delete(idData string) (entity.RecybotCore, error) {
	data := model.Recybot{}
	err := rb.db.Where("id = ?", idData).Delete(&data).Error
	if err != nil {
		return entity.RecybotCore{}, err
	}
	result := entity.ModelRecybotToCoreRecybot(data)
	return result, err
}
