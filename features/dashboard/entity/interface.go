package entity

import (
	report "recything/features/report/entity"
	user "recything/features/user/entity"
	voucher "recything/features/voucher/entity"
	"recything/utils/dashboard"
)

type DashboardRepositoryInterface interface {
	CountUserActive() ([]user.UsersCore, []report.ReportCore, error)
	CountUserActiveLastMonth() ([]user.UsersCore, []report.ReportCore, error)
	CountVoucherExchanges() ([]voucher.ExchangeVoucherCore, []voucher.ExchangeVoucherCore, error)
	CountReports() (int64, int64, error)
	CountTrashExchanges() (int64, int64, error)
	CountScaleTypes() (int64, int64, error)
}

type DashboardServiceInterface interface {
	CountUserActive() (dashboard.GetCountUser, dashboard.GetCountExchangeVoucher, dashboard.GetCountReporting, dashboard.GetCountTrashExchange, dashboard.GetCountScaleType, error)
}
