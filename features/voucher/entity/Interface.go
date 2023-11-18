package entity

type VoucherRepositoryInterface interface {
	Create(recybot VoucherCore) error
	GetAll() ([]VoucherCore, error)
	GetById(idVoucher string) (VoucherCore, error)
	Update(idVoucher string, recybot VoucherCore) error
	Delete(idVoucher string) error
}

type VoucherServiceInterface interface {
	Create(data VoucherCore) error
	GetAll() ([]VoucherCore, error)
	GetById(idVoucher string) (VoucherCore, error)
	UpdateData(idVoucher string, data VoucherCore) error 
	DeleteData(idVoucher string) error
}
