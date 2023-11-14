package response

import "time"

type AdminRespon struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at" gorm:"type:DATETIME(0)"` 
}
