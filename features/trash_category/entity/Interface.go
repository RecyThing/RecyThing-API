package entity

import "recything/utils/pagination"

type TrashCategoryRepositoryInterface interface {
	Create(data TrashCategoryCore) error
	Update(idTrash string, data TrashCategoryCore) (TrashCategoryCore, error)
	Delete(idTrash string) error
	GetById(idTrash string) (TrashCategoryCore, error)
	FindAll(page, limit int, trashType string) ([]TrashCategoryCore, pagination.PageInfo, error)
}

type TrashCategoryServiceInterface interface {
	CreateCategory(data TrashCategoryCore) error
	UpdateCategory(idTrash string, data TrashCategoryCore) (TrashCategoryCore, error)
	DeleteCategory(idTrash string) error
	GetAllCategory(page, trashType, limit string) ([]TrashCategoryCore, pagination.PageInfo, error)
	GetById(idTrash string) (TrashCategoryCore, error)
}
