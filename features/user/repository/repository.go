package repository

import (
	"errors"
	"recything/features/user/entity"
	"recything/features/user/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UsersRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// Register implements entity.UsersRepositoryInterface.
func (ur *userRepository) Register(data entity.UsersCore) error {
	request := entity.UsersCoreToUsersModel(data)

	err := ur.db.Create(&request).Error
	if err != nil {
		return err
	}

	return nil
}

// GetById implements entity.UsersRepositoryInterface.
func (ur *userRepository) GetById(id string) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	err := ur.db.Where("id = ?", id).First(&dataUsers).Error
	if err != nil {
		return entity.UsersCore{}, err
	}

	dataResponse := entity.UsersModelToUsersCore(dataUsers)
	return dataResponse, nil
}

func (ur *userRepository) FindByEmail(email string) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	err := ur.db.Where("email = ?", email).First(&dataUsers).Error

	if err != nil {
		return entity.UsersCore{}, err
	}

	dataResponse := entity.UsersModelToUsersCore(dataUsers)
	return dataResponse, nil
}

// UpdateById implements entity.UsersRepositoryInterface.
func (ur *userRepository) UpdateById(id string, data entity.UsersCore) error {

	request := entity.UsersCoreToUsersModel(data)
	err := ur.db.Where("id = ?", id).Updates(&request).Error
	if err != nil {
		return err
	}

	return nil
}

// ForgetPassword implements entity.UsersRepositoryInterface.
func (ur *userRepository) UpdatePassword(id string, data entity.UsersCore) error {

	request := entity.UsersCoreToUsersModel(data)

	err := ur.db.Where("id = ?", id).Updates(&request).Error
	if err != nil {
		return err
	}

	return nil
}

// GetByVerificationToken implements entity.UsersRepositoryInterface.
func (ur *userRepository) GetByVerificationToken(token string) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	err := ur.db.Where("verification_token = ?", token).First(&dataUsers).Error
	if err != nil {
		return entity.UsersCore{}, err
	}

	userToken := entity.UsersModelToUsersCore(dataUsers)
	return userToken, nil
}

// UpdateIsVerified implements entity.UsersRepositoryInterface.
func (ur *userRepository) UpdateIsVerified(id string, isVerified bool) error {
	dataUser := model.Users{}

	errFind := ur.db.First(&dataUser, id).Error
	if errFind != nil {
		return errFind
	}

	dataUser.IsVerified = isVerified

	errSave := ur.db.Save(&dataUser).Error
	if errSave != nil {
		return errSave
	}

	return nil
}

// SendOTP implements entity.UsersRepositoryInterface.
func (ur *userRepository) SendOTP(emailUser string, otp string, expiry int64) (data entity.UsersCore, err error) {
	dataUsers := model.Users{}

	errData := ur.db.Where("email = ?", emailUser).First(&dataUsers).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return entity.UsersCore{}, errors.New("pengguna tidak ditemukan")
		}
		return entity.UsersCore{}, errData
	}

	dataUsers.Otp = otp
	dataUsers.OtpExpiration = expiry

	errUpdate := ur.db.Save(&dataUsers).Error
	if errUpdate != nil {
		return entity.UsersCore{}, errUpdate
	}

	dataResponse := entity.UsersModelToUsersCore(dataUsers)

	return dataResponse, nil
}

// VerifyOTP implements entity.UsersRepositoryInterface.
func (ur *userRepository) VerifyOTP(otp string) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	err := ur.db.Where("otp = ?", otp).First(&dataUsers).Error
	if err != nil {
		return entity.UsersCore{}, err
	}

	dataResponse := entity.UsersModelToUsersCore(dataUsers)
	return dataResponse, nil
}

// ResetOTP implements entity.UsersRepositoryInterface.
func (ur *userRepository) ResetOTP(otp string) (data entity.UsersCore, err error) {
	dataUsers := model.Users{}

	errData := ur.db.Where("otp = ?", otp).First(&dataUsers).Error
	if errData != nil {
		return entity.UsersCore{}, errData
	}

	dataUsers.Otp = ""
	dataUsers.OtpExpiration = 0

	errUpdate := ur.db.Save(&dataUsers).Error
	if errUpdate != nil {
		return entity.UsersCore{}, errUpdate
	}

	dataResponse := entity.UsersModelToUsersCore(dataUsers)
	return dataResponse, nil
}

// ForgetPassword implements entity.UsersRepositoryInterface.
func (ur *userRepository) NewPassword(otp string, data entity.UsersCore) (entity.UsersCore, error) {
	dataUsers := model.Users{}

	errData := ur.db.Where("otp = ?", otp).First(&dataUsers).Error
	if errData != nil {
		return entity.UsersCore{}, errData
	}

	errUpdate := ur.db.Model(&dataUsers).Updates(entity.UsersCoreToUsersModel(data))
	if errUpdate != nil {
		return entity.UsersCore{}, errUpdate.Error
	}

	dataResponse := entity.UsersModelToUsersCore(dataUsers)

	return dataResponse, nil
}
