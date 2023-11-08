package entity

import "recything/features/user/entity"

type AdminRepositoryInterface interface {
	Insert(data AdminCore) (AdminCore, error)
	SelectAll() ([]AdminCore, error)
	SelectById(adminId string) (AdminCore, error)
	Update(adminId string, data AdminCore) error
	Delete(adminId string) error
	FindByEmailANDPassword(email, password string) (AdminCore, error)
	//Manage Users
	SelectAllUsers() ([]entity.UsersCore, error)
	SelectByIdUsers(adminId string) (entity.UsersCore, error)
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
	GetAllUsers() ([]entity.UsersCore, error)
	GetByIdUsers(adminId string) (entity.UsersCore, error)
	DeleteUsers(adminId string) error
}
