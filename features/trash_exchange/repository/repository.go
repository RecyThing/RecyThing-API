package repository

import (
	"errors"
	"recything/features/trash_exchange/entity"
	"recything/features/trash_exchange/model"
	"recything/utils/constanta"
	"recything/utils/pagination"

	"gorm.io/gorm"
)

type trashExchangeRepository struct {
	db *gorm.DB
}

func NewTrashExchangeRepository(db *gorm.DB) entity.TrashExchangeRepositoryInterface {
	return &trashExchangeRepository{
		db: db,
	}
}

// CreateTrashExchange implements entity.TrashExchangeRepositoryInterface.
func (ter *trashExchangeRepository) CreateTrashExchange(data entity.TrashExchangeCore) (entity.TrashExchangeCore, error) {
	request := entity.TrashExchangeCoreToTrashExchangeModel(data)

	tx := ter.db.Create(&request)
	if tx.Error != nil {
		return entity.TrashExchangeCore{}, tx.Error
	}

	dataResponse := entity.TrashExchangeModelToTrashExchangeCore(request)
	return dataResponse, nil
}

// CreateTrashExchangeDetails implements entity.TrashExchangeRepositoryInterface.
func (ter *trashExchangeRepository) CreateTrashExchangeDetails(data entity.TrashExchangeDetailCore) (entity.TrashExchangeDetailCore, error) {
	request := entity.TrashExchangeDetailCoreToTrashExchangeDetailModel(data)

	tx := ter.db.Save(&request)
	if tx.Error != nil {
		return entity.TrashExchangeDetailCore{}, tx.Error
	}

	dataResponse := entity.TrashExchangeDetailModelToTrashExchangeDetailCore(request)
	return dataResponse, nil
}

// DeleteTrashExchangeById implements entity.TrashExchangeRepositoryInterface.
func (ter *trashExchangeRepository) DeleteTrashExchangeById(id string) error {
	TrashExchange := model.TrashExchange{}

	tx := ter.db.Unscoped().Where("id = ?", id).Delete(&TrashExchange)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// GetAllTrashExchange implements entity.TrashExchangeRepositoryInterface.
func (*trashExchangeRepository) GetAllTrashExchange(page int, limit int, search string) ([]entity.TrashExchangeCore, pagination.PageInfo, error) {
	panic("unimplemented")
}

// GetTrashExchangeById implements entity.TrashExchangeRepositoryInterface.
func (ter *trashExchangeRepository) GetTrashExchangeById(id string) (entity.TrashExchangeCore, error) {
	trashExchange := model.TrashExchange{}

	tx := ter.db.Preload("TrashExchangeDetails").Where("id = ?", id).First(&trashExchange)
	if tx.Error != nil {
		return entity.TrashExchangeCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.TrashExchangeCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dropPointId := entity.TrashExchangeModelToTrashExchangeCoreForGetById(trashExchange)
	return dropPointId, nil
}
