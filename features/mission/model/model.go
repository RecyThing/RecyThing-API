package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mission struct {
	ID              string `gorm:"type:varchar(255)"`
	Title           string `gorm:"not null;unique"`
	Status          string `gorm:"type:enum('Aktif', 'Melewati Tenggat');default:'aktif'"`
	AdminID         string
	MissionImage    string
	Point           int
	Description     string
	StartDate       string
	EndDate         string
	ClaimedMissions []ClaimedMission `gorm:"foreignKey:MissionID"`
	MissionStages   []MissionStage   `gorm:"foreignKey:MissionID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type MissionStage struct {
	ID          string
	Title       string
	Description string
	MissionID   string `gorm:"type:varchar(255)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (m *Mission) BeforeCreate(tx *gorm.DB) (err error) {
	trimmedUuid := strings.ReplaceAll(uuid.New().String(), "-", "")[:15]
	uppercasedUUID := strings.ToUpper(trimmedUuid)
	m.ID = "MIS-" + uppercasedUUID
	return nil
}

func (ms *MissionStage) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	ms.ID = newUuid.String()
	return nil
}

func (m *Mission) BeforeSave(tx *gorm.DB) (err error) {
	var mission Mission
	if tx.Model(&Mission{}).First(&mission, "id = ?", m.ID).Error != nil {
		return nil 
	}
	m.MissionStages = mission.MissionStages

	return nil
}

type ClaimedMission struct {
	ID        string `gorm:"type:varchar(255);primaryKey"`
	UserID    string `gorm:"type:varchar(255);index"`
	MissionID string `gorm:"type:varchar(255);index"`
	Claimed   bool   `gorm:"default:true"`
	CreatedAt time.Time
}

func (cm *ClaimedMission) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	cm.ID = newUuid.String()
	return nil
}

