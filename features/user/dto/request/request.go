package request

type UserRegister struct {
	Fullname        string `json:"fullname" valid:"required~harap lengkapi nama"`
	Email           string `json:"email" valid:"required~harap lengkapi email,email~format email tidak valid"`
	Password        string `json:"password" valid:"required~harap lengkapi password,minstringlength(8)~password minimal 8 karakter"`
	ConfirmPassword string `json:"confirm_password" valid:"required~harap konfirmasi password"`
}

type UserLogin struct {
	Email    string `json:"email" valid:"required~harap lengkapi email,email~format email tidak valid"`
	Password string `json:"password" valid:"required~harap lengkapi password"`
}

type UserUpdate struct {
	Fullname    string `json:"fullname"`
	Phone       string `json:"phone" valid:"numeric~format nomor telepon tidak valid,minstringlength(11)~nomor minimal 11 digit,maxstringlength(15)~nomor maximal 15 digit"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	Purpose     string `json:"purpose"`
}

type UserNewPassword struct {
	Email           string `json:"email" valid:"required~harap lengkapi email,email~format email tidak valid"`
	Password        string `json:"password" valid:"required~harap lengkapi password,minstringlength(8)~password minimal 8 karakter"`
	ConfirmPassword string `json:"confirm_password" valid:"required~harap konfirmasi password"`
}

type UserUpdatePassword struct {
	Password        string `json:"password" valid:"required~harap lengkapi password"`
	NewPassword     string `json:"new_password" valid:"required~harap lengkapi password baru,minstringlength(8)~password minimal 8 karakter"`
	ConfirmPassword string `json:"confirm_password" valid:"required~harap konfirmasi password"`
}
type UserSendOTP struct {
	Email string `json:"email" valid:"required~harap lengkapi email,email~format email tidak valid"`
}

type UserVerifyOTP struct {
	Email string `json:"email" valid:"required~harap lengkapi email,email~format email tidak valid"`
	Otp   string `json:"otp" valid:"required~kode otp wajib di isi!"`
}
