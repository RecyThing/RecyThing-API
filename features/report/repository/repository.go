package repository

import (
	"recything/features/report/entity"
	"recything/features/report/model"

	"gorm.io/gorm"
)

type reportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *reportRepository {
	return &reportRepository{db: db}
}

func (report *reportRepository) Insert(reportInput entity.ReportCore) (entity.ReportCore, error) {
	dataReport := entity.ReportCoreToReportModel(reportInput)
	if err := report.db.Create(&dataReport).Error; err != nil {
		return entity.ReportCore{}, err
	}


	ReportCreated := entity.ReportModelToReportCore(dataReport)
	return ReportCreated, nil
}

func (report *reportRepository) SelectById(iDReport string) (entity.ReportCore, error) {
	reportModel := model.Report{}
	err := report.db.Where("id = ?", iDReport).Preload("Images").First(&reportModel).Error
	if err != nil {
		return entity.ReportCore{}, err
	}
	dataReport := entity.ReportModelToReportCore(reportModel)
	return dataReport, nil
}
