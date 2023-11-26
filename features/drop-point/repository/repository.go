package repository

import (
	"errors"
	"recything/features/drop-point/entity"
	"recything/features/drop-point/model"
	"recything/utils/constanta"
	"recything/utils/pagination"

	"gorm.io/gorm"
)

type dropPointRepository struct {
	db *gorm.DB
}

func NewDropPointRepository(db *gorm.DB) entity.DropPointRepositoryInterface {
	return &dropPointRepository{
		db: db,
	}
}

// Create implements entity.DropPointRepositoryInterface.
func (dpr *dropPointRepository) CreateDropPoint(data entity.DropPointCore) (entity.DropPointCore, error) {
	request := entity.DropPointCoreToDropPointModel(data)

	tx := dpr.db.Create(&request)
	if tx.Error != nil {
		return entity.DropPointCore{}, tx.Error
	}

	dataResponse := entity.DropPointModelToDropPointCore(request)
	return dataResponse, nil
}

// DeleteDropPoint implements entity.DropPointRepositoryInterface.
func (dpr *dropPointRepository) DeleteDropPointById(id string) error {
	dropPointData := model.DropPoint{}

	tx := dpr.db.Unscoped().Where("id = ?", id).Delete(&dropPointData)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// GetAllDropPoint implements entity.DropPointRepositoryInterface.
func (dpr *dropPointRepository) GetAllDropPoint(page, limit int, name, address string) ([]entity.DropPointCore, pagination.PageInfo, error) {
	dropPoint := []model.DropPoint{}

	offset := (page - 1) * limit
	query := dpr.db.Model(&model.DropPoint{}).Preload("OperationalSchedules")

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	if address != "" {
		query = query.Where("address LIKE ?", "%"+address+"%")
	}


	var totalCount int64
	tx := query.Count(&totalCount).Find(&dropPoint)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, tx.Error
	}

	query = query.Offset(offset).Limit(limit)

	tx = query.Find(&dropPoint)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, tx.Error
	}

	dropPointCores := []entity.DropPointCore{}
	for _, dropPointModel := range dropPoint {
		dropPointCore := entity.DropPointModelToDropPointCore(dropPointModel)
		dropPointCores = append(dropPointCores, dropPointCore)
	}

	// Menghitung informasi paginasi
	pageInfo := pagination.CalculateData(int(totalCount), limit, page)

	return dropPointCores, pageInfo, nil

}

// GetById implements entity.DropPointRepositoryInterface.
func (dpr *dropPointRepository) GetDropPointById(id string) (entity.DropPointCore, error) {
	dropPoint := model.DropPoint{}

	tx := dpr.db.Preload("OperationalSchedules").Where("id = ?", id).First(&dropPoint)
	if tx.Error != nil {
		return entity.DropPointCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.DropPointCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dropPointId := entity.DropPointModelToDropPointCore(dropPoint)
	return dropPointId, nil
}

// UpdateById implements entity.DropPointRepositoryInterface.
func (dpr *dropPointRepository) UpdateDropPointById(id string, data entity.DropPointCore) (entity.DropPointCore, error) {
	dropPointData := model.DropPoint{}
	operationalData := model.OperationalSchedules{}

	// Perbarui data DropPoint
	tx := dpr.db.Where("id = ?", id).First(&dropPointData)
	if tx.Error != nil {
		return entity.DropPointCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.DropPointCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	errUpdate := dpr.db.Model(&dropPointData).Updates(entity.DropPointCoreToDropPointModel(data))
	if errUpdate.Error != nil {
		return entity.DropPointCore{}, errUpdate.Error
	}

	// Hapus data OperationalSchedules yang ada
	tx = dpr.db.Unscoped().Where("drop_point_id = ?", id).Delete(&operationalData)
	if tx.Error != nil {
		return entity.DropPointCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.DropPointCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	// Tambahkan data OperationalSchedules yang baru
	for _, schedule := range data.OperationalSchedules {
		newOperationalData := entity.OperationalSchedulesCoreToOperationalSchedulesModel(schedule)
		newOperationalData.DropPointId = id

		err := dpr.db.Create(&newOperationalData)
		if err.Error != nil {
			return entity.DropPointCore{}, err.Error
		}
	}

	dataResponse := entity.DropPointModelToDropPointCore(dropPointData)

	return dataResponse, nil
}
