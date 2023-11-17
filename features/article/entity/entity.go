package entity

import "time"

type ArticleCore struct {
	ID        string
	Title     string
	Image     string
	Content   string
	Category  string
	PostDate  string
	Like      int
	Share     int
	CreatedAt time.Time
	UpdatedAt time.Time
}
