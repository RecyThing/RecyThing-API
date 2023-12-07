package service

import (
	"math"
	"recything/features/dashboard/entity"
	"recything/utils/dashboard"
)

type dashboardService struct {
	dashboardRepository entity.DashboardRepositoryInterface
}

func NewDashboardService(dashboardRepo entity.DashboardRepositoryInterface) entity.DashboardServiceInterface {
	return &dashboardService{
		dashboardRepository: dashboardRepo,
	}
}

// CountUserActive implements entity.DashboardUsecaseInterface.
func (ds *dashboardService) CountUserActive() (dashboard.GetCountUser, dashboard.GetCountExchangeVoucher, dashboard.GetCountReporting, dashboard.GetCountTrashExchange, dashboard.GetCountScaleType, error) {
	users, reports, err := ds.dashboardRepository.CountUserActive()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, err
	}

	usersLastMonth, reportsLastMonth, err := ds.dashboardRepository.CountUserActiveLastMonth()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, err
	}

	penggunaAktif := make(map[string]struct{})
	for _, u := range users {
		penggunaAktif[u.Id] = struct{}{}
	}

	for _, r := range reports {
		if _, exist := penggunaAktif[r.UserId]; !exist {
			penggunaAktif[r.UserId] = struct{}{}
		}
	}

	totalAktifBulanIni := len(penggunaAktif)

	penggunaAktifBulanLalu := make(map[string]struct{})
	for _, u := range usersLastMonth {
		penggunaAktifBulanLalu[u.Id] = struct{}{}
	}

	for _, r := range reportsLastMonth {
		if _, exist := penggunaAktifBulanLalu[r.UserId]; !exist {
			penggunaAktifBulanLalu[r.UserId] = struct{}{}
		}
	}

	totalAktifBulanLalu := len(penggunaAktifBulanLalu)

	var persentasePerubahan float64
	if totalAktifBulanLalu > 0 {
		persentasePerubahan = float64(totalAktifBulanIni-totalAktifBulanLalu) / float64(totalAktifBulanLalu) * 100
	} else {
		persentasePerubahan = 0
	}

	var status string
	if persentasePerubahan > 0 {
		status = "naik"
	} else if persentasePerubahan < 0 {
		status = "turun"
	} else {
		status = "tetap"
	}

	persentasePerubahan = math.Abs(persentasePerubahan)

	totalVoucherThisMonth, totalVoucherLastMonth, err := ds.dashboardRepository.CountVoucherExchanges()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, err
	}

	totalReportsThisMonth, totalReportsLastMonth, err := ds.dashboardRepository.CountReports()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, err
	}

	totalTrashThisMonth, totalTrashLastMonth, err := ds.dashboardRepository.CountTrashExchanges()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, err
	}

	totalLargeScale, totalSmallScale, err := ds.dashboardRepository.CountScaleTypes()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, err
	}

	totalThisMonth := len(totalVoucherThisMonth)
	totalLastMonth := len(totalVoucherLastMonth)

	trashResult := dashboard.MapToGetCountTrashExchange(totalTrashThisMonth, totalTrashLastMonth)
	reportResult := dashboard.MapToGetCountReporting(totalReportsThisMonth, totalReportsLastMonth)
	voucherResult := dashboard.MapToGetCountExchangeVoucher(totalThisMonth, totalLastMonth)
	scalaResult := dashboard.MapToGetCountScaleTypePercentage(totalLargeScale, totalSmallScale)
	result := dashboard.MapToGetCountUser(totalAktifBulanIni, persentasePerubahan, status)
	return result, voucherResult, reportResult, trashResult, scalaResult, nil
}
