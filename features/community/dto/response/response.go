package response

import "time"

type CommunityResponse struct {
	Name        string    `json:"name"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
}

type CommunityResponseForDetails struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	MaxMembers  int       `json:"max_members"`
	Image       string    `json:"image"`
	CreatedAt   time.Time `json:"created_at"`
}