package response

import (
	"time"
)

type ArticleCreateResponse struct {
	Id          string    `json:"Id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Image       string    `json:"image,omitempty"`
	Content     string    `json:"content,omitempty"`
	Category_id []string  `json:"category,omitempty"`
	Like        int       `json:"like,omitempty"`
	Share       int       `json:"share,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// type TrashCategoryResponse struct {
// 	ArticleID       string
// 	TrashCategoryID string
// }
