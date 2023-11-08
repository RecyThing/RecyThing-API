package service

import (
	"errors"
	"recything/features/user/entity"
	"recything/utils/email"
	"recything/utils/helper"
	"recything/utils/jwt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepo entity.UsersRepositoryInterface
	validate *validator.Validate
}

func NewUserService(userRepo entity.UsersRepositoryInterface) entity.UsersUsecaseInterface {
	return &userService{
		userRepo: userRepo,
		validate: validator.New(),
	}
}

// ForgetPassword implements entity.UsersUsecaseInterface.
func (us *userService) ForgetPassword(id string, updated entity.UsersCore) (data entity.UsersCore, err error) {
	if id == "" {
		return entity.UsersCore{}, errors.New("invalid id")
	}

	if updated.Password != updated.ConfirmPassword {
		return entity.UsersCore{}, errors.New("confirm password does not match")
	}

	if len(updated.Password) < 8 {
		return entity.UsersCore{},errors.New("your password is too short, must be at least 8 characters")
	}
	
	hashedPassword, errHash := helper.HashPassword(updated.Password)
	if errHash != nil {
		return entity.UsersCore{},errors.New("error hash password")
	}
	updated.Password = hashedPassword

	updatePassword, err := us.userRepo.ForgetPassword(id, updated)
	if err != nil {
		return entity.UsersCore{}, err
	}

	return updatePassword, nil
}

// GetById implements entity.UsersUsecaseInterface.
func (us *userService) GetById(id string) (entity.UsersCore, error) {
	if id == "" {
		return entity.UsersCore{}, errors.New("invalid id")
	}

	idUser, err := us.userRepo.GetById(id)
	return idUser, err
}

// GetByVerificationToken implements entity.UsersUsecaseInterface.
func (us *userService) VerifyUser(token string) (bool, error) {
	if token == "" {
		return false, errors.New("invalid token")
	}

	user, err := us.userRepo.GetByVerificationToken(token)
	if err != nil {
		return false, errors.New("failed to get user")
	}

	if user.IsVerified {
		return true, nil
	}

	err = us.userRepo.UpdateIsVerified(user.Id, true)
	if err != nil {
		return false, errors.New("failed to activate the user")
	}

	return false, nil
}

// Login implements entity.UsersUsecaseInterface.
func (us *userService) Login(email string, password string) (entity.UsersCore, string, error) {
	if email == "" || password == "" {
		return entity.UsersCore{}, "", errors.New("email dan password harus diisi")
	}

	dataLogin, err := us.userRepo.Login(email, password)
	if err != nil {
		return entity.UsersCore{}, "", err
	}

	if helper.CompareHash(dataLogin.Password, password) {
		token, err := jwt.CreateToken(dataLogin.Id, "")
		if err != nil {
			return entity.UsersCore{}, "", err
		}
		return dataLogin, token, nil
	}
	return entity.UsersCore{}, "", errors.New("login failed")
}

// Register implements entity.UsersUsecaseInterface.
func (us *userService) Register(data entity.UsersCore) error {
	errValidate := us.validate.Struct(data)
	if errValidate != nil {
		return errValidate
	}

	emailExists, errEmail := us.userRepo.EmailExists(data.Email)
	if errEmail != nil {
		return errors.New("failed to check if email exists")
	}
	
	if emailExists {
		return errors.New("email already in use")
	}

	if data.Password != data.ConfirmPassword {
		return errors.New("confirm password does not match")
	}

	if len(data.Password) < 8 {
		return errors.New("your password is too short, must be at least 8 characters")
	}

	hashedPassword, errHash := helper.HashPassword(data.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}
	data.Password = hashedPassword

	uniqueToken := email.GenerateUniqueToken()
	data.VerificationToken = uniqueToken

	err := us.userRepo.Register(data)
	if err != nil {
		return err
	}

	email.SendVerificationEmail(data.Email, uniqueToken)
	return nil
}

// UpdateById implements entity.UsersUsecaseInterface.
func (us *userService) UpdateById(id string, updated entity.UsersCore) (data entity.UsersCore, err error) {
	if id == "" {
		return entity.UsersCore{}, errors.New("invalid id")
	}

	if updated.DateOfBirth != "" {
		if _, parseErr := time.Parse("2006-01-02", updated.DateOfBirth); parseErr != nil {
			return entity.UsersCore{}, errors.New("error, date must be in the format 'yyyy-mm-dd'")
		}
	}

	if updated.Phone != "" {
		phoneRegex := `^(?:\+62|0)[0-9-]+$`
		match, _ := regexp.MatchString(phoneRegex, updated.Phone)
		if !match {
			return entity.UsersCore{}, errors.New("error, phone number format not valid")
		}
	}

	updateData, err := us.userRepo.UpdateById(id, updated)
	if err != nil {
		return entity.UsersCore{}, err
	}

	return updateData, nil
}

// UpdateIsVerified implements entity.UsersUsecaseInterface.
func (us *userService) UpdateIsVerified(id string, isVerified bool) error {
	if id == "" {
		return errors.New("user id is required")
	}

	return us.userRepo.UpdateIsVerified(id, isVerified)
}
