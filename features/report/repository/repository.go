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

func NewReportRepository(db *gorm.DB) entity.ReportRepositoryInterface {
	return &reportRepository{db: db}
}

// ReadAllReport implements entity.ReportRepositoryInterface.
func (report *reportRepository) ReadAllReport(idUser string) ([]entity.ReportCore, error) {
	dataReport := []model.Report {}

	tx := report.db.Where("users_id = ?", idUser).Find(&dataReport)
	if tx.Error != nil {

		return nil, tx.Error
	}

	mapData := make([]entity.ReportCore, len(dataReport))
	for i, value := range dataReport {
		mapData[i] = entity.ReportModelToReportCore(value)
	}

	return mapData, nil
}


// UploadProof implements entity.ReportRepositoryInterface.
func (*reportRepository) UploadProof(id string, data entity.ReportCore, image *multipart.FileHeader) (purchases entity.ReportCore, err error) {
	panic("unimplemented")
}

func (report *reportRepository) Insert(reportInput entity.ReportCore) (entity.ReportCore, error) {
	dataReport := entity.ReportCoreToReportModel(reportInput)

	tx := report.db.Create(&dataReport)
	if tx.Error != nil {
		return entity.ReportCore{}, tx.Error
	}

	dataResponse := entity.ReportModelToReportCore(dataReport)
	return dataResponse, nil
}

func (report *reportRepository) SelectById(iDReport string) (entity.ReportCore, error) {
	dataReports := model.Report{}

	tx := report.db.Where("id = ?", iDReport).Preload("Images").First(&dataReports)
	if tx.Error != nil {
		return entity.ReportCore{}, tx.Error
	}

	dataResponse := entity.ReportModelToReportCore(dataReports)
	return dataResponse, nil
}
