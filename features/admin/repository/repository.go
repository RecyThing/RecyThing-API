package repository

import (
	"errors"
	"recything/features/admin/entity"
	"recything/features/admin/model"

	report "recything/features/report/entity"
	reportModel "recything/features/report/model"

	user "recything/features/user/entity"
	userModel "recything/features/user/model"
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

	err := ar.db.Create(&dataAdmins).Error
	if err != nil {
		return entity.AdminCore{}, err
	}

	dataResponse := entity.AdminModelToAdminCore(dataAdmins)
	return dataResponse, nil
}

func (ar *AdminRepository) SelectAll() ([]entity.AdminCore, error) {
	dataAdmins := []model.Admin{}

	err := ar.db.Where("role = ? ", helper.ADMIN).Find(&dataAdmins).Error
	if err != nil {
		return nil, err
	}

	dataResponse := entity.ListAdminModelToAdminCore(dataAdmins)
	return dataResponse, nil
}

func (ar *AdminRepository) SelectById(adminId string) (entity.AdminCore, error) {
	dataAdmins := model.Admin{}

	err := ar.db.Where("id = ? AND role = ?", adminId, helper.ADMIN).First(&dataAdmins).Error
	if err != nil {
		return entity.AdminCore{}, err
	}

	dataResponse := entity.AdminModelToAdminCore(dataAdmins)
	return dataResponse, nil
}

func (ar *AdminRepository) Update(adminId string, data entity.AdminCore) error {
	dataAdmins := entity.AdminCoreToAdminModel(data)

	err := ar.db.Where("id = ?", adminId).Updates(&dataAdmins).Error
	if err != nil {
		return err
	}

	return nil
}

func (ar *AdminRepository) Delete(adminId string) error {
	dataAdmins := model.Admin{}

	result, _ := ar.SelectById(adminId)
	if result.Role == helper.SUPERADMIN {
		return errors.New("can`t delete")
	}

	err := ar.db.Where("id = ? AND role = ?", adminId, helper.ADMIN).Delete(&dataAdmins).Error
	if err != nil {
		return err
	}

	return nil
}

func (ar *AdminRepository) FindByEmail(email string) error {
	dataAdmins := model.Admin{}

	err := ar.db.Where("email = ?", email).First(&dataAdmins).Error
	if err != nil {
		return err
	}

	return nil
}

func (ar *AdminRepository) FindByEmailANDPassword(data entity.AdminCore) (entity.AdminCore, error) {
	dataAdmins := model.Admin{}

	err := ar.db.Where("email = ?", data.Email).First(&dataAdmins).Error
	if err != nil {
		return entity.AdminCore{}, err
	}

	if comparePass := helper.CompareHash(dataAdmins.Password, data.Password); !comparePass {
		return entity.AdminCore{}, errors.New("password tidak sama")
	}

	dataResponse := entity.AdminModelToAdminCore(dataAdmins)
	return dataResponse, nil
}

// Manage Users
func (ar *AdminRepository) GetAllUsers() ([]user.UsersCore, error) {
	dataUsers := []userModel.Users{}

	err := ar.db.Find(&dataUsers).Error
	if err != nil {
		return nil, err
	}

	dataResponse := user.ListUserModelToUserCore(dataUsers)
	return dataResponse, nil
}

func (ar *AdminRepository) GetByIdUser(userId string) (user.UsersCore, error) {
	dataUsers := userModel.Users{}

	err := ar.db.Where("id = ?", userId).Find(&dataUsers).Error
	if err != nil {
		return user.UsersCore{}, err
	}

	dataResponse := user.UsersModelToUsersCore(dataUsers)
	return dataResponse, nil
}

func (ar *AdminRepository) DeleteUsers(userId string) error {
	dataUsers := userModel.Users{}

	err := ar.db.Where("id = ?", userId).Delete(&dataUsers).Error
	if err != nil {
		return err
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

// // UpdateStatusReport implements entity.AdminRepositoryInterface.
func (ar *AdminRepository) UpdateStatusReport(id ,status string) (report.ReportCore, error) {
	dataReports := reportModel.Report{}

	errData := ar.db.Where("id = ?", id).First(&dataReports).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return report.ReportCore{}, errors.New("data tidak ditemukan")
		}
		return report.ReportCore{}, errData
	}

	dataReports.Status = status
	err := ar.db.Save(&dataReports).Error
	if err != nil {
		return report.ReportCore{}, err
	}

	dataResponse := report.ReportModelToReportCore(dataReports)
	return dataResponse, nil
}
