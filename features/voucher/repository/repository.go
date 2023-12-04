package repository

import (
	"errors"
	"log"
	"mime/multipart"
	"recything/features/voucher/entity"
	"recything/features/voucher/model"
	"recything/utils/constanta"
	"recything/utils/pagination"
	"recything/utils/storage"

	"gorm.io/gorm"
)

type voucherRepository struct {
	db *gorm.DB
}

func NewVoucherRepository(db *gorm.DB) entity.VoucherRepositoryInterface {
	return &voucherRepository{
		db: db,
	}
}

func (vr *voucherRepository) Create(image *multipart.FileHeader, data entity.VoucherCore) error {
	input := entity.CoreVoucherToModelVoucher(data)

	imageURL, errUpload := storage.UploadThumbnail(image)
	if errUpload != nil {
		return errUpload
	}

	input.Image = imageURL
	log.Println(input)
	tx := vr.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// CreateExchangeVoucher implements entity.VoucherRepositoryInterface.
func (vr *voucherRepository) CreateExchangeVoucher(idUser string, data entity.ExchangeVoucherCore) error {
	input := entity.CoreExchangeVoucherToModelExchangeVoucher(data)

	input.IdUser = idUser
	tx := vr.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (vr *voucherRepository) GetAll(page, limit int, search string) ([]entity.VoucherCore, pagination.PageInfo, int, error) {
	dataVouchers := []model.Voucher{}
	offsetInt := (page - 1) * limit

	totalCount, err := vr.GetCount(search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	if search == "" {
		tx := vr.db.Limit(limit).Offset(offsetInt).Find(&dataVouchers)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	if search != "" {
		tx := vr.db.Where("reward_name LIKE ? or point LIKE ? ", "%"+search+"%", "%"+search+"%").Limit(limit).Offset(offsetInt).Find(&dataVouchers)
		if tx.Error != nil {
			return nil, pagination.PageInfo{}, 0, tx.Error
		}
	}

	dataResponse := entity.ListModelVoucherToCoreVoucher(dataVouchers)
	paginationInfo := pagination.CalculateData(totalCount, limit, page)

	return dataResponse, paginationInfo, totalCount, nil
}

func (vr *voucherRepository) GetCount(search string) (int, error) {
	var totalCount int64

	if search == "" {
		tx := vr.db.Model(&model.Voucher{}).Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}
	}

	if search != "" {
		tx := vr.db.Model(&model.Voucher{}).Where("reward_name LIKE ? or point LIKE ? ", "%"+search+"%", "%"+search+"%").Count(&totalCount)
		if tx.Error != nil {
			return 0, tx.Error
		}

	}
	return int(totalCount), nil
}

func (vr *voucherRepository) GetById(idVoucher string) (entity.VoucherCore, error) {
	dataVouchers := model.Voucher{}

	tx := vr.db.Where("id = ?", idVoucher).First(&dataVouchers)
	if tx.Error != nil {
		return entity.VoucherCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.VoucherCore{}, tx.Error
	}

	result := entity.ModelVoucherToCoreVoucher(dataVouchers)
	return result, nil
}

func (vr *voucherRepository) Update(idVoucher string, image *multipart.FileHeader, data entity.VoucherCore) error {
	input := entity.CoreVoucherToModelVoucher(data)
	dataVoucher := model.Voucher{}

	tx := vr.db.Where("id = ?", idVoucher).First(&dataVoucher)
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
		input.Image = dataVoucher.Image
	}

	tx = vr.db.Where("id = ?", idVoucher).Updates(&input)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

func (vr *voucherRepository) Delete(idVoucher string) error {
	request := model.Voucher{}

	tx := vr.db.Where("id = ?", idVoucher).Delete(&request)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}
