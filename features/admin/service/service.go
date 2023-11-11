package service

import (
	"errors"
	"recything/features/admin/entity"
	user "recything/features/user/entity"
	"recything/utils/jwt"
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

	err := as.AdminRepository.FindByEmail(data.Email)
	if err == nil{
		return entity.AdminCore{},errors.New("email sudah ada, gunakan email lain")
	}

	if data.ConfirmPassword != data.Password {
		return entity.AdminCore{}, errors.New("password tidak sesuai")
	}

	result, err := as.AdminRepository.Insert(data)
	if err != nil {
		return entity.AdminCore{}, errors.New("gagal membuat data admin")
	}

	return result, nil
}

func (as *AdminService) GetAll() ([]entity.AdminCore, error) {

	result, err := as.AdminRepository.SelectAll()

	if err != nil {
		return nil, errors.New("gagal mengambil semua data admin")
	}

	return result, nil
}

func (as *AdminService) GetById(adminId string) (entity.AdminCore, error) {

	result, err := as.AdminRepository.SelectById(adminId)
	if err != nil {
		return entity.AdminCore{}, errors.New("data admin tidak ada")
	}

	return result, nil
}

func (as *AdminService) UpdateById(adminId string, data entity.AdminCore) error {

	err := as.AdminRepository.Update(adminId, data)
	if err != nil {
		return errors.New("gagal melakukan update data admin")
	}

	return nil
}

func (as *AdminService) DeleteById(adminId string) error {

	err := as.AdminRepository.Delete(adminId)
	if err != nil {
		return errors.New("gagal menghapus data admin")
	}

	return nil
}

func (as *AdminService) FindByEmailANDPassword(data entity.AdminCore) (entity.AdminCore, string, error) {

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

	data, err := as.AdminRepository.SelectAllUsers()
	if err != nil {
		return nil, errors.New("")
	}

	return data, nil
}

func (as *AdminService) GetByIdUsers(userId string) (user.UsersCore, error) {

	data, err := as.AdminRepository.SelectByIdUsers(userId)

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
