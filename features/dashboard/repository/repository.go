package repository

import (
	"recything/features/dashboard/entity"
	report "recything/features/report/entity"
	voucher "recything/features/voucher/entity"
	modelReport "recything/features/report/model"
	modelTrash "recything/features/trash_exchange/model"
	user "recything/features/user/entity"
	modelUser "recything/features/user/model"
	modelVoucher "recything/features/voucher/model"
	"time"

	"gorm.io/gorm"
)

type dashboardRepository struct {
	db *gorm.DB
}

func NewDashboardRepository(db *gorm.DB) entity.DashboardRepositoryInterface {
	return &dashboardRepository{
		db: db,
	}
}

// CountUserActive implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountUserActive() ([]user.UsersCore, []report.ReportCore, error) {
	now := time.Now()

	// Cari pengguna yang diupdate dalam bulan ini
	users := []modelUser.Users{}
	err := dr.db.Where("MONTH(updated_at) = ? AND YEAR(updated_at) = ?", now.Month(), now.Year()).Find(&users).Error
	if err != nil {
		return nil, nil, err
	}

	// Cari laporan yang dibuat dalam bulan ini
	reports := []modelReport.Report{}
	err = dr.db.Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", now.Month(), now.Year()).Find(&reports).Error
	if err != nil {
		return nil, nil, err
	}

	// Memetakan data model ke core
	mappedUsers := user.ListUserModelToUserCore(users)
	mappedReports := report.ListReportModelToReportCore(reports)

	return mappedUsers, mappedReports, nil
}

// CountUserActiveLastMonth implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountUserActiveLastMonth() ([]user.UsersCore, []report.ReportCore, error) {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)

	// Cari pengguna yang diupdate dalam bulan ini
	users := []modelUser.Users{}
	err := dr.db.Where("MONTH(updated_at) = ? AND YEAR(updated_at) = ?", lastMonth.Month(), lastMonth.Year()).Find(&users).Error
	if err != nil {
		return nil, nil, err
	}

	// Cari laporan yang dibuat dalam bulan ini
	reports := []modelReport.Report{}
	err = dr.db.Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", lastMonth.Month(), lastMonth.Year()).Find(&reports).Error
	if err != nil {
		return nil, nil, err
	}

	// Memetakan data model ke core
	mappedUsers := user.ListUserModelToUserCore(users)
	mappedReports := report.ListReportModelToReportCore(reports)

	return mappedUsers, mappedReports, nil
}

// CountVoucherExchanges implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountVoucherExchanges() ([]voucher.ExchangeVoucherCore, []voucher.ExchangeVoucherCore, error) {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)

	// Ambil data pertukaran voucher bulan ini
	var exchangesThisMonth []modelVoucher.ExchangeVoucher
	if err := dr.db.Model(&modelVoucher.ExchangeVoucher{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", now.Month(), now.Year()).
		Find(&exchangesThisMonth).Error; err != nil {
		return nil, nil, err
	}

	// Ambil data pertukaran voucher bulan lalu
	var exchangesLastMonth []modelVoucher.ExchangeVoucher
	if err := dr.db.Model(&modelVoucher.ExchangeVoucher{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", lastMonth.Month(), lastMonth.Year()).
		Find(&exchangesLastMonth).Error; err != nil {
		return nil, nil, err
	}

	// Konversi dari model ke core
	coreThisMonth := voucher.ListModelExchangeVoucherToCoreExchangeVoucher(exchangesThisMonth)
	coreLastMonth := voucher.ListModelExchangeVoucherToCoreExchangeVoucher(exchangesLastMonth)

	return coreThisMonth, coreLastMonth, nil
}

// CountReports implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountReports() (int64, int64, error) {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)

	// Hitung total pelaporan bulan ini
	var totalThisMonth int64
	if err := dr.db.Model(&modelReport.Report{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", now.Month(), now.Year()).
		Count(&totalThisMonth).Error; err != nil {
		return 0, 0, err
	}

	// Hitung total pelaporan bulan lalu
	var totalLastMonth int64
	if err := dr.db.Model(&modelReport.Report{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", lastMonth.Month(), lastMonth.Year()).
		Count(&totalLastMonth).Error; err != nil {
		return 0, 0, err
	}

	return totalThisMonth, totalLastMonth, nil
}

// CountTrashExchanges implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountTrashExchanges() (int64, int64, error) {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)

	// Hitung total TrashExchange bulan ini
	var totalThisMonth int64
	if err := dr.db.Model(&modelTrash.TrashExchange{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", now.Month(), now.Year()).
		Count(&totalThisMonth).Error; err != nil {
		return 0, 0, err
	}

	// Hitung total TrashExchange bulan lalu
	var totalLastMonth int64
	if err := dr.db.Model(&modelTrash.TrashExchange{}).
		Where("MONTH(created_at) = ? AND YEAR(created_at) = ?", lastMonth.Month(), lastMonth.Year()).
		Count(&totalLastMonth).Error; err != nil {
		return 0, 0, err
	}

	return totalThisMonth, totalLastMonth, nil
}

// CountScaleTypes implements entity.DashboardRepositoryInterface.
func (dr *dashboardRepository) CountScaleTypes() (int64, int64, error) {
	var totalLargeScale int64
	if err := dr.db.Model(&modelReport.Report{}).
		Where("scale_type = ?", "skala besar").
		Count(&totalLargeScale).Error; err != nil {
		return 0, 0, err
	}

	// Hitung total pelaporan skala kecil
	var totalSmallScale int64
	if err := dr.db.Model(&modelReport.Report{}).
		Where("scale_type = ?", "skala kecil").
		Count(&totalSmallScale).Error; err != nil {
		return 0, 0, err
	}

	return totalLargeScale, totalSmallScale, nil
}
