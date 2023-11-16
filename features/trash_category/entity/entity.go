package entity

import "time"

type TrashCategoryCore struct {
	ID        string 
	TrashType string 
	Point     int    
	Satuan    string 
	CreatedAt time.Time
	UpdatedAt time.Time
}
type PagnationInfo struct {
	Limit       int
	CurrentPage int
	LastPage    int
}