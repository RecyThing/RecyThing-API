package entity

import "recything/features/user/model"

func UsersCoreToUsersModel(mainData UsersCore) model.Users {
	return model.Users{
		Email:             mainData.Email,
		Password:          mainData.Password,
		Fullname:          mainData.Fullname,
		Phone:             mainData.Phone,
		Address:           mainData.Address,
		DateOfBirth:       mainData.DateOfBirth,
		Purpose:           mainData.Purpose,
		Point:             mainData.Point,
		IsVerified:        mainData.IsVerified,
		VerificationToken: mainData.VerificationToken,
	}
}

func UsersModelToUsersCore(mainData model.Users) UsersCore {
	return UsersCore{
		Id:                mainData.Id,
		Email:             mainData.Email,
		Password:          mainData.Password,
		Fullname:          mainData.Fullname,
		Phone:             mainData.Phone,
		Address:           mainData.Address,
		DateOfBirth:       mainData.DateOfBirth,
		Purpose:           mainData.Purpose,
		Point:             mainData.Point,
		IsVerified:        mainData.IsVerified,
		VerificationToken: mainData.VerificationToken,
		CreatedAt:         mainData.CreatedAt,
		UpdatedAt:         mainData.UpdatedAt,
	}
}
