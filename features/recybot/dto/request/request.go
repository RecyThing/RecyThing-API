package request

type RecybotRequest struct {
	Category string `json:"category" form:"category"`
	Question string `json:"question" form:"question"`
}

