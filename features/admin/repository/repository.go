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

func (ar *AdminRepository) Insert(data entity.AdminCore) (entity.AdminCore, error) {
	dataCreate := entity.AdminCoreToAdminModel(data)

	tx := ar.db.Create(&dataCreate)
	if tx.Error != nil {
		return entity.AdminCore{}, tx.Error
	}

	adminData := entity.AdminModelToAdminCore(dataCreate)
	return adminData, nil
}

func (ar *AdminRepository) SelectAll() ([]entity.AdminCore, error) {
	dataAdmin := []model.Admin{}

	tx := ar.db.Where("role = ? ", helper.ADMIN).Find(&dataAdmin)
	if tx.Error != nil {
		return nil, tx.Error
	}

	result := entity.ListAdminModelToAdminCore(dataAdmin)
	return result, nil
}

func (ar *AdminRepository) SelectById(adminId string) (entity.AdminCore, error) {
	dataAdmin := model.Admin{}

	tx := ar.db.Where("id = ? AND role = ?", adminId, helper.ADMIN).First(&dataAdmin)
	if tx.Error != nil {
		return entity.AdminCore{}, tx.Error
	}

	result := entity.AdminModelToAdminCore(dataAdmin)
	return result, nil
}

func (ar *AdminRepository) Update(adminId string, data entity.AdminCore) error {
	request := entity.AdminCoreToAdminModel(data)

	tx := ar.db.Where("id = ?", adminId).Updates(&request)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (ar *AdminRepository) Delete(adminId string) error {
	dataAdmin := model.Admin{}

	result, _ := ar.SelectById(adminId)
	if result.Role == helper.SUPERADMIN {
		return errors.New("can`t delete")
	}

	tx := ar.db.Where("id = ? AND role = ?", adminId, helper.ADMIN).Delete(&dataAdmin)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (ar *AdminRepository) FindByEmail(email string) error {
	dataAdmin := model.Admin{}

	tx := ar.db.Where("email = ?", email).First(&dataAdmin)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (ar *AdminRepository) FindByEmailANDPassword(data entity.AdminCore) (entity.AdminCore, error) {
	dataAdmin := model.Admin{}

	tx := ar.db.Where("email = ?", data.Email).First(&dataAdmin)
	if tx.Error != nil {
		return entity.AdminCore{}, tx.Error
	}

	if comparePass := helper.CompareHash(dataAdmin.Password, data.Password); !comparePass {
		return entity.AdminCore{}, errors.New("password tidak sama")
	}

	adminCore := entity.AdminModelToAdminCore(dataAdmin)
	return adminCore, nil
}

// Manage Users
func (ar *AdminRepository) SelectAllUsers() ([]user.UsersCore, error) {
	dataUser := []userModel.Users{}

	tx := ar.db.Find(&dataUser)
	if tx.Error != nil {
		return nil, tx.Error
	}

	var dataAllUser []user.UsersCore = user.ListUserModelToUserCore(dataUser)
	return dataAllUser, nil
}

func (ar *AdminRepository) SelectByIdUsers(userId string) (user.UsersCore, error) {
	dataUser := userModel.Users{}

	tx := ar.db.Where("id = ?", userId).Find(&dataUser)
	if tx.Error != nil {
		return user.UsersCore{}, tx.Error
	}

	data := user.UsersModelToUsersCore(dataUser)
	return data, nil
}

func (ar *AdminRepository) DeleteUsers(userId string) error {
	dataUser := userModel.Users{}

	tx := ar.db.Where("id = ?", userId).Delete(&dataUser)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// GetByStatusReport implements entity.AdminRepositoryInterface.
func (ar *AdminRepository) GetByStatusReport(status string) ([]report.ReportCore, error) {
	var dataReports []reportModel.Report
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
func (ar *AdminRepository) UpdateStatusReport(id string, status string) (report.ReportCore, error) {
	var usersData reportModel.Report

    errData := ar.db.Where("id = ?", id).First(&usersData).Error
    if errData != nil {
        if errors.Is(errData, gorm.ErrRecordNotFound) {
            return report.ReportCore{}, errors.New("data tidak ditemukan")
        }
        return report.ReportCore{}, errData
    }

    usersData.Status = status
    err := ar.db.Save(&usersData).Error
    if err != nil {
        return report.ReportCore{}, err
    }

	data := report.ReportModelToReportCore(usersData)

	return data, nil
}