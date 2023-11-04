package dto

import "recything/features/user/entity"

type UserRegister struct {
	Id              string `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RequestRegister(dataRequest UserRegister) entity.UsersCore {
	return entity.UsersCore {
		Id:              dataRequest.Id,
		Username:        dataRequest.Username,
		Email:           dataRequest.Email,
		Password:        dataRequest.Password,
		ConfirmPassword: dataRequest.ConfirmPassword,
	}
}
