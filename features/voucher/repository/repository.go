package repository

import (
	"errors"
	"recything/features/voucher/entity"
	"recything/features/voucher/model"
	"recything/utils/constanta"

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

func (vr *voucherRepository) Create(recybot entity.VoucherCore) error {
	input := entity.CoreVoucherToModelVoucher(recybot)


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
		return entity.VoucherCore{}, errors.New(constanta.ERROR_DATA_ID)
	}

	result := entity.ModelVoucherToCoreVoucher(dataVouchers)
	return result, nil
}

func (vr *voucherRepository) Update(idVoucher string, recybot entity.VoucherCore) error {
	request := entity.CoreVoucherToModelVoucher(recybot)

	tx := vr.db.Where("id = ?", idVoucher).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_ID)
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
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}
