package request

import "recything/features/admin/entity"

func AdminRequestToAdminCore(admin AdminRequest) entity.AdminCore {
	return entity.AdminCore{
		Name:     admin.Name,
		Email:    admin.Email,
		Password: admin.Password,
	}
}
