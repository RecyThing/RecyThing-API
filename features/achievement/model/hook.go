package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (a *Achievement) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	a.Id = newUuid.String()

	return nil
}