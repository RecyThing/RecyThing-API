package entity

import "mime/multipart"

type ReportRepositoryInterface interface {
	Insert(reportInput ReportCore, image *multipart.FileHeader) (ReportCore, error)
	SelectById(idReport string) (ReportCore, error)
	ReadAllReport(idUser string) ([]ReportCore, error)
}

type ReportServiceInterface interface {
	Create(reportInput ReportCore, userId string, image *multipart.FileHeader) (ReportCore, error)
	ReadAllReport(idUser string) ([]ReportCore, error)
	SelectById(idReport string) (ReportCore, error)
}
