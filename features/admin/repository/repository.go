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
	"recything/utils/pagination"

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

func (ar *AdminRepository) SelectAll(page, limit int, fullName string) ([]entity.AdminCore,pagination.PageInfo, error) {
	dataAdmins := []model.Admin{}
	offsetInt := (page - 1) * limit

	var totalCount int64
	if fullName == "" {
		tx := ar.db.Model(&model.Admin{}).Where("role = ? ", constanta.ADMIN).Count(&totalCount)
		if tx.Error != nil {
			return nil, pagination.PageInfo{},tx.Error
		}

		tx = ar.db.Limit(limit).Offset(offsetInt).Where("role = ? ", constanta.ADMIN).Find(&dataAdmins)
		if tx.Error != nil {
			return nil,pagination.PageInfo{}, tx.Error
		}

		
	}
	if fullName != "" {

		tx := ar.db.Model(&model.Admin{}).Where("role = ? AND fullname LIKE ?", constanta.ADMIN, "%"+fullName+"%").Count(&totalCount)
		if tx.Error != nil {
			return nil, pagination.PageInfo{},tx.Error
		}

		tx = ar.db.Limit(limit).Offset(offsetInt).Where("role = ? AND fullname LIKE ?", constanta.ADMIN, "%"+fullName+"%").Find(&dataAdmins)
		if tx.Error != nil {
			return nil,pagination.PageInfo{}, tx.Error
		}
		
	}

	dataResponse := entity.ListAdminModelToAdminCore(dataAdmins)
	paginationInfo := pagination.CalculateData(int(totalCount), limit, page)

	return dataResponse,paginationInfo, nil
}

func (ar *AdminRepository) SelectById(adminId string) (entity.AdminCore, error) {
	dataAdmins := model.Admin{}

	tx := ar.db.Where("id = ? AND role = ?", adminId, constanta.ADMIN).First(&dataAdmins)
	if tx.Error != nil {
		return entity.AdminCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return entity.AdminCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
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
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
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
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
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
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
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
		return entity.AdminCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
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
		return user.UsersCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
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
		return errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	return nil
}

// GetByStatusReport implements entity.AdminRepositoryInterface.
func (ar *AdminRepository) GetAllReport(status, name, id string, page, limit int) ([]report.ReportCore, pagination.PageInfo, error) {
	dataReports := []reportModel.Report{}

	offset := (page - 1) * limit
	query := ar.db.Model(&reportModel.Report{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if name != "" {
		query = query.Joins("JOIN users AS u ON reports.users_id = u.id").
			Where("u.fullname LIKE ?", "%"+name+"%")
	}

	if id != "" {
		query = query.Where("id = ?", id)
	}

	var totalCount int64
	if err := query.Count(&totalCount).Error; err != nil {
		return nil, pagination.PageInfo{}, err
	}

	query = query.Offset(offset).Limit(limit)

	if err := query.Find(&dataReports).Error; err != nil {
		return nil, pagination.PageInfo{}, err
	}

	dataAllReport := report.ListReportModelToReportCore(dataReports)
	paginationInfo := pagination.CalculateData(int(totalCount), limit, page)

	return dataAllReport, paginationInfo, nil
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
		return report.ReportCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dataResponse := report.ReportModelToReportCore(dataReports)
	return dataResponse, nil
}

func (ar *AdminRepository) GetReportById(id string) (report.ReportCore, error) {
	dataReports := reportModel.Report{}

	tx := ar.db.Preload("Images").Where("id = ?", id).First(&dataReports)
	if tx.Error != nil {
		return report.ReportCore{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return report.ReportCore{}, errors.New(constanta.ERROR_DATA_NOT_FOUND)
	}

	dataResponse := report.ReportModelToReportCore(dataReports)
	return dataResponse, nil
}
