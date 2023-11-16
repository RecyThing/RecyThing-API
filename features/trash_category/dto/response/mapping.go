package response

import (
	"recything/features/trash_category/entity"
)

func CoreTrashCategoryToReponseTrashCategory(trash entity.TrashCategoryCore) TrashCategory {
	return TrashCategory{
		ID:        trash.ID,
		TrashType: trash.TrashType,
		Point:     trash.Point,
		Satuan:    trash.Satuan,
		CreatedAt: trash.CreatedAt,
		UpdatedAt: trash.UpdatedAt,
	}
}

func ListCoreTrashCategoryToReponseTrashCategory(trash []entity.TrashCategoryCore) []TrashCategory {
	list := []TrashCategory{}
	for _, v := range trash {
		result := CoreTrashCategoryToReponseTrashCategory(v)
		list = append(list, result)
	}
	return list
}
