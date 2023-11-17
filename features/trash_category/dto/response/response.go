package response

import "time"

type TrashCategory struct {
	ID        string    `json:"id"`
	TrashType string    `json:"trash_type"`
	Point     int       `json:"point"`
	Satuan    string    `json:"satuan"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated"`
}
