package response

type UserCreateResponse struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}

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
	Point       int    `json:"point"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Purpose     string `json:"purpose"`
}

type UserResponseManageUsers struct {
	Id       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Point    int    `json:"point"`
}

type UserResponseDetailManageUsers struct {
	Id          string `json:"id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	DateOfBirth string `json:"date_of_birth"`
	Point       int    `json:"point"`
	Purpose     string `json:"purpose"`
	Address     string `json:"address"`
}
