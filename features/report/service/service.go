package service

import (
	"errors"
	"fmt"
	"time"

	"mime/multipart"
	"recything/features/report/entity"
)

type reportService struct {
	ReportRepository entity.ReportRepositoryInterface
}

func NewReportService(report entity.ReportRepositoryInterface) entity.ReportServiceInterface {
	return &reportService{
		ReportRepository: report,
	}
}

// ReadAllReport implements entity.ReportServiceInterface.
func (rc *reportService) ReadAllReport(idUser string) ([]entity.ReportCore, error) {
	if idUser == "" {
		return []entity.ReportCore{}, errors.New("pengguna tidak ditemukan")
	}

	reports, err := rc.ReportRepository.ReadAllReport(idUser)
	if err != nil {
		return []entity.ReportCore{}, errors.New("gagal mendapatkan data")
	}

	return reports, nil
}

// SelectById implements entity.ReportRepositoryInterface.
func (rc *reportService) SelectById(idReport string) (entity.ReportCore, error) {
	if idReport == "" {
		return entity.ReportCore{}, errors.New("id tidak cocok")
	}

	reportData, err := rc.ReportRepository.SelectById(idReport)
	if err != nil {
		return entity.ReportCore{}, errors.New(err.Error())
	}

	return reportData, nil
}

func (report *reportService) Create(reportInput entity.ReportCore, userId string, images []*multipart.FileHeader) (entity.ReportCore, error) {

	if reportInput.ReportType == "Pelanggaran Sampah" {

		fmt.Println("service : ", reportInput.InsidentDate)
		if _, parseErr := time.Parse("2006-01-02", reportInput.InsidentDate); parseErr != nil {
			return entity.ReportCore{}, errors.New("error, tanggal harus dalam format 'yyyy-mm-dd'")
		}

		if _, errHour := time.Parse("15:04", reportInput.InsidentTime); errHour != nil {
			return entity.ReportCore{}, errors.New("error, jam harus dalam format 'hh:mm'")
		}
	}

	for _, image := range images {
        if image != nil && image.Size > 20*1024*1024 {
            return entity.ReportCore{}, errors.New("ukuran file tidak boleh lebih dari 20 MB")
        }
    }

	reportInput.UserId = userId
	createdReport, errinsert := report.ReportRepository.Insert(reportInput, images)
	if errinsert != nil {
		return entity.ReportCore{}, errinsert
	}

	return createdReport, nil
}
