package request

import (
	"recything/features/trash_category/entity"
)

func RequestTrashCategoryToCoreTrashCategory(trash TrashCategory) entity.TrashCategoryCore {
	return entity.TrashCategoryCore{
		TrashType: trash.TrashType,
		Point:     trash.Point,
		Satuan:    trash.Satuan,
	}
}


