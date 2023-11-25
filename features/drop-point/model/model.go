package model

import (
	"time"

	"gorm.io/gorm"
)

type DropPoint struct {
	Id                   string `gorm:"primary key"`
	Name                 string `gorm:"not null"`
	Address              string `gorm:"not null"`
	Latitude             string `gorm:"not null"`
	Longitude            string `gorm:"not null"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt         `gorm:"index"`
	OperationalSchedules []OperationalSchedules `gorm:"foreignKey:DropPointId;constraint:OnDelete:CASCADE;"`
}

type OperationalSchedules struct {
	Id          string `gorm:"primary key"`
	DropPointId string `gorm:"index"`
	Days        string
	Open        string
	Close       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
