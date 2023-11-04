package entity

type AdminRepositoryInterface interface {
	Insert(data AdminCore) error
	SelectAll() ([]AdminCore, error)
	SelectById(id_admin, role string) (AdminCore, error)
	Update(id_admin string, data AdminCore) error
	Delete(id_admin string) error
}

type AdminServiceInterface interface {
	Create(data AdminCore) error
	GetAll() ([]AdminCore, error)
	GetById(id_admin, role string) (AdminCore, error)
	UpdateById(id_admin, role string, data AdminCore) error
	DeleteById(id_admin string) error
}
