package response

type UserLoginResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserResponseProfile struct {
	Email       string `json:"email"`
	Fullname    string `json:"fullname"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	DateOfBirth string `json:"date_of_birth"`
	Purpose     string `json:"purpose"`
	Point       int    `json:"point"`
}

type ResponseManageUsers struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
