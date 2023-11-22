package entity

import (
	report "recything/features/report/entity"
	user "recything/features/user/entity"
)

type AdminRepositoryInterface interface {
	Create(data AdminCore) (AdminCore, error)
	SelectAll() ([]AdminCore, error)
	SelectById(adminId string) (AdminCore, error)
	Update(adminId string, data AdminCore) error
	Delete(adminId string) error
	FindByEmail(email string) error
	FindByEmailANDPassword(data AdminCore) (AdminCore, error)
	//Manage Users
	GetAllUsers() ([]user.UsersCore, error)
	GetByIdUser(userId string) (user.UsersCore, error)
	DeleteUsers(adminId string) error
	// Manage Reporting
	GetByStatusReport(status string) ([]report.ReportCore, error)
	UpdateStatusReport(id, status, reason string) (report.ReportCore, error)
	GetReportById(id string) (report.ReportCore, error)
}

type AdminServiceInterface interface {
	Create(data AdminCore) (AdminCore, error)
	GetAll() ([]AdminCore, error)
	GetById(adminId string) (AdminCore, error)
	UpdateById(adminId string, data AdminCore) error
	DeleteById(adminId string) error
	FindByEmailANDPassword(data AdminCore) (AdminCore, string, error)
	//Manage Users
	GetAllUsers() ([]user.UsersCore, error)
	GetByIdUsers(adminId string) (user.UsersCore, error)
	DeleteUsers(adminId string) error
	// Manage Reporting
	GetByStatusReport(status string) (data []report.ReportCore, err error)
	UpdateStatusReport(id, status, reason string) (report.ReportCore, error)
	GetReportById(id string) (report.ReportCore, error)
}
