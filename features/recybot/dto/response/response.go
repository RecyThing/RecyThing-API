package response

import "time"

type RecybotResponse struct {
	ID        string
	Category  string
	Question  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
