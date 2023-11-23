package request

type DropPointRequest struct {
	Name                 string                        `json:"name"`
	Address              string                        `json:"address"`
	Latitude             string                        `json:"latitude"`
	Longitude            string                        `json:"longitude"`
	OperationalSchedules []OperationalSchedulesRequest `json:"operational_schedules"`
}

type OperationalSchedulesRequest struct {
	Days  string `json:"days"`
	Open  string `json:"open"`
	Close string `json:"close"`
}
