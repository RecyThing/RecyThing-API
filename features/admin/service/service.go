package service

import (
	"errors"
	"recything/features/admin/entity"
	report "recything/features/report/entity"
	user "recything/features/user/entity"
	"recything/utils/constanta"
	"recything/utils/helper"
	"recything/utils/jwt"
	"recything/utils/validation"
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
		return entity.AdminCore{}, errEmpty
	}

	errEmail := validation.EmailFormat(data.Email)
	if errEmail != nil {
		return entity.AdminCore{}, errEmail
	}

	errLength := validation.MinLength(data.Password, 8)
	if errLength != nil {
		return entity.AdminCore{}, errLength
	}

	errFind := as.AdminRepository.FindByEmail(data.Email)
	if errFind == nil {
		return entity.AdminCore{}, errors.New("email sudah ada, gunakan email lain")
	}

	if data.ConfirmPassword != data.Password {
		return entity.AdminCore{}, errors.New("password tidak sesuai")
	}

	dataAdmins, errCreate := as.AdminRepository.Create(data)
	if errCreate != nil {
		return entity.AdminCore{}, errors.New("gagal membuat data admin")
	}

	return dataAdmins, nil
}

func (as *AdminService) GetAll() ([]entity.AdminCore, error) {

	dataAdmins, err := as.AdminRepository.SelectAll()
	if err != nil {
		return nil, errors.New("gagal mengambil semua data admin")
	}

	return dataAdmins, nil
}

func (as *AdminService) GetById(adminId string) (entity.AdminCore, error) {

	dataAdmins, err := as.AdminRepository.SelectById(adminId)
	if err != nil {
		return entity.AdminCore{}, errors.New("data admin tidak ada")
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
		return errors.New("gagal melakukan update data admin")
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
		return user.UsersCore{}, errors.New("")
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
func (as *AdminService) GetByStatusReport(status string) (data []report.ReportCore, err error) {
	switch status {
	case "perlu ditinjau", "diterima", "ditolak":
		data, err = as.AdminRepository.GetByStatusReport(status)
	case "":
		data, err = as.AdminRepository.GetByStatusReport("")
	default:
		return nil, errors.New("status tidak valid")
	}

	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateStatusReport implements entity.AdminServiceInterface.
func (as *AdminService) UpdateStatusReport(id string, status string, reason string) (report.ReportCore, error) {
	if id == "" {
		return report.ReportCore{}, errors.New("id tidak valid")
	}

	if status == "" {
		return report.ReportCore{}, errors.New("status tidak valid")
	}

	if status == "ditolak" && reason == "" {
		return report.ReportCore{}, errors.New("alasan harus diisi saat menolak laporan")
	}

	dataStatus, err := as.AdminRepository.GetReportById(id)
	if err != nil {
		return report.ReportCore{}, errors.New("gagal mengambil data laporan")
	}

	if dataStatus.Status == "diterima" || dataStatus.Status == "ditolak" {
		return report.ReportCore{}, errors.New("status sudah diterima atau ditolak, tidak bisa update data lagi")
	}

	data, err := as.AdminRepository.UpdateStatusReport(id, status, reason)
	if err != nil {
		return report.ReportCore{}, errors.New("gagal update status")
	}

	return data, nil
}

// GetReportById implements entity.AdminServiceInterface.
func (as *AdminService) GetReportById(id string) (report.ReportCore, error) {
	if id == "" {
		return report.ReportCore{}, errors.New(constanta.ERROR_ID_INVALID)
	}

	idReport, err := as.AdminRepository.GetReportById(id)
	return idReport, err
}
