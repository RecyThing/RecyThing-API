package entity

type TrashCategoryRepositoryInterface interface {
	Create(data TrashCategoryCore)  error
	Update(idTrash string, data TrashCategoryCore) (TrashCategoryCore, error)
	Delete(idTrash string) error
	GetAll(page string, limit string) ([]TrashCategoryCore, PagnationInfo, error)
	GetById(idTrash string) (TrashCategoryCore, error)
}

type TrashCategoryServiceInterface interface {
	CreateCategory(data TrashCategoryCore) error
	UpdateCategory(idTrash string, data TrashCategoryCore) (TrashCategoryCore, error)
	DeleteCategory(idTrash string) error
	GetAllCategory(page, limit string) ([]TrashCategoryCore, PagnationInfo, error)
	GetById(idTrash string) (TrashCategoryCore, error)
}
