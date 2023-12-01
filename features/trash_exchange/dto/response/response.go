package response

import "time"

type TrashExchangeResponse struct {
	Name                 string                       `json:"name"`
	EmailUser            string                       `json:"email"`
	Address              string                       `json:"address"`
	TotalUnit            float64                      `json:"total_unit"`
	TotalPoint           int                          `json:"total_point"`
	CreatedAt            time.Time                    `json:"created_at"`
	TrashExchangeDetails []TrashExchangeDetailRespose `json:"trash_exchange_details"`
}

type TrashExchangeDetailRespose struct {
	TrashType   string `json:"trash_type"`
	Unit        string `json:"unit"`
	TotalPoints int    `json:"total_points"`
}
