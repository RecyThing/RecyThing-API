package entity

import (
	"time"

)

type AchievementCore struct {
	Id          string 
	Name        string 
	TargetPoint int 
	TotalUser   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
