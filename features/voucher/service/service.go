package service

import (
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

func (vs *voucherService) Create(data entity.VoucherCore) error {

	errEmpty := validation.CheckDataEmpty(data.RewardName, data.Image,data.Description,data.StartDate,data.EndDate,data.StartDate,data.EndDate)
	if errEmpty != nil {
		return errEmpty
	}

	errCreate := vs.voucherRepository.Create(data)
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

func (vs *voucherService) UpdateData(idVoucher string, data entity.VoucherCore) error {

	errEmpty := validation.CheckDataEmpty(data.RewardName, data.Image,data.Description,data.StartDate,data.EndDate,data.StartDate,data.EndDate)
	if errEmpty != nil {
		return errEmpty
	}

	err := vs.voucherRepository.Update(idVoucher, data)
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