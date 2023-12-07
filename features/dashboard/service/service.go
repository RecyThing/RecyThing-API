package service

import (
	"recything/features/dashboard/entity"
	"recything/utils/dashboard"
	"recything/utils/helper"
	"time"
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
func (ds *dashboardService) Dashboard() (dashboard.GetCountUser, dashboard.GetCountExchangeVoucher, dashboard.GetCountReporting, dashboard.GetCountTrashExchange, dashboard.GetCountScaleType, []dashboard.UserRanking, error) {
	users, reports, err := ds.dashboardRepository.CountUserActive()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, err
	}

	usersLastMonth, reportsLastMonth, err := ds.dashboardRepository.CountUserActiveLastMonth()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, err
	}

	userResult, err := dashboard.CalculateAndMapUserStats(users, usersLastMonth, reports, reportsLastMonth)
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, err
	}

	voucherThisMonth, voucherLastMonth, err := ds.dashboardRepository.CountVoucherExchanges()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, err
	}

	totalReportsThisMonth, totalReportsLastMonth, err := ds.dashboardRepository.CountReports()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, err
	}

	totalTrashThisMonth, totalTrashLastMonth, err := ds.dashboardRepository.CountTrashExchanges()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, err
	}

	totalLargeScale, totalSmallScale, err := ds.dashboardRepository.CountScaleTypes()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, err
	}

	pointUsers, err := ds.dashboardRepository.GetUserRanking()
	if err != nil {
		return dashboard.GetCountUser{}, dashboard.GetCountExchangeVoucher{}, dashboard.GetCountReporting{}, dashboard.GetCountTrashExchange{}, dashboard.GetCountScaleType{}, nil, err
	}

	userRanking := dashboard.MapUserRanking(pointUsers)
	trashResult := dashboard.MapToGetCountTrashExchange(len(totalTrashThisMonth), len(totalTrashLastMonth))
	reportResult := dashboard.MapToGetCountReporting(len(totalReportsThisMonth), len(totalReportsLastMonth))
	voucherResult := dashboard.MapToGetCountExchangeVoucher(len(voucherThisMonth), len(voucherLastMonth))
	scalaResult := dashboard.MapToGetCountScaleTypePercentage(len(totalLargeScale), len(totalSmallScale))

	return userResult, voucherResult, reportResult, trashResult, scalaResult, userRanking, nil
}

// CountWeeklyTrashAndScalaTypes implements entity.DashboardServiceInterface.
func (ds *dashboardService) CountWeeklyTrashAndScalaTypes() ([]dashboard.WeeklyStats, error) {
	trashAndScalaTypes, err := ds.dashboardRepository.CountWeeklyTrashAndScalaTypes()
	if err != nil {
		return nil, err
	}

	// Menghitung jumlah minggu dalam bulan ini
	startOfMonth := time.Now().AddDate(0, 0, -time.Now().Day()+1)
	year, month, _ := startOfMonth.Date()
	weeksInMonth := helper.GetWeeksInMonth(year, month)
	weeklyStats := make([]dashboard.WeeklyStats, weeksInMonth)

	for i := 0; i < weeksInMonth; i++ {
		// Menghitung awal dan akhir minggu
		weekStartDate := startOfMonth.AddDate(0, 0, 7*i)
		weekEndDate := startOfMonth.AddDate(0, 0, 7*(i+1))

		// Filter data yang berada dalam rentang waktu minggu ini
		filteredData := dashboard.FilterDataByDate(trashAndScalaTypes, weekStartDate, weekEndDate)

		// Hitung jumlah data trash_type dan scala_type
		trashCount, scalaCount := dashboard.CountTrashAndScalaTypes(filteredData)

		// Set nilai WeeklyStats
		weeklyStats[i].Week = i + 1
		weeklyStats[i].Trash = trashCount
		weeklyStats[i].Scala = scalaCount
	}

	return weeklyStats, nil
}

