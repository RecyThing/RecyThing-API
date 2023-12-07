package dashboard

import (
	"fmt"
	"math"
	report "recything/features/report/entity"
	user "recything/features/user/entity"
	"time"
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

type UserRanking struct {
	Name  string
	Email string
	Point int
}

type WeeklyStats struct {
	Week  int
	Trash int
	Scala int
}

func CalculateAndMapUserStats(users, usersLastMonth []user.UsersCore, reports, reportsLastMonth []report.ReportCore) (GetCountUser, error) {
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

	persentasePerubahanInt := int(math.Round(persentasePerubahan))

	result := GetCountUser{
		TotalPenggunaAktif: fmt.Sprintf("%d", totalAktifBulanIni),
		Persentase:         fmt.Sprintf("%d", persentasePerubahanInt),
		Status:             status,
	}

	return result, nil
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

	persentasePerubahanInt := int(math.Round(persentasePerubahanVoucher))

	// Buat map hasil untuk pertukaran voucher
	result := GetCountExchangeVoucher{
		TotalPenukaran: fmt.Sprintf("%d", totalThisMonth),
		Persentase:     fmt.Sprintf("%d", persentasePerubahanInt),
		Status:         statusVoucher,
	}

	return result
}

func MapToGetCountReporting(totalThisMonth int, totalLastMonth int) GetCountReporting {
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

	persentasePerubahanInt := int(math.Round(persentasePerubahanReporting))

	// Buat map hasil untuk pertukaran voucher
	result := GetCountReporting{
		TotalReporting: fmt.Sprintf("%d", totalThisMonth),
		Persentase:     fmt.Sprintf("%d", persentasePerubahanInt),
		Status:         statusReporting,
	}

	return result
}

// MapToGetCountTrashExchange membuat objek GetCountTrashExchange dari total TrashExchange.
func MapToGetCountTrashExchange(totalThisMonth int, totalLastMonth int) GetCountTrashExchange {
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

	persentasePerubahanInt := int(math.Round(persentasePerubahanTrash))
	
	// Buat map hasil untuk pertukaran voucher
	result := GetCountTrashExchange{
		TotalTrashExchange: fmt.Sprintf("%d", totalThisMonth),
		Persentase:         fmt.Sprintf("%d", persentasePerubahanInt),
		Status:             statusTrash,
	}

	return result
}

// MapToGetCountScaleTypePercentage membuat objek GetCountScaleType dengan persentase pelaporan skala besar dan skala kecil.
func MapToGetCountScaleTypePercentage(totalLargeScale int, totalSmallScale int) GetCountScaleType {
	totalReports := totalLargeScale + totalSmallScale

	var percentageLargeScale float64
	if totalReports > 0 {
		percentageLargeScale = float64(totalLargeScale) / float64(totalReports) * 100
	}

	var percentageSmallScale float64
	if totalReports > 0 {
		percentageSmallScale = float64(totalSmallScale) / float64(totalReports) * 100
	}

	persentaseLargeScalePerubahanInt := int(math.Round(percentageLargeScale))
	persentaseSmallScalePerubahanInt := int(math.Round(percentageSmallScale))

	// Buat map hasil untuk persentase pelaporan skala besar dan skala kecil
	result := GetCountScaleType{
		PersentaseLargeScale: fmt.Sprintf("%d", persentaseLargeScalePerubahanInt),
		PersentaseSmallScale: fmt.Sprintf("%d", persentaseSmallScalePerubahanInt),
	}

	return result
}

func MapUserRanking(users []user.UsersCore) []UserRanking {
	var userRanking []UserRanking
	for _, user := range users {
		userRanking = append(userRanking, UserRanking{
			Name:  user.Fullname,
			Email: user.Email,
			Point: user.Point,
		})
	}
	return userRanking
}

// Function untuk memfilter data berdasarkan range tanggal
func FilterDataByDate(data []report.ReportCore, startDate, endDate time.Time) []report.ReportCore {
    var filteredData []report.ReportCore
    for _, entry := range data {
        if entry.CreatedAt.After(startDate) && entry.CreatedAt.Before(endDate) {
            filteredData = append(filteredData, entry)
        }
    }
    return filteredData
}

// Function untuk menghitung jumlah data trash_type dan scala_type
func CountTrashAndScalaTypes(data []report.ReportCore) (int, int) {
    var trashCount, scalaCount int
    for _, entry := range data {
        if entry.TrashType != "" {
            trashCount++
        }
        if entry.ScaleType != "" {
            scalaCount++
        }
    }
    return trashCount, scalaCount
}