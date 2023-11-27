package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mission struct {
	ID           string
	Name         string `gorm:"not null;unique"`
	Creator      string `gorm:"not null"`
	Status       string `gorm:"type:enum('aktif', 'melewati tenggat');not null"`
	AdminID      string
	MissionImage string
	Point        int
	Description  string
	StartDate    string
	EndDate      string
	MissionStages []MissionStages
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type MissionStages struct {
	ID          string
	Title       string
	Description string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (m *Mission) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	m.ID = newUuid.String()
	return nil
}

func (ms *MissionStages) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	ms.ID = newUuid.String()
	return nil
}
