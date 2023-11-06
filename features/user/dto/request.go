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

type UserUpdate struct {
	Username    string `json:"username"`
	Fullname    string `json:"fullname"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	Purpose     string `json:"purpose"`
}

type UserForgetPassword struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func RequestRegister(dataRequest UserRegister) entity.UsersCore {
	return entity.UsersCore{
		Id:              dataRequest.Id,
		Username:        dataRequest.Username,
		Email:           dataRequest.Email,
		Password:        dataRequest.Password,
		ConfirmPassword: dataRequest.ConfirmPassword,
	}
}

func RequestUpdate(dataUpdate UserUpdate) entity.UsersCore {
	return entity.UsersCore{
		Username:    dataUpdate.Username,
		Fullname:    dataUpdate.Fullname,
		Phone:       dataUpdate.Phone,
		DateOfBirth: dataUpdate.DateOfBirth,
		Address:     dataUpdate.Address,
		Purpose:     dataUpdate.Purpose,
	}
}

func RequestForgetPassword(dataUpdate UserForgetPassword) entity.UsersCore {
	return entity.UsersCore{
		Password: dataUpdate.Password,
		ConfirmPassword: dataUpdate.ConfirmPassword,
	}
}
