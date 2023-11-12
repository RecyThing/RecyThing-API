package service

import (
	"errors"
	"recything/features/user/entity"
	"recything/utils/email"
	"recything/utils/helper"
	"recything/utils/jwt"
	"time"
)

type userService struct {
	userRepo entity.UsersRepositoryInterface
}

func NewUserService(userRepo entity.UsersRepositoryInterface) entity.UsersUsecaseInterface {
	return &userService{
		userRepo: userRepo,
	}
}

// Register implements entity.UsersUsecaseInterface.
func (us *userService) Register(data entity.UsersCore) error {

	_,err := us.userRepo.FindByEmail(data.Email)
	if err == nil {
		return errors.New("email sudah digunakan, silahkan gunakan yang lain")
	}

	if data.Password != data.ConfirmPassword {
		return errors.New("password tidak sama")
	}

	hashedPassword, err := helper.HashPassword(data.Password)
	if err != nil {
		return errors.New("error hash password")
	}

	data.Password = hashedPassword
	uniqueToken := email.GenerateUniqueToken()
	data.VerificationToken = uniqueToken

	err = us.userRepo.Register(data)
	if err != nil {
		return err
	}

	email.SendVerificationEmail(data.Email, uniqueToken)

	return nil
}


// Login implements entity.UsersUsecaseInterface.
func (us *userService) Login(email, password string) (entity.UsersCore, string, error) {
	
	dataUser,errEmail := us.userRepo.FindByEmail(email)
	if errEmail != nil {
		return entity.UsersCore{},"",errors.New("email belum terdaftar")
	}

	if !dataUser.IsVerified {
		return entity.UsersCore{}, "", errors.New("akun belum terverifikasi")
	}



	comparePass := helper.CompareHash(dataUser.Password, password)
	if !comparePass {
		return entity.UsersCore{}, "", errors.New("password salah")
	}

	token, err := jwt.CreateToken(dataUser.Id, "")
	if err != nil {
		return entity.UsersCore{}, "", errors.New("gagal mendapatkan generate token")
	}
	return dataUser, token, nil
}

// GetById implements entity.UsersUsecaseInterface.
func (us *userService) GetById(id string) (entity.UsersCore, error) {
	if id == "" {
		return entity.UsersCore{}, errors.New("invalid id")
	}

	dataUser, err := us.userRepo.GetById(id)
	if err != nil {
		return entity.UsersCore{},errors.New("data user tidak ada")
	}
	return dataUser, nil
}


// UpdateById implements entity.UsersUsecaseInterface.
func (us *userService) UpdateById(id string, data entity.UsersCore) error{
	if id == "" {
		return errors.New("invalid id")
	}

	_, errGet := us.userRepo.GetById(id)
	if errGet != nil {
		return errors.New("data user tidak ada")
	}

	if data.DateOfBirth != "" {
		if _, errParse := time.Parse("2006-01-02", data.DateOfBirth); errParse != nil {
			return errors.New("error, tanggal harus dalam format 'yyyy-mm-dd'")
		}
	}

	if data.Phone != "" {
		phone := helper.PhoneNumberValid(data.Phone)
		if !phone {
			return errors.New("nomor telepon tidak valid")
		}
	}

	err := us.userRepo.UpdateById(id, data)
	if err != nil {
		return  err
	}

	return nil
}

// UpdatePassword implements entity.UsersUsecaseInterface.
func (us *userService) UpdatePassword(id string, data entity.UsersCore)  error {
	if id == "" {
		return errors.New("invalid id")
	}

	result ,err := us.GetById(id)
	if err != nil {
		return errors.New("data tidak ada")
	}

	ComparePass := helper.CompareHash(result.Password,data.Password) 
	if !ComparePass {
		return errors.New("password lama tidak sesuai") 
	}

	if data.NewPassword != data.ConfirmPassword {
		return errors.New("password tidak sama")
	}

	HashPassword, errHash := helper.HashPassword(data.NewPassword)
	if errHash != nil {
		return errors.New("error hash password")
	}
	data.Password = HashPassword

	err = us.userRepo.UpdatePassword(id, data)
	if err != nil {
		return err
	}

	return nil
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


// UpdateIsVerified implements entity.UsersUsecaseInterface.
func (us *userService) UpdateIsVerified(id string, isVerified bool) error {
	if id == "" {
		return errors.New("user id is required")
	}

	return us.userRepo.UpdateIsVerified(id, isVerified)
}

// SendOTP implements entity.UsersUsecaseInterface.
func (us *userService) SendOTP(emailUser string) error {

	otp, errGenerate := email.GenerateOTP(4)
	if errGenerate != nil {
		return errors.New("generate otp gagal")
	}

	expired := time.Now().Add(5 * time.Minute).Unix()

	_, errSend := us.userRepo.SendOTP(emailUser, otp, expired)
	if errSend != nil {
		return errSend
	}

	if email.ContainsLowerCase(otp) {
		return errors.New("otp tidak boleh mengandung huruf kecil")
	}

	email.SendOTPEmail(emailUser, otp)
	return nil
}

// VerifyOTP implements entity.UsersUsecaseInterface.
func (us *userService) VerifyOTP(otp string) (string, error) {
	dataUsers, err := us.userRepo.VerifyOTP(otp)
	if err != nil {
		return "", errors.New("otp tidak ditemukan")
	}

	if dataUsers.OtpExpiration <= time.Now().Unix() {
		return "", errors.New("otp sudah kadaluwarsa")
	}

	if dataUsers.Otp != otp {
		return "", errors.New("otp tidak valid")
	}

	token, err := jwt.CreateTokenVerifikasi(otp)
	if err != nil {
		return "", errors.New("token gagal dibuat")
	}

	return token, nil
}

// ForgetPassword implements entity.UsersUsecaseInterface.
func (us *userService) NewPassword(otp string, data entity.UsersCore) error {
	if data.Password != data.ConfirmPassword {
		return errors.New("password tidak sama")
	}

	HashPassword, errHash := helper.HashPassword(data.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}
	data.Password = HashPassword

	_, errNew := us.userRepo.NewPassword(otp, data)
	if errNew != nil {
		return errors.New("otp tidak ditemukan")
	}

	_, errReset := us.userRepo.ResetOTP(otp)
	if errReset != nil {
		return errors.New("gagal mengatur ulang OTP")
	}

	return nil
}
