package entity

import (
	"mime/multipart"
)

type ReportRepositoryInterface interface {
	Insert(reportInput ReportCore) (ReportCore, error)
	SelectById(idReport string) (ReportCore, error)
	ReadAllReport(idUser string) ([]ReportCore, error)
	UploadProof(id string, data ReportCore, image *multipart.FileHeader) (purchases ReportCore, err error)
}

type ReportServiceInterface interface {
	Create(reportInput ReportCore, userId string) (ReportCore, error)
	ReadAllReport(idUser string) ([]ReportCore, error)
	SelectById(idReport string) (ReportCore, error)
	UploadProof(id string, data ReportCore, image *multipart.FileHeader) (purchases ReportCore, err error)
}
