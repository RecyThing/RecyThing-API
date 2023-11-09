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

// UpdatePassword implements entity.UsersUsecaseInterface.
func (us *userService) UpdatePassword(id string, updated entity.UsersCore) (data entity.UsersCore, err error) {
	if id == "" {
		return entity.UsersCore{}, errors.New("invalid id")
	}

	existingUser, err := us.userRepo.GetById(id)
    if err != nil {
        return entity.UsersCore{}, err
    }

	if helper.CompareHash(updated.Password, existingUser.Password) {
        return entity.UsersCore{}, errors.New("password lama tidak benar")
    }

	if updated.NewPassword != updated.ConfirmPassword {
		return entity.UsersCore{}, errors.New("password tidak sama")
	}

	if len(updated.Password) < 8 {
		return entity.UsersCore{}, errors.New("password anda terlalu pendek, minimal 8 karakter untuk password")
	}

	hashedPassword, errHash := helper.HashPassword(updated.NewPassword)
	if errHash != nil {
		return entity.UsersCore{}, errors.New("error hash password")
	}
	updated.Password = hashedPassword

	updatePassword, err := us.userRepo.UpdatePassword(id, updated)
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
		return false, errors.New("gagal mendapatkan data")
	}

	if user.IsVerified {
		return true, nil
	}

	err = us.userRepo.UpdateIsVerified(user.Id, true)
	if err != nil {
		return false, errors.New("gagal untuk mengaktifkan user")
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
	return entity.UsersCore{}, "", errors.New("login gagal")
}

// Register implements entity.UsersUsecaseInterface.
func (us *userService) Register(data entity.UsersCore) error {
	errValidate := us.validate.Struct(data)
	if errValidate != nil {
		return errValidate
	}

	emailExists, errEmail := us.userRepo.EmailExists(data.Email)
	if errEmail != nil {
		return errors.New("gagal mengecek bahwa email telah ada")
	}

	if emailExists {
		return errors.New("email telah digunakan")
	}

	if data.Password != data.ConfirmPassword {
		return errors.New("password tidak sama")
	}

	if len(data.Password) < 8 {
		return errors.New("password anda terlalu pendek, minimal 8 karakter untuk password")
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
			return entity.UsersCore{}, errors.New("error, tanggal harus dalam format 'yyyy-mm-dd'")
		}
	}

	if updated.Phone != "" {
		phoneRegex := `^(?:\+62|0)[0-9-]+$`
		match, _ := regexp.MatchString(phoneRegex, updated.Phone)
		if !match {
			return entity.UsersCore{}, errors.New("error, format nomor telepon tidak valid")
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

// SendOTP implements entity.UsersUsecaseInterface.
func (us *userService) SendOTP(emailUser string) error {
	otp, err := email.GenerateOTP(4)
	if err != nil {
		return errors.New("generate otp gagal")
	}

	expiration := time.Now().Add(15 * time.Minute)
	_, err = us.userRepo.SendOTP(emailUser, otp, expiration)
	if err != nil {
		return err
	}

	email.SendOTPEmail(emailUser, otp)
	return nil
}

// VerifyOTP implements entity.UsersUsecaseInterface.
func (us *userService) VerifyOTP(emailUser string, otp string) (string, error) {
	user, err := us.userRepo.VerifyOTP(emailUser, otp)
	if err != nil {
		return "", err
	}

	if user.OtpExpiration.Before(time.Now()) {
		return "", errors.New("otp sudah kadaluwarsa")
	}

	if user.Otp != otp {
		return "", errors.New("otp tidak valid")
	}

	token, err := jwt.CreateTokenVerifikasi(emailUser)
	if err != nil {
		return "", errors.New("token gagal dibuat")
	}

	_, err = us.userRepo.ResetOTP(emailUser)
	if err != nil {
		return "", errors.New("gagal mengatur ulang OTP")
	}

	return token, nil
}

// ForgetPassword implements entity.UsersUsecaseInterface.
func (us *userService) ForgetPassword(emailUser string, updated entity.UsersCore) error {
	if updated.Password != updated.ConfirmPassword {
		return errors.New("password tidak sama")
	}


	if len(updated.Password) < 8 {
		return errors.New("password anda terlalu pendek, minimal 8 karakter untuk password")
	}

	hashedPassword, errHash := helper.HashPassword(updated.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}
	updated.Password = hashedPassword

	_, err := us.userRepo.ForgetPassword(emailUser, updated)
	if err != nil {
		return err
	}

	return  nil
}