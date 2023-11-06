package dto

import "recything/features/user/entity"

type UserLoginResponse struct {
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

func LoginResponse(email string) UserLoginResponse {
	return UserLoginResponse{
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
