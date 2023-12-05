package model

import (
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
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (m *Mission) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	m.ID = newUuid.String()
	return nil
}

func (ms *MissionStage) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	ms.ID = newUuid.String()
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

type UploadMissionTask struct {
	ID          string `gorm:"type:varchar(255);primaryKey" `
	UserID      string `gorm:"type:varchar(255);index" `
	MissionID   string `gorm:"type:varchar(255)" `
	Description string
	Images      []ImageUploadMission
	Status      string    `gorm:"type:enum('disetujui','ditolak','perlu tinjauan');default:'perlu tinjauan'"`
	CreatedAt   time.Time `gorm:"type:DATETIME(0)" `
	UpdatedAt   time.Time `gorm:"type:DATETIME(0)" `
}

type ImageUploadMission struct {
	ID                  string `gorm:"primaryKey" `
	UploadMissionTaskID string `gorm:"type:varchar(255);index" `
	Image               string
	CreatedAt           time.Time `gorm:"type:DATETIME(0)" `
}

func (cm *ImageUploadMission) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	cm.ID = newUuid.String()
	return nil
}

func (cm *UploadMissionTask) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	cm.ID = newUuid.String()
	return nil
}
