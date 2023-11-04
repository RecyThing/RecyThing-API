package entity

import (
	"recything/features/admin/dto"
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
		adminModel := AdminModelToAdminCore(admin)
		listAdmin = append(listAdmin, adminModel)
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


func AdminRequestToAdminCore(admin dto.AdminRequest) AdminCore {
	return AdminCore{
		Name:     admin.Name,
		Email:    admin.Email,
		Password: admin.Password,
	}
}

func AdminCoreToAdminResponse(admin AdminCore) dto.AdminRespon {
	return dto.AdminRespon{
		Name:     admin.Name,
		Email:    admin.Email,
	}
}
