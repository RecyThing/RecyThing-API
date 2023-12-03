package request

type TrashExchangeRequest struct {
	Name                 string                       `json:"name"`
	EmailUser            string                       `json:"email"`
	Address              string                       `json:"address"`
	TrashExchangeDetails []TrashExchangeDetailRequest `json:"trash_exchange_details"`
}

type TrashExchangeDetailRequest struct {
	TrashType string `json:"trash_type"`
	Unit      string `json:"unit"`
}
