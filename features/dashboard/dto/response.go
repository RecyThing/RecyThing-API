package dto

type GetCountUserResponse struct {
	TotalPenggunaAktif string `json:"total_pengguna_aktif"`
	Persentase         string `json:"persentase"`
	Status             string `json:"status"`
}

type GetCountExchangeVoucherResponse struct {
	TotalPenukaran string `json:"total_penukaran"`
	Persentase     string `json:"persentase"`
	Status         string `json:"status"`
}

type GetCountReportingResponse struct {
	TotalReporting string `json:"total_report"`
	Persentase     string `json:"persentase"`
	Status         string `json:"status"`
}

type GetCountTrashExchangeResponse struct {
	TotalTrashExchange string `json:"total_recycle"`
	Persentase         string `json:"persentase"`
	Status             string `json:"status"`
}

type GetCountScaleTypeResponse struct {
	PersentaseLargeScale string `json:"large_scale"`
	PersentaseSmallScale string `json:"small_scale"`
}
