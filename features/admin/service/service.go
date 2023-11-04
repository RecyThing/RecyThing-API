package service

import (
	"errors"
	"recything/features/admin/entity"
)

type AdminService struct {
	AdminRepository entity.AdminRepositoryInterface
}

func NewAdminService(admin entity.AdminRepositoryInterface) *AdminService{
	return &AdminService{AdminRepository: admin}
}

func (admin *AdminService) Create(data entity.AdminCore) error {

	err := admin.AdminRepository.Insert(data)
	if err != nil {
		return errors.New("")
	}

	return err
}

func (admin *AdminService) GetAll() ([]entity.AdminCore, error) {

	data, err := admin.AdminRepository.SelectAll()
	if err != nil {
		return nil, errors.New("")
	}

	return data, err
}

func (admin *AdminService) GetById(id_admin, role string) (entity.AdminCore, error) {

	data, err := admin.AdminRepository.SelectById(id_admin, role)

	if data == (entity.AdminCore{}) {
		return entity.AdminCore{}, errors.New("null")
	}

	if err != nil {
		return entity.AdminCore{}, errors.New("")
	}

	return data, nil
}

func (admin *AdminService) UpdateById(id_admin, role string, data entity.AdminCore) error {

	err := admin.AdminRepository.Update(id_admin, data)
	if err != nil {
		return errors.New("")
	}

	return err
}

func (admin *AdminService) DeleteById(id_admin string) error {

	err := admin.AdminRepository.Delete(id_admin)
	if err != nil {
		return errors.New("")
	}

	return err
}
