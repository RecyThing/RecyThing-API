package request

type TrashCategory struct {
	TrashType string `json:"trash_type"`
	Point     int    `json:"point"`
	Unit      string `json:"unit"`
}
