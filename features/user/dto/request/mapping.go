package request

import "recything/features/user/entity"

func UsersRequestRegisterToUsersCore(data UserRegister) entity.UsersCore {
	return entity.UsersCore{
		Id:              data.Id,
		Fullname:        data.Fullname,
		Email:           data.Email,
		Password:        data.Password,
		ConfirmPassword: data.ConfirmPassword,
	}
}

func UsersRequestUpdateToUsersCore(data UserUpdate) entity.UsersCore {
	return entity.UsersCore{
		Fullname:    data.Fullname,
		Phone:       data.Phone,
		DateOfBirth: data.DateOfBirth,
		Address:     data.Address,
		Purpose:     data.Purpose,
	}
}

func UsersRequestForgetPasswordToUsersCore(data UserForgetPassword) entity.UsersCore {
	return entity.UsersCore{
		Password:        data.Password,
		ConfirmPassword: data.ConfirmPassword,
	}
}

func UsersRequestLoginToUsersCore(data UserLogin) entity.UsersCore {
	return entity.UsersCore{
		Fullname: data.Email,
		Password: data.Password,
	}
}
