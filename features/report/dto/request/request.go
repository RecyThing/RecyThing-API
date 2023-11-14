package request

type ReportRubbishRequest struct {
	ReportType           string  `form:"report_type"`
	Longitude            float64 `form:"longitude"`
	Latitude             float64 `form:"latitude"`
	Location             string  `form:"location"`
	AddressPoint         string  `form:"address_point"`
	Status               string  `form:"status"`
	TrashType            string  `form:"trash_type"`
	ScaleType            string  `form:"scale_type"`
	InsidentDate         string  `form:"insident_date"`
	InsidentTime         string  `form:"insident_time"`
	CompanyName          string  `form:"company_name"`
	DangerousWaste       bool    `form:"dangerous_waste"`
	RejectionDescription string  `form:"rejection_description"`
	Description          string  `form:"description"`
	Images               []ImageRequest
}

type ImageRequest struct {
	Image string `json:"image"`
}

type UpdateStatusReportRubbish struct {
	Status               string `json:"status"`
	RejectionDescription string `json:"rejection_description"`
}
