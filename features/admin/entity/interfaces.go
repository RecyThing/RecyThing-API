package entity

type AdminRepositoryInterface interface {
	Insert(data AdminCore) error
	GetAll() ([]AdminCore, error)
	GetById(id_admin, role string) (AdminCore, error)
}

type AdminServiceInterface interface {
	Create(data AdminCore) error
}
