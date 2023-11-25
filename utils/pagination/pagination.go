package pagination

import (
	"fmt"

	"math"
)

type PageInfo struct {
	Limit       int `json:"limit"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
}

func CalculateData(totalCount, limitInt, pageInt int) PageInfo {
	lastPage := int(math.Ceil(float64(totalCount) / float64(limitInt)))

	paginationInfo := PageInfo{
		Limit:       limitInt,
		CurrentPage: pageInt,
		LastPage:    lastPage,
	}
	return paginationInfo
}

func PaginationMessage(paginationInfo PageInfo, totalData int) string {
	limit := paginationInfo.Limit
	currentPage := paginationInfo.CurrentPage

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	startIndex := (currentPage-1)*limit + 1
	endIndex := min(startIndex+limit-1, totalData)

	responseMessage := fmt.Sprintf("menampilkan data %d sampai %d dari %d data", startIndex, endIndex, totalData)
	return responseMessage
}
