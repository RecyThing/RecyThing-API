package entity

import (
	"recything/features/admin/model"
)

func AdminModelToAdminCore(admin model.Admin) AdminCore {
	return AdminCore{
		Id:        admin.Id,
		Name:      admin.Name,
		Role:      admin.Role,
		Email:     admin.Email,
		Password:  admin.Password,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}

}

func ListAdminModelToAdminCore(admins []model.Admin) []AdminCore {
	listAdmin := []AdminCore{}
	for _, admin := range admins {
		adminCore := AdminModelToAdminCore(admin)
		listAdmin = append(listAdmin, adminCore)
	}
	return listAdmin
}

func AdminCoreToAdminModel(admin AdminCore) model.Admin {
	return model.Admin{
		Id:        admin.Id,
		Name:      admin.Name,
		Role:      admin.Role,
		Email:     admin.Email,
		Password:  admin.Password,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}

func ListAdminCoreToAdminModel(admins []AdminCore) []model.Admin {
	listAdmin := []model.Admin{}
	for _, admin := range admins {
		adminModel := AdminCoreToAdminModel(admin)
		listAdmin = append(listAdmin, adminModel)
	}
	return listAdmin
}

