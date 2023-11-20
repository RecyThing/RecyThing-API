package request

type ArticleRequest struct {
	Title   string `form:"title"`
	Image   string `form:"image"`
	Content string `form:"content"`
	// Category []request.TrashCategory
	Category string `form:"category"`
}
