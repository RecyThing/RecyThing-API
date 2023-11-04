package repository

import (
	"recything/features/admin/entity"
	"recything/features/admin/model"

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

func (admin *AdminRepository) Insert(data entity.AdminCore) error {

	dataCreate := entity.AdminCoreToAdminModel(data)
	if err := admin.db.Create(&dataCreate).Error; err != nil {
		return err
	}

	return nil
}

func (admin *AdminRepository) SelectAll() ([]entity.AdminCore, error) {
	dataAdmin := []model.Admin{}

	if err := admin.db.Find(&dataAdmin).Error; err != nil {
		return nil, err
	}

	var dataAllAdmin []entity.AdminCore = entity.ListAdminModelToAdminCore(dataAdmin)
	return dataAllAdmin, nil
}

func (admin *AdminRepository) SelectById(id_admin, role string) (entity.AdminCore,error) {
	dataAdmin := model.Admin{}

	if err := admin.db.Where("id = ? AND role = ? ", id_admin, role).Find(&dataAdmin).Error; err != nil {
		return entity.AdminCore{},err
	}

	data := entity.AdminModelToAdminCore(dataAdmin)
	return data,nil
}

func (admin *AdminRepository) Update(id_admin string,data entity.AdminCore) (error) {
	
	dataAdmin := entity.AdminCoreToAdminModel(data)
	if err := admin.db.Where("id = ?", id_admin).Updates(&dataAdmin).Error; err != nil {
		return err
	}

	return nil
}

func (admin *AdminRepository) Delete(id_admin string) error {
	dataAdmin := model.Admin{}

	if err := admin.db.Delete(&dataAdmin, id_admin).Error; err != nil {
		return err
	}

	return nil
}