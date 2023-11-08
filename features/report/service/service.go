package service

import (
	"errors"

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
		return []entity.ReportCore{}, errors.New("user not found")
	}

	reports, err := rc.ReportRepository.ReadAllReport(idUser)
	if err != nil {
		return []entity.ReportCore{}, errors.New("error get data")
	}

	return reports, nil
}

// SelectById implements entity.ReportRepositoryInterface.
func (rc *reportService) SelectById(idReport string) (entity.ReportCore, error) {
	if idReport == "" {
		return entity.ReportCore{}, errors.New("invalid id")
	}

	reportData, err := rc.ReportRepository.SelectById(idReport)
	if err != nil {
		return entity.ReportCore{}, errors.New("failed to read report")
	}

	return reportData, nil
}

// UploadProof implements entity.ReportRepositoryInterface.
func (*reportService) UploadProof(id string, data entity.ReportCore, image *multipart.FileHeader) (purchases entity.ReportCore, err error) {
	panic("unimplemented")
}

func (report *reportService) Create(reportInput entity.ReportCore, userId string) (entity.ReportCore, error) {

	reportInput.UserId = userId
	createdReport, errinsert := report.ReportRepository.Insert(reportInput)
	if errinsert != nil {
		return entity.ReportCore{}, errinsert
	}

	return createdReport, nil
}
