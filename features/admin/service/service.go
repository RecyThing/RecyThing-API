package service

import (
	"errors"
	"recything/features/admin/entity"
	"recything/utils/jwt"
)

type AdminService struct {
	AdminRepository entity.AdminRepositoryInterface
}

func NewAdminService(admin entity.AdminRepositoryInterface) *AdminService {
	return &AdminService{
		AdminRepository: admin,
	}
}

func (admin *AdminService) Create(data entity.AdminCore) (entity.AdminCore, error) {

	dataAdmin, err := admin.AdminRepository.Insert(data)
	if err != nil {
		return entity.AdminCore{}, err
	}

	return dataAdmin, nil
}

func (admin *AdminService) GetAll() ([]entity.AdminCore, error) {

	data, err := admin.AdminRepository.SelectAll()
	if err != nil {
		return nil, errors.New("")
	}

	return data, nil
}

func (admin *AdminService) GetById(adminId string) (entity.AdminCore, error) {

	data, err := admin.AdminRepository.SelectById(adminId)

	if data == (entity.AdminCore{}) {
		return entity.AdminCore{}, errors.New("null")
	}

	if err != nil {
		return entity.AdminCore{}, errors.New("")
	}

	return data, nil
}

func (admin *AdminService) UpdateById(adminId string, data entity.AdminCore) error {

	err := admin.AdminRepository.Update(adminId, data)
	if err != nil {
		return errors.New("")
	}

	return nil
}

func (admin *AdminService) DeleteById(adminId string) error {

	err := admin.AdminRepository.Delete(adminId)
	if err != nil {
		return err
	}

	return nil
}

func (admin *AdminService) FindByEmailANDPassword(email, password string) (entity.AdminCore, string, error) {
	data, err := admin.AdminRepository.FindByEmailANDPassword(email, password)
	if err != nil {
		return entity.AdminCore{}, "", errors.New("Gagal woy")
	}

	token, errToken := jwt.CreateToken(data.Id, data.Role)
	if errToken != nil {
		return entity.AdminCore{}, "", errToken
	}

	return data, token, nil
}

//Manage Users

func (admin *AdminService) GetAllUsers() ([]entity.AdminCore, error){
	
	data, err := admin.AdminRepository.SelectAll()
	if err != nil {
		return nil, errors.New("")
	}

	return data, nil
}

func (admin *AdminService) GetByIdUsers(userId string) (entity.AdminCore, error){
	
	data, err := admin.AdminRepository.SelectById(userId)

	if data == (entity.AdminCore{}) {
		return entity.AdminCore{}, errors.New("null")
	}

	if err != nil {
		return entity.AdminCore{}, errors.New("")
	}

	return data, nil
}

func (admin *AdminService) DeleteUsers(userId string) error{
	
	err := admin.AdminRepository.Delete(userId)
	if err != nil {
		return err
	}

	return nil
}