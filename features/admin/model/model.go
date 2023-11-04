package model

import (
	"recything/utils/helper"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	Id        string         `gorm:"primary key"`
	Name      string         `json:"name" form:"name"`
	Role      string         `gorm:"type:enum('admin', 'super_admin');default:'admin'"`
	Email     string         `json:"email" form:"email"`
	Password  string         `json:"password" form:"password"`
	Status    string         `gorm:"type:enum('active', 'nonactive');default:'active'"`
	CreatedAt time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	a.Id = newUuid.String()

	a.Password, _ = helper.HashPassword(a.Password)
	return nil
}

func (a *Admin) BeforeUpdate(tx *gorm.DB) (err error) {
	a.Password, _ = helper.HashPassword(a.Password)
	return nil
}

// nama saya
