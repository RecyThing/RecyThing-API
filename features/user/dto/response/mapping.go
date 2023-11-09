package response

import "recything/features/user/entity"

func UsersCoreToLoginResponse(data entity.UsersCore) UserLoginResponse {
	return UserLoginResponse{
		Fullname: data.Fullname,
		Email:    data.Email,
	}
}

func UsersCoreToResponseProfile(data entity.UsersCore) UserResponseProfile {
	return UserResponseProfile{
		Email:       data.Email,
		Fullname:    data.Fullname,
		Phone:       data.Phone,
		Address:     data.Address,
		DateOfBirth: data.DateOfBirth,
		Purpose:     data.Purpose,
		Point:       data.Point,
	}
}

func UsersCoreToResponseUsers(data entity.UsersCore) ResponseManageUsers {
	return ResponseManageUsers{
		Id:       data.Id,
		Fullname: data.Fullname,
		Email:    data.Email,
		Phone:    data.Phone,
	}
}

func UsersCoreToResponseUsersList(dataCore []entity.UsersCore) []ResponseManageUsers {
	var result []ResponseManageUsers
	for _, v := range dataCore {
		result = append(result, UsersCoreToResponseUsers(v))
	}
	return result
}