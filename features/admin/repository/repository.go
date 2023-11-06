package repository

import (
	"errors"
	"recything/features/admin/entity"
	user "recything/features/user/entity"
	"recything/features/admin/model"
	userModel "recything/features/user/model"
	"recything/utils/helper"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db: db}
}

func (admin *AdminRepository) Insert(data entity.AdminCore) (entity.AdminCore, error) {

	dataCreate := entity.AdminCoreToAdminModel(data)
	if err := admin.db.Create(&dataCreate).Error; err != nil {
		return entity.AdminCore{}, err
	}
	adminData := entity.AdminModelToAdminCore(dataCreate)
	return adminData, nil
}

func (admin *AdminRepository) SelectAll() ([]entity.AdminCore, error) {
	dataAdmin := []model.Admin{}
	if err := admin.db.Find(&dataAdmin).Error; err != nil {
		return nil, err
	}

	var dataAllAdmin []entity.AdminCore = entity.ListAdminModelToAdminCore(dataAdmin)
	return dataAllAdmin, nil
}

func (admin *AdminRepository) SelectById(adminId string) (entity.AdminCore, error) {
	dataAdmin := model.Admin{}

	if err := admin.db.Where("id = ?", adminId).Find(&dataAdmin).Error; err != nil {
		return entity.AdminCore{}, err
	}

	data := entity.AdminModelToAdminCore(dataAdmin)
	return data, nil
}

func (admin *AdminRepository) Update(adminId string, data entity.AdminCore) error {

	dataAdmin := entity.AdminCoreToAdminModel(data)
	if err := admin.db.Where("id = ?", adminId).Updates(&dataAdmin).Error; err != nil {
		return err
	}

	return nil
}

func (admin *AdminRepository) Delete(adminId string) error {
	dataAdmin := model.Admin{}

	adminCore, _ := admin.SelectById(adminId)
	if adminCore.Role == helper.SUPERADMIN {
		return errors.New("can`t delete")
	}

	if err := admin.db.Where("id = ? AND role = ?", adminId, helper.ADMIN).Delete(&dataAdmin).Error; err != nil {

		return err
	}

	return nil
}

func (admin *AdminRepository) FindByEmailANDPassword(email, password string) (entity.AdminCore, error) {
	var err error
	adminModel := model.Admin{}

	if err = admin.db.Where("email = ?", email).First(&adminModel).Error; err != nil {
		return entity.AdminCore{}, err
	}

	if comparePass := helper.CompareHash(adminModel.Password, password); !comparePass {
		return entity.AdminCore{}, err
	}

	adminCore := entity.AdminModelToAdminCore(adminModel)
	return adminCore, nil
}

//Manage Users
func (admin *AdminRepository) SelectAllUsers() ([]user.UsersCore, error){
	dataUser := []userModel.Users{}
	if err := admin.db.Find(&dataUser).Error; err != nil {
		return nil, err
	}

	var dataAllUser []user.UsersCore = user.ListUserModelToUserCore(dataUser)
	return dataAllUser, nil
}

func (admin *AdminRepository) SelectByIdUsers(userId string) (user.UsersCore, error){
	dataUser := userModel.Users{}

	if err := admin.db.Where("id = ?", userId).Find(&dataUser).Error; err != nil {
		return user.UsersCore{}, err
	}

	data := user.UsersModelToUsersCore(dataUser)
	return data, nil
}

func (admin *AdminRepository) DeleteUsers(userId string) error{
	dataUser := userModel.Users{}

	if err := admin.db.Where("id = ?", userId).Delete(&dataUser).Error; err != nil {

		return err
	}

	return nil
}