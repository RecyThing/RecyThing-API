package request

type AdminRequest struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type AdminLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
