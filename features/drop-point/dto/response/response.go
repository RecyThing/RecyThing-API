package response

type DropPointResponse struct {
	Name                 string                         `json:"name"`
	Address              string                         `json:"address"`
	Latitude             string                         `json:"latitude"`
	Longitude            string                         `json:"longitude"`
	OperationalSchedules []OperationalSchedulesResponse `json:"operational_schedules"`
}

type OperationalSchedulesResponse struct {
	Days  string `json:"days"`
	Open  string `json:"open"`
	Close string `json:"close"`
}
