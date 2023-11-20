package request

import (
	"recything/features/trash_category/entity"
)

func RequestTrashCategoryToCoreTrashCategory(trash TrashCategory) entity.TrashCategoryCore {
	return entity.TrashCategoryCore{
		TrashType: trash.TrashType,
		Point:     trash.Point,
		Unit:      trash.Unit,
	}
}

func TrashCategoryrequestToTrashCategoryCore(category TrashCategory) entity.TrashCategoryCore {
	return entity.TrashCategoryCore{
		TrashType: category.TrashType,
	}
}

func ListTrashCategoryRequestToTrashCategoryCore(category []TrashCategory) []entity.TrashCategoryCore {
	listTrashCategory := []entity.TrashCategoryCore{}
	for _, v := range category {
		TrashCategory := TrashCategoryrequestToTrashCategoryCore(v)
		listTrashCategory = append(listTrashCategory, TrashCategory)
	}

	return listTrashCategory
}
