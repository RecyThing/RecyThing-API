package request

type ArticleRequest struct {
	Title    string `form:"title"`
	Image    string `form:"image"`
	Content  string `form:"content"`
	Category string `form:"category"`
}
