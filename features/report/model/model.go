package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	//all
	Id          string `gorm:"primary key"`
	ReportType  string `gorm:"type:enum('rubbish','littering')"`
	UsersId     string `gorm:"type:varchar(191);index"`
	Longitude   float64
	Latitude    float64
	Location    string
	Description string
	Images      []Image `gorm:"foreignKey:ReportId"`

	//rubbish only
	TrashType string `gorm:"type:enum('dry waste','Wet waste');default:Null"`

	//littering only
	ScaleType    string `gorm:"type:enum('big','small');default:Null"`
	InsidentTime string

	//all
	CreatedAt time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Image struct {
	ID        string `gorm:"primaryKey"`
	ReportId  string `gorm:"index;foreignKey:Id"`
	Image     string
	CreatedAt time.Time      `gorm:"type:DATETIME(0)"`
	UpdatedAt time.Time      `gorm:"type:DATETIME(0)"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (r *Report) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	r.Id = newUuid.String()
	return nil
}
func (i *Image) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	i.ID = newUuid.String()
	return nil
}
