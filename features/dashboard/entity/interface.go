package entity

import (
	report "recything/features/report/entity"
	user "recything/features/user/entity"
	voucher "recything/features/voucher/entity"
	trash "recything/features/trash_exchange/entity"
	"recything/utils/dashboard"
)

type DashboardRepositoryInterface interface {
	CountUserActive() ([]user.UsersCore, []report.ReportCore, error)
	CountUserActiveLastMonth() ([]user.UsersCore, []report.ReportCore, error)
	CountVoucherExchanges() ([]voucher.ExchangeVoucherCore, []voucher.ExchangeVoucherCore, error)
	CountReports() ([]report.ReportCore, []report.ReportCore, error)
	CountTrashExchanges() ([]trash.TrashExchangeCore, []trash.TrashExchangeCore, error)
	CountScaleTypes() ([]report.ReportCore, []report.ReportCore, error)
	GetUserRanking() ([]user.UsersCore, error)
	CountWeeklyTrashAndScalaTypes() ([]report.ReportCore, error)
}

type DashboardServiceInterface interface {
	Dashboard() (dashboard.GetCountUser, dashboard.GetCountExchangeVoucher, dashboard.GetCountReporting, dashboard.GetCountTrashExchange, dashboard.GetCountScaleType, []dashboard.UserRanking, error)
	CountWeeklyTrashAndScalaTypes() ([]dashboard.WeeklyStats, error)
}
