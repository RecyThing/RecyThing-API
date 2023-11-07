package service

import "recything/features/report/entity"

type reportService struct {
	ReportRepository entity.ReportRepositoryInterface
}

func NewReportService(report entity.ReportRepositoryInterface) *reportService {
	return &reportService{
		ReportRepository: report,
	}
}

func (report *reportService) Create(reportInput entity.ReportCore, userId string) (entity.ReportCore, error) {

	reportInput.UserId = userId
	createdReport, err := report.ReportRepository.Insert(reportInput)
	if err != nil {
		return entity.ReportCore{}, err
	}
	
	return createdReport, nil
}
