package request

type CommunityRequest struct {
	Name        string `form:"name"`
	Description string `form:"description"`
	Location    string `form:"location"`
	Max_Members int    `form:"max_members"`
	Image       string `form:"image"`
}
