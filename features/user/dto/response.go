package dto

import "recything/features/user/entity"

type UserLoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserResponseProfile struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	Purpose     string `json:"purpose"`
	Point       int    `json:"point"`
}

func LoginResponse(username, email string) UserLoginResponse {
	return UserLoginResponse{
		Username: username,
		Email:    email,
	}
}

func ResponseProfile(user entity.UsersCore) UserResponseProfile {
	return UserResponseProfile{
		Username:    user.Username,
		Email:       user.Email,
		Fullname:    user.Fullname,
		Phone:       user.Phone,
		Address:     user.Address,
		DateOfBirth: user.DateOfBirth,
		Purpose:     user.Purpose,
		Point:       user.Point,
	}
}
