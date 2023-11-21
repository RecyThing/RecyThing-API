package entity

type TrashCategoryRepositoryInterface interface {
	Create(data TrashCategoryCore) error
	Update(idTrash string, data TrashCategoryCore) (TrashCategoryCore, error)
	Delete(idTrash string) error
	GetById(idTrash string) (TrashCategoryCore, error)
	FindAll(page, limit int, trashType string)([]TrashCategoryCore, PagnationInfo, error)
}

type TrashCategoryServiceInterface interface {
	CreateCategory(data TrashCategoryCore) error
	UpdateCategory(idTrash string, data TrashCategoryCore) (TrashCategoryCore, error)
	DeleteCategory(idTrash string) error
	GetAllCategory(page, trashType, limit string) ([]TrashCategoryCore, PagnationInfo, error)
	GetById(idTrash string) (TrashCategoryCore, error)
}
