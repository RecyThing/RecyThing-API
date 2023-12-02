package repository

import (
	"errors"
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

func (rb *recybotRepository) FindAll(page, limit int, filter, search string) ([]entity.RecybotCore, pagination.PageInfo, int, error) {
	dataRecybots := []model.Recybot{}

	offsetInt := (page - 1) * limit
	totalCount, err := rb.GetCount(filter, search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	paginationQuery := rb.db.Limit(limit).Offset(offsetInt)

	if filter == "" || search == "" {
		tx := paginationQuery.Find(&dataRecybots)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	if filter != "" {
		tx := paginationQuery.Where("category LIKE ?", "%"+filter+"%").Find(&dataRecybots)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	if search != "" {
		tx := paginationQuery.Where("category LIKE ? or question LIKE ? ", "%"+search+"%", "%"+search+"%").Find(&dataRecybots)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	result := entity.ListModelRecybotToCoreRecybot(dataRecybots)
	paginationInfo := pagination.CalculateData(int(totalCount), limit, page)
	return result, paginationInfo, totalCount, nil

}

func (rb *recybotRepository) GetCount(filter, search string) (int, error) {
	var totalCount int64
	model := rb.db.Model(&model.Recybot{})
	if filter == "" || search == "" {
		err := model.Count(&totalCount).Error
		if err != nil {
			return 0, err
		}

	}
	if filter != "" {
		tx := model.Where("category LIKE ?", "%"+filter+"%").Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}
	}

	if search != "" {
		tx := model.Where("category LIKE ? or question LIKE ? ", "%"+search+"%", "%"+search+"%").Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}

	}
	return int(totalCount), nil

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
