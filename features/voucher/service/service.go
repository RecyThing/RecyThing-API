package service

import (
	"errors"
	"mime/multipart"
	user "recything/features/user/entity"
	"recything/features/voucher/entity"
	"recything/utils/pagination"
	"recything/utils/validation"
)

type voucherService struct {
	voucherRepository entity.VoucherRepositoryInterface
	userRepository    user.UsersRepositoryInterface
}

func NewVoucherService(voucher entity.VoucherRepositoryInterface, user user.UsersRepositoryInterface) entity.VoucherServiceInterface {
	return &voucherService{
		voucherRepository: voucher,
		userRepository:    user,
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

func (vs *voucherService) GetAll(page, limit, search string) ([]entity.VoucherCore, pagination.PageInfo, int, error) {
	pageInt, limitInt, err := validation.ValidateParamsPagination(page, limit)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}
	data, pagnationInfo, count, err := vs.voucherRepository.GetAll(pageInt, limitInt, search)
	if err != nil {
		return nil, pagination.PageInfo{}, 0, err
	}

	return data, pagnationInfo, count, nil
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

// CreateExchangeVoucher implements entity.VoucherServiceInterface.
func (vs *voucherService) CreateExchangeVoucher(idUser string, data entity.ExchangeVoucherCore) error {
	errEmpty := validation.CheckDataEmpty(data.IdVoucher, data.Phone)
	if errEmpty != nil {
		return errEmpty
	}

	userData, err := vs.userRepository.GetById(idUser)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	voucherData, err := vs.voucherRepository.GetById(data.IdVoucher)
	if err != nil {
		return errors.New("voucher tidak ditemukan")
	}

	if userData.Point <= voucherData.Point {
        return errors.New("point tidak cukup")
    }

	userData.Point -= voucherData.Point

	// Update user
	err = vs.userRepository.UpdateById(userData.Id, userData)
	if err != nil {
		return errors.New("gagal memperbarui nilai point pengguna")
	}

	err = vs.voucherRepository.CreateExchangeVoucher(idUser, data)
	if err != nil {
		return err
	}
	return nil
}


func (vs *voucherService) GetAllExchange() ([]entity.ExchangeVoucherCore, error) {

	dataExchange,errGet := vs.voucherRepository.GetAllExchange()
	if errGet != nil {
		return []entity.ExchangeVoucherCore{},errGet
	}

	return dataExchange,nil
}

func (vs *voucherService) GetByIdExchange(idExchange string) (entity.ExchangeVoucherCore, error){

	dataExchange ,errGet := vs.voucherRepository.GetByIdExchange(idExchange)
	if errGet != nil {
		return entity.ExchangeVoucherCore{},errGet
	}

	return dataExchange,nil
}