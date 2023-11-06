package dto

import "recything/features/user/entity"

type UserLoginResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserResponseProfile struct {
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	Purpose     string `json:"purpose"`
	Point       int    `json:"point"`
}

func LoginResponse(fullname, email string) UserLoginResponse {
	return UserLoginResponse{
		Fullname: fullname,
		Email:    email,
	}
}

func ResponseProfile(user entity.UsersCore) UserResponseProfile {
	return UserResponseProfile{
		Email:       user.Email,
		Fullname:    user.Fullname,
		Phone:       user.Phone,
		Address:     user.Address,
		DateOfBirth: user.DateOfBirth,
		Purpose:     user.Purpose,
		Point:       user.Point,
	}
}

type ResponseManageUsers struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func UsersCoreToResponseUsers(user entity.UsersCore) ResponseManageUsers {
	return ResponseManageUsers{
		Id:       user.Id,
		Fullname: user.Fullname,
		Email:    user.Email,
		Phone:    user.Phone,
	}
}

func UsersCoreToResponseUsersList(dataCore []entity.UsersCore) []ResponseManageUsers {
	var result []ResponseManageUsers
	for _, v := range dataCore {
		result = append(result, UsersCoreToResponseUsers(v))
	}
	return result
}
