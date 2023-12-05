package model

import "time"

type DailyPoint struct {
	Id          int `sql:"AUTO_INCREMENT" gorm:"primary key"`
	Point       int
	Description string
}

type DailyCache struct {
	Id           string
	UserID       string
	DailyPointID int
	CreatedAt    time.Time
}
