package response

import "recything/features/admin/entity"

func AdminCoreToAdminResponse(admin entity.AdminCore) AdminRespon {
	return AdminRespon{
		ID:        admin.Id,
		Name:      admin.Name,
		Email:     admin.Email,
		Status:    admin.Status,
		CreatedAt: admin.CreatedAt,
		UpdatedAt: admin.UpdatedAt,
	}
}

func ListAdminCoreToAdminResponse(admins []entity.AdminCore) []AdminRespon {
	listAdmin := []AdminRespon{}
	for _, admin := range admins {
		adminResp := AdminCoreToAdminResponse(admin)
		listAdmin = append(listAdmin, adminResp)
	}
	return listAdmin
}