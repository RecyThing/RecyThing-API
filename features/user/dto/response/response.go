package response

type UserLoginResponse struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

type UserResponseProfile struct {
	Id          string `json:"id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Purpose     string `json:"purpose"`
	Point       int    `json:"point"`
}

type ResponseManageUsers struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
