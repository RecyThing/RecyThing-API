package dashboard

import (
	"fmt"
	"math"
)

type GetCountUser struct {
	TotalPenggunaAktif string
	Persentase         string
	Status             string
}

type GetCountExchangeVoucher struct {
	TotalPenukaran string
	Persentase     string
	Status         string
}

type GetCountReporting struct {
	TotalReporting string
	Persentase     string
	Status         string
}

type GetCountTrashExchange struct {
	TotalTrashExchange string
	Persentase         string
	Status             string
}

type GetCountScaleType struct {
	PersentaseLargeScale string
	PersentaseSmallScale string
}

func MapToGetCountUser(totalAktifBulanIni int, persentasePerubahan float64, status string) GetCountUser {
	return GetCountUser{
		TotalPenggunaAktif: fmt.Sprintf("%d", totalAktifBulanIni),
		Persentase:         fmt.Sprintf("%.2f", persentasePerubahan),
		Status:             status,
	}
}

func MapToGetCountExchangeVoucher(totalThisMonth, totalLastMonth int) GetCountExchangeVoucher {
	var persentasePerubahanVoucher float64
	if totalLastMonth > 0 {
		persentasePerubahanVoucher = float64(totalThisMonth-totalLastMonth) / float64(totalLastMonth) * 100
	} else {
		persentasePerubahanVoucher = 0
	}

	var statusVoucher string
	if persentasePerubahanVoucher > 0 {
		statusVoucher = "naik"
	} else if persentasePerubahanVoucher < 0 {
		statusVoucher = "turun"
	} else {
		statusVoucher = "tetap"
	}

	persentasePerubahanVoucher = math.Abs(persentasePerubahanVoucher)

	// Buat map hasil untuk pertukaran voucher
	result := GetCountExchangeVoucher{
		TotalPenukaran: fmt.Sprintf("%d", totalThisMonth),
		Persentase:     fmt.Sprintf("%.2f", persentasePerubahanVoucher),
		Status:         statusVoucher,
	}

	return result
}

func MapToGetCountReporting(totalThisMonth int64, totalLastMonth int64) GetCountReporting {
	var persentasePerubahanReporting float64
	if totalLastMonth > 0 {
		persentasePerubahanReporting = float64(totalThisMonth-totalLastMonth) / float64(totalLastMonth) * 100
	} else {
		persentasePerubahanReporting = 0
	}

	var statusReporting string
	if persentasePerubahanReporting > 0 {
		statusReporting = "naik"
	} else if persentasePerubahanReporting < 0 {
		statusReporting = "turun"
	} else {
		statusReporting = "tetap"
	}

	persentasePerubahanReporting = math.Abs(persentasePerubahanReporting)

	// Buat map hasil untuk pertukaran voucher
	result := GetCountReporting{
		TotalReporting: fmt.Sprintf("%d", totalThisMonth),
		Persentase:     fmt.Sprintf("%.2f", persentasePerubahanReporting),
		Status:         statusReporting,
	}

	return result
}

// MapToGetCountTrashExchange membuat objek GetCountTrashExchange dari total TrashExchange.
func MapToGetCountTrashExchange(totalThisMonth int64, totalLastMonth int64) GetCountTrashExchange {
	var persentasePerubahanTrash float64
	if totalLastMonth > 0 {
		persentasePerubahanTrash = float64(totalThisMonth-totalLastMonth) / float64(totalLastMonth) * 100
	} else {
		persentasePerubahanTrash = 0
	}

	var statusTrash string
	if persentasePerubahanTrash > 0 {
		statusTrash = "naik"
	} else if persentasePerubahanTrash < 0 {
		statusTrash = "turun"
	} else {
		statusTrash = "tetap"
	}

	persentasePerubahanTrash = math.Abs(persentasePerubahanTrash)

	// Buat map hasil untuk pertukaran voucher
	result := GetCountTrashExchange{
		TotalTrashExchange: fmt.Sprintf("%d", totalThisMonth),
		Persentase:         fmt.Sprintf("%.2f", persentasePerubahanTrash),
		Status:             statusTrash,
	}

	return result
}

// MapToGetCountScaleTypePercentage membuat objek GetCountScaleType dengan persentase pelaporan skala besar dan skala kecil.
func MapToGetCountScaleTypePercentage(totalLargeScale int64, totalSmallScale int64) GetCountScaleType {
	totalReports := totalLargeScale + totalSmallScale

	var percentageLargeScale float64
	if totalReports > 0 {
		percentageLargeScale = float64(totalLargeScale) / float64(totalReports) * 100
	}

	var percentageSmallScale float64
	if totalReports > 0 {
		percentageSmallScale = float64(totalSmallScale) / float64(totalReports) * 100
	}

	// Buat map hasil untuk persentase pelaporan skala besar dan skala kecil
	result := GetCountScaleType{
		PersentaseLargeScale: fmt.Sprintf("%.2f", percentageLargeScale),
		PersentaseSmallScale: fmt.Sprintf("%.2f", percentageSmallScale),
	}

	return result
}
