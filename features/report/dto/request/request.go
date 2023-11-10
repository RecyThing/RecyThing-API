package request

type ReportRubbishRequest struct {
	ReportType     string         `json:"report_type" form:"report_type"`
	Longitude      float64        `json:"longitude" form:"longitude"`
	Latitude       float64        `json:"latitude" form:"latitude"`
	Location       string         `json:"location" form:"location"`
	TrashType      string         `json:"trash_type" form:"trash_type"`
	Description    string         `json:"description" form:"description"`
	ScaleType      string         `json:"scale_type" form:"scale_type"`
	InsidentTime   string         `json:"insident_time" form:"insident_time"`
	CompanyName    string         `json:"company_name" form:"company_name"`
	DangerousWaste bool           `json:"dangerous_waste" form:"dangerous_waste"`
	Images         []ImageRequest `json:"images" form:"images"`
}

type ImageRequest struct {
	Image string `json:"image"`
}
