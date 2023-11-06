package entity

import (
	"recything/features/user/model"
)

func UsersCoreToUsersModel(mainData UsersCore) model.Users {
	return model.Users{
		Username:          mainData.Username,
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

func ListUserCoreToUserModel(mainData []UsersCore) []model.Users {
	listUser := []model.Users{}
	for _, user := range mainData {
		userModel := UsersCoreToUsersModel(user)
		listUser = append(listUser, userModel)
	}
	return listUser
}

func UsersModelToUsersCore(mainData model.Users) UsersCore {
	return UsersCore{
		Id:                mainData.Id,
		Username:          mainData.Username,
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

func ListUserModelToUserCore(mainData []model.Users) []UsersCore {
	listUser := []UsersCore{}
	for _, user := range mainData {
		userModel := UsersModelToUsersCore(user)
		listUser = append(listUser, userModel)
	}
	return listUser
}