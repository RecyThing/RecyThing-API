package entity

import "time"

type AdminCore struct {
	Id              string
	Name            string
	Role            string
	Email           string
	Password        string
	ConfirmPassword string
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
