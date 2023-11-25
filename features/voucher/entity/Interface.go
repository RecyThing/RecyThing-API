package entity

import "mime/multipart"

type VoucherRepositoryInterface interface {
	Create(image *multipart.FileHeader,recybot VoucherCore) error
	GetAll() ([]VoucherCore, error)
	GetById(idVoucher string) (VoucherCore, error)
	Update(idVoucher string, image *multipart.FileHeader,data VoucherCore) error
	Delete(idVoucher string) error
}

type VoucherServiceInterface interface {
	Create(image *multipart.FileHeader,data VoucherCore) error
	GetAll() ([]VoucherCore, error)
	GetById(idVoucher string) (VoucherCore, error)
	UpdateData(idVoucher string, image *multipart.FileHeader,data VoucherCore) error
	DeleteData(idVoucher string) error
}
