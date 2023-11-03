package repository

import (
	"recything/features/admin/entity"
	"recything/features/admin/model"

	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) entity.AdminRepositoryInterface {
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


func (admin *AdminRepository) GetAll() ([]entity.AdminCore, error) {
	dataAdmin := []model.Admin{}

	if err := admin.db.Find(&dataAdmin).Error; err != nil {
		return nil, err
	}

	var dataAllAdmin []entity.AdminCore = entity.ListAdminModelToAdminCore(dataAdmin)
	return dataAllAdmin, nil
}


func (admin *AdminRepository) GetById(id_admin, role string) (entity.AdminCore, error) {
	dataAdmin := model.Admin{}

	if err := admin.db.Where("id = ? AND role = ? ", id_admin,role).Find(&dataAdmin).Error; err != nil {
		return entity.AdminCore{}, err
	}

	if err := admin.db.Find(&dataAdmin).Error; err != nil {
		return entity.AdminCore{}, err
	}

	data := entity.AdminModelToAdminCore(dataAdmin)
	return data, nil
}

