package request

type ArticleRequest struct {
	Title       string   `form:"title"`
	Image       string   `form:"image"`
	Content     string   `form:"content"`
	Category_id []string `form:"category_id"`
}

// type ArticleTrashCategoryRequest struct {
// 	ArticleID       string
// 	TrashCategoryID string
// }
