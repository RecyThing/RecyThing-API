package service

import (
	"errors"
	"recything/features/admin/entity"
	report "recything/features/report/entity"
	user "recything/features/user/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"recything/utils/pagination"
	"recything/utils/validation"
	"strconv"
)

type AdminService struct {
	AdminRepository entity.AdminRepositoryInterface
}

func NewAdminService(ar entity.AdminRepositoryInterface) entity.AdminServiceInterface {
	return &AdminService{
		AdminRepository: ar,
	}
}

func (as *AdminService) Create(data entity.AdminCore) (entity.AdminCore, error) {

	errEmpty := validation.CheckDataEmpty(data.Fullname, data.Email, data.Password, data.ConfirmPassword)
	if errEmpty != nil {
		return entity.AdminCore{}, errors.New(constanta.ERROR_EMPTY)
	}

	errEmail := validation.EmailFormat(data.Email)
	if errEmail != nil {
		return entity.AdminCore{}, errors.New(constanta.ERROR_FORMAT_EMAIL)
	}

	errLength := validation.MinLength(data.Password, 8)
	if errLength != nil {
		return entity.AdminCore{}, errors.New(constanta.ERROR_LENGTH_PASSWORD)
	}

	errFind := as.AdminRepository.FindByEmail(data.Email)
	if errFind == nil {
		return entity.AdminCore{}, errors.New(constanta.ERROR_EMAIL_EXIST)
	}

	if data.ConfirmPassword != data.Password {
		return entity.AdminCore{}, errors.New(constanta.ERROR_CONFIRM_PASSWORD)
	}

	dataAdmins, errCreate := as.AdminRepository.Create(data)
	if errCreate != nil {
		return entity.AdminCore{}, errors.New("gagal membuat data admin")
	}

	return dataAdmins, nil
}

func (as *AdminService) GetAll(limit, page, fullName string) ([]entity.AdminCore, pagination.PageInfo, error) {

	var limitInt int
	var pageInt int
	var err error
	if limit == "" {
		limitInt = 10
	}
	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return nil, pagination.PageInfo{}, errors.New("limit harus berupa angka")
		}
	}

	if page == "" {
		pageInt = 1
	}
	if page != "" {
		pageInt, err = strconv.Atoi(page)
		if err != nil {
			return nil, pagination.PageInfo{}, errors.New("page harus berupa angka")
		}
	}

	if pageInt <= 0 {
		pageInt = 1
	}

	maxLimit := 10

	if limitInt <= 0 || limitInt > maxLimit {
		limitInt = maxLimit
	}

	dataAdmins, pagnationInfo, err := as.AdminRepository.SelectAll(pageInt, limitInt, fullName)
	if err != nil {
		return nil, pagination.PageInfo{}, errors.New("gagal mengambil semua data admin")
	}

	return dataAdmins, pagnationInfo, nil
}

func (as *AdminService) GetById(adminId string) (entity.AdminCore, error) {

	dataAdmins, err := as.AdminRepository.SelectById(adminId)
	if err != nil {
		return entity.AdminCore{}, err
	}

	return dataAdmins, nil
}

func (as *AdminService) UpdateById(adminId string, data entity.AdminCore) error {

	if data.Email != "" {
		errEmail := validation.EmailFormat(data.Email)
		if errEmail != nil {
			return errEmail
		}

	}

	if data.Password != "" {
		errLength := validation.MinLength(data.Password, 8)
		if errLength != nil {
			return errLength
		}

		HashPassword, errHash := helper.HashPassword(data.Password)
		if errHash != nil {
			return errors.New("error hash password")
		}
		data.Password = HashPassword

	}

	err := as.AdminRepository.Update(adminId, data)
	if err != nil {
		return err
	}

	return nil
}

func (as *AdminService) DeleteById(adminId string) error {

	err := as.AdminRepository.Delete(adminId)
	if err != nil {
		return err
	}

	return nil
}

