package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (dropPoint *DropPoint) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	dropPoint.Id = newUuid.String()

	return nil
}

func (operational *OperationalSchedules) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	operational.Id = newUuid.String()

	return nil
}
