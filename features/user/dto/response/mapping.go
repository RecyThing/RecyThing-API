package response

import "recything/features/user/entity"

func UsersCoreToUsersCreateResponse(data entity.UsersCore) UserCreateResponse {
	return UserCreateResponse{
		Id:       data.Id,
		Fullname: data.Fullname,
		Email:    data.Email,
	}
}

func UsersCoreToLoginResponse(data entity.UsersCore, token string) UserLoginResponse {
	return UserLoginResponse{
		Id:       data.Id,
		Fullname: data.Fullname,
		Email:    data.Email,
		Token:    token,
	}
}

func UsersCoreToResponseProfile(data entity.UsersCore) UserResponseProfile {
	return UserResponseProfile{
		Id:          data.Id,
		Fullname:    data.Fullname,
		Email:       data.Email,
		DateOfBirth: data.DateOfBirth,
		Phone:       data.Phone,
		Address:     data.Address,
		Purpose:     data.Purpose,
		Point:       data.Point,
	}
}

func UsersCoreToResponseManageUsers(data entity.UsersCore) UserResponseManageUsers {
	return UserResponseManageUsers{
		Id:       data.Id,
		Fullname: data.Fullname,
		Email:    data.Email,
		Point:    data.Point,
	}
}

func UsersCoreToResponseManageUsersList(dataCore []entity.UsersCore) []UserResponseManageUsers {
	var result []UserResponseManageUsers
	for _, v := range dataCore {
		result = append(result, UsersCoreToResponseManageUsers(v))
	}
	return result
}

func UsersCoreToResponseDetailManageUsers(data entity.UsersCore) UserResponseDetailManageUsers {
	return UserResponseDetailManageUsers{
		Id:          data.Id,
		Fullname:    data.Fullname,
		Email:       data.Email,
		Point:       data.Point,
		Address:     data.Address,
		DateOfBirth: data.DateOfBirth,
		Purpose:     data.Purpose,
	}
}