func (as *AdminService) FindByEmailANDPassword(data entity.AdminCore) (entity.AdminCore, string, error) {

	errEmpty := validation.CheckDataEmpty(data.Email, data.Password)
	if errEmpty != nil {
		return entity.AdminCore{}, "", errEmpty
	}

	errEmail := validation.EmailFormat(data.Email)
	if errEmail != nil {
		return entity.AdminCore{}, "", errEmail
	}

	data, err := as.AdminRepository.FindByEmailANDPassword(data)
	if err != nil {
		return entity.AdminCore{}, "", errors.New("data tidak ada")
	}

	token, errToken := jwt.CreateToken(data.Id, data.Role)
	if errToken != nil {
		return entity.AdminCore{}, "", errors.New("gagal membuat token session")
	}
	return data, token, nil
}

//Manage Users

func (as *AdminService) GetAllUsers() ([]user.UsersCore, error) {

	data, err := as.AdminRepository.GetAllUsers()
	if err != nil {
		return nil, errors.New("")
	}

	return data, nil
}

func (as *AdminService) GetByIdUsers(userId string) (user.UsersCore, error) {

	data, err := as.AdminRepository.GetByIdUser(userId)

	if data == (user.UsersCore{}) {
		return user.UsersCore{}, errors.New("null")
	}

	if err != nil {
		return user.UsersCore{}, err
	}

	return data, nil
}

func (as *AdminService) DeleteUsers(userId string) error {

	err := as.AdminRepository.DeleteUsers(userId)
	if err != nil {
		return err
	}

	return nil
}

// Manage Reporting
// GetByStatusReport implements entity.AdminServiceInterface.
func (as *AdminService) GetAllReport(status, name, id, page, limit string) (data []report.ReportCore, paginationInfo pagination.PageInfo, err error) {
	pageInt, limitInt, validationErr := validation.ValidateTypePaginationParameter(limit, page)
	if validationErr != nil {
		return nil, pagination.PageInfo{}, validationErr
	}

	pageValid, limitValid := validation.ValidatePaginationParameters(pageInt, limitInt)

	validStatus := map[string]bool{
		"perlu ditinjau": true,
		"diterima":       true,
		"ditolak":        true,
	}

	if _, ok := validStatus[status]; status != "" && !ok {
		return nil, pagination.PageInfo{}, errors.New("status tidak valid")
	}

	if status != "" || name != "" || id != "" {
		data, paginationInfo, err = as.AdminRepository.GetAllReport(status, name, id, pageValid, limitValid)
	} else {
		data, paginationInfo, err = as.AdminRepository.GetAllReport("", "", "", pageValid, limitValid)
	}

	if err != nil {
		return nil, pagination.PageInfo{}, err
	}

	return data, paginationInfo, nil
}

// UpdateStatusReport implements entity.AdminServiceInterface.
func (as *AdminService) UpdateStatusReport(id string, status string, reason string) (report.ReportCore, error) {

	errEmpty := validation.CheckDataEmpty(status)
	if errEmpty != nil {
		return report.ReportCore{}, errEmpty
	}

	if status == "diterima" && reason != "" {
		return report.ReportCore{}, errors.New("tidak perlu memberikan alasan laporan")
	}

	if status == "ditolak" && reason == "" {
		return report.ReportCore{}, errors.New("alasan harus dilengkapi saat menolak laporan")
	}

	dataStatus, err := as.AdminRepository.GetReportById(id)
	if err != nil {
		return report.ReportCore{}, err
	}

	if dataStatus.Status == "diterima" || dataStatus.Status == "ditolak" {
		return report.ReportCore{}, errors.New("status sudah diterima atau ditolak")
	}

	data, err := as.AdminRepository.UpdateStatusReport(id, status, reason)
	if err != nil {
		return report.ReportCore{}, errors.New("gagal update status")
	}

	return data, nil
}

// GetReportById implements entity.AdminServiceInterface.
func (as *AdminService) GetReportById(id string) (report.ReportCore, error) {
	idReport, err := as.AdminRepository.GetReportById(id)
	if err != nil {
		return report.ReportCore{}, err
	}
	return idReport, err
}
