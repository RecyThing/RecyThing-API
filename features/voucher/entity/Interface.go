package entity

import (
	"mime/multipart"
	"recything/utils/pagination"
)

type VoucherRepositoryInterface interface {
	Create(image *multipart.FileHeader,recybot VoucherCore) error
	GetAll(page, limit int, search string) ([]VoucherCore, pagination.PageInfo, int, error)
	GetCount(search string) (int, error)
	GetById(idVoucher string) (VoucherCore, error)
	Update(idVoucher string, image *multipart.FileHeader,data VoucherCore) error
	Delete(idVoucher string) error
	CreateExchangeVoucher(idUser string, data ExchangeVoucherCore) error
	GetAllExchange() ([]ExchangeVoucherCore, error) 
	GetByIdExchange(idExchange string) (ExchangeVoucherCore, error)
}

type VoucherServiceInterface interface {
	Create(image *multipart.FileHeader,data VoucherCore) error
	GetAll(page, limit, search string ) ([]VoucherCore,pagination.PageInfo, int,error)
	GetById(idVoucher string) (VoucherCore, error)
	UpdateData(idVoucher string, image *multipart.FileHeader,data VoucherCore) error
	DeleteData(idVoucher string) error
	CreateExchangeVoucher(idUser string, data ExchangeVoucherCore) error
	GetAllExchange() ([]ExchangeVoucherCore, error) 
	GetByIdExchange(idExchange string) (ExchangeVoucherCore, error)
}
