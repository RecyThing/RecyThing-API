package repository

import (
	"errors"

	"recything/features/admin/entity"
	"recything/features/admin/model"

	report "recything/features/report/entity"
	reportModel "recything/features/report/model"

	user "recything/features/user/entity"
	userModel "recything/features/user/model"
	"recything/utils/constanta"
	"recything/utils/helper"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) entity.AdminRepositoryInterface {
	return &AdminRepository{
		db: db,
	}
}

func (ar *AdminRepository) Create(data entity.AdminCore) (entity.AdminCore, error) {
	dataAdmins := entity.AdminCoreToAdminModel(data)

	tx := ar.db.Create(&dataAdmins)
	if tx.Error != nil {
		return entity.AdminCore{}, tx.Error
	}

	dataResponse := entity.AdminModelToAdminCore(dataAdmins)
	return dataResponse, nil
}

func (ar *AdminRepository) SelectAll() ([]entity.AdminCore, error) {
	dataAdmins := []model.Admin{}

	tx := ar.db.Where("role = ? ", constanta.ADMIN).Find(&dataAdmins)
	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("role tidak ditemukan")
	}

	dataResponse := entity.ListAdminModelToAdminCore(dataAdmins)
	return dataResponse, nil
}

func (ar *AdminRepository) SelectById(adminId string) (entity.AdminCore, error) {
	dataAdmins := model.Admin{}

	tx := ar.db.Where("id = ? AND role = ?", adminId, constanta.ADMIN).First(&dataAdmins)
	if tx.Error != nil {
		return entity.AdminCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.AdminCore{}, errors.New(constanta.ERROR_ID_ROLE)
	}

	dataResponse := entity.AdminModelToAdminCore(dataAdmins)
	return dataResponse, nil
}

func (ar *AdminRepository) Update(adminId string, data entity.AdminCore) error {
	dataAdmins := entity.AdminCoreToAdminModel(data)

	tx := ar.db.Where("id = ?", adminId).Updates(&dataAdmins)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}

func (ar *AdminRepository) Delete(adminId string) error {
	dataAdmins := model.Admin{}

	result, _ := ar.SelectById(adminId)
	if result.Role == constanta.SUPERADMIN {
		return errors.New("tidak bisa menghapus super admin")
	}

	tx := ar.db.Where("id = ? AND role = ?", adminId, constanta.ADMIN).Delete(&dataAdmins)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_ID_ROLE)
	}

	return nil
}

func (ar *AdminRepository) FindByEmail(email string) error {
	dataAdmins := model.Admin{}

	tx := ar.db.Where("email = ?", email).First(&dataAdmins)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_EMAIL)
	}

	return nil
}

func (ar *AdminRepository) FindByEmailANDPassword(data entity.AdminCore) (entity.AdminCore, error) {
	dataAdmins := model.Admin{}

	tx := ar.db.Where("email = ?", data.Email).First(&dataAdmins)
	if tx.Error != nil {
		return entity.AdminCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.AdminCore{}, errors.New(constanta.ERROR_DATA_EMAIL)
	}

	if comparePass := helper.CompareHash(dataAdmins.Password, data.Password); !comparePass {
		return entity.AdminCore{}, errors.New(constanta.ERROR_PASSWORD)
	}

	dataResponse := entity.AdminModelToAdminCore(dataAdmins)
	return dataResponse, nil
}

// Manage Users
func (ar *AdminRepository) GetAllUsers() ([]user.UsersCore, error) {
	dataUsers := []userModel.Users{}

	tx := ar.db.Find(&dataUsers)
	if tx.Error != nil {
		return nil, tx.Error
	}

	dataResponse := user.ListUserModelToUserCore(dataUsers)
	return dataResponse, nil
}

func (ar *AdminRepository) GetByIdUser(userId string) (user.UsersCore, error) {
	dataUsers := userModel.Users{}

	tx := ar.db.Where("id = ?", userId).Find(&dataUsers)
	if tx.Error != nil {
		return user.UsersCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return user.UsersCore{}, errors.New(constanta.ERROR_DATA_ID)
	}

	dataResponse := user.UsersModelToUsersCore(dataUsers)
	return dataResponse, nil
}

func (ar *AdminRepository) DeleteUsers(userId string) error {
	dataUsers := userModel.Users{}

	tx := ar.db.Where("id = ?", userId).Delete(&dataUsers)
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New(constanta.ERROR_DATA_ID)
	}

	return nil
}

// GetByStatusReport implements entity.AdminRepositoryInterface.
func (ar *AdminRepository) GetByStatusReport(status string) ([]report.ReportCore, error) {
	dataReports := []reportModel.Report{}
	var result *gorm.DB

	if status != "" {
		result = ar.db.Where("status = ?", status).Find(&dataReports)
	} else {
		result = ar.db.Find(&dataReports)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	dataAllReport := report.ListReportModelToReportCore(dataReports)
	return dataAllReport, nil
}

// UpdateStatusReport implements entity.AdminRepositoryInterface.
func (ar *AdminRepository) UpdateStatusReport(id, status, reason string) (report.ReportCore, error) {
	dataReports := reportModel.Report{}

	errData := ar.db.Where("id = ?", id).First(&dataReports)
	if errData.Error != nil {
		return report.ReportCore{}, errData.Error
	}

	dataReports.Status = status
	dataReports.RejectionDescription = reason
	tx := ar.db.Save(&dataReports)
	if tx.Error != nil {
		return report.ReportCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return report.ReportCore{}, errors.New(constanta.ERROR_DATA_ID)
	}

	dataResponse := report.ReportModelToReportCore(dataReports)
	return dataResponse, nil
}

func (ar *AdminRepository) GetReportByID(id string) (report.ReportCore, error) {
    dataReports := reportModel.Report{}

    tx := ar.db.Where("id = ?", id).First(&dataReports)
    if tx.Error != nil {
        return report.ReportCore{}, tx.Error
    }

    if tx.RowsAffected == 0 {
        return report.ReportCore{}, errors.New(constanta.ERROR_DATA_ID)
    }
	
	dataResponse := report.ReportModelToReportCore(dataReports)
    return dataResponse, nil
}