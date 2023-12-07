package dto

import "recything/utils/dashboard"

func MapToGetCountUserResponse(input dashboard.GetCountUser) GetCountUserResponse {
	return GetCountUserResponse{
		TotalPenggunaAktif: input.TotalPenggunaAktif,
		Persentase:         input.Persentase,
		Status:             input.Status,
	}
}

func MapToGetCountExchangeVoucherResponse(input dashboard.GetCountExchangeVoucher) GetCountExchangeVoucherResponse {
	return GetCountExchangeVoucherResponse{
		TotalPenukaran: input.TotalPenukaran,
		Persentase:     input.Persentase,
		Status:         input.Status,
	}
}

func MapToGetCountReportingResponse(input dashboard.GetCountReporting) GetCountReportingResponse {
	return GetCountReportingResponse{
		TotalReporting: input.TotalReporting,
		Persentase:     input.Persentase,
		Status:         input.Status,
	}
}

func MapToGetCountTrashExchangeResponse(input dashboard.GetCountTrashExchange) GetCountTrashExchangeResponse {
	return GetCountTrashExchangeResponse{
		TotalTrashExchange: input.TotalTrashExchange,
		Persentase:         input.Persentase,
		Status:             input.Status,
	}
}

func MapToGetCountPersentaseScalaReportingResponse(input dashboard.GetCountScaleType) GetCountScaleTypeResponse {
	return GetCountScaleTypeResponse{
		PersentaseLargeScale: input.PersentaseLargeScale,
		PersentaseSmallScale: input.PersentaseSmallScale,
	}
}

func MapToCombinedResponse(
	userActiveResult dashboard.GetCountUser,
	voucherResult dashboard.GetCountExchangeVoucher,
	reportResult dashboard.GetCountReporting,
	trashExchangeResult dashboard.GetCountTrashExchange,
	scalaResult dashboard.GetCountScaleType,
) map[string]interface{} {
	userActiveResponse := MapToGetCountUserResponse(userActiveResult)
	voucherResponse := MapToGetCountExchangeVoucherResponse(voucherResult)
	reportResponse := MapToGetCountReportingResponse(reportResult)
	trashExchangeResponse := MapToGetCountTrashExchangeResponse(trashExchangeResult)
	scalaResponse := MapToGetCountPersentaseScalaReportingResponse(scalaResult)

	combinedResponse := map[string]interface{}{
		"user_active":     userActiveResponse,
		"voucher":         voucherResponse,
		"report":          reportResponse,
		"recycle":         trashExchangeResponse,
		"scala_persentase": scalaResponse,
	}

	return combinedResponse
}
