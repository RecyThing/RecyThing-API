package entity

import (
	"time"
)

type CommunityCore struct {
	Id          string
	Name        string
	Description string
	Location    string
	Members     int
	MaxMembers  int
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
