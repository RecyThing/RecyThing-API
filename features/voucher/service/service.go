package service

import (
	"mime/multipart"
	"recything/features/voucher/entity"
	"recything/utils/pagination"
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

func (vs *voucherService) GetAll(page, limit, search string ) ([]entity.VoucherCore,pagination.PageInfo, int,error) {
	pageInt, limitInt, err := validation.ValidateParamsPagination(page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0,err 
	}
	data, pagnationInfo,count, err := vs.voucherRepository.GetAll(pageInt,limitInt,search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return data, pagnationInfo,count, nil
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
