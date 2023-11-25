package service

import (
	"mime/multipart"
	"recything/features/voucher/entity"
	"recything/utils/validation"
)

type voucherService struct {
	voucherRepository entity.VoucherRepositoryInterface
}

func NewVoucherService(voucher entity.VoucherRepositoryInterface) entity.VoucherServiceInterface {
	return &voucherService{
		voucherRepository: voucher,
	}
}

func (vs *voucherService) Create(image *multipart.FileHeader, data entity.VoucherCore) error {

	errEmpty := validation.CheckDataEmpty(data.RewardName, data.Point, data.Description, data.StartDate, data.EndDate)
	if errEmpty != nil {
		return errEmpty
	}

	errDate := validation.ValidateDate(data.StartDate, data.EndDate)
	if errDate != nil {
		return errDate
	}

	errCreate := vs.voucherRepository.Create(image, data)
	if errCreate != nil {
		return errCreate
	}
	return nil
}

func (vs *voucherService) GetAll() ([]entity.VoucherCore, error) {
	result, err := vs.voucherRepository.GetAll()
	if err != nil {
		return result, err
	}

	return result, nil
}

func (vs *voucherService) GetById(idVoucher string) (entity.VoucherCore, error) {
	result, err := vs.voucherRepository.GetById(idVoucher)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (vs *voucherService) UpdateData(idVoucher string, image *multipart.FileHeader, data entity.VoucherCore) error {

	errEmpty := validation.CheckDataEmpty(data.RewardName, data.Point, data.Description, data.StartDate, data.EndDate)
	if errEmpty != nil {
		return errEmpty
	}

	errDate := validation.ValidateDate(data.StartDate, data.EndDate)
	if errDate != nil {
		return errDate
	}

	err := vs.voucherRepository.Update(idVoucher, image, data)
	if err != nil {
		return err
	}

	return nil
}

func (vs *voucherService) DeleteData(idVoucher string) error {

	err := vs.voucherRepository.Delete(idVoucher)
	if err != nil {
		return err
	}
	return nil
}
