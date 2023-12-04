package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)


func (r *Voucher) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	r.Id = newUuid.String()
	return nil
}

func (r *ExchangeVoucher) BeforeCreate(tx *gorm.DB) (err error) {
	newUuid := uuid.New()
	r.Id = newUuid.String()
	return nil
}