package entity

type AdminRepositoryInterface interface {
	Insert(data AdminCore) (AdminCore, error)
	SelectAll() ([]AdminCore, error)
	SelectById(adminId string) (AdminCore, error)
	Update(adminId string, data AdminCore) error
	Delete(adminId string) error
	FindByEmailANDPassword(email, password string) (AdminCore, error)
	//Manage Users
	SelectAllUsers() ([]AdminCore, error)
	SelectByIdUsers(adminId string) (AdminCore, error)
	DeleteUsers(adminId string) error
}

type AdminServiceInterface interface {
	Create(data AdminCore) (AdminCore, error)
	GetAll() ([]AdminCore, error)
	GetById(adminId string) (AdminCore, error)
	UpdateById(adminId string, data AdminCore) error
	DeleteById(adminId string) error
	FindByEmailANDPassword(email, password string) (AdminCore, string, error)
	//Manage Users
	GetAllUsers() ([]AdminCore, error)
	GetByIdUsers(adminId string) (AdminCore, error)
	DeleteUsers(adminId string) error
}
