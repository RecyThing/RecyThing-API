package repository

import (
	"errors"
	"log"
	"recything/features/recybot/entity"
	"recything/features/recybot/model"
	"recything/utils/constanta"
	"recything/utils/pagination"

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

	tx := rb.db.Create(&input)
	if tx.Error != nil {
		return entity.RecybotCore{}, tx.Error
	}

	result := entity.ModelRecybotToCoreRecybot(input)
	return result, nil
}

func (rb *recybotRepository) FindAll(page, limit int, category string) ([]entity.RecybotCore, pagination.PageInfo, int, error) {
	dataRecybots := []model.Recybot{}
	var totalCount int64
	offsetInt := (page - 1) * limit
	log.Println("offsetInt", offsetInt)
	log.Println("page repo ", page)
	log.Println("limit repo", limit)

	if category == "" {

		err := rb.db.Model(&model.Recybot{}).Count(&totalCount).Error
		if err != nil {
			return nil, pagination.PageInfo{}, 0, err
		}

		tx := rb.db.Limit(limit).Offset(offsetInt).Find(&dataRecybots)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	if category != "" {
		tx := rb.db.Model(&model.Recybot{}).Where("category LIKE ?", "%"+category+"%").Count(&totalCount)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}

		tx = rb.db.Limit(limit).Offset(offsetInt).Where("category LIKE ?", "%"+category+"%").Find(&dataRecybots)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	result := entity.ListModelRecybotToCoreRecybot(dataRecybots)
	paginationInfo := pagination.CalculateData(int(totalCount), limit, page)
	return result, paginationInfo, int(totalCount), nil

}

func (rb *recybotRepository) GetAll() ([]entity.RecybotCore, error) {
	dataRecybots := []model.Recybot{}

	tx := rb.db.Find(&dataRecybots)
	if tx.Error != nil {
		return []entity.RecybotCore{}, tx.Error
	}

	result := entity.ListModelRecybotToCoreRecybot(dataRecybots)
	return result, nil
}

func (rb *recybotRepository) GetById(idData string) (entity.RecybotCore, error) {
	dataRecybots := model.Recybot{}

	tx := rb.db.Where("id = ?", idData).First(&dataRecybots)
	if tx.Error != nil {
		return entity.RecybotCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.RecybotCore{}, errors.New(constanta.ERROR_DATA_ID)
	}

	result := entity.ModelRecybotToCoreRecybot(dataRecybots)
	return result, nil
}

func (rb *recybotRepository) Update(idData string, recybot entity.RecybotCore) (entity.RecybotCore, error) {
	data := entity.CoreRecybotToModelRecybot(recybot)

	tx := rb.db.Where("id = ?", idData).Updates(&data)
	if tx.Error != nil {
		return entity.RecybotCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.RecybotCore{}, errors.New(constanta.ERROR_DATA_ID)
	}

	result := entity.ModelRecybotToCoreRecybot(data)
	return result, nil
}

func (rb *recybotRepository) Delete(idData string) error {
	data := model.Recybot{}

	tx := rb.db.Where("id = ?", idData).Delete(&data)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}
