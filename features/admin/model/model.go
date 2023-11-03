package model

import "time"

type Admin struct {
	Id        string    `gorm:"primary key"`
	Name      string    `json:"name" form:"name"`
	Role      string    `gorm:"type:enum('admin', 'super_admin');default:'admin'"`
	Email     string    `json:"email" form:"email"`
	Password  string    `json:"password" form:"password"`
	Status    string    `gorm:"type:enum('aktif, 'tidak aktif');default:'aktif'"`
	CreatedAt time.Time `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time `gorm:"type:DATETIME(0)"`
	Delete    time.Time `gorm:"type:DATETIME(0)"`
}
