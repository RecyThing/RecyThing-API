package dto

type ReportRubbishRequest struct {
	ReportType  string  `json:"report_type"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Location    string  `json:"location"`
	TrashType   string  `json:"trash_type"`
	Description string  `json:"description"`
	Images      []ImageRequest `json:"images" form:"images"`
}

type ImageRequest struct {
	Image string `json:"image"`
}
