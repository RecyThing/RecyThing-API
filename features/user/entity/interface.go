package entity

type UsersRepositoryInterface interface {
	Register(data UsersCore) error
	Login(email, password string) (UsersCore, error)
	GetById(id string) (UsersCore, error)
	UpdateById(id string, updated UsersCore) (data UsersCore, err error)
	GetByVerificationToken(token string) (UsersCore, error)
	UpdateIsVerified(id string, isVerified bool) error
}

type UsersUsecaseInterface interface {
	Register(data UsersCore) error
	Login(email, password string) (UsersCore, string, error)
	GetById(id string) (UsersCore, error)
	UpdateById(id string, updated UsersCore) (data UsersCore, err error)
	VerifyUser(token string) (bool, error)
	UpdateIsVerified(id string, isVerified bool) error
}
