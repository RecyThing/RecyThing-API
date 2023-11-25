package repository

import (
	"errors"
	"mime/multipart"
	"recything/features/voucher/entity"
	"recything/features/voucher/model"
	"recything/utils/constanta"
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
	tx := vr.db.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (vr *voucherRepository) GetAll() ([]entity.VoucherCore, error) {
	dataVouchers := []model.Voucher{}

	tx := vr.db.Find(&dataVouchers)
	if tx.Error != nil {
		return []entity.VoucherCore{}, tx.Error
	}

	dataResponse := entity.ListModelVoucherToCoreVoucher(dataVouchers)
	return dataResponse, nil
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

	imageURL, errUpload := storage.UploadThumbnail(image)
	if errUpload != nil {
		return errUpload
	}

	input.Image = imageURL

	tx := vr.db.Where("id = ?", idVoucher).Updates(&input)
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
