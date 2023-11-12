package repository

import (
	"errors"
	"recything/features/user/entity"
	"recything/features/user/model"

	"github.com/google/uuid"
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

// ForgetPassword implements entity.UsersRepositoryInterface.
func (ur *userRepository) UpdatePassword(id string, updated entity.UsersCore) (data entity.UsersCore, err error) {
	var usersData model.Users

	errData := ur.db.Where("id = ?", id).First(&usersData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return entity.UsersCore{}, errors.New("pengguna tidak ditemukan")
		}
		return entity.UsersCore{}, errData
	}

	errUpdate := ur.db.Model(&usersData).Updates(entity.UsersCoreToUsersModel(updated))
	if errUpdate != nil {
		return entity.UsersCore{}, errUpdate.Error
	}
	data = entity.UsersModelToUsersCore(usersData)

	return data, nil
}

// ForgetPassword implements entity.UsersRepositoryInterface.
func (ur *userRepository) ForgetPassword(otp string, updated entity.UsersCore) (data entity.UsersCore, err error) {
	var usersData model.Users

    errData := ur.db.Where("otp = ?", otp).First(&usersData).Error
    if errData != nil {
        if errors.Is(errData, gorm.ErrRecordNotFound) {
            return entity.UsersCore{}, errors.New("otp tidak ditemukan")
        }
        return entity.UsersCore{}, errData
    }

    errUpdate := ur.db.Model(&usersData).Updates(entity.UsersCoreToUsersModel(updated))
	if errUpdate != nil {
		return entity.UsersCore{}, errUpdate.Error
	}
	data = entity.UsersModelToUsersCore(usersData)

	return data, nil
}

// GetById implements entity.UsersRepositoryInterface.
func (ur *userRepository) GetById(id string) (entity.UsersCore, error) {
	var userData model.Users

	result := ur.db.Where("id = ?", id).First(&userData)
	if result.Error != nil {
		return entity.UsersCore{}, result.Error
	}

	var userById = entity.UsersModelToUsersCore(userData)
	return userById, nil
}

// GetByVerificationToken implements entity.UsersRepositoryInterface.
func (ur *userRepository) GetByVerificationToken(token string) (entity.UsersCore, error) {
	var userData model.Users
	result := ur.db.Where("verification_token = ?", token).First(&userData)
	if result.Error != nil {
		return entity.UsersCore{}, result.Error
	}

	var userToken = entity.UsersModelToUsersCore(userData)
	return userToken, nil
}

// Login implements entity.UsersRepositoryInterface.
func (ur *userRepository) Login(email string, password string) (entity.UsersCore, error) {
	var data model.Users

	tx := ur.db.Where("email = ?", email).First(&data)
	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	dataMain := entity.UsersModelToUsersCore(data)
	return dataMain, nil
}

// Register implements entity.UsersRepositoryInterface.
func (ur *userRepository) Register(data entity.UsersCore) error {
	dataInput := entity.UsersCoreToUsersModel(data)

	tx := ur.db.Create(&dataInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// UpdateById implements entity.UsersRepositoryInterface.
func (ur *userRepository) UpdateById(id string, updated entity.UsersCore) (data entity.UsersCore, err error) {
	var usersData model.Users

	errData := ur.db.Where("id = ?", id).First(&usersData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return entity.UsersCore{}, errors.New("pengguna tidak ditemukan")
		}
		return entity.UsersCore{}, errData
	}

	errUpdate := ur.db.Model(&usersData).Updates(entity.UsersCoreToUsersModel(updated))
	if errUpdate != nil {
		return entity.UsersCore{}, errUpdate.Error
	}
	data = entity.UsersModelToUsersCore(usersData)

	return data, nil
}

// UpdateIsVerified implements entity.UsersRepositoryInterface.
func (ur *userRepository) UpdateIsVerified(id string, isVerified bool) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	var user model.Users
	result := ur.db.First(&user, uuidID)
	if result.Error != nil {
		return result.Error
	}

	user.IsVerified = isVerified
	result = ur.db.Save(&user)

	return result.Error
}

// EmailExists implements entity.UsersRepositoryInterface.
func (ur *userRepository) EmailExists(email string) (bool, error) {
	var user model.Users
	result := ur.db.Select("id").Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// SendOTP implements entity.UsersRepositoryInterface.
func (ur *userRepository) SendOTP(emailUser string, otp string, expiry int64) (data entity.UsersCore, err error) {
	var usersData model.Users

	errData := ur.db.Where("email = ?", emailUser).First(&usersData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return entity.UsersCore{}, errors.New("pengguna tidak ditemukan")
		}
		return entity.UsersCore{}, errData
	}

	usersData.Otp = otp
	usersData.OtpExpiration = expiry

	errUpdate := ur.db.Save(&usersData).Error
	if errUpdate != nil {
		return entity.UsersCore{}, errUpdate
	}
	data = entity.UsersModelToUsersCore(usersData)

	return data, nil
}

// VerifyOTP implements entity.UsersRepositoryInterface.
func (ur *userRepository) VerifyOTP(otp string) (entity.UsersCore, error) {
	var user model.Users
	tx := ur.db.Where("otp = ?", otp).First(&user)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return entity.UsersCore{}, errors.New("otp tidak ditemukan")
		}
		return entity.UsersCore{}, tx.Error
	}

	dataMain := entity.UsersModelToUsersCore(user)
	return dataMain, nil
}

// ResetOTP implements entity.UsersRepositoryInterface.
func (ur *userRepository) ResetOTP(otp string) (data entity.UsersCore, err error) {
	var usersData model.Users
	errData := ur.db.Where("otp = ?", otp).First(&usersData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return entity.UsersCore{}, errData
		}
		return entity.UsersCore{}, errData
	}

	usersData.Otp = ""
	usersData.OtpExpiration = 0

	errUpdate := ur.db.Save(&usersData).Error
	if errUpdate != nil {
		return entity.UsersCore{}, errUpdate
	}

	data = entity.UsersModelToUsersCore(usersData)
	return data, nil
}
