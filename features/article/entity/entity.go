package entity

import (
	"time"
)

type ArticleCore struct {
	ID          string
	Title       string
	Image       string
	Content     string
	Category_id []string
	Like        int
	Share       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// type CategoryCore struct{
// 	ArticleID string
// 	TrashCategoryID string
// }
