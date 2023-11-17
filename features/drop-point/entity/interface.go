package entity

type DropPointRepositoryInterface interface {
	CreateDropPoint(data DropPointCore) (DropPointCore, error)
	UpdateDropPointById(id string, data DropPointCore) (DropPointCore, error)
	GetDropPointById(id string) (DropPointCore, error)
	GetAllDropPoint() ([]DropPointCore, error)
	DeleteDropPointById(id string) error
}

type DropPointServiceInterface interface {
	CreateDropPoint(data DropPointCore) (DropPointCore, error)
	UpdateDropPointById(id string, data DropPointCore) (DropPointCore, error)
	GetDropPointById(id string) (DropPointCore, error)
	GetAllDropPoint() ([]DropPointCore, error)
	DeleteDropPointById(id string) error
}
