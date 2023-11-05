package dto

import "recything/features/user/entity"

type UserLoginResponse struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserResponseProfile struct {
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
}

func LoginResponse(id, email, token string) UserLoginResponse {
	return UserLoginResponse{
		Id:    id,
		Email: email,
		Token: token,
	}
}

func ResponseProfile(user entity.UsersCore) UserResponseProfile {
	return UserResponseProfile{
		Fullname:    user.Fullname,
		Email:       user.Email,
		Phone:       user.Phone,
		Address:     user.Address,
		DateOfBirth: user.DateOfBirth,
	}
}