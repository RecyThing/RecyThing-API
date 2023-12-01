package entity

import (
	report "recything/features/report/entity"
	user "recything/features/user/entity"
	"recything/utils/pagination"
)

type AdminRepositoryInterface interface {
	Create(data AdminCore) (AdminCore, error)
	SelectAll(page, limit int, search string) ([]AdminCore, pagination.PageInfo, int, error)
	SelectById(adminId string) (AdminCore, error)
	Update(adminId string, data AdminCore) error
	Delete(adminId string) error
	FindByEmail(email string) error
	FindByEmailANDPassword(data AdminCore) (AdminCore, error)
	GetCount(fullName, role string) (int, error)
	//Manage Users
	GetAllUsers( search string, page, limit int) ([]user.UsersCore,  pagination.PageInfo, int, error)
	GetByIdUser(userId string) (user.UsersCore, error)
	DeleteUsers(adminId string) error
	// Manage Reporting
	GetAllReport(status, search string, page, limit int) ([]report.ReportCore, pagination.PageInfo, int,error)
	UpdateStatusReport(id, status, reason string) (report.ReportCore, error)
	GetReportById(id string) (report.ReportCore, error)
}

type AdminServiceInterface interface {
	Create(data AdminCore) (AdminCore, error)
	GetAll(page, limit, search string) ([]AdminCore, pagination.PageInfo, int, error)
	GetById(adminId string) (AdminCore, error)
	UpdateById(adminId string, data AdminCore) error
	DeleteById(adminId string) error
	FindByEmailANDPassword(data AdminCore) (AdminCore, string, error)
	//Manage Users
	GetAllUsers(search, page, limit string) ([]user.UsersCore, pagination.PageInfo, int, error)
	GetByIdUsers(adminId string) (user.UsersCore, error)
	DeleteUsers(adminId string) error
	// Manage Reporting
	GetAllReport(status, search, page, limit string) ([]report.ReportCore, pagination.PageInfo, int,error)
	UpdateStatusReport(id, status, reason string) (report.ReportCore, error)
	GetReportById(id string) (report.ReportCore, error)
}
