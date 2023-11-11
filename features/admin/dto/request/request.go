package request

type AdminRequest struct {
	Name            string `json:"name" valid:"required~harap lengkapi nama"`
	Email           string `json:"email" valid:"required~harap lengkapi email,email~format email tidak valid"`
	Password        string `json:"password" valid:"required~harap lengkapi password,minstringlength(8)~password minimal 8 karakter"`
	ConfirmPassword string `json:"confirm_password" valid:"required~harap konfirmasi password"`
}

type RequestLogin struct {
	Email    string `json:"email" valid:"required~harap lengkapi email,email~format email tidak valid"`
	Password string `json:"password"  valid:"required~harap lengkapi password"`
}
