package entity

import "time"

type UsersRepositoryInterface interface {
	Register(data UsersCore) error
	Login(email, password string) (UsersCore, error)
	GetById(id string) (UsersCore, error)
	UpdateById(id string, updated UsersCore) (data UsersCore, err error)
	GetByVerificationToken(token string) (UsersCore, error)
	UpdateIsVerified(id string, isVerified bool) error
	ForgetPassword(email string, updated UsersCore) (data UsersCore, err error)
	UpdatePassword(id string, updated UsersCore) (data UsersCore, err error)
	EmailExists(email string) (bool, error)
	SendOTP(emailUser string, otp string, expiry time.Time) (data UsersCore, err error)
	VerifyOTP(emailUser string, otp string) (data UsersCore, err error)
	ResetOTP(emailUser string) (data UsersCore, err error)
}

type UsersUsecaseInterface interface {
	Register(data UsersCore) error
	Login(email, password string) (UsersCore, string, error)
	GetById(id string) (UsersCore, error)
	UpdateById(id string, updated UsersCore) (data UsersCore, err error)
	VerifyUser(token string) (bool, error)
	UpdateIsVerified(id string, isVerified bool) error
	ForgetPassword(emailUser string, updated UsersCore) error
	UpdatePassword(id string, updated UsersCore) (data UsersCore, err error)
	SendOTP(emailUser string) error
	VerifyOTP(emailUser string, otp string) (string, error)
}
