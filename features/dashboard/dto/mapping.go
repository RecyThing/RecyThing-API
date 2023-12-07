package dto

import "recything/utils/dashboard"

func MapToGetCountUserResponse(input dashboard.GetCountUser) GetCountUserResponse {
	return GetCountUserResponse{
		TotalUserActive: input.TotalPenggunaAktif,
		Percentage:      input.Persentase,
		Status:          input.Status,
	}
}

func MapToGetCountExchangeVoucherResponse(input dashboard.GetCountExchangeVoucher) GetCountExchangeVoucherResponse {
	return GetCountExchangeVoucherResponse{
		TotalExchange: input.TotalPenukaran,
		Percentage:    input.Persentase,
		Status:        input.Status,
	}
}

func MapToGetCountReportingResponse(input dashboard.GetCountReporting) GetCountReportingResponse {
	return GetCountReportingResponse{
		TotalReporting: input.TotalReporting,
		Percentage:     input.Persentase,
		Status:         input.Status,
	}
}

func MapToGetCountTrashExchangeResponse(input dashboard.GetCountTrashExchange) GetCountTrashExchangeResponse {
	return GetCountTrashExchangeResponse{
		TotalTrashExchange: input.TotalTrashExchange,
		Percentage:         input.Persentase,
		Status:             input.Status,
	}
}

func MapToGetCountPersentaseScalaReportingResponse(input dashboard.GetCountScaleType) GetCountScaleTypeResponse {
	return GetCountScaleTypeResponse{
		PercentageLargeScale: input.PersentaseLargeScale,
		PercentageSmallScale: input.PersentaseSmallScale,
	}
}

func MapToGetWeeklyStatsResponse(input dashboard.WeeklyStats) WeeklyStatsResponse {
	return WeeklyStatsResponse{
		Week:      input.Week,
		TrashType: input.Trash,
		ScaleType: input.Scala,
	}
}

func ListMapToWeeklyStatsResponses(stats []dashboard.WeeklyStats) []WeeklyStatsResponse {
	var responses []WeeklyStatsResponse
	for _, stat := range stats {
		responses = append(responses, MapToGetWeeklyStatsResponse(stat))
	}
	return responses
}

func MapToGetUserRankingResponse(input dashboard.UserRanking) UserRankingResponse {
	return UserRankingResponse{
		Name:  input.Name,
		Email: input.Email,
		Point: input.Point,
	}
}

func ListMapToGetUserRankingResponse(rankingResult []dashboard.UserRanking) []UserRankingResponse {
	var rankingResponse []UserRankingResponse
	for _, userRanking := range rankingResult {
		rankingResponse = append(rankingResponse, MapToGetUserRankingResponse(userRanking))
	}
	return rankingResponse
}

func MapToCombinedResponse(
	userActiveResult dashboard.GetCountUser,
	voucherResult dashboard.GetCountExchangeVoucher,
	reportResult dashboard.GetCountReporting,
	trashExchangeResult dashboard.GetCountTrashExchange,
	scalaResult dashboard.GetCountScaleType,
	rankingResult []dashboard.UserRanking,
) map[string]interface{} {
	userActiveResponse := MapToGetCountUserResponse(userActiveResult)
	voucherResponse := MapToGetCountExchangeVoucherResponse(voucherResult)
	reportResponse := MapToGetCountReportingResponse(reportResult)
	trashExchangeResponse := MapToGetCountTrashExchangeResponse(trashExchangeResult)
	scalaResponse := MapToGetCountPersentaseScalaReportingResponse(scalaResult)
	rankingResponse := ListMapToGetUserRankingResponse(rankingResult)

	combinedResponse := map[string]interface{}{
		"user_active": userActiveResponse,
		"exchange":    voucherResponse,
		"report":      reportResponse,
		"recycle":     trashExchangeResponse,
		"scale":       scalaResponse,
		"ranking":     rankingResponse,
	}

	return combinedResponse
}
