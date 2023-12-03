package repository

import (
	"errors"
	"mime/multipart"
	"recything/features/community/entity"
	"recything/features/community/model"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/storage"

	"gorm.io/gorm"
)

type communityRepository struct {
	db *gorm.DB
}

func NewCommunityRepository(db *gorm.DB) entity.CommunityRepositoryInterface {
	return &communityRepository{
		db: db,
	}
}

// CreateCommunity implements entity.CommunityRepositoryInterface.
func (cr *communityRepository) CreateCommunity(image *multipart.FileHeader, data entity.CommunityCore) error {
	request := entity.CoreCommunityToModelCommunity(data)

	imageURL, errUpload := storage.UploadThumbnail(image)
	if errUpload != nil {
		return errUpload
	}
	request.Image = imageURL

	tx := cr.db.Create(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteCommunityById implements entity.CommunityRepositoryInterface.
func (cr *communityRepository) DeleteCommunityById(id string) error {
	request := model.Community{}

	tx := cr.db.Where("id = ?", id).Delete(&request)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// GetAllCommunity implements entity.CommunityRepositoryInterface.
func (cr *communityRepository) GetAllCommunity(page int, limit int, search string) ([]entity.CommunityCore, pagination.PageInfo, int, error) {
	communityExchange := []model.Community{}

	offset := (page - 1) * limit
	query := cr.db.Model(&model.Community{})

	if search != "" {
		query = query.Where("name LIKE ? OR location LIKE ?", "%"+search+"%", "%"+search+"%")
	}
	
	var totalCount int64
	tx := query.Count(&totalCount)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	query = query.Order("members DESC")
	query = query.Offset(offset).Limit(limit)

	tx = query.Find(&communityExchange)
	if tx.Error != nil {
		return nil, pagination.PageInfo{}, 0, tx.Error
	}

	response := entity.ListModelCommunityToCoreCommunity(communityExchange)
	pageInfo := pagination.CalculateData(int(totalCount), limit, page)
	return response, pageInfo, int(totalCount), nil
}

// GetCommunityById implements entity.CommunityRepositoryInterface.
func (cr *communityRepository) GetCommunityById(id string) (entity.CommunityCore, error) {
	communityExchange := model.Community{}

	tx := cr.db.Where("id = ?", id).First(&communityExchange)
	if tx.Error != nil {
		return entity.CommunityCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.CommunityCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	result := entity.ModelCommunityToCoreCommunity(communityExchange)
	return result, nil
}

// UpdateCommunityById implements entity.CommunityRepositoryInterface.
func (cr *communityRepository) UpdateCommunityById(id string, image *multipart.FileHeader, data entity.CommunityCore) error {
	dataCommunity := model.Community{}

	input := entity.CoreCommunityToModelCommunity(data)

	tx := cr.db.Where("id = ?", id).First(&dataCommunity)
	if tx.Error != nil {
		return tx.Error
	}
	
	if image != nil {
		imageURL, errUpload := storage.UploadThumbnail(image)
		if errUpload != nil {
			return errUpload
		}
		input.Image = imageURL
	} else {
		input.Image = dataCommunity.Image
	}

	tx = cr.db.Where("id = ?", id).Updates(&input)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

func (cr *communityRepository) GetByName(name string) (entity.CommunityCore, error) {
	dataCommunity := model.Community{}

	tx := cr.db.Where("name = ?", name).First(&dataCommunity)

	if tx.RowsAffected == 0 {
		return entity.CommunityCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	if tx.Error != nil {
		return entity.CommunityCore{}, tx.Error
	}

	result := entity.ModelCommunityToCoreCommunity(dataCommunity)
	return result, nil
}