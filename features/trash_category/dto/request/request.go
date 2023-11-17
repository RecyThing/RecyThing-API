package request

type TrashCategory struct {
	TrashType string `json:"trash_type"`
	Point     int    `json:"point"`
	Satuan    string `json:"satuan"`
}
