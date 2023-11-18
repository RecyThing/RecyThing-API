package helper

import (
	"math"
	"recything/features/trash_category/entity"
)

func CalculatePagination(totalCount, limitInt, pageInt int) entity.PagnationInfo {
	lastPage := int(math.Ceil(float64(totalCount) / float64(limitInt)))

	paginationInfo := entity.PagnationInfo{
		Limit:       limitInt,
		CurrentPage: pageInt,
		LastPage:    lastPage,
	}
	return paginationInfo
}
