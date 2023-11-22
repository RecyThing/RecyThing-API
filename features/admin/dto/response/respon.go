package response

import "time"

type AdminRespon struct {
	ID        string    `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"type:DATETIME(0)"`
}

type AdminResponseLogin struct {
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
