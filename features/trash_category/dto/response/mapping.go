package response

import (
	"recything/features/trash_category/entity"
)

func CoreTrashCategoryToReponseTrashCategory(trash entity.TrashCategoryCore) TrashCategory {
	return TrashCategory{
		ID:        trash.ID,
		TrashType: trash.TrashType,
		Point:     trash.Point,
		Unit:    trash.Unit,
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

func CoreTrashCategoryToReponseTrashCategoryArticle(trash entity.TrashCategoryCore) TrashCategoryArticle {
	return TrashCategoryArticle{
		ID:        trash.ID,
		TrashType: trash.TrashType,
	}
}

func ListCoreTrashCategoryToReponseTrashCategoryArticle(trash []entity.TrashCategoryCore) []TrashCategoryArticle {
	list := []TrashCategoryArticle{}
	for _, v := range trash {
		result := CoreTrashCategoryToReponseTrashCategoryArticle(v)
		list = append(list, result)
	}
	return list
}
