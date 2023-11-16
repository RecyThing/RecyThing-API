package request

import "recything/features/admin/entity"

func AdminRequestToAdminCore(data AdminRequest) entity.AdminCore {
	return entity.AdminCore{
		Fullname:        data.Fullname,
		Email:           data.Email,
		Password:        data.Password,
		ConfirmPassword: data.ConfirmPassword,
	}
}

func RequestLoginToAdminCore(data AdminLogin) entity.AdminCore {
	return entity.AdminCore{
		Email:    data.Email,
		Password: data.Password,
	}
}
