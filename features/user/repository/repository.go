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

// GetById implements entity.UsersRepositoryInterface.
func (userRep *userRepository) GetById(id string) (entity.UsersCore, error) {
	var userData model.Users

	// Gunakan Preload untuk memuat data pickup terkait.
	result := userRep.db.Where("id = ?", id).First(&userData)
	if result.Error != nil {
		return entity.UsersCore{}, result.Error
	}

	var userById = entity.UsersModelToUsersCore(userData)
	return userById, nil
}

// GetByVerificationToken implements entity.UsersRepositoryInterface.
func (userRep *userRepository) GetByVerificationToken(token string) (entity.UsersCore, error) {
	var userData model.Users
	result := userRep.db.Where("verification_token = ?", token).First(&userData)
	if result.Error != nil {
		return entity.UsersCore{}, result.Error
	}

	var userToken = entity.UsersModelToUsersCore(userData)
	return userToken, nil
}

// Login implements entity.UsersRepositoryInterface.
func (userRep *userRepository) Login(email string, password string) (entity.UsersCore, error) {
	var data model.Users

	tx := userRep.db.Where("email = ? AND password = ?", email, password).First(&data)
	if tx.Error != nil {
		return entity.UsersCore{}, tx.Error
	}

	dataMain := entity.UsersModelToUsersCore(data)

	return dataMain, nil

}

// Register implements entity.UsersRepositoryInterface.
func (userRep *userRepository) Register(data entity.UsersCore) error {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	dataInput := entity.UsersCoreToUsersModel(data)
	dataInput.Id = newUUID.String()
	// uniqueToken := email.GenerateUniqueToken()
	// dataInput.VerificationToken = uniqueToken

	tx := userRep.db.Create(&dataInput)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// UpdateById implements entity.UsersRepositoryInterface.
func (userRep *userRepository) UpdateById(id string, updated entity.UsersCore) (data entity.UsersCore, err error) {
	var usersData model.Users

	errData := userRep.db.Where("id = ?", id).First(&usersData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return entity.UsersCore{}, errors.New("user not found")
		}
		return entity.UsersCore{}, errData
	}
	userRep.db.Model(&usersData).Updates(entity.UsersCoreToUsersModel(updated))
	data = entity.UsersModelToUsersCore(usersData)

	return data, nil
}

// UpdateIsVerified implements entity.UsersRepositoryInterface.
func (userRep *userRepository) UpdateIsVerified(id string, isVerified bool) error {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	var user model.Users
	result := userRep.db.First(&user, uuidID)
	if result.Error != nil {
		return result.Error
	}

	user.IsVerified = isVerified
	result = userRep.db.Save(&user)

	return result.Error
}
