package repository

import (
	"mime/multipart"
	"recything/features/report/entity"
	"recything/features/report/model"

	"gorm.io/gorm"
)

type reportRepository struct {
	db *gorm.DB
}

// ReadAllReport implements entity.ReportRepositoryInterface.
func (report *reportRepository) ReadAllReport(idUser string) ([]entity.ReportCore, error) {
	var dataReport []model.Report

	errData := report.db.Where("users_id = ?", idUser).Find(&dataReport).Error
	if errData != nil {

		return nil, errData
	}

	mapData := make([]entity.ReportCore, len(dataReport))
	for i, value := range dataReport {
		mapData[i] = entity.ReportModelToReportCore(value)
	}

	return mapData, nil
}

func NewReportRepository(db *gorm.DB) entity.ReportRepositoryInterface {
	return &reportRepository{db: db}
}

// UploadProof implements entity.ReportRepositoryInterface.
func (*reportRepository) UploadProof(id string, data entity.ReportCore, image *multipart.FileHeader) (purchases entity.ReportCore, err error) {
	panic("unimplemented")
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
